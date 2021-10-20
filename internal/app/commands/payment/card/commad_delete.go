package card

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *DemoCardCommander) Delete(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	cardId, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	ok, err := c.cardService.Remove(cardId)

	if !ok {
		_, err = c.bot.Send(tgbotapi.NewMessage(
			inputMsg.Chat.ID, fmt.Sprintf("DemoCardCommander.Delete: %s", err)))
		if err != nil {
			log.Printf("DemoCardCommander.Get: error sending reply message to chat - %v", err)
		}
		return
	}

	_, err = c.bot.Send(tgbotapi.NewMessage(
		inputMsg.Chat.ID, "Done!"))

	if err != nil {
		log.Printf("DemoCardCommander.Get: error sending reply message to chat - %v", err)
	}
}
