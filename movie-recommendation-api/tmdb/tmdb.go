package tmdb

import (
	"io"
	"net/http"
	"encoding/json"
)

type Client struct {
	APIKey string
}

func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

// RequestAuthenticaton initiates an authentication session with TMDB
func (c *Client) RequestAuthentication() (map[string]interface{}, error) {
	url := "https://api.themoviedb.org/3/trending/movie/week?api_key="

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}    
	err = json.Unmarshal(body, &result)     
	 if err != nil {                 
		return  nil, err    
	}

	return result, nil
}
