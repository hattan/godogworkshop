package utils

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

func ConfigureLogger(logFileName string) *os.File {
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE, 0666)
	var writer io.Writer
	if err != nil {
		//panic(fmt.Sprintf("error opening file: %v", err))
		writer = os.Stdout
		fmt.Println("error opening file")
	} else {
		writer = f
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(writer, nil)))
	return f
}
