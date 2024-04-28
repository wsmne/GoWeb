package logger

import (
	"log/slog"
	"web-server/tlog"
)

type LogConfig struct {
	LogDir string
}

var Log *slog.Logger = tlog.NewProductionLogger(tlog.LogConfig{LogDir: "logs"})
