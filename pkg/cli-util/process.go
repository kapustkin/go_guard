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
		Short: "Сброс бакета с данными",
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

	var rootCmd = &cobra.Command{Use: "go-guard"}
	rootCmd.AddCommand(cmdReset)
	rootCmd.PersistentFlags().StringVarP(&server, "server", "s", "", "go-guard application host")
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

func printCommandResult(result bool) {
	if result {
		log.Printf("Команда выполнена успешно")
	} else {
		log.Printf("Команда НЕ выполнена")
	}
}
