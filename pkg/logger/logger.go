package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func InitLogger() {
	// create new logger with json handler
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // default logging level
	})

	// assign to global varible
	Log = slog.New(handler)

	// set the logger as the default logger
	slog.SetDefault(Log)

}

// Info logs an informational message
func Info(msg string, keysAndValues ...interface{}) {
	Log.Info(msg, keysAndValues...)
}

// Warn logs a warning message
func Warn(msg string, keysAndValues ...interface{}) {
	Log.Warn(msg, keysAndValues...)
}

// Error logs an error message
func Error(msg string, err error, keysAndValues ...interface{}) {
	Log.Error(msg, append(keysAndValues, "error", err)...)
}
