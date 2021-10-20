package card

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const limit int = 2

func (c *DemoCardCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the cards: \n\n"

	cards, err := c.cardService.List(0, uint64(limit))
	if err != nil {
		_, err = c.bot.Send(tgbotapi.NewMessage(
			inputMessage.Chat.ID, fmt.Sprintf("DemoCardCommander.List: %s", err)))
		if err != nil {
			log.Printf("DemoCardCommander.List: error sending reply message to chat - %v", err)
		}
		return
	}
	for _, p := range cards {
		outputMsgText += p.Number
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: limit,
	})

	callbackPath := path.CallbackPath{
		Domain:       domain,
		Subdomain:    subDomain,
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoCardCommander.List: error sending reply message to chat - %v", err)
	}
}
