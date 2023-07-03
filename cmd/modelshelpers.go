package cmd

import (
	"context"
	"encoding/json"
	"net/http"
)

type ModelResponse struct {
	Data []Model `json:"data"`
}

type Model struct {
	ID         string            `json:"id"`
	Object     string            `json:"object"`
	Created    int               `json:"created"`
	OwnedBy    string            `json:"owned_by"`
	Permission []ModelPermission `json:"permission"`
}

type ModelPermission struct {
	ID                 string  `json:"id"`
	Object             string  `json:"object"`
	Created            int64   `json:"created"`
	AllowCreateEngine  bool    `json:"allow_create_engine"`
	AllowSampling      bool    `json:"allow_sampling"`
	AllowLogprobs      bool    `json:"allow_logprobs"`
	AllowSearchIndices bool    `json:"allow_search_indices"`
	AllowView          bool    `json:"allow_view"`
	AllowFineTuning    bool    `json:"allow_fine_tuning"`
	Organization       string  `json:"organization"`
	Group              *string `json:"group"`
	IsBlocking         bool    `json:"is_blocking"`
}

const ModelURL = "https://api.openai.com/v1/models"

// GetModels use the "OPENAI_API_KEY" environment variable as an "Authorization: Bearer" header and issue a get request on the url
// return the response body as a []Model
func GetModels(ctx context.Context, url, apiKey string) []Model {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "frybot")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var modelResponse ModelResponse
	if err := json.NewDecoder(resp.Body).Decode(&modelResponse); err != nil {
		panic(err)
	}
	return modelResponse.Data
}
