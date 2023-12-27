package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *Commander) List(inMes *tgbotapi.Message) {
	args := inMes.CommandArguments()

	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Panicln("wrong args", args)
	}

	outputMsg := "Here all the products: \n\n"
	products := c.productService.List()

	for _, p := range products {
		outputMsg += p.Title
		outputMsg += fmt.Sprintf("arg %v", arg)
		outputMsg += "\n"
	}

	msg := tgbotapi.NewMessage(inMes.Chat.ID, outputMsg)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1", "1"),
		),
	)
	var resp tgbotapi.Message
	resp, err = c.bot.Send(msg)

	if err != nil {
		log.Fatalf("Ошибка при отправке ответа: %s", resp)
	}
}
