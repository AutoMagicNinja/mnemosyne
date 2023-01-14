package mnemosyne

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// re-export so downstream consumers do not need to import zapcore directly
const (
	DebugLevel   = zapcore.DebugLevel
	InfoLevel    = zapcore.InfoLevel
	WarnLevel    = zapcore.WarnLevel
	ErrorLevel   = zapcore.ErrorLevel
	DPanicLevel  = zapcore.DPanicLevel
	PanicLevel   = zapcore.PanicLevel
	FatalLevel   = zapcore.FatalLevel
	InvalidLevel = zapcore.InvalidLevel

	DefaultLogLevel = WarnLevel
)

var (
	zapLogEncoderConfig   zapcore.EncoderConfig
	zapCoreJSONEncoder    zapcore.Encoder
	zapCoreConsoleEncoder zapcore.Encoder
	selectedEncoder       *zapcore.Encoder

	globalLogger   zap.SugaredLogger
	globalLogLevel zap.AtomicLevel
)

// Use init() function in order to guarantee the logger is created every* time
func init() {
	globalLogLevel = zap.NewAtomicLevel()
	globalLogLevel.SetLevel(DefaultLogLevel)

	zapLogEncoderConfig = zap.NewProductionEncoderConfig()
	zapLogEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapLogEncoderConfig.TimeKey = "@timestamp"
	zapLogEncoderConfig.CallerKey = "caller"
	zapLogEncoderConfig.FunctionKey = "function"
	zapLogEncoderConfig.StacktraceKey = "stacktrace"
	zapLogEncoderConfig.MessageKey = "message"
	zapLogEncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	zapCoreJSONEncoder = zapcore.NewJSONEncoder(zapLogEncoderConfig)
	zapCoreConsoleEncoder = zapcore.NewConsoleEncoder(zapLogEncoderConfig)

	selectedEncoder = &zapCoreConsoleEncoder // default to human encoding

	globalLogger = *zap.New(
		zapcore.NewCore(*selectedEncoder, zapcore.AddSync(os.Stdout), globalLogLevel),
		zap.AddCaller(), zap.AddStacktrace(zapcore.DPanicLevel),
	).Sugar()
}

func GetRawLogger() *zap.SugaredLogger {
	return &globalLogger
}

func GetConfig() *zapcore.EncoderConfig {
	return &zapLogEncoderConfig
}

// Call ResetEncoderConfigs() after modifying the EncoderConfig returned by GetConfig()
func ResetEncoderConfigs() {
	zapCoreJSONEncoder = zapcore.NewJSONEncoder(zapLogEncoderConfig)
	zapCoreConsoleEncoder = zapcore.NewConsoleEncoder(zapLogEncoderConfig)
}

func ResetEncoder() {
	globalLogger = *zap.New(
		zapcore.NewCore(*selectedEncoder, zapcore.AddSync(os.Stdout), globalLogLevel),
	).Sugar()
}

// Lowers the level that stack traces are added and sets development mode using the builtin zap.Option
func EnableDevelopmentMode() {
	globalLogger = *globalLogger.WithOptions(
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.Development(),
	)
}

func UseJSONEncoder() {
	selectedEncoder = &zapCoreJSONEncoder
	ResetEncoder()
}

func UseConsoleEncoder() {
	selectedEncoder = &zapCoreConsoleEncoder
	ResetEncoder()
}

func SetLevel(logLevel zapcore.Level) {
	globalLogger.Sync()
	globalLogLevel.SetLevel(logLevel)
}
