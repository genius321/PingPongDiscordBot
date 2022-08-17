package main

import (
	"fmt"

	"github.com/genius321/PingPongDiscordBot/bot"
	"github.com/genius321/PingPongDiscordBot/config"
)

func main() {

	fmt.Println("Reading config file...")

	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Starting bot...")

	err = bot.Start()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running! Press CTRL-C to exit.")

	<-make(chan struct{})
}
