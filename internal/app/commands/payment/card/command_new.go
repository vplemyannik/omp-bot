package card

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/payment"
	"log"
	"time"
)

type CardNewData struct {
	OwnerId        uint64    `json:"owner_id"`
	PaymentSystem  string    `json:"payment_system"`
	Number         string    `json:"number"`
	HolderName     string    `json:"holder_name"`
	ExpirationDate time.Time `json:"expiration_date"`
	CvcCvv         string    `json:"cvc_cvv"`
}

func (c *DemoCardCommander) New(inputMsg *tgbotapi.Message) {
	parsedData := CardNewData{}
	args := inputMsg.CommandArguments()
	err := json.Unmarshal([]byte(args), &parsedData)
	if err != nil {
		log.Printf("DemoCardCommander.CallbackNewData: "+
			"error reading json data for type CallbackNewData from "+
			"input string %v - %v", inputMsg.Text, err)
		return
	}

	card := payment.Card{
		OwnerId:        parsedData.OwnerId,
		PaymentSystem:  parsedData.PaymentSystem,
		Number:         parsedData.Number,
		HolderName:     parsedData.HolderName,
		ExpirationDate: parsedData.ExpirationDate,
		CvcCvv:         parsedData.CvcCvv,
	}

	id, err := c.cardService.Create(card)

	if err != nil {
		_, err = c.bot.Send(tgbotapi.NewMessage(
			inputMsg.Chat.ID, fmt.Sprintf("DemoCardCommander.New: %s", err)))
		if err != nil {
			log.Printf("DemoCardCommander.List: error sending reply message to chat - %v", err)
		}
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("Card create with id: %d", id),
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoCardCommander.New: error sending reply message to chat - %v", err)
	}
}
