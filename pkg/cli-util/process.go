package cli

import (
	"log"

	"github.com/kapustkin/go_guard/pkg/cli-util/commands"
	"github.com/kapustkin/go_guard/pkg/cli-util/config"
	"github.com/spf13/cobra"
)

// Run entry point of app
func Run() {
	//read config
	conf := config.Init()

	var server string

	var rootCmd = &cobra.Command{Use: "go-guard-cli"}

	rootCmd.PersistentFlags().StringVarP(&server, "server", "s", "", "application host:port")

	rootCmd.AddCommand(
		commands.ResetCmd(server, conf.Host),
		commands.ListCmd(server, conf.Host),
		commands.ParamsCmd(server, conf.Host))

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("application error: %v", err)
	}
}
