package logger

import (
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var (
	log     *logrus.Logger
	PFCPLog *logrus.Entry
)

func init() {
	log = logrus.New()
	log.SetReportCaller(false)

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	}

	PFCPLog = log.WithFields(logrus.Fields{"component": "LIB", "category": "PFCP"})
}

func GetLogger() *logrus.Logger {
	return log
}

func SetLogLevel(level logrus.Level) {
	PFCPLog.Infoln("set log level :", level)
	log.SetLevel(level)
}

func SetReportCaller(enable bool) {
	PFCPLog.Infoln("set report call :", enable)
	log.SetReportCaller(enable)
}
