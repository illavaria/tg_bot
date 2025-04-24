# Simple Telegram Echo Bot

A simple Telegram bot written in Go that replies "hi" to every message.

## Prerequisites

- Go 1.21 or later
- A Telegram Bot Token (get it from [@BotFather](https://t.me/botfather))

## Setup

1. Clone this repository
2. Copy `.env` file and set your bot token:
   ```bash
   cp .env.example .env
   # Edit .env and add your bot token
   ```
3. Install dependencies:
   ```bash
   go mod download
   ```

## Running the Bot

```bash
go run main.go
```

The bot will start and respond "hi" to every message it receives. You can stop the bot by pressing Ctrl+C.

## Testing the Bot

1. Start a chat with your bot on Telegram
2. Send any message
3. The bot will reply with "hi"