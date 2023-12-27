package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *Commander) SearchAnime(inMes *tgbotapi.Message) {
	//searchAnime := inMes.CommandArguments()

	outputMsg := "Searched anime: anime \n\n"

	msg := tgbotapi.NewMessage(inMes.Chat.ID, outputMsg)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Rate", "Rate"),
			tgbotapi.NewInlineKeyboardButtonData("Year", "Year"),
		),
	)

	resp, err := c.bot.Send(msg)

	if err != nil {
		log.Fatalf("Ошибка при отправке ответа: %s", resp)
	}
}
