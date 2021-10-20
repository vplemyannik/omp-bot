package card

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/payment/card"
)

type CardCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type DemoCardCommander struct {
	bot         *tgbotapi.BotAPI
	cardService service.CardService
}

func NewCardCommander(
	bot *tgbotapi.BotAPI,
) *DemoCardCommander {
	subdomainService := service.NewDummyCardService()

	return &DemoCardCommander{
		bot:         bot,
		cardService: subdomainService,
	}
}

func (c *DemoCardCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	case "new":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoCardCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *DemoCardCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
