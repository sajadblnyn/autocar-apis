package logging

import "github.com/sajadblnyn/autocar-apis/config"

type Logger interface {
	Init(cfg *config.Config)

	Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})

	Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Infof(template string, args ...interface{})

	Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Warnf(template string, args ...interface{})

	Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Errorf(template string, args ...interface{})

	Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Fatalf(template string, args ...interface{})
}

func NewLogger(cfg *config.Config) Logger {
	if cfg.Logger.Logger == "zap" {
		return newZapLogger(cfg)
	} else if cfg.Logger.Logger == "zerolog" {
		return newZeroLog(cfg)
	}
	panic("no valid logger specified in config")
}
