package card

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *DemoCardCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("DemoCardCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}
	outputMsgText := "Here all the cards: \n\n"

	cards, err := c.cardService.List(uint64(parsedData.Offset), uint64(limit))
	if err != nil {
		log.Printf("DemoCardCommander.CallbackList: %s", err)
		return
	}
	for _, p := range cards {
		outputMsgText += p.Number
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	if len(cards) < limit {
		return
	}

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: parsedData.Offset + limit,
	})

	newCallbackPath := path.CallbackPath{
		Domain:       domain,
		Subdomain:    subDomain,
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", newCallbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoCardCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
