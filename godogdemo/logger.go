package main

import (
	"fmt"
	"log/slog"
	"os"
)

func configureLogger(logFileName string) *os.File {
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v", err))
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(f, nil)))
	return f
}
