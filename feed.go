package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mmcdole/gofeed"
	"log"
	"time"
)

var fp = gofeed.NewParser()

// Parse and return rss feed. Use http://lorem-rss.herokuapp.com/feed for debug.
func parseFeed() *gofeed.Feed {
	feed, err := fp.ParseURL("https://itfy.org/forums/python-help/index.rss")
	if err != nil {
		log.Println("ParseURL err:", err)
	}
	return feed
}

// Check for new posts and send it to chatIds.
func checkFeed(oldFeed, newFeed *gofeed.Feed, ids map[int64]struct{}, bot *tgbotapi.BotAPI) {
	if oldFeed == nil || oldFeed == newFeed {
		return
	}
	for _, newI := range newFeed.Items {
		if find := func(item *gofeed.Item) bool {
			for _, oldI := range oldFeed.Items {
				if newI.Link == oldI.Link {
					return true
				}
			}
			return false
		}(newI); !find {
			for id := range ids {
				msg := tgbotapi.NewMessage(id, fmt.Sprintf(`<a href="%s">%s</a>`, newI.Link, newI.Title))
				msg.ParseMode = "HTML"
				if _, err := bot.Send(msg); err != nil {
					log.Println("bot.Send err:", err)
				}
				time.Sleep(config.SendTimeoutMsec * time.Millisecond)
			}
		}
	}
}
