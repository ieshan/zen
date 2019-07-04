package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"sync"
)

type Logger struct {
	*logrus.Logger
}

var log *Logger
var once sync.Once

func New() *Logger {
	return &Logger{
		logrus.New(),
	}
}

func InitLogger(output io.Writer) {
	once.Do(func() {
		log = New()
		log.SetOutput(output)
		logLevel, err := logrus.ParseLevel(viper.GetString("log.level"))
		log.ErrorFatal("Error parsing log.level=" + viper.GetString("log.level"), err)
		log.SetLevel(logLevel)
		if viper.GetString("log.formatter") == "json" {
			log.SetFormatter(&logrus.JSONFormatter{})
		} else {
			log.SetFormatter(&logrus.TextFormatter{})
		}
	})
}

func GetLogger() *Logger {
	return log
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

func (logger *Logger) ErrorFatal(message string, err error)  {
	if err != nil {
		logger.Fatalln(message, err)
	}
}
