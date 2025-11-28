package fayil

import (
	bot2 "anime_filimlar_reklama/fayil/bot"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

func Bot() {
	pref := tele.Settings{
		Token:  "8033878705:AAGBNTCWzzvSt6BZAd9NQuQw8Ejq0-FbfK8", // tokeningizni shu yerga yozing
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	b.Handle(tele.OnText, bot2.HandleText)
	b.Handle(tele.OnPhoto, bot2.HandlePhoto)
	b.Handle(tele.OnVideo, bot2.HandleVideo)
	log.Println("✅ Bot ishga tushdi!")

	b.Start()
}
