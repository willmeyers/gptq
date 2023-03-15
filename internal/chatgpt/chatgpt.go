package chatgpt

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

// GPTQParameters helps format the prompt that will be sent to ChatGPT
type GPTQParameters struct {
	Input        string // piped input that gptq will place inside [[gptqdata]] tags for ChatGPT to shape
	OutFormat    string // desired output format (text, json, csv, html...)
	Instructions string // simple instruction that ChatGPT will use to help reshape data
}

// secretKey attempts to retrieve and return an OpenAI API key
func secretKey() string {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		// If missing envvar, attempt to get OPENAI_API_KEY from .openairc
		dat, err := os.ReadFile("~/.openairc")
		if err != nil {
			fmt.Fprintf(os.Stderr, "OPENAI_API_KEY is not configured. Did you export OPEN_API_KEY or add your key to .openairc?\n")
			os.Exit(1)
		}

		return string(dat)
	}

	return apiKey
}

// Creates a new ChatGPT chat and executes the prompt based on given parameters.
func Exec(parameters GPTQParameters) (string, error) {
	apiKey := secretKey()
	client := openai.NewClient(apiKey)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a command line utility that reshapes data from one structured format to another. I want you to only reply with the with the reshaped output, and nothing else. Do not write an explanation, do not type commands, do not respond in a personable matter, and do not respond in codeblocks. If you cannot reshape the data just respond with only 'Could not execute instruction'. I only want you to respond with the reshaped output. Do not use unnecessary whitespace for the sake of making the output pretty. You will receive messages that have the unshaped data, the desired output format (this could be html, json, text, csv, etc...), and a single instruction from the user that describes how they want the data shaped. Those instructions and parameters will be placed inside [[gptq]][[/gptq]] brackets so to not interfere with the data. The inputted data (for you to reshape) will be inclosed in [[gptqdata]][[/gptqdata]] brackets. Because the inputted data can potentially be large, it can be split into multiple messages.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("[[gptq]]You instruction: %s[[/gptq]]\n[[gptq]]Your output format: %s[[/gptq]]\n[[gptqdata]]%s[[/gptqdata]]", parameters.Instructions, parameters.OutFormat, parameters.Input),
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
