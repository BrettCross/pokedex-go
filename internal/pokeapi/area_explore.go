package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreArea(area string) (RespAreaEncounters, error) {
	url := baseURL + "/location-area/" + area

	// check cache before sending request
	cachedEntry, exists := c.cache.Get(url)
	if exists {
		// unmarshal cached data and return
		cachedArea := RespAreaEncounters{}
		err := json.Unmarshal(cachedEntry, &cachedArea)
		if err != nil {
			return RespAreaEncounters{}, err
		}
		return cachedArea, nil
	}
	
	// set up the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespAreaEncounters{}, err
	}

	// get the response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespAreaEncounters{}, err
	}

	defer resp.Body.Close()

	// read response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespAreaEncounters{}, err	
	}

	// add data to cache
	c.cache.Add(url, data)

	// unmarshal into struct and return
	areaResp := RespAreaEncounters{}
	err = json.Unmarshal(data, &areaResp)
	if err != nil {
		return RespAreaEncounters{}, err
	}

	return areaResp, nil
}