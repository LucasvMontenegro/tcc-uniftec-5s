package environment

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var Env EnvVars

const (
	OptionalKey = "optional"
)

type EnvVars struct {
	DbHost     string `env:"DB_HOST"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PSWD"`
	DbName     string `env:"DB_NAME"`
	DbPort     string `env:"DB_PORT"`

	Environment string `env:"ENVIRONMENT"`
	LogLevel    string `env:"LOG_LEVEL"`
	HttpPort    string `env:"HTTP_PORT"`
}

func LoadEnv(filenames ...string) {
	log.Info().Msgf("Loading env variables %s", filenames[0])

	err := godotenv.Load(filenames...)

	err = loadEnvVars(&Env)
	if err != nil {
		log.Fatal().Stack().Err(err).Send()
	}

	log.Debug().Interface("env", Env).Msg("Successfully loaded env variables")
}

func loadEnvVars(vars interface{}) error {

	pointr := reflect.ValueOf(vars)
	values := pointr.Elem()
	typeOfValues := values.Type()

	for i := 0; i < values.NumField(); i++ {
		value := values.Field(i).String()
		field := pointr.Elem().Field(i)
		fieldName := typeOfValues.Field(i).Name

		fieldKey := fieldName
		optional := false

		tag := typeOfValues.Field(i).Tag.Get("env")
		if tag != "" {
			tagOpts := strings.Split(tag, ",")
			fieldKey = tagOpts[0]
			keys := tagOpts[1:]
			for _, key := range keys {
				if key == OptionalKey {
					optional = true
					break
				}
			}
		}

		if field.CanSet() && value == "" {
			field.SetString(os.Getenv(fieldKey))
		}

		if !optional && field.String() == "" {
			return fmt.Errorf(`env "%s", fieldname "%s" must be defined`, fieldKey, fieldName)
		}
	}

	return nil
}
