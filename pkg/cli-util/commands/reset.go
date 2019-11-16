package commands

import (
	"log"

	"github.com/kapustkin/go_guard/pkg/cli-util/handlers"
	"github.com/spf13/cobra"
)

func ResetCmd(server, deafultServer string) *cobra.Command {
	var address, login string

	var cmdReset = &cobra.Command{
		Use:   "reset",
		Short: "Reset address and login data",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if server == "" {
				server = deafultServer
			}
			reset(server, address, login)
		},
	}

	cmdReset.Flags().StringVarP(&address, "address", "a", "", "address to reset")
	cmdReset.Flags().StringVarP(&login, "login", "l", "", "login to reset")
	_ = cmdReset.MarkFlagRequired("address")
	_ = cmdReset.MarkFlagRequired("login")

	return cmdReset
}

func reset(server, address, login string) {
	res, err := handlers.Reset(server, address, login)
	if err != nil {
		log.Fatalf("%v", err)
	}

	printCommandResult(res)
}
