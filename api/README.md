Prompt: please respond with only the code, no commentary: please write a readme for this package, including enviornment variales. 

Response:
$!FILENAME=README.md
# OpenAI API Go Client

This is a Go client for the OpenAI API. It currently supports the following endpoints:

- [Chat API](https://beta.openai.com/docs/api-reference/chat)
- [Completion API](https://beta.openai.com/docs/api-reference/completions)
- [Image API](https://beta.openai.com/docs/api-reference/images)
- [Models API](https://beta.openai.com/docs/api-reference/models)

## Usage

To use this client, you will need an OpenAI API key. You can sign up for an API key [here](https://beta.openai.com/signup/).

### Environment Variables

This client requires the following environment variables to be set:

- `OPENAI_API_KEY`: Your OpenAI API key

### Examples

#### Chat API

```go
package main

import (
	"fmt"

	"github.com/openai/api-go"
)

func main() {
	response, err := api.AnalyzeCode(api.GPT4, "Hello, how are you?", "", 0.5)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response.String())
}
```

#### Completion API

```go
package main

import (
	"fmt"

	"github.com/openai/api-go"
)

func main() {
	cr := api.CompletionRequest{
		Model:       api.Davinci,
		Prompt:      "Hello, how are you?",
		MaxTokens:   10,
		Temperature: 0.5,
	}

	response, err := api.Complete(cr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response.Choices[0].Text)
}
```

#### Image API

```go
package main

import (
	"fmt"

	"github.com/openai/api-go"
)

func main() {
	imageURLs, err := api.ImageGeneration("A painting of a cat", 1, api.ImageSize256)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(imageURLs[0])
}
```

#### Models API

```go
package main

import (
	"fmt"

	"github.com/openai/api-go"
)

func main() {
	models, err := api.GetAvailableModels()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, model := range models {
		fmt.Println(model.ID)
	}
}
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)