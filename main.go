package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/joho/godotenv"
	"tg_bot/utils"
)

const helpMessage = "Available commands:\n" +
	"/hello - Get a personal greeting\n" +
	"/story - Generate a short sci-fi story\n" +
	"/help - Show this help message\n" +
	"Any text message - Get reversed version of your text"

func getGreeting(user *gotgbot.User) string {
	name := user.FirstName
	if user.Username != "" {
		name = "@" + user.Username
	}
	return fmt.Sprintf("👋 Hello, %s! Nice to meet you!", name)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is not set")
	}

	// Create bot
	b, err := gotgbot.NewBot(token, nil)
	if err != nil {
		log.Fatal("Error creating bot: ", err)
	}

	// Create dispatcher
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Printf("Error handling update: %v", err)
			return ext.DispatcherActionNoop
		},
	})

	// Handler for all commands (must be first)
	dispatcher.AddHandler(handlers.NewMessage(
		func(msg *gotgbot.Message) bool {
			return msg.Text != "" && strings.HasPrefix(msg.Text, "/") && 
				!strings.HasPrefix(msg.Text, "/start") && 
				!strings.HasPrefix(msg.Text, "/help") && 
				!strings.HasPrefix(msg.Text, "/story") &&
				!strings.HasPrefix(msg.Text, "/hello")
		},
		func(b *gotgbot.Bot, ctx *ext.Context) error {
			msg := "Unknown command.\n\n" + helpMessage
			_, err := ctx.EffectiveChat.SendMessage(b, msg, nil)
			return err
		},
	))

	// Add handler for /start and /help commands
	dispatcher.AddHandler(handlers.NewCommand("start", func(b *gotgbot.Bot, ctx *ext.Context) error {
		greeting := getGreeting(ctx.EffectiveUser)
		_, err := ctx.EffectiveChat.SendMessage(b, greeting+"\n\n"+helpMessage, nil)
		return err
	}))
	
	dispatcher.AddHandler(handlers.NewCommand("help", func(b *gotgbot.Bot, ctx *ext.Context) error {
		_, err := ctx.EffectiveChat.SendMessage(b, helpMessage, nil)
		return err
	}))

	// Add handler for /hello command
	dispatcher.AddHandler(handlers.NewCommand("hello", func(b *gotgbot.Bot, ctx *ext.Context) error {
		greeting := getGreeting(ctx.EffectiveUser)
		_, err := ctx.EffectiveChat.SendMessage(b, greeting, nil)
		return err
	}))

	// Add handler for /story command
	dispatcher.AddHandler(handlers.NewCommand("story", func(b *gotgbot.Bot, ctx *ext.Context) error {
		story, err := utils.GenerateSciFiStory()
		if err != nil {
			log.Printf("Error generating story: %v", err)
			_, err := ctx.EffectiveChat.SendMessage(b, "Sorry, I couldn't generate a story right now. Try again later.", nil)
			return err
		}
		_, err = ctx.EffectiveChat.SendMessage(b, story, nil)
		return err
	}))

	// Add handler for text messages (non-commands)
	dispatcher.AddHandler(handlers.NewMessage(
		func(msg *gotgbot.Message) bool { 
			return !strings.HasPrefix(msg.Text, "/")
		},
		func(b *gotgbot.Bot, ctx *ext.Context) error {
			reversed := utils.ReverseString(ctx.EffectiveMessage.Text)
			_, err := ctx.EffectiveChat.SendMessage(b, reversed, nil)
			return err
		},
	))

	// Create updater
	updater := ext.NewUpdater(dispatcher, nil)

	// Start receiving updates
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
	})
	if err != nil {
		log.Fatal("Error starting polling: ", err)
	}
	log.Printf("Bot @%s started", b.User.Username)

	// Idle until shutdown
	updater.Idle()
} 