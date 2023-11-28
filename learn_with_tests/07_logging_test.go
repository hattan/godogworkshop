package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"testing"

	"log/slog"

	"github.com/stretchr/testify/assert"
)

func Test_Log_TextInfo(t *testing.T) {
	// Arrange
	message := "hello, world"
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, nil))

	// Act
	logger.Info(message)

	// Assert
	output := buf.String()
	assert.Contains(t, output, fmt.Sprintf("level=INFO msg=\"%s\"", message))
}

func Test_Log_JsonInfo(t *testing.T) {
	// Arrange
	type LogMessage struct {
		Level string `json:"level"`
		Msg   string `json:"msg"`
		Time  string `json:"time"`
	}
	message := "hello, world"
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))

	// Act
	logger.Info(message)

	// Assert
	var log LogMessage
	output := buf.String()
	err := json.Unmarshal([]byte(output), &log)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, "INFO", log.Level)
}

func Test_Log_TextInfo_AdditionalAttrributes(t *testing.T) {
	// Arrange
	message := "hello, world"
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, nil))

	// Act
	logger.Info(message, "runId", "123", "requestId", "456")

	// Assert
	output := buf.String()
	assert.Contains(t, output, fmt.Sprintf("level=INFO msg=\"%s\"", message))
	assert.Contains(t, output, "runId=123", message)
	assert.Contains(t, output, "requestId=456", message)
}

func Test_Log_JsonInfo_AdditionalAttrributes(t *testing.T) {
	// Arrange
	type LogMessage struct {
		Level     string `json:"level"`
		Msg       string `json:"msg"`
		Time      string `json:"time"`
		UserId    string `json:"userId"`
		RequestId string `json:"requestId"`
	}
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))

	// Act
	logger.Info("cerritos", "userId", "NX01", "requestId", "1701-D")

	// Assert
	var log LogMessage
	output := buf.String()
	err := json.Unmarshal([]byte(output), &log)
	if err != nil {
		assert.Fail(t, err.Error())
		return
	}

	assert.Equal(t, "INFO", log.Level)
	assert.Equal(t, "cerritos", log.Msg)
	assert.Equal(t, "NX01", log.UserId)
	assert.Equal(t, "1701-D", log.RequestId)
}
func Test_Log_ChildLoggers(t *testing.T) {
	// Arrange
	buildInfo, _ := debug.ReadBuildInfo()
	f := CreateFileHandle("test.log")
	defer func() {
		f.Close()
		//os.Remove("test.log")
	}()

	logger := slog.New(slog.NewTextHandler(f, nil))

	child := logger.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
			slog.String("go_version", buildInfo.GoVersion),
		),
	)

	// Act
	child.Info("hello")
	child.Warn("Some warning")
	child.Error("Something went wrong")

	// Assert
	assert.FileExists(t, "test.log")
}
