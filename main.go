package main

import (
	"gobot/bot"
)

func main() {

	bot.Start()

	<-make(chan struct{})
	return
}