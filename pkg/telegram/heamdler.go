package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"regexp"
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
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Введите вместо английской 'C' русскую")
		b.bot.Send(msg)
	}
}

func (b Bot) selectGropu(message *tgbotapi.Message) {
	log.Printf("Selecting group for message: %s", message.Text)

	if searchGroup(message.Text) == true {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Группа найдена")
		b.bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Группа не найдена\n"+
			"Попробуйте снова")
		b.bot.Send(msg)
	}
}
