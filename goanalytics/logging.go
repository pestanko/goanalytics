package goanalytics

import (
	"os"

	"github.com/sirupsen/logrus"
)

/*
 * CreateLogger - Creates instace of the logger
 */
func CreateLogger() *logrus.Logger {
	log := logrus.New()
	log.Out = os.Stderr
	return log
}
