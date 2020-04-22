package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mmcdole/gofeed"

	"log"
	"time"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.TelegramApiToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = config.TelegramApiDebug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	var oldFeed *gofeed.Feed
	var newFeed *gofeed.Feed
	var chatIds = make(map[int64]struct{})
	ticker := time.NewTicker(config.RssUpdateTimeSec * time.Second)
	for {
		select {
		case <-ticker.C:
			newFeed = parseFeed()
			checkFeed(oldFeed, newFeed, chatIds, bot)
			oldFeed = newFeed
		case update := <-updates:
			if update.Message == nil { // ignore any non-Message Updates
				return
			}
			//if int64(update.Message.From.ID) != update.Message.Chat.ID { // allow only private conversations
			//	continue
			//}
			//if !update.Message.IsCommand() { // ignore any non-command Messages
			//	continue
			//}
			chatIds[update.Message.Chat.ID] = struct{}{}
		}
	}
}
