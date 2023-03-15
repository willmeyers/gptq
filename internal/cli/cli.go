package cli

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/willmeyers/gptq/internal/chatgpt"
)

func scanInput() string {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = input + scanner.Text()
	}

	return input
}

// Run the gptq command. Will block until shutdown or a termination signal is received.
func Run() {
	outFormat := flag.String("f", "text", "Desired output format (text, json, html, table, ...etc)")
	flag.Parse()

	instructions := flag.Arg(0)
	if instructions == "" {
		fmt.Fprintf(os.Stderr, "Missing instructions\n")
		os.Exit(1)
	}

	input := scanInput()
	parameters := chatgpt.GPTQParameters{
		Instructions: instructions,
		Input:        input,
		OutFormat:    *outFormat,
	}

	result, err := chatgpt.Exec(parameters)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad response: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
