package global

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type LogType struct {
	Name string
	Path string
}



var logMap = map[string]*zap.Logger{}
var(
	MainLog *zap.Logger = nil
)

func init() {
	MainLog = InitLog("/data/logs/main.log")
}

func InitLog(path string)*zap.Logger {
	hook := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    300,
		MaxAge:     10,
		MaxBackups: 1,
		LocalTime:  true,
		Compress:   false,
	}

	ec := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	zc := zapcore.NewCore(
		zapcore.NewJSONEncoder(ec),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),zapcore.AddSync(hook)),
		zap.InfoLevel,
	)
	logger := zap.New(zc)
	return logger
}
