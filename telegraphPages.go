package main

import (
	"github.com/StarkBotsIndustries/telegraph/v2"
	"log"
)

func newTelegraphPage(query string) (page *telegraph.Page, err error) {
	page, err = telegraph.CreatePage(telegraph.CreatePageOpts{
		Title: query,
		Content: []telegraph.Node{
			"Wait a bit...",
		},
		AccessToken: telegraphAccessToken,
	})
	return
}

// подменяет контент на странице Телеграф сгенерированным текстом
func addCompletionToPage(query string, page *telegraph.Page) (err error) {
	if _, err = telegraph.EditPage(telegraph.EditPageOpts{
		AccessToken: telegraphAccessToken,
		Path:        page.Path,
		Title:       page.Title,
		Content: []telegraph.Node{
			generateLongCompletion(query),
		},
	}); err != nil {
		log.Println(err.Error())
		return err
	} else {
		return nil
	}
}
