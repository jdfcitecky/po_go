package utils

import (
	"os"
	"path/filepath"
	"po_go/conf"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

var logToFile *logrus.Logger

//the name of log
var loggerFile string

func SetLogFile(file string) {
	loggerFile = file
}

//init
func init() {
	SetLogFile(filepath.Join(conf.Conf.MyLog.Path, conf.Conf.MyLog.Name))
}

func Log() *logrus.Logger {
	//save as file
	if conf.Conf.MyLog.Model == "file" {
		return logFile()
	} else {
		//control panel output
		if log == nil {
			log = logrus.New()
			log.Out = os.Stdout
			log.Formatter = &logrus.JSONFormatter{TimestampFormat: "2020-06-13 15:04:05"}
			log.SetLevel(logrus.DebugLevel)
		}
	}
	return log
}

func logFile() *logrus.Logger {

	if logToFile == nil {
		logToFile = logrus.New()

		logToFile.SetLevel(logrus.DebugLevel)

		// set rotatelogs return logWriter
		logWriter, _ := rotatelogs.New(
			// the name after divide
			loggerFile+"_%Y%m%d.log",

			// set max time to record
			rotatelogs.WithMaxAge(30*24*time.Hour),

			// set time to divide
			rotatelogs.WithRotationTime(24*time.Hour),
		)

		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}

		//set the format of time
		lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
			TimestampFormat: "2020-06-13 15:04:05",
		})

		// create a Hook
		logToFile.AddHook(lfHook)
	}
	return logToFile
}
