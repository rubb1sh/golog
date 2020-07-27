package main

import "github.com/rubb1sh/golog/log"

func main() {
	filename := "error.log"
	log.SetToFile(filename)
}
