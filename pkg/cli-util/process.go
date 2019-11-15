package cli

import (
	"log"

	"github.com/kapustkin/go_guard/pkg/cli-util/config"
	"github.com/kapustkin/go_guard/pkg/cli-util/handlers"
	"github.com/spf13/cobra"
)

// Run entry point of app
func Run() {
	//read config
	conf := config.Init()

	var server, address, login string

	var cmdReset = &cobra.Command{
		Use:   "reset",
		Short: "Reset address and login data",
		Long: `Команда удаляет данны бакетов с адресом и логином, что позволяет выполнить 
		запросы в рамках установленных лимитов`,
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if server == "" {
				server = conf.Host
			}
			reset(server, address, login)
		},
	}

	cmdReset.Flags().StringVarP(&address, "address", "a", "", "address to reset")
	cmdReset.Flags().StringVarP(&login, "login", "l", "", "login to reset")
	_ = cmdReset.MarkFlagRequired("address")
	_ = cmdReset.MarkFlagRequired("login")

	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "List allow/block networks",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if server == "" {
				server = conf.Host
			}
			listAll(server)
		},
	}

	var cmdAdd = &cobra.Command{
		Use:   "add",
		Short: "Add network to list",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
	}

	var cmdAddAllow = &cobra.Command{
		Use:   "allow",
		Short: "Add network to white list",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if server == "" {
				server = conf.Host
			}
			listAdd(server, args[0], true)
		},
	}

	var cmdAddBlock = &cobra.Command{
		Use:   "denied",
		Short: "Add network to black list",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if server == "" {
				server = conf.Host
			}
			listAdd(server, args[0], false)
		},
	}

	var cmdRemove = &cobra.Command{
		Use:   "remove",
		Short: "Remove network from list",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
	}

	var cmdRemoveAllow = &cobra.Command{
		Use:   "allow",
		Short: "Remove network from white list",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if server == "" {
				server = conf.Host
			}
			listRemove(server, args[0], true)
		},
	}

	var cmdRemoveBlock = &cobra.Command{
		Use:   "denied",
		Short: "Remove network from black list",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if server == "" {
				server = conf.Host
			}
			listRemove(server, args[0], false)
		},
	}

	var cmdParams = &cobra.Command{
		Use:   "params",
		Short: "Return current params",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if server == "" {
				server = conf.Host
			}
			listParams(server)
		},
	}

	var rootCmd = &cobra.Command{Use: "go-guard-cli"}
	rootCmd.AddCommand(cmdReset, cmdList, cmdParams)
	rootCmd.PersistentFlags().StringVarP(&server, "server", "s", "", "go-guard application host")

	cmdList.AddCommand(cmdAdd, cmdRemove)
	cmdAdd.AddCommand(cmdAddAllow, cmdAddBlock)
	cmdRemove.AddCommand(cmdRemoveAllow, cmdRemoveBlock)
	//nolint
	rootCmd.Execute()
}

func reset(server, address, login string) {
	res, err := handlers.Reset(server, address, login)
	if err != nil {
		log.Fatalf("%v", err)
	}
	printCommandResult(res)
}

func listAll(server string) {
	res, err := handlers.GetAllList(server)
	if err != nil {
		log.Fatalf("%v", err)
	}
	printDataResult(res)
}

func listAdd(server, network string, isWhite bool) {
	res, err := handlers.AddToList(server, network, isWhite)
	if err != nil {
		log.Fatalf("%v", err)
	}
	printCommandResult(res)
}

func listRemove(server, network string, isWhite bool) {
	res, err := handlers.RemoveFromList(server, network, isWhite)
	if err != nil {
		log.Fatalf("%v", err)
	}
	printCommandResult(res)
}

func listParams(server string) {
	res, err := handlers.GetParams(server)
	if err != nil {
		log.Fatalf("%v", err)
	}
	printParamsResult(server, res)
}

func printParamsResult(server string, data *handlers.RespParams) {
	log.Printf("Paramters for server %v", server)
	log.Printf("K = %v", data.Data.K)
	log.Printf("M = %v", data.Data.M)
	log.Printf("N = %v", data.Data.N)
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
