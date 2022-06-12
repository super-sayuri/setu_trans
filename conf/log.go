package conf

import (
	"context"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func initLog(conf *LogConfig) error {
	if conf == nil {
		conf = &LogConfig{
			Level: "INFO",
		}
	}
	if strings.ToUpper(conf.Format) == "JSON" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}
	if strings.ToUpper(conf.Output) == "FILE" {
		if len(conf.Path) == 0 {
			conf.Path = "syr.log"
		}
		f, err := os.OpenFile(conf.Path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		logrus.SetOutput(f)
	} else {
		logrus.SetOutput(os.Stdout)
	}
	level, err := logrus.ParseLevel(conf.Level)
	if err != nil {
		return err
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)
	return nil
}

func GetLog(ctx context.Context) *logrus.Entry {
	l := logrus.NewEntry(logrus.StandardLogger())
	if ctx.Value("job") != nil {
		l = l.WithField("job", ctx.Value("job"))
	}
	if ctx.Value("msgId") != nil {
		l = l.WithField("msgId", ctx.Value("msgId"))
	}
	return l
}
