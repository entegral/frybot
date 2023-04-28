package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// ImageRequestBody is the request body for the OpenAI Image Generation API
type ImageRequestBody struct {
	Prompt         string `json:"prompt"`
	N              int    `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
}

// ImageSizes is a type for the image sizes enum
type ImageSizes string

const (
	// ImageSize256 256x256
	ImageSize256 ImageSizes = "256x256"
	// ImageSize512 512x512
	ImageSize512 ImageSizes = "512x512"
	// ImageSize1024 1024x1024
	ImageSize1024 ImageSizes = "1024x1024"
)

// ImageGeneration is a function that generates images from a text prompt
func ImageGeneration(prompt string, n int, size ImageSizes) ([]string, error) {
	requestBody := ImageRequestBody{
		Prompt:         prompt,
		N:              n,
		Size:           string(size),
		ResponseFormat: "url",
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", ev.ImageAPIURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
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
	logrus.Println(string(bodyBytes))
	var response map[string]interface{}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}
	logrus.Println(string(bodyBytes))
	data := response["data"].([]interface{})
	imageURLs := make([]string, len(data))
	for i, item := range data {
		imageURLs[i] = item.(map[string]interface{})["url"].(string)
	}

	return imageURLs, nil
}
