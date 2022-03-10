package bot

import (
	"io/ioutil"
	"github.com/bwmarrin/discordgo"
)

func History(s *discordgo.Session, m *discordgo.MessageCreate, args []string) (error, string) {
	contents, err := ioutil.ReadFile("history/" + m.Author.ID + ".txt")
	if err != nil {
			panic(err)
	}

	_, err = s.ChannelMessageSend(m.ChannelID, "You have previously searched: \n" + string(contents))
	if err != nil {
		return err, "Failure"
	}

	return err, "Success"
}