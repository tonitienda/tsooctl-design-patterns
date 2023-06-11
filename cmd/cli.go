package main

import (
	"fmt"
	"os"
	"tsoobame/tsooctl/pkg/contracts"
	"tsoobame/tsooctl/pkg/pods"
)

var commandStrategy = map[string]map[string]contracts.Command{
	"pods": {
		"get": pods.GetPods{},
	},
}

func printHelp() {
	fmt.Println("tsooclt help")

	for noun, commands := range commandStrategy {
		for verb, command := range commands {
			fmt.Printf("\t%s %s: %s\n", verb, noun, command.Help().Message)
		}
	}
	fmt.Println()
}

func main() {

	if len(os.Args) < 3 {
		printHelp()
		os.Exit(1)
	}

	verb := os.Args[1]
	noun := os.Args[2]

	commands, ok := commandStrategy[noun]

	if !ok {
		printHelp()
		os.Exit(1)
	}

	command, ok := commands[verb]

	if !ok {
		printHelp()
		os.Exit(1)
	}

	// TODO - Pass extra args
	result := command.Execute()

	fmt.Println(result.Message)
}
