package main

import (
	"scratch/cmd"
	"scratch/src/log"
)

func init() {
	log.NewLogger()
}
func main() {
	cmd.Execute()

}
