package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/func0x/pokedexcli/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{
		Cache:   pokecache.NewCache(5 * time.Second),
		Pokedex: make(map[string]Pokemon),
	}

	cliCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Displays information about pokemon in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays your Pokedex",
			callback:    commandPokedex,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
	}

	fmt.Println("Welcome to the Pokedex!")

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		words := strings.Fields(strings.ToLower(line))
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]
		args := words[1:]

		if cmdFunc, ok := cliCommands[cmdName]; ok {
			err := cmdFunc.callback(cfg, args)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command:", cmdName)
		}
	}
}
