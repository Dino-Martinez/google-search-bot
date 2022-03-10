package bot

import "github.com/bwmarrin/discordgo"

func Help(s *discordgo.Session, m *discordgo.MessageCreate, args []string) (error, string) {
	_, err := s.ChannelMessageSend(m.ChannelID, "To search google for top results, type !search <query> \nTo see your own search history, type !history")

	if err != nil {
		return err, "Failure"
	}

	return err, "Success"
}