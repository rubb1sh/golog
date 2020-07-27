package log

import (
	"fmt"
	"log"
	"os"
)

// if error log output to os.stderr
func SetToFile(filename string) bool {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		log.SetOutput(os.Stderr)
		return false
	}
	log.SetOutput(file)
	return true
}
