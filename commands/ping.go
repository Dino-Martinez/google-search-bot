package bot

import "github.com/bwmarrin/discordgo"

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) (error, string) {
	_, err := s.ChannelMessageSend(m.ChannelID, "ping")
	if err != nil {
		return err, "Failure"
	}

	return err, "Success"
}