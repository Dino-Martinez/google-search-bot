package bot

import (
	"fmt"
	"os"
	"strings"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	
	"gobot/commands"
)

var BotId string
var goBot *discordgo.Session

func Start() string {
	err := godotenv.Load(".env")

	if err != nil {
			panic("Error loading .env file")
	}

	goBot, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))

	if err != nil {
		panic(err)
	}

	u, err := goBot.User("@me")

	if err != nil {
		panic(err)
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		panic(err)
	}
	fmt.Println("Bot is running!")
	return "Bot is running!"
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}
	if m.Content[0:1] != "!" {
		return
	}

	content := strings.Fields(m.Content[1:])
	command := content[0]
	args := content[1:]

	var err error 
	var result string

	if command == "ping" {
		err, result = bot.Ping(s, m)
	}
	if command == "help" {
		err, result = bot.Help(s, m, args)
	}
	if command == "search" {
		err, result = bot.Search(s, m, args)
	}
	if command == "history" {
		err, result = bot.History(s, m, args)
	}

	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
