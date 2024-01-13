package main

import (
	"fmt"

	"github.com/joomcode/errorx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const logFileName = "math-examples.log"

var (
	ErrPrinterNamespace   = errorx.NewNamespace("printer")
	ErrFailBuildZapLogger = errorx.NewType(ErrPrinterNamespace, "FailBuildZapLogger")
)

type Printer interface {
	Print(format string, values ...any)
	Println(format string, values ...any)
	PrintUserInput(value string)
}

type LogPrinter struct {
	fileLogger    *zap.Logger
	consoleLogger *zap.Logger
}

func NewLogPrinter() (Printer, error) {
	fileLogger, err := createLogger(logFileName)
	if err != nil {
		return nil, err
	}

	consoleLogger, err := createLogger("stderr")
	if err != nil {
		return nil, err
	}

	return &LogPrinter{
		fileLogger:    fileLogger,
		consoleLogger: consoleLogger,
	}, nil
}

func createLogger(output string) (*zap.Logger, error) {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.LevelKey = zapcore.OmitKey
	encoderConfig.CallerKey = zapcore.OmitKey
	encoderConfig.TimeKey = zapcore.OmitKey
	encoderConfig.SkipLineEnding = true
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{output},
		ErrorOutputPaths: []string{output},
	}

	logger, err := config.Build()
	if err != nil {
		return nil, ErrFailBuildZapLogger.WrapWithNoMessage(err)
	}

	return logger, nil
}

func (p *LogPrinter) Print(format string, values ...any) {
	result := fmt.Sprintf(format, values...)
	p.fileLogger.Info(result)
	p.consoleLogger.Info(result)
}

func (p *LogPrinter) Println(format string, values ...any) {
	p.Print(format+"\n", values...)
}

func (p *LogPrinter) PrintUserInput(value string) {
	p.fileLogger.Info(value)
}
