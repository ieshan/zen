package logger

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"io"
	"sync"
)

type Logger struct {
	*logrus.Logger
}

var logInstance *Logger
var logOnce sync.Once

func GetLogger() *Logger {
	logOnce.Do(func() {
		logInstance = &Logger{
			logrus.New(),
		}
	})
	return logInstance
}

func (logger *Logger) ErrorDebug(message string, err error)  {
	if err != nil {
		logger.Errorln(message, err)
	}
}

func (logger *Logger) ErrorDebugReturn(message string, err error) error {
	if err != nil {
		logger.Errorln(message, err)
	}
	return err
}

func (logger *Logger) ErrorFatal(message string, err error) {
	if err != nil {
		logger.Fatalln(message, err)
	}
}

func (logger *Logger) SetLogLevel(level string) {
	logLevel, err := logrus.ParseLevel(level)
	logger.ErrorFatal("Log level error", err)
	logger.Logger.SetLevel(logLevel)
}

func (logger *Logger) Output() io.Writer {
	return logger.Out
}

func (logger *Logger) Prefix() string {
	return ""
}

func (logger *Logger) SetPrefix(p string) {
	// Do nothing
}

func (logger *Logger) Level() log.Lvl {
	return toEchoLevel(logger.Logger.Level)
}

func (logger *Logger) SetLevel(v log.Lvl) {
	logger.Logger.Level = toLogrusLevel(v)
}

func (logger *Logger) SetHeader(h string) {
	// Do nothing
}

func (logger *Logger) Printj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	logger.Logger.Println(string(b))
}

func (logger *Logger) Debugj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	logger.Logger.Debugln(string(b))
}

func (logger *Logger) Infoj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	logger.Logger.Infoln(string(b))
}

func (logger *Logger) Warnj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	logger.Logger.Warnln(string(b))
}

func (logger *Logger) Errorj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	logger.Logger.Errorln(string(b))
}

func (logger *Logger) Fatalj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	logger.Logger.Fatalln(string(b))
}

func (logger *Logger) Panicj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	logger.Logger.Panicln(string(b))
}

// To logrus.Level
func toLogrusLevel(level log.Lvl) logrus.Level {
	switch level {
	case log.DEBUG:
		return logrus.DebugLevel
	case log.INFO:
		return logrus.InfoLevel
	case log.WARN:
		return logrus.WarnLevel
	case log.ERROR:
		return logrus.ErrorLevel
	}

	return logrus.InfoLevel
}

// To Echo.log.lvl
func toEchoLevel(level logrus.Level) log.Lvl {
	switch level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.InfoLevel:
		return log.INFO
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	}

	return log.OFF
}
