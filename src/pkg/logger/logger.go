package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func Init(environment string) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.TimeKey = "timestamp"

	var encoder zapcore.Encoder
	if environment == "development" {
		encoder = zapcore.NewConsoleEncoder(config)
	} else {
		encoder = zapcore.NewJSONEncoder(config)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		zapcore.InfoLevel,
	)

	log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	log.Fatal(message, fields...)
}

func WithFields(fields map[string]interface{}) *zap.Logger {
	var zapFields []zap.Field
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return log.With(zapFields...)
}
