# Luna

🌙 **Luna** is your AI-powered Linux command-line companion. Luna uses Google's Gemini to answer your Linux questions, explain command outputs, and debug scripts directly in your terminal.

## Features

- Ask Luna anything about Linux and receive concise, actionable answers.
- Pipe command output to Luna for instant explanations or debugging.
- Configure with your own Google AI Studio API Key and model.
- Markdown rendering for beautiful, readable responses.

## Installation

### Quick Install (Recommended)

You can install Luna directly from the main branch using Go:

```sh
go install github.com/sabrek15/luna@latest
```

This will place the `luna` binary in your `$GOPATH/bin` or `$HOME/go/bin` directory.

> Make sure your Go environment is set up ([Download Go](https://golang.org/dl/)). Requires Go 1.18+.

### Manual Install

1. **Clone the repository:**
   ```sh
   git clone https://github.com/sabrek15/luna.git
   cd luna
   ```

2. **Build the binary:**
   ```sh
   go build -o luna .
   ```

3. **(Optional) Install to your `$PATH`:**
   ```sh
   sudo mv luna /usr/local/bin/luna
   ```

### Install Required Go Libraries

Luna depends on several Go libraries for CLI, configuration, and markdown rendering. These libraries are automatically installed when you run `go install` or `go build`, but you can install them manually if needed:

```sh
go get github.com/spf13/cobra@latest
go get github.com/spf13/viper@latest
go get -u github.com/google/generative-ai-go/genai
go get github.com/charmbracelet/glamour@latest
```

> These commands will add the packages to your Go module and download their latest versions.

## Getting a Gemini API Key

Luna requires a Gemini API key from Google AI Studio.

1. **Visit [Google AI Studio](https://aistudio.google.com/app/apikey)**
2. Sign in with your Google account.
3. Create a new API key for Gemini.
4. Copy your API key.

Set your API key in Luna:
```sh
luna config --set-key YOUR_API_KEY
```

For more details, see [Google AI Studio Documentation](https://aistudio.google.com/app/apikey).

## Configuration

Set your API key as shown above.

You can also change the Gemini model (default: `gemini-2.5-flash`):

```sh
luna config --set-model gemini-2.5-flash
```

View your current configuration:

```sh
luna config --show
```

## Usage

### Ask Luna a question

```sh
luna "How do I list all files recursively?"
```

### Explain command output

Pipe output from another command:

```sh
ls -l | luna "What do these permissions mean?"
```

### General Command Structure

```sh
luna "<your question>"
<some_command> | luna "<question about the piped input>"
```

## Commands

- `luna "<question>"` — Ask Luna any Linux-related question.
- `luna config --set-key <YOUR_API_KEY>` — Set your Google AI Studio API Key.
- `luna config --set-model <MODEL_NAME>` — Set the Gemini model.
- `luna config --show` — Show your current configuration.
- `luna config` — Shows help for configuration commands.

## Repository Structure

```
.
├── cmd/                # CLI entry points (root, config commands)
│   ├── root.go         # Main command logic ('luna' command)
│   └── config.go       # Configuration management
├── internal/
│   ├── ai/             # Gemini AI integration
│   │   └── gemini.go   # Generates content via Gemini API
│   ├── config/         # Configuration file logic
│   │   └── config.go   # Load and save config, config struct
│   └── tui/            # Terminal user interface
│       └── renderer.go # Markdown rendering for terminal
├── main.go             # Main application entrypoint
├── README.md           # This file
```

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

---

Feel free to reach out for help or suggestions!