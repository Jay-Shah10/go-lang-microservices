package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger = zerolog.New(os.Stdout)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func Info(message string) {
	log.Info().Msg(message)

}

func Debug(message string) {
	log.Debug().Msg(message)
}

func Error(message string, err error) {
	log.Err(err).Msg(message)
}
