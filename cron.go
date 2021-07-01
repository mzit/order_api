package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

func main() {
	log.Println("Starting...")

	c := cron.New(cron.WithSeconds()) //WithSeconds 支持秒级定时器
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		//models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		//models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
