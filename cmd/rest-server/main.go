package main

import (
	rest "github.com/kapustkin/go_guard/pkg/rest-server"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := rest.Run()
	if err != nil {
		log.Fatalf("unhandled app exception %v", err)
	}
}
