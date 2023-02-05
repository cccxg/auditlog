package logger

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/soducool/auditlog/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.SugaredLogger

func Init() (err error) {
	// Catching potential error during log init.
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("%v", p)
		}
	}()

	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "ts",
		NameKey:        "name",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stack",
		SkipLineEnding: false,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05 Mon"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	})

	infoWriter := getWriter("info")
	debugWriter := getWriter("debug")
	warnWriter := getWriter("warn")
	errorWriter := getWriter("error")

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(debugWriter), debugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
	)

	logger = zap.New(core, zap.AddCaller()).Sugar()
	Debugf("init logger")

	return
}

func getWriter(level string) *lumberjack.Logger {
	fileName := fmt.Sprintf("%s.log", level)
	return &lumberjack.Logger{
		Filename:   filepath.Join(config.Config.Log.Path, fileName),
		MaxSize:    config.Config.Log.MaxSize,
		MaxAge:     config.Config.Log.MaxAge,
		MaxBackups: config.Config.Log.MaxBackups,
		LocalTime:  config.Config.Log.LocalTime,
		Compress:   config.Config.Log.Compress,
	}
}

func Info(args ...interface{})  { logger.Info(args) }
func Debug(args ...interface{}) { logger.Debug(args) }
func Warn(args ...interface{})  { logger.Warn(args) }
func Error(args ...interface{}) { logger.Error(args) }

func Infof(template string, args ...interface{})  { logger.Infof(template, args...) }
func Debugf(template string, args ...interface{}) { logger.Debugf(template, args...) }
func Warnf(template string, args ...interface{})  { logger.Warnf(template, args...) }
func Errorf(template string, args ...interface{}) { logger.Errorf(template, args...) }
