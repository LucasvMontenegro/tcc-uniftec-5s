package database

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	pgsql "gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func (c Config) String() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		c.Host,
		c.User,
		c.Password,
		c.Database,
		c.Port,
	)
}

func NewPostgres(config Config) (*gorm.DB, error) {
	log.Info().Msg("starting database connection")

	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
	}

	db, err := gorm.Open(pgsql.Open(config.String()), &gormConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateUp(gormInstance *gorm.DB, database string, path string) error {
	log.Info().Msg("starting database migration up")

	sqlInstance, err := gormInstance.DB()
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(sqlInstance, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		path,
		database,
		driver,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

var ErrUniqueViolation = "unique_violation"
var ErrCodes = map[string]string{
	ErrUniqueViolation: "23505",
}
