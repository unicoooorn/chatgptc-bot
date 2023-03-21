package main

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
	"time"
)

// формирует плашку "Ответить коротко"
func newInlineArticle(update *tgbotapi.Update) (inlineArticle tgbotapi.InlineQueryResultArticle) {

	start := time.Now()
	log.Println(update.InlineQuery.Query, " sent ", time.Now())
	inlineResponse := generateShortCompletion(update.InlineQuery.Query) // openai api ***
	inlineArticle = tgbotapi.NewInlineQueryResultArticle("0", "Ответить коротко",
		"Q: "+update.InlineQuery.Query+"\n\n"+"A: "+inlineResponse)
	inlineArticle.Description = inlineResponse

	elapsed := time.Since(start)
	log.Println(update.InlineQuery.Query, " took ", elapsed)
	return
}

// формирует плашку "Ответить развёрнуто"
func newDummyPageArticle(update *tgbotapi.Update) (pageArticle tgbotapi.InlineQueryResultArticle) {
	dummyPage, _ := newTelegraphPage(update.InlineQuery.Query)
	pageArticle = tgbotapi.NewInlineQueryResultArticle(
		"1",
		"Ответить развёрнуто",
		"",
	)
	pageArticle.InputMessageContent = tgbotapi.InputTextMessageContent{
		Text:      "[" + update.InlineQuery.Query + "](" + dummyPage.URL + ")",
		ParseMode: "Markdown",
	}
	pageArticle.Description = dummyPage.URL
	queryToPage[update.InlineQuery.Query] = dummyPage // мапит query к странице - нужно для подмены контента
	return
}

func inlineQueryHandler(update tgbotapi.Update) error {
	dummyPageArticle := newDummyPageArticle(&update)
	inlineResultArticle := newInlineArticle(&update)
	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: update.InlineQuery.ID,
		Results:       []interface{}{inlineResultArticle, dummyPageArticle},
		CacheTime:     0,
	}
	if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
		log.Println(err.Error())
		return err
	} else {
		return nil
	}
}
