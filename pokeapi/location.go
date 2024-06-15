package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationAreaResponse struct {
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ListLocations(pageUrl *string) (LocationArea, error) {
	url := URL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	// Check if location data is in cache
	if val, ok := c.cache.GetEntry(url); ok {
		location := LocationArea{}
		// using Unmarshal to store val in location [LocationArea struct]
		err := json.Unmarshal(val, &location)
		if err != nil {
			return LocationArea{}, err
		}
		return location, nil
	}

	// sending GET request to url
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}
	// getting response (res)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer res.Body.Close()
	// data is waht we get by reading all data from Response Body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	// new variable of type LocationArea{}
	locations := LocationArea{}
	// storing (after reading)data in locations
	err = json.Unmarshal(data, &locations)

	if err != nil {
		return LocationArea{}, err
	}
	// adding entry into cache for future references
	c.cache.AddEntry(url, data)
	return locations, nil

}

func (c *Client) ExploreArea(location string) ([]string, error) {
	url := URL + "/location-area/" + location
	if val, ok := c.cache.GetEntry(url); ok {
		var response LocationAreaResponse
		err := json.Unmarshal(val, &response)
		if err != nil {
			return nil, err
		}

		pokemonNames := parsePokemonNames(response)
		return pokemonNames, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response LocationAreaResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	c.cache.AddEntry(url, data)
	pokemonNames := parsePokemonNames(response)
	return pokemonNames, nil
}

func parsePokemonNames(response LocationAreaResponse) []string {
	var names []string
	for _, encounter := range response.Encounters {
		names = append(names, encounter.Pokemon.Name)
	}
	return names
}
