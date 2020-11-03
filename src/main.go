package main

import (
	"./bot"
	"./config"
	"fmt"
)
// Main function, loads up the bot config system
func main() {
	// Checks if the config file can be read
	err := config.ReadConfig()
	// If we encounter an error, escape
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Start the bot
	bot.Start()
	
	<-make(chan struct{})
	return
}
