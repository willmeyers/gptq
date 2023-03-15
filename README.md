# gptq

gptq is an experimental command-line tool for helping reshape structured data into more structured data.

## Installing

`gptq` is recommended to be run with `>=1.19`.

#### Building from source

Simply checkout the repo and run `make build`.

## Usage

```
gptq -h
```

A basic example:

```bash
cat users.csv | gptq 'json list of users sorted by id' -f json | jq > users.json
```

### Disclaimer

This is a fun toy. Do not ever use this in a production environment or
pipe in any sensitive data you don't want OpenAI to have.
