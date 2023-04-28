package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Models string

const (
	// Davinci is a bit slower than gpt3.5 turbo but has a higher token limit of ~8k
	Davinci Models = "text-davinci-003"

	// GPT3Turbo is a reasonably new and cheap gpt model, but has a lower token limit of ~4k
	GPT3Turbo Models = "gpt-3.5-turbo"

	// GPT4 is the most expensive model, but has a token limit of ~8k and quality analysis
	GPT4 Models = "gpt-4"
)

// Model is a struct representing an individual model from the API
type Model struct {
	ID         string        `json:"id"`
	Object     string        `json:"object"`
	OwnedBy    string        `json:"owned_by"`
	Permission []interface{} `json:"permission"`
}

// ModelListResponse is a struct representing the list of models returned by the API
type ModelListResponse struct {
	Data   []Model `json:"data"`
	Object string  `json:"object"`
}

// GetAvailableModels is a function that retrieves the list of available models from the API
func GetAvailableModels() ([]Model, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", ev.ModelsAPIURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ev.OpenAIAPIKey))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// logrus.Println(string(bodyBytes))

	var response ModelListResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
