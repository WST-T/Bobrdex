package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
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

func GetLocationAreas(url string) (LocationAreaResp, error) {
	if url == "" {
		url = baseURL
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

	var locationResp LocationAreaResp
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return LocationAreaResp{}, err
	}
	return locationResp, nil
}
