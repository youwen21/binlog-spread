package applog

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

var defaultLogger *logrus.Logger
var once sync.Once

func initDefault() {
	defaultLogger = logrus.New()

	logPath := os.Getenv("APP_LOG_PATH")
	if logPath == "" {
		rootPath, _ := os.Getwd()
		logPath = rootPath + "/storage/logs/"
	}
	logFile := logPath + "/default.log.%Y%m%d%H"

	writer, _ := rotatelogs.New(
		logFile,
		rotatelogs.WithLinkName(logPath+"/default.log"),
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationSize(1024*1024*1024),
	)
	defaultLogger.Out = writer
}

func Default() *logrus.Logger {
	once.Do(initDefault)
	return defaultLogger
}
