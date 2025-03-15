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
