Pokémon CLI

A simple terminal program in Go to explore, catch Pokémon, and view your Pokedex.

Features
	•	Explore locations and find Pokémon.
	•	Catch Pokémon based on their BaseExperience.
	•	View your Pokedex with all caught Pokémon.
	•	Check Pokémon details: height, weight, stats, and types.

Install
	1.	Install Go (1.21+).
	2.	Clone the repo:

git clone https://github.com/YourRepo/pokedexcli.git
cd pokedexcli

	3.	Download dependencies:

go mod tidy

	4.	Run the program:

go run main.go

Usage
	•	explore <location> – find and try to catch Pokémon.
	•	pokedex – list all caught Pokémon.
	•	inspect <pokemon> – see detailed info of a Pokémon.

Example
> explore canalave-city-area
Exploring canalave-city-area...
Found Pokemon:
- tentacool
- tentacruel
- staryu
- magikarp
- gyarados
- wingull
- pelipper
- shellos
- gastrodon
- finneon
- lumineon

> catch pidgey
Pidgey was caught!

> pokedex
  - pidgey

> inspect pidgey
Name: Pidgey
Height: 3
Weight: 18
Stats:
  - hp: 40
  - attack: 45
  - defense: 40
  - special-attack: 35
  - special-defense: 35
  - speed: 56
Types:
  - normal
  - flying

Project structure

pokedexcli/
├─ main.go          # Entry point
├─ commands.go      # CLI commands
├─ utils.go         # Helper functions
├─ internal/
│  ├─ pokecache/    # Cache for API data


Requirements
	•	Go 1.21+
