# Bobrdex: Your Command-Line Pokémon Adventure (⌒‿⌒)

![image](https://github.com/user-attachments/assets/b5498d09-6baa-4145-8077-f5435a438ad6)

Welcome to Bobrdex, an adorable command-line Pokédex application that lets you explore the Pokémon world, catch your favorite Pokémon, and build your own collection! ʕ•ᴥ•ʔ

## Features (づ｡◕‿‿◕｡)づ

✨ **Explore Locations:** Discover different areas in the Pokémon world and find out which Pokémon live there.

✨ **Catch Pokémon:** Throw Pokéballs and try to catch wild Pokémon (harder ones have higher escape rates!).

✨ **Inspect Pokémon:** View detailed information about Pokémon you've caught, including stats and types.

✨ **Track Your Collection:** Keep track of all the Pokémon you've caught in your personal Pokédex.

## Commands (◠‿◠✿)

| Command | Description |
|---------|-------------|
| `help` | Show available commands |
| `exit` | Exit the Pokédex |
| `map` | Display 20 location areas |
| `mapb` | Display previous 20 location areas |
| `explore <location>` | Explore a location area for Pokémon |
| `catch <pokemon>` | Try to catch a Pokémon |
| `inspect <pokemon>` | View details about a caught Pokémon |
| `pokedex` | View all your caught Pokémon |

## How to Use (｡♥‿♥｡)

1. **Start your adventure:**
   ```
   ./Bobrdex
   ```

2. **Explore locations to find Pokémon:**
   ```
   Pokedex > map
   Pokedex > explore canalave-city-area
   ```

3. **Catch Pokémon you discover:**
   ```
   Pokedex > catch pikachu
   ```

4. **View your collection:**
   ```
   Pokedex > pokedex
   ```

5. **Inspect your caught Pokémon:**
   ```
   Pokedex > inspect pikachu
   ```

## Example Session (⁀ᗢ⁀)

```
Welcome to the Pokedex!
Type 'help' to see available commands

Pokedex > map
Location areas:
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Pokemon encounters:
  tentacool
  tentacruel
  magikarp
  gyarados
  ...

Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!

Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats:
  -hp: 20
  -attack: 10
  -defense: 55
  -special-attack: 15
  -special-defense: 20
  -speed: 80
Types:
  - water

Pokedex > pokedex
Caught Pokemon:
  magikarp (water)

Total caught: 1
```

## Developing (ﾉ◕ヮ◕)ﾉ*:･ﾟ✧

This project uses:
- Go 1.24
- The [PokeAPI](https://pokeapi.co/) for Pokémon data
- A custom caching system for improved performance

## License (✿◠‿◠)

MIT License - See LICENSE file for details.

---

Made with ❤️ and Go. Happy Pokémon catching! (づ･ω･)づ
