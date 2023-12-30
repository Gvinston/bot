package commands

import (
	"github.com/Gvinston/bot/internal/service/searcher"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *Commander) SearchAnime(inMes *tgbotapi.Message) {
	searchAnime := inMes.CommandArguments()

	ss := searcher.SearcherService{}
	anime := ss.Search(searchAnime)
	var outputMsg string

	if anime.Link == "" {
		outputMsg = "Аниме не найдено"
	} else {
		outputMsg = "Найденное аниме: " + anime.Link + " \n\n"
	}

	msg := tgbotapi.NewMessage(inMes.Chat.ID, outputMsg)

	resp, err := c.bot.Send(msg)

	if err != nil {
		log.Fatalf("Ошибка при отправке ответа: %s", resp)
	}
}
