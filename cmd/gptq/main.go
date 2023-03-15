/*
gptq can help glue other cli programs together
It uses OpenAI's ChatGPT to take piped in data and shape it into something useful.

It's experimental and wonky. Use at your own risk!

Usage:

	gptq '[instruction]' [flags]

The flags are:

	-f
		The desired format that the inputted data should be converted to.
		Any value can be entered here: text, json, csv, html... etc.

Currently gptq only reads from standard input and writes to standard output.
If you're interested in adding more features: come check us out at https://github.com/willmeyers/gptq.
*/
package main

import "github.com/willmeyers/gptq/internal/cli"

func main() {
	cli.Run()
}
