package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gocolly/colly/v2"
	"log"
	"regexp"
	"strings"
)

func (b Bot) updateCommand(message *tgbotapi.Message) {
	log.Printf("[%s]:- Command: %s", message.From.UserName, message.Text)
	switch message.Command() {
	case start:
		msg := tgbotapi.NewMessage(message.Chat.ID, startMessage)
		b.bot.Send(msg)
	case help:
		msg := tgbotapi.NewMessage(message.Chat.ID, helpMessage)
		b.bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, unknownCommand)
		b.bot.Send(msg)
	}
}

func (b Bot) updateMessage(message *tgbotapi.Message) {
	log.Printf("[%s]:- %s", message.From.UserName, message.Text)

	re := regexp.MustCompile(groupExample)
	if re.MatchString(message.Text) {
		b.selectGropu(message)
	}
	switch message.Text {
	case weekOverUnder:
		b.scrapingISTU(message)
	}

}

func (b Bot) selectGropu(message *tgbotapi.Message) {
	log.Printf("[%s] искал группу: %s", message.From.UserName, message.Text)

	sg, _, _, _ := searchGroup(message.Text)

	if sg == true {
		msg := tgbotapi.NewMessage(message.Chat.ID, theGroupHasBeenFound)
		b.bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, theGroupWasNotfound)
		b.bot.Send(msg)
	}
}

func (b Bot) scrapingISTU(message *tgbotapi.Message) {

	c := colly.NewCollector()

	var textValues []string

	c.OnHTML("span.site-header-top-element-text", func(e *colly.HTMLElement) {
		textValue := strings.TrimSpace(e.Text)
		textValues = append(textValues, textValue)
	})

	err := c.Visit("https://istu.ru/")
	if err != nil {
		log.Fatal(err)
	}

	for i, textValue := range textValues {
		if i == 2 {
			msg := tgbotapi.NewMessage(message.Chat.ID, textValue)
			b.bot.Send(msg)
		}
	}
}
