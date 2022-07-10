package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type ILogger interface {
	Trace(message any, args ...any)
	Debug(message any, args ...any)
	Info(message any, args ...any)
	Warn(message any, args ...any)
	Error(message any, args ...any)
	Fatal(message any, args ...any)
}

type Logger struct {
	logger *zerolog.Logger
}

// New creates and configures logger
func New(level string) *Logger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	case "trace":
		l = zerolog.TraceLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC1123}

	skipFrameCount := 3
	logger := zerolog.New(output).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &Logger{
		logger: &logger,
	}
}

func (l *Logger) Trace(message any, args ...any) {
	l.msg("trace", message, args...)
}

func (l *Logger) Debug(message any, args ...any) {
	l.msg("debug", message, args...)
}

func (l *Logger) Info(message any, args ...any) {
	l.msg("info", message, args...)
}

func (l *Logger) Warn(message any, args ...any) {
	l.msg("warn", message, args...)
}

func (l *Logger) Error(message any, args ...any) {
	l.msg("error", message, args...)
}

func (l *Logger) Fatal(message any, args ...any) {
	l.msg("fatal", message, args...)
}

func (l *Logger) log(level, message string, args ...any) {
	switch level {
	case "trace":
		if len(args) == 0 {
			l.logger.Trace().Msg(message)
		} else {
			l.logger.Trace().Msgf(message, args...)
		}
	case "debug":
		if len(args) == 0 {
			l.logger.Debug().Msg(message)
		} else {
			l.logger.Debug().Msgf(message, args...)
		}
	case "warn":
		if len(args) == 0 {
			l.logger.Warn().Msg(message)
		} else {
			l.logger.Warn().Msgf(message, args...)
		}
	case "error":
		if len(args) == 0 {
			l.logger.Error().Msg(message)
		} else {
			l.logger.Error().Msgf(message, args...)
		}
	case "fatal":
		if len(args) == 0 {
			l.logger.Fatal().Msg(message)
		} else {
			l.logger.Fatal().Msgf(message, args...)
		}
	default:
		if len(args) == 0 {
			l.logger.Info().Msg(message)
		} else {
			l.logger.Info().Msgf(message, args...)
		}
	}
}

func (l *Logger) msg(level string, message any, args ...any) {
	switch msg := message.(type) {
	case error:
		l.log(level, msg.Error(), args...)
	case string:
		l.log(level, msg, args...)
	default:
		l.log(level, fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}
