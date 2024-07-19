package log

import (
	"log/slog"
	"os"
)

// Init initializes the global log/slog logger with the given log level and sets output format to JSON.
func Init(level slog.Level) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))

	slog.SetDefault(logger)
}
