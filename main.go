package main

import (
	"github.com/StarkBotsIndustries/telegraph/v2"
	"github.com/sashabaranov/go-openai"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"log"
	"time"
)

var chatIDToMessages = make(map[int64][]openai.ChatCompletionMessage)
var bot, err = tgbotapi.NewBotAPI("")
var telegraphAccessToken = ""
var queryToPage = make(map[string]*telegraph.Page)
var userIDToDebounced = make(map[int]func(f func()))

const debouncedTiming = 3000 * time.Millisecond

func main() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	if updates, err := bot.GetUpdatesChan(u); err == nil {
		for update := range updates {
			go updateHandler(update)
		}
	} else {
		log.Fatal(err.Error())
	}
}
