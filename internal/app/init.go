package app

import (
	"fmt"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/app/repository"
	"github.com/tcc-uniftec-5s/internal/app/service"
	usecase "github.com/tcc-uniftec-5s/internal/app/use_case"
	"github.com/tcc-uniftec-5s/internal/infra/database"
	"github.com/tcc-uniftec-5s/internal/infra/environment"
	logConfig "github.com/tcc-uniftec-5s/internal/infra/log"
	server "github.com/tcc-uniftec-5s/internal/interface/http"
	"github.com/tcc-uniftec-5s/internal/interface/http/controller"
	"github.com/tcc-uniftec-5s/internal/token"
)

type TracerLogger struct {
}

func (t TracerLogger) Log(msg string) {
	log.Info().Msg(msg)
}

func Init(rootdir string) {
	fmt.Println(rootdir)
	environment.LoadEnv(filepath.Join(rootdir, ".env"))
	logConfig.ConfigZeroLog(environment.Env.LogLevel)

	pgService, err := database.NewPostgres(database.Config{
		Host:     environment.Env.DbHost,
		Port:     environment.Env.DbPort,
		User:     environment.Env.DbUser,
		Password: environment.Env.DbPassword,
		Database: environment.Env.DbName,
	})
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("fail to start postgres connection")
	}

	migrationsPath := "file://" + filepath.Join(rootdir, "internal/infra/database/migration")
	err = database.MigrateUp(pgService, environment.Env.DbName, migrationsPath)
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("fail to migrate database")
	}

	txHandler := repository.NewTxHandler(pgService)
	credentialRepository := repository.NewCredentialRepository(pgService)
	accountRepository := repository.NewAccountRepository(pgService)
	userRepository := repository.NewUserRepository(pgService)
	sessionRepository := repository.NewSessionRepository(pgService)
	editionRepository := repository.NewEdition(pgService)
	prizeRepository := repository.NewPrize(pgService)
	teamRepository := repository.NewTeam(pgService)

	jwtMaker, err := token.NewJWTMaker(environment.Env.JWTSigningKey)
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("error generating new jwt maker")
	}

	credentialFactory := service.NewCredentialFactory(credentialRepository, jwtMaker)
	accountFactory := service.NewAccountFactory(accountRepository)
	userFactory := service.NewUserFactory(userRepository)
	sessionFactory := service.NewSessionFactory(sessionRepository)
	editionFactory := service.NewEditionFactory(editionRepository)
	prizeFactory := service.NewPrizeFactory(prizeRepository)
	teamFactory := service.NewTeamFactory(teamRepository)

	signupUseCase := usecase.NewSignup(txHandler, credentialFactory, accountFactory, userFactory)
	loginUseCase := usecase.NewLogin(txHandler, credentialFactory, sessionFactory)
	resetPasswordUseCase := usecase.NewResetPassword(txHandler, credentialFactory)
	createEditionUseCase := usecase.NewCreateEdition(txHandler, editionFactory, prizeFactory)
	listTeamlessUsersUseCase := usecase.NewListTeamlessUsers(txHandler, userFactory)
	createTeamUseCase := usecase.NewCreateTeam(txHandler, teamFactory, editionFactory)

	httpServer := server.New(
		fmt.Sprintf(":%s", "3000"),
		"tcc-uniftec-5s",
		environment.Env.JWTSigningKey,
	)

	accessValidator := controller.NewAccessValidator()

	controllers := []controller.Router{
		controller.NewSignupController(httpServer.Instance, signupUseCase),
		controller.NewLoginController(httpServer.Instance, loginUseCase),
		controller.NewResetPasswordController(httpServer.Instance, resetPasswordUseCase),
		controller.NewEdition(httpServer.Instance, httpServer.Restricted, accessValidator, createEditionUseCase),
		controller.NewUser(httpServer.Instance, listTeamlessUsersUseCase),
		controller.NewTeam(httpServer.Instance, httpServer.Restricted, accessValidator, createTeamUseCase),
	}

	registerControllersRoutes(controllers)

	go func() {
		if err := httpServer.Start(); err != nil {
			log.Fatal().Stack().Err(err).Msg("fail to start http server")
		}
	}()
}

func registerControllersRoutes(controllers []controller.Router) {
	for _, c := range controllers {
		c.RegisterRoutes()
	}
}
