package commands

import (
	"log"

	"github.com/kapustkin/go_guard/pkg/cli-util/handlers"
	"github.com/spf13/cobra"
)

func ParamsCmd(server, deafultServer string) *cobra.Command {
	var n, m, k int

	var cmdParams = &cobra.Command{
		Use:   "params",
		Short: "Manage parameters",
		Args:  cobra.MinimumNArgs(1),
	}

	var cmdParamsList = &cobra.Command{
		Use:   "list",
		Short: "List parameters",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if server == "" {
				server = deafultServer
			}
			listParams(server)
		},
	}

	var cmdParamsAdd = &cobra.Command{
		Use:   "add",
		Short: "Add new parameters",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if server == "" {
				server = deafultServer
			}
			paramsAdd(server, n, m, k)
		},
	}

	cmdParamsAdd.Flags().IntVarP(&n, "n", "n", 0, "limit for login")
	cmdParamsAdd.Flags().IntVarP(&m, "m", "m", 0, "limit for password")
	cmdParamsAdd.Flags().IntVarP(&k, "k", "k", 0, "limit for address")
	_ = cmdParamsAdd.MarkFlagRequired("n")
	_ = cmdParamsAdd.MarkFlagRequired("m")
	_ = cmdParamsAdd.MarkFlagRequired("k")

	cmdParams.AddCommand(cmdParamsList, cmdParamsAdd)

	return cmdParams
}

func listParams(server string) {
	res, err := handlers.GetParams(server)
	if err != nil {
		log.Fatalf("%v", err)
	}

	printParamsResult(server, res)
}

func paramsAdd(server string, n, m, k int) {
	res, err := handlers.SetParams(server, n, m, k)
	if err != nil {
		log.Fatalf("%v", err)
	}

	printCommandResult(res)
}
