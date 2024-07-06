package logging

import (
	"sync"

	"github.com/sajadblnyn/autocar-apis/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once
var zapSyncLogger *zap.SugaredLogger

var logLevels map[string]zapcore.Level = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"warn":  zapcore.WarnLevel,
	"info":  zapcore.InfoLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

type zapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func newZapLogger(cfg *config.Config) *zapLogger {
	zapLogger := &zapLogger{cfg: cfg}
	zapLogger.Init(cfg)
	return zapLogger
}

func (l *zapLogger) getLogLevel() zapcore.Level {
	level, exisits := logLevels[l.cfg.Logger.Level]
	if !exisits {
		return zapcore.DebugLevel
	}
	return level
}

func (l *zapLogger) Init(cfg *config.Config) {
	once.Do(func() {
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.Logger.FilePath,
			MaxSize:    1,
			MaxAge:     5,
			MaxBackups: 10,
			LocalTime:  true,
			Compress:   true,
		})
		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(config), w, l.getLogLevel())
		logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()

		zapSyncLogger = logger.With("AppName", "AutoCar", "Logger", "zapp")

	})

	l.logger = zapSyncLogger
}

func (l *zapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	l.logger.Debugw(msg, prepareLogKeys(extra, cat, sub)...)
}
func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *zapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Infow(msg, prepareLogKeys(extra, cat, sub)...)

}
func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)

}

func (l *zapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Warnw(msg, prepareLogKeys(extra, cat, sub)...)

}
func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)

}

func (l *zapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Errorw(msg, prepareLogKeys(extra, cat, sub)...)

}
func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)

}

func (l *zapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Fatalw(msg, prepareLogKeys(extra, cat, sub)...)

}
func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func prepareLogKeys(extra map[ExtraKey]interface{}, cat Category, sub SubCategory) []interface{} {
	if extra == nil {
		extra = map[ExtraKey]interface{}{}
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub

	params := mapToZapParams(extra)
	return params
}
