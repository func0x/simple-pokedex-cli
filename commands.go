package main

import (
	"fmt"
	"math/rand/v2"
	"os"
)

func commandExit(cfg *Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, args []string) error {
	fmt.Println("this is help message")
	return nil
}

func commandMap(cfg *Config, args []string) error {
	url := cfg.MapNext
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	response, err := fetchLocations(cfg.Cache, url)
	if err != nil {
		return err
	}

	cfg.MapNext = response.Next
	cfg.MapPrevious = response.Previous

	for _, d := range response.Results {
		fmt.Println(d.Name)
	}

	return nil
}

func commandMapB(cfg *Config, args []string) error {
	if cfg.MapPrevious == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	response, err := fetchLocations(cfg.Cache, cfg.MapPrevious)
	if err != nil {
		return err
	}

	cfg.MapNext = response.Next
	cfg.MapPrevious = response.Previous

	for _, d := range response.Results {
		fmt.Println(d.Name)
	}

	return nil
}

func commandExplore(cfg *Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Usage: explore <location>")
		return nil
	}
	location := args[0]

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", location)
	fmt.Printf("Exploring %s...", location)

	response, err := fetchExplore(cfg.Cache, url)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, encounter := range response.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Usage: catch <pokemon>")
		return nil
	}
	pokemon := args[0]

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemon)
	fmt.Printf("Throwing a Pokeball at %s...", pokemon)

	response, err := fetchCatch(cfg.Cache, url)
	if err != nil {
		return err
	}

	chance := rand.IntN(100)

	if chance < response.BaseExperience {
		fmt.Printf("%s escaped!\n", pokemon)
	} else {
		fmt.Printf("%s was caught!\n", pokemon)
		p := Pokemon{
			Name:           response.Name,
			BaseExperience: response.BaseExperience,
			Height:         response.Height,
			Weight:         response.Weight,
		}

		p.Stats = make([]Stat, 0, len(response.Stats))
		for _, s := range response.Stats {
			p.Stats = append(p.Stats, Stat{
				Name: s.Stat.Name,
				Base: s.BaseStat,
			})
		}

		p.Types = make([]string, 0, len(response.Types))
		for _, t := range response.Types {
			p.Types = append(p.Types, t.Type.Name)
		}

		cfg.Pokedex[p.Name] = p
	}

	return nil
}

func commandPokedex(cfg *Config, args []string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("You haven't caught any Pokemon yet!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.Pokedex {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}

func commandInspect(cfg *Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Usage: inspect <pokemon>")
		return nil
	}
	name := args[0]

	p, ok := cfg.Pokedex[name]
	if !ok {
		fmt.Println("You have not caught that Pokémon")
		return nil
	}

	PrintPokemonInfo(p)
	return nil
}
