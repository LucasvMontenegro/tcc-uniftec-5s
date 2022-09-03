package app

import (
	"fmt"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/tcc-uniftec-5s/internal/app/service"
	"github.com/tcc-uniftec-5s/internal/infra/constants"
	"github.com/tcc-uniftec-5s/internal/infra/database"
	"github.com/tcc-uniftec-5s/internal/infra/database/repository"
	"github.com/tcc-uniftec-5s/internal/infra/environment"
	logConfig "github.com/tcc-uniftec-5s/internal/infra/log"
	api "github.com/tcc-uniftec-5s/internal/interface/http"
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

	migrationsPath := "file://" + filepath.Join(rootdir, "internal/infra/database/migrations")
	err = database.MigrateUp(pgService, environment.Env.DbName, migrationsPath)
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("fail to migrate database")
	}

	sampleRepository := repository.NewSampleRepository(pgService)
	sampleService := service.NewSampleService(sampleRepository)
	apiService := api.NewService(fmt.Sprintf(":%s", environment.Env.HttpPort), constants.ServiceName, sampleService)

	go func() {
		if err := apiService.StartServer(); err != nil {
			log.Fatal().Stack().Err(err).Msg("fail to start http server")
		}
	}()
}
