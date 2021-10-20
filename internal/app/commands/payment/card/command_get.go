package card

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DemoCardCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	card, err := c.cardService.Describe(idx)
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}

	year, month, _ := card.ExpirationDate.Date()
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("OwnerId: %d \n", card.OwnerId)+
			fmt.Sprintf("HolderName: %s \n", card.HolderName)+
			fmt.Sprintf("PaymentSystem: %s \n", card.PaymentSystem)+
			fmt.Sprintf("Number: %s \n", card.Number)+
			fmt.Sprintf("Cvc/Cvv: %s \n", card.CvcCvv)+
			fmt.Sprintf("ExpirationDate: %d/%d \n", year, int(month)),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoCardCommander.Get: error sending reply message to chat - %v", err)
	}
}
