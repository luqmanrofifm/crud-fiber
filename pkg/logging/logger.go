package logging

import (
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()
	//Log.SetFormatter(&logrus.JSONFormatter{})
	//Log.SetOutput(os.Stdout)
}
