package main

import (
	"github.com/PennState/go-additional-properties/pkg/generator"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := generator.Run()
	if err != nil {
		log.Error(err)
	}
}
