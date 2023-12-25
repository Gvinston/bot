package main

import (
	"github.com/Gvinston/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)

	productService := product.NewService()

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	for len(updates) != 0 {
		<-updates
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Command() == "help" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "/help - help")

			bot.Send(msg)
			continue
		}
		if update.Message.Command() == "list" {
			outputMsg := "Here all the products: \n\n"
			products := productService.List()

			for _, p := range products {
				outputMsg += p.Title
				outputMsg += "\n"
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, outputMsg)

			bot.Send(msg)
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You wrote: "+update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID
		// Add logic here
		bot.Send(msg)
	}
}
