# pokedex-go

A command-line Pokédex application written in Go. This tool allows you to interact with Pokémon data, catch Pokémon, and manage your own Pokédex using the [PokéAPI](https://pokeapi.co/).

## Features

- Search for Pokémon by name
- Catch Pokémon with a randomized chance
- Maintain a local Pokédex of caught Pokémon
- Simple command-line interface

## Requirements

- Go 1.18 or newer

## Installation

Clone the repository:

```bash
git clone https://github.com/brettcross/pokedex-go.git
cd pokedex-go
```

Build the project:

```bash
go build -o pokedex-go
```

## Usage

Run the application:

```bash
./pokedex-go
```

### Commands
- `help`: Display available commands
- `exit`: close the Pokédex REPL
- `map`: Display names of next 20 locations
- `mapb`: Display names of previous 20 locations
- `explore <area_name>`: Display available Pokémon at given area
- `catch <pokemon_name>`: Attempt to catch a Pokémon by name.
- `inspect <pokemon_name>`: Display given Pokémon data
- `pokedex`: Display all previously caught pokemon

Example:

```bash
./pokedex-go catch pikachu
```

## Configuration

The application uses the PokéAPI to fetch Pokémon data. No additional configuration is required.



## License

MIT
