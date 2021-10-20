package card

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *DemoCardCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		fmt.Sprintf("/help__%s__%s — print list of commands\n", domain, subDomain)+
			fmt.Sprintf("/list__%s__%s — get a list of your entity (💎: with pagination via telegram keyboard)\n", domain, subDomain)+
			fmt.Sprintf("/delete__%s__%s — delete an existing entity\n", domain, subDomain)+
			fmt.Sprintf("/new__%s__%s — create a new entity\n", domain, subDomain)+
			fmt.Sprintf("/edit__%s__%s — edit a entity\n", domain, subDomain),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoCardCommander.Help: error sending reply message to chat - %v", err)
	}
}
