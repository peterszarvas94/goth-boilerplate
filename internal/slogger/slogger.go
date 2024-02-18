package slogger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/samber/slog-multi"
)

type logger struct {
	Debug func(msg string, args ...any)
	Info  func(msg string, args ...any)
	Warn  func(msg string, args ...any)
	Error func(msg string, args ...any)
	Fatal func(msg string, args ...any)
}

const (
	levelFatal = slog.Level(12)
)

var levelNames = map[slog.Leveler]string{
	levelFatal: "FATAL",
}

var LogLevelFlag string

func Get() logger {
	var minLevel slog.Level

	switch strings.ToUpper(LogLevelFlag) {
	case "DEBUG":
		minLevel = slog.LevelDebug
	case "INFO":
		minLevel = slog.LevelInfo
	case "WARNING":
		minLevel = slog.LevelWarn
	case "ERROR":
		minLevel = slog.LevelError
	case "FATAL":
		minLevel = levelFatal
	default:
		minLevel = slog.LevelInfo
	}

	now := time.Now().UTC().Format("2006-01-02")

	logDir := "logs"

	// check if the logs directory exists
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		// create the logs directory in /tmp
		// this is so that we don't got any errors in tests
		os.Mkdir("/tmp/logs", 0755)
		logDir = "/tmp/logs"
	}

	logFile := fmt.Sprintf("server-%s.log", now)
	logPath := filepath.Join(logDir, logFile)

	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
	}

	// TODO: include request id
	all := newJSONHandler(file, minLevel)
	error := newJSONHandler(os.Stderr, slog.LevelError)

	s := *slog.New(slogmulti.Fanout(all, error))

	var ctx = context.Background()

	return logger{
		Debug: s.Debug,
		Info:  s.Info,
		Warn:  s.Warn,
		Error: s.Error,
		Fatal: func(msg string, args ...any) {
			s.Log(ctx, levelFatal, msg, args...)
			os.Exit(1)
		},
	}
}

func newJSONHandler(w io.Writer, level slog.Level) slog.Handler {
	return slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
		// search the custom log level name, like "FATAL" in this case
		ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
			if attr.Key == slog.LevelKey {
				level := attr.Value.Any().(slog.Level)
				levelLabel, exists := levelNames[level]
				if !exists {
					levelLabel = level.String()
				}

				attr.Value = slog.StringValue(levelLabel)
			}

			return attr
		},
	})
}
