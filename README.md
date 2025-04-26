# Telegram Bot with Text Reversal and Story Generation

A versatile Telegram bot written in Go that offers multiple features including message reversal and AI-powered story generation.

## Features

### Commands
- `/hello` - Get a personalized greeting
- `/story` - Generate a unique sci-fi story using OpenAI's API
- `/help` - Display all available commands
- Any text message - Get its reversed version

### Text Reversal Capabilities
- Reverses any text message while preserving character encoding (UTF-8 support)
- Works with:
  - Multiple languages (English, Russian, etc.)
  - Emojis and special characters
  - Numbers and symbols
  - Spaces and punctuation

### Other Features
- Personalized greetings using user's Telegram name or username
- AI-powered story generation
- Helpful command suggestions for unknown commands

## Prerequisites

- Go 1.21 or later
- A Telegram Bot Token (get it from [@BotFather](https://t.me/botfather))
- OpenAI API Key (for story generation feature)

## Setup

1. Clone this repository
2. Copy `.env` file and set your tokens:
   ```bash
   cp .env.example .env
   # Edit .env and add your:
   # - TELEGRAM_BOT_TOKEN
   # - OPENAI_API_KEY
   ```
3. Install dependencies:
   ```bash
   go mod download
   ```

## Running the Bot

```bash
go run main.go
```

## Usage Examples

1. Send `/hello` to get a personalized greeting
2. Send `/story` to receive a unique AI-generated sci-fi story
3. Send any text message to get its reversed version:
   - "hello" → "olleh"
   - "привет" → "тевирп"
   - "hello world" → "dlrow olleh"
4. Send `/help` to see all available commands

## Testing

To run the tests:
```bash
cd utils
go test -v
```

The bot will start and respond to commands and messages accordingly. You can stop the bot by pressing Ctrl+C.

## Error Handling

The bot includes robust error handling for various scenarios:
- Invalid commands show a help message with available commands
- OpenAI API quota exceeded notifications
- Network and API error handling
- Graceful shutdown support