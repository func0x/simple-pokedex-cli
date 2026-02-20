package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/func0x/pokedexcli/internal/pokecache"
)

func fetchLocations(cache *pokecache.Cache, url string) (*LocationResponse, error) {

	if data, ok := cache.Get(url); ok {
		fmt.Println("CACHE HIT")
		var response LocationResponse
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, fmt.Errorf("error decoding JSON: %v", err)
		}
		return &response, nil
	}

	fmt.Println("CACHE MISS")

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("network error: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	cache.Add(url, body)

	var response LocationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return &response, nil
}

func fetchExplore(cache *pokecache.Cache, url string) (*ExploreResponse, error) {

	if data, ok := cache.Get(url); ok {
		fmt.Println("CACHE HIT")

		var response ExploreResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil, err
		}
		return &response, nil
	}

	fmt.Println("CACHE MISS")

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("network error: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	cache.Add(url, body)

	var response ExploreResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return &response, nil
}

func fetchCatch(cache *pokecache.Cache, url string) (*PokemonResponse, error) {
	if data, ok := cache.Get(url); ok {
		fmt.Println("CACHE HIT")

		var response PokemonResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil, err
		}
		return &response, nil
	}

	fmt.Println("CACHE MISS")

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("network error: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	cache.Add(url, body)

	var response PokemonResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	return &response, nil
}

func PrintPokemonInfo(p Pokemon) {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)

	if len(p.Stats) > 0 {
		fmt.Println("Stats:")
		for _, s := range p.Stats {
			fmt.Printf("  - %s: %d\n", s.Name, s.Base)
		}
	}

	if len(p.Types) > 0 {
		fmt.Println("Types:")
		for _, t := range p.Types {
			fmt.Printf("  - %s\n", t)
		}
	}
}
