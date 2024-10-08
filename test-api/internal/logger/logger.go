package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func Init() {
	Log = logrus.New()

	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	Log.SetLevel(logrus.InfoLevel)

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.SetOutput(file)
	} else {
		Log.Info("Failed to log to file, using default stderr")
	}
}
