package main

import (
	"alfmigcli/cmd/cluster"
	"alfmigcli/cmd/node"
	"alfmigcli/cmd/user"
	"alfmigcli/cmd/workflow"
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) > 0 {
		rute(argsWithoutProg[0], os.Args[2:])
	} else {
		fmt.Println("Please select a service.")
		t := prompt.Input("> ", completer)
		rute(strings.Split(t, " ")[0], strings.Split(t, " ")[1:])
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "user -login <user> <pass>", Description: "User Login"},
		{Text: "user -describe ", Description: "User infonrmation"},
		{Text: "cluster -add <name> <url>", Description: "Add a new configuration cluster"},
		{Text: "cluster -describe", Description: "Describe all configuration clusters"},
		{Text: "cluster -visible <cluster position>", Description: "Set visibility for a cluster"},
		{Text: "cluster -invisible <cluster position>", Description: "Remove visibility for a cluster"},
		{Text: "cluster -current <cluster position>", Description: "Set to current workspace cluster"},
		{Text: "cluster -delete <cluster position>", Description: "Remove a cluster configuration"},
		{Text: "cluster -ticket ", Description: "Get current ticket"},
		{Text: "workflow -getalldef", Description: "Get all definitions"},
		{Text: "node -spath <path>", Description: "Node path search"},
		{Text: "node -sid <id> <tenant>", Description: "Node id search"},
		{Text: "node -schildren <id> ", Description: "Node path search"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func rute(command string, commands []string) {
	switch command {
	case "user":
		user.Exec(commands)
	case "node":
		node.Exec(commands)
	case "cluster":
		cluster.Exec(commands)
	case "workflow":
		workflow.Exec(commands)
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
