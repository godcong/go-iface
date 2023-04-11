package generator

import (
	"os"

	"golang.org/x/exp/slog"
)

var (
	logOpts = slog.HandlerOptions{AddSource: true}
	log     = slog.New(logOpts.NewJSONHandler(os.Stderr))
)
