package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	// check cache
	cachedEntry, exists := c.cache.Get(url)
	if exists {
		// unmarshal cached data and return
		cachedArea := Pokemon{}
		err := json.Unmarshal(cachedEntry, &cachedArea)
		if err != nil {
			return Pokemon{}, err
		}
		return cachedArea, nil
	}

	// set up request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// get response from request 
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	// read response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// add to cache 
	c.cache.Add(url, data)

	// unmarshal data
	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	
	return pokemon, nil
}