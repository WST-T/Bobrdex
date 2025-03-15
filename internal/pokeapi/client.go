package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/WST-T/Bobrdex/internal/pokecache"
)

type Config struct {
	Next          string
	Previous      string
	CaughtPokemon map[string]Pokemon
}

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
}

type LocationAreaResp struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
type LocationAreaDetailResp struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

const baseURL = "https://pokeapi.co/api/v2/location-area/"
const pokemonURL = "https://pokeapi.co/api/v2/pokemon/"

var cache = pokecache.NewCache(5 * time.Minute)

func GetLocationAreas(url string) (LocationAreaResp, error) {
	if url == "" {
		url = baseURL
	}

	if cachedData, found := cache.Get(url); found {
		var locationResp LocationAreaResp
		err := json.Unmarshal(cachedData, &locationResp)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locationResp, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	cache.Add(url, body)

	var locationResp LocationAreaResp
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return LocationAreaResp{}, err
	}
	return locationResp, nil
}

func GetLocationAreaDetails(locationAreaName string) (LocationAreaDetailResp, error) {
	url := baseURL + locationAreaName

	if cachedData, found := cache.Get(url); found {
		var locationDetailResp LocationAreaDetailResp
		err := json.Unmarshal(cachedData, &locationDetailResp)
		if err != nil {
			return LocationAreaDetailResp{}, err
		}
		return locationDetailResp, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return LocationAreaDetailResp{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaDetailResp{}, err
	}

	cache.Add(url, body)

	var locationDetailResp LocationAreaDetailResp
	err = json.Unmarshal(body, &locationDetailResp)
	if err != nil {
		return LocationAreaDetailResp{}, err
	}
	return locationDetailResp, nil
}

func GetPokemon(pokemonName string) (Pokemon, error) {
	url := pokemonURL + strings.ToLower(pokemonName)

	if cachedData, found := cache.Get(url); found {
		var pokemon Pokemon
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return Pokemon{}, fmt.Errorf("Pokemon not found: %s", pokemonName)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	cache.Add(url, body)

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
