package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// log varíavel de log
	log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var erro error

	if log, erro = logConfig.Build(); erro != nil {
		panic(erro)
	}

}

//GetLogger retorna o gerenciador de logs
func GetLogger() {
	return
}

//Info log de informação
func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}

//Error log de informação
func Error(msg string, erro error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", erro))
	log.Error(msg, tags...)
	log.Sync()
}
