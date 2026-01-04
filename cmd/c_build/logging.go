package main

import (
	"log/slog"
	"os"
)

func InitLogger(level slog.Level) {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     level,
		AddSource: level == slog.LevelDebug, // debug 模式显示文件/行号
	})

	slog.SetDefault(slog.New(handler))
}
