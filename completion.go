package main

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
)
import "context"

var client = openai.NewClient("")

func generateShortCompletion(query string) string {
	if resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Your answers must be very short. Your answers must be less than 15 words",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	); err != nil {
		log.Println(err.Error(), "Unable to find answer")
		return "Unable to find answer"
	} else {
		return resp.Choices[0].Message.Content
	}
}

func generateLongCompletion(query string) string {
	if resp, _ := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Provide detailed answer if possible",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	); len(resp.Choices) < 1 {
		return "Не получилось ответить на ваш вопрос"
	} else {
		return resp.Choices[0].Message.Content
	}
}

func generateCompletionWithHistory(message string, chatID int64) (string, error) {
	chatIDToMessages[chatID] = append(
		chatIDToMessages[chatID],
		openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: message,
		},
	)
	fmt.Println("request was sent")
	if resp, error := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: chatIDToMessages[chatID],
		},
	); err != nil || len(resp.Choices) < 1 {
		return "Не получилось ответить на ваш вопрос", error
	} else {
		return resp.Choices[0].Message.Content, nil
	}
}
