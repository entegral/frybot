# Frybot

Frybot is an AI language model that can assist users with answering questions about code in their current working directory (`cwd`). The `cwd` may be a combination of more than one file, but must not exceed the token length for the model you have selected. 

`Prompt` - loads the `cwd` into the model and allows you to ask questions about the code. 

`Chat` - currently does not load the `cwd` into the model because the token lengths provided by OpenAI models dont currently support large amounts of text. This command is a work in progress and should eventually allow you to chat with Frybot about the code in your `cwd`, probably when the `gpt-4-32k` model is released.

## Getting Started

To use Frybot, you will need to have an OpenAI API key. You can sign up for an API key on the [OpenAI website](https://beta.openai.com/signup/).

### Prerequisites

- Go 1.19 or later

### Installing - build yourself

1. Clone the Frybot repository:

   ```
   git clone https://github.com/openai/frybot.git
   ```

2. Change into the Frybot directory:

   ```
   cd frybot
   ```

3. Set your OpenAI API key as an environment variable:

   ```
   export OPENAI_API_KEY=<your-api-key>
   ```

4. Run the Frybot conversation:

   ```
   go run main.go
   ```

5. Optionally, you could run `make build` instead. This command will build the binary for your operating system and place it in the GOPATH bin directory. **Please ensure you have the GOPATH bin directory in your PATH environment variable.**

## Usage


### Chat

The `chat` command initializes a chat session with the bot. Type `exit` to exit or press `crtl + c`. You may also say `save this conversation` to save the conversation to a file named `frybot_conversation.txt` in your `cwd`.

```bash
frybot chat
```

### Prompt

The `prompt` command prompts the bot with a one-off question about local files.

Frybot can assist you with answering questions about provided code with the `prompt` command. Simply navigate to a directory with files in it and run `frybot prompt "<your question here>"`. The tool will combine all files in the directory into a single context, then send it along with your prompt to the OpenAI API. The response will be printed to the console.

You may also save the response to a file with the `-s` flag. This will save the response to a file named `frybot_response.txt` in your `cwd`.

```bash

```bash
frybot prompt -p "How do I find a User record using platformservices?" -t "path/to/file"
```

The `-p` flag is required and specifies the prompt for the bot to answer. The `-t` flag is optional and specifies the file to provide context to the prompt.
