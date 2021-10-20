package card

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/payment"
	"log"
	"strconv"
	"strings"
)

func (c *DemoCardCommander) Edit(inputMsg *tgbotapi.Message) {
	parsedData := CardNewData{}
	args := inputMsg.CommandArguments()
	argsArr := strings.SplitN(args, " ", 2)
	err := json.Unmarshal([]byte(argsArr[1]), &parsedData)
	if err != nil {
		log.Printf("DemoCardCommander.CallbackNewData: "+
			"error reading json data for type CallbackNewData from "+
			"input string %v - %v", inputMsg.Text, err)
		return
	}

	idx, err := strconv.ParseUint(argsArr[0], 10, 64)
	if err != nil {
		log.Println("wrong args", args)
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

	err = c.cardService.Update(idx, card)

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
		fmt.Sprintf("Card with id=%d has been updated", idx),
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoCardCommander.Edit: error sending reply message to chat - %v", err)
	}
}
