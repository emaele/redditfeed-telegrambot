package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"

	"github.com/emaele/redditfeed-telegrambot/bot"
	conf "github.com/emaele/redditfeed-telegrambot/config"
)

var (
	config           conf.Config
	debug            bool
	err              error
	configFilePath   string
	redditConfigPath string
)

func main() {

	setCLIParams()

	config, err = conf.ReadConfig(configFilePath)
	if err != nil {
		log.Panic(err)
	}

	tBot, err := tgbotapi.NewBotAPI(config.TelegramTokenBot)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Authorized on account %s", tBot.Self.UserName)
	}

	tBot.Debug = debug

	if rBot, err := reddit.NewBotFromAgentFile(redditConfigPath, 0); err != nil {
		fmt.Println("Failed to create reddit bot", err)
	} else {
		cfg := graw.Config{Subreddits: config.Sources}
		handler := &bot.PostingBot{RBot: rBot, TBot: tBot, Config: config}
		if _, wait, err := graw.Run(handler, rBot, cfg); err != nil {
			fmt.Println("Failed to start graw run: ", err)
		} else {
			fmt.Println("graw run failed:", wait())
		}
	}
}

func setCLIParams() {
	flag.BoolVar(&debug, "debug", false, "activate all the debug features")
	flag.StringVar(&configFilePath, "config", "./config.toml", "configuration file path")
	flag.StringVar(&redditConfigPath, "reddit", "./redditConfig.agent", "reddit configuration file path")
	flag.Parse()
}
