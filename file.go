package golog

import (
	"fmt"
	"log"
	"os"
)

// if error log output to os.stderr
func SetToFile(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		log.SetOutput(os.Stderr)
		return err
	}
	log.SetOutput(file)
	return nil
}
