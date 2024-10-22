package azzurro

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	// The base URL for the API
	baseURL string
	// The API key to use for authentication
	apiAuth  string
	clientID string

	httpClient *http.Client
}

// NewClient creates a new client with the given API key, client ID and default base URL
func NewClient(apiAuth string, clientID string) *Client {
	return NewClientWithBaseURL(apiAuth, clientID, "https://third.zcsazzurroportal.com:19003")
}

// NewClientWithBaseURL creates a new client with the given API key and base URL
func NewClientWithBaseURL(apiAuth string, clientID string, baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		apiAuth:    apiAuth,
		clientID:   clientID,
		httpClient: http.DefaultClient,
	}
}

// FetchRealtimeData fetches the realtime data for the thing with the given ID
func (c *Client) FetchRealtimeData(thingID string) (RealtimeDataResponse, error) {
	cmd := RealtimeDataRequest{
		RealtimeData: RealtimeData{
			Command: "realtimeData",
			Params: Params{
				ThingKey:       thingID,
				RequiredValues: "*",
			},
		},
	}

	// Encode cmd to JSON
	jsonData, err := json.Marshal(cmd)
	if err != nil {
		return RealtimeDataResponse{}, err
	}

	// Create a new request with JSON-encoded cmd as the body
	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return RealtimeDataResponse{}, err
	}

	// Add the API key to the request
	req.Header.Add("Authorization", c.apiAuth)
	req.Header.Add("Client", c.clientID)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RealtimeDataResponse{}, err
	}
	defer resp.Body.Close()

	// Decode the response
	var data RealtimeDataResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return RealtimeDataResponse{}, err
	}

	// Return the data
	return data, nil
}
