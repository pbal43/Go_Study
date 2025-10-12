package lessons

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Lesson11() {
	// логирование

	//zero-log

	//log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02T15:04:05.999Z07:00"}).With().Timestamp().Logger()

	zerolog.SetGlobalLevel(zerolog.InfoLevel) // если установим так, то дебага не будет,
	// так как есть иерархия логов и покажутся только такой и более высокие уровни:
	// debug
	// info
	// warn
	// error
	// fatal

	log.Info().Msg("Info log")

	hello := "value"

	log.Info().Str("Info value", hello).Send()

	log.Error().Err(fmt.Errorf("testError")).Msg("Error log")

	log.Debug().Msg("Debug log")

	log.Warn().Msg("Warn log")

	log.Fatal().Msg("Fatal log") // exit status 1
}
