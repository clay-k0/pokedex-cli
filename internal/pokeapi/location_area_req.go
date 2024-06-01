package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(url *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if url != nil {
		fullURL = *url
	}

	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		var locationAreas LocationAreasResp
		err := json.Unmarshal(data, &locationAreas)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationAreas, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	var locationAreas LocationAreasResp
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreas, nil
}
