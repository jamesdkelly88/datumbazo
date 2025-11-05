package logging

import (
	"io"
	"log/slog"
)

func SetupLogger(w io.Writer, l slog.Level) {
	logger := slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: l.Level(),
	}))
	slog.SetDefault(logger)
}
