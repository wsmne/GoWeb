package tlog

import (
	"log/slog"
)

type LogConfig struct {
	LogDir string
}

func NewProductionLogger(cfg LogConfig) *slog.Logger {
	handler := newZlogHandler(cfg.LogDir)
	return slog.New(handler)
}
