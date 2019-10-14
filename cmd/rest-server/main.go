package main

import (
	"log"

	rest "github.com/kapustkin/go_guard/pkg/rest-server"
	h "github.com/kapustkin/go_guard/pkg/utils/helpers"
)

func main() {
	if h.HasError(rest.Run()) {
		log.Fatalf("unhandled app exception")
	}
}
