package app

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gometr/internal/handlers"
	"gometr/internal/infrastructure/config"
	"log"
	"time"
)

func timeFormat(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		loc, _ = time.LoadLocation("Local")
	}
	enc.AppendString(t.In(loc).Format("2006-01-02 15:04:05.000"))
}

func (a *App) Bootstrap() error {
	conf := zap.NewProductionConfig()
	conf.EncoderConfig.EncodeTime = timeFormat

	logger, err := conf.Build()
	if err != nil {
		log.Fatal("failed to initialize logger")
	}
	a.log = logger

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	a.cfg = cfg

	a.log.Info("gometr is bootstrapping")

	a.handler = handlers.NewHandler(a.log)

	a.log.Info("gometr bootstrapped")
	return nil
}
