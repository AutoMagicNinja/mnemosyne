package mnemosyne

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// This is used by wrapper functions to get the correct caller when using this library's wrapper functions
	// Otherwise, the caller will always be this logging library, instead of the actual (downstream) caller
	optionSkipWrapper zap.Option = zap.AddCallerSkip(1)
)

// Convenience wrappers to enable this package to be used as a global logger
// Note that we have to use the optionSkipWrapper zap.Option because of this
// Otherwise, the caller is always going to be this module instead of the actual caller

func DPanic(args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).DPanic(args...)
}
func DPanicf(template string, args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).DPanicf(template, args...)
}
func DPanicln(args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).DPanicln(args...)
}
func DPanicw(msg string, keysAndValues ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).DPanicw(msg, keysAndValues...)
}
func Debug(args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Debug(args...)
}
func Debugf(template string, args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Debugf(template, args...)
}
func Debugln(args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Debugln(args...)
}
func Debugw(msg string, keysAndValues ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Debugw(msg, keysAndValues...)
}
func Error(args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Error(args...)
}
func Errorf(template string, args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Errorf(template, args...)
}
func Errorln(args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Errorln(args...)
}
func Errorw(msg string, keysAndValues ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Errorw(msg, keysAndValues...)
}
func Fatal(args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Fatal(args...)
}
func Fatalf(template string, args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Fatalf(template, args...)
}
func Fatalln(args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Fatalln(args...)
}
func Fatalw(msg string, keysAndValues ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Fatalw(msg, keysAndValues...)
}
func Info(args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Info(args...)
}
func Infof(template string, args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Infof(template, args...)
}
func Infoln(args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Infoln(args...)
}
func Infow(msg string, keysAndValues ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Infow(msg, keysAndValues...)
}
func Level() zapcore.Level {
	return globalLogger.Level()
}
func Panic(args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Panic(args...)
}
func Panicf(template string, args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Panicf(template, args...)
}
func Panicln(args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Panicln(args...)
}
func Panicw(msg string, keysAndValues ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Panicw(msg, keysAndValues...)
}
func Sync() error {
	return globalLogger.Sync()
}
func Warn(args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Warn(args...)
}
func Warnf(template string, args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Warnf(template, args...)
}
func Warnln(args ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Warnln(args...)
}
func Warnw(msg string, keysAndValues ...interface{}) {
	globalLogger.WithOptions(optionSkipWrapper).Warnw(msg, keysAndValues...)
}
func WithOptions(opts ...zap.Option) {
	globalLogger = *globalLogger.WithOptions(opts...)
}
