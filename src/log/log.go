package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

var once sync.Once

func NewLogger() *logrus.Logger {
	once.Do(func() {

		Logger = logrus.New()

		// Output to stdout instead of the default stderr
		// Can be any io.Writer, see below for File example
		Logger.SetOutput(os.Stdout)
		// Only logrus the warning severity or above.
		Logger.SetLevel(logrus.DebugLevel)
		Logger.SetReportCaller(true)
		// logrus as JSON instead of the default ASCII formatter.
		Logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05", CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		}})
	})

	return Logger
}
