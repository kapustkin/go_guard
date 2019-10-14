package main

import (
	"log"

	rest "github.com/kapustkin/go_guard/pkg/rest-server"
	"github.com/kapustkin/go_guard/pkg/utils/helpers"
)

func main() {
	if helpers.HasError(rest.Run()) {
		log.Fatalf("unhandled app exception")
	}
}
