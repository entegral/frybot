Prompt: please only respond with file contents: can you please create a readme for using the code i have provided?

Response:
$!FILENAME=README.md
# Frybot

Frybot is a command-line interface (CLI) tool that uses OpenAI's GPT-3 to answer questions about local files. It is designed to help developers quickly find answers to questions they have about their codebase.

## Installation

To install Frybot, clone this repository and run `go install` from the root directory.

## Usage

Frybot has two commands: `chat` and `prompt`.

### Chat

The `chat` command initializes a chat session with the bot. Use `ctrl + c` to exit.

```bash
frybot chat
```

### Prompt

The `prompt` command prompts the bot with a one-off question about local files.

```bash
frybot prompt -p "How do I find a User record using platformservices?" -t "path/to/file"
```

The `-p` flag is required and specifies the prompt for the bot to answer. The `-t` flag is optional and specifies the file to provide context to the prompt.

## Contributing

Contributions are welcome! Please see the [contributing guidelines](CONTRIBUTING.md) for more information.

## License

This project is licensed under the [MIT License](LICENSE).