package main

import (
	"os"
	"scratch/cmd"

	"github.com/sirupsen/logrus"
)

func init() {
	// logrus as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only logrus the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}
func main() {
	cmd.Execute()

}
