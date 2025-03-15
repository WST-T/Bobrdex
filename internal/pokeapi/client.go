package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/WST-T/Bobrdex/internal/pokecache"
)

type Config struct {
	Next     string
	Previous string
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
