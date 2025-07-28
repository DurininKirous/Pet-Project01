package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
//	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Log *zap.Logger

func InitLogger() {
/*	lumberjackLogger := &lumberjack.Logger{
		Filename: "/home/durininkirous/Pet-Project/app/logs/app.log",
		MaxSize:  10, 
		MaxBackups: 5,
		MaxAge: 28,
		Compress: true,
	}
*/
	writeSyncer := zapcore.AddSync(os.Stdout)

	encoderCfg := zapcore.EncoderConfig {
		TimeKey: "time",
		LevelKey: "level",
		NameKey: "logger",
		CallerKey: "caller",
		MessageKey: "msg",
		StacktraceKey: "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.CapitalColorLevelEncoder,
        EncodeTime:     zapcore.ISO8601TimeEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
        zapcore.NewConsoleEncoder(encoderCfg), 
        writeSyncer,
        zapcore.InfoLevel,
    )

    Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

