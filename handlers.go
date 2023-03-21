package main

import (
	"github.com/StarkBotsIndustries/telegraph/v2"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
)

func chosenResultHandler(query string, page *telegraph.Page) error {
	if addCompletionToPage(query, page) != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func messageHandler(update tgbotapi.Update) (err error) {
	var chatCompletion string
	if chatCompletion, err = generateCompletionWithHistory(
		update.Message.Text,
		update.Message.Chat.ID,
	); err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, chatCompletion)
	msg.ParseMode = "markdown"
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
	return nil
}

func updateHandler(update tgbotapi.Update) {
	switch {
	case update.Message != nil:
		messageHandler(update)
		break
	case update.InlineQuery != nil && update.InlineQuery.Query != "":
		/*		if _, ok := userIDToDebounced[update.InlineQuery.From.ID]; !ok {
					userIDToDebounced[update.InlineQuery.From.ID] = debounce.New(debouncedTiming)
				}
				userIDToDebounced[update.InlineQuery.From.ID](func() { inlineQueryHandler(update) })*/
		inlineQueryHandler(update)
		break
	case update.ChosenInlineResult != nil:
		chosenResultHandler(update.ChosenInlineResult.Query, queryToPage[update.ChosenInlineResult.Query])
		break
	}
	return
}
