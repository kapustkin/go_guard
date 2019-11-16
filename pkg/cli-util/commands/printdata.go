package commands

import (
	"log"

	"github.com/kapustkin/go_guard/pkg/cli-util/handlers"
)

func printParamsResult(server string, data *handlers.RespParams) {
	log.Printf("Parameters for server %v", server)
	log.Printf("Login(N)\t= %v", data.Data.N)
	log.Printf("Password(M)\t= %v", data.Data.M)
	log.Printf("Address(K)\t= %v", data.Data.K)
}

func printDataResult(data *handlers.RespData) {
	log.Printf("|-------------------|---------------|")
	log.Printf("|      network\t|    status\t|")
	log.Printf("|-------------------|---------------|")

	for _, item := range data.Data {
		status := "denied"
		if item.IsWhite {
			status = "allow"
		}

		log.Printf("|%v\t|    %v\t|\r\n", item.Network, status)
	}

	log.Printf("|-------------------|---------------|")
}

func printCommandResult(data *handlers.RespData) {
	if data.Success {
		log.Printf("success")
	} else {
		log.Fatalf("error: %v", data.Description)
	}
}
