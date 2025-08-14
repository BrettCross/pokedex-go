package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// check if url in cache before making request
	cachedEntry, exists := c.cache.Get(url)
	if exists {
		// unmarshal cachedEntry and return it
		cachedLocation := RespShallowLocations{}
		err := json.Unmarshal(cachedEntry, &cachedLocation)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return cachedLocation, nil
	}
	

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// add to cache here
	c.cache.Add(url, data)

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)

	if err != nil {
		return RespShallowLocations{}, err
	}
	return locationsResp, nil 
}