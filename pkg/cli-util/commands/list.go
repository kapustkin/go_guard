package commands

import (
	"log"

	"github.com/kapustkin/go_guard/pkg/cli-util/handlers"
	"github.com/spf13/cobra"
)

func ListCmd(server, defaultServer string) *cobra.Command {
	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "Manage allow/block lists",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if server == "" {
				server = defaultServer
			}
			listAll(server)
		},
	}

	cmdList.AddCommand(
		listAddCmd(server, defaultServer),
		listRemoveCmd(server, defaultServer))

	return cmdList
}

//nolint:dupl
func listAddCmd(server, defaultServer string) *cobra.Command {
	var cmdAdd = &cobra.Command{
		Use:   "add",
		Short: "Add network to list",
		Args:  cobra.MinimumNArgs(1),
	}

	var cmdAddAllow = &cobra.Command{
		Use:   "allow",
		Short: "Add network to white list",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if server == "" {
				server = defaultServer
			}
			listAdd(server, args[0], true)
		},
	}

	var cmdAddBlock = &cobra.Command{
		Use:   "denied",
		Short: "Add network to black list",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if server == "" {
				server = defaultServer
			}
			listAdd(server, args[0], false)
		},
	}

	cmdAdd.AddCommand(cmdAddAllow, cmdAddBlock)

	return cmdAdd
}

//nolint:dupl
func listRemoveCmd(server, defaultServer string) *cobra.Command {
	var cmdRemove = &cobra.Command{
		Use:   "remove",
		Short: "Remove network from list",
		Args:  cobra.MinimumNArgs(1),
	}

	var cmdRemoveAllow = &cobra.Command{
		Use:   "allow",
		Short: "Remove network from white list",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if server == "" {
				server = defaultServer
			}
			listRemove(server, args[0], true)
		},
	}

	var cmdRemoveBlock = &cobra.Command{
		Use:   "denied",
		Short: "Remove network from black list",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if server == "" {
				server = defaultServer
			}
			listRemove(server, args[0], false)
		},
	}

	cmdRemove.AddCommand(cmdRemoveAllow, cmdRemoveBlock)

	return cmdRemove
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
