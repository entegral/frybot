package api

import (
	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

// EnvVars is the struct for the environment variables used in this application
type EnvVars struct {
	OpenAIAPIKey     string `env:"OPENAI_API_KEY,notEmpty"`
	ChatAPIURL       string `env:"CHAT_API_URL" envDefault:"https://api.openai.com/v1/chat/completions"`
	CompletionAPIURL string `env:"COMPLETION_API_URL" envDefault:"https://api.openai.com/v1/completions"`
	ImageAPIURL      string `env:"IMAGE_API_URL" envDefault:"https://api.openai.com/v1/images/generations"`
	ModelsAPIURL     string `env:"MODELS_API_URL" envDefault:"https://api.openai.com/v1/models"`
}

var ev EnvVars

func init() {
	if err := env.Parse(&ev); err != nil {
		logrus.Fatalln(err)
	}
}
