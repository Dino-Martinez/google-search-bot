package bot

import (
	"strings"
	"os"
	"github.com/rocketlaunchr/google-search"
	"github.com/bwmarrin/discordgo"
)

type SearchResult struct {
	Title string
	URL string
}

func appendHistory(author string, data string) {
	f, err := os.OpenFile("history/" + author + ".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
			panic(err)
	}
	if _, err := f.Write([]byte(data + "\n")); err != nil {
			panic(err)
	}
	if err := f.Close(); err != nil {
			panic(err)
	}
}

func searchGoogle(query string, limit int) []SearchResult {
	opts := googlesearch.SearchOptions{Limit:limit}
	googleResults, err := googlesearch.Search(nil, query, opts)
	if err != nil {
		panic(err)
	}
	numResults := len(googleResults)
	var results []SearchResult
	for i := 0; i < numResults; i++ {
		results = append(results, SearchResult{
			googleResults[i].Title,
			googleResults[i].URL})
	}

	return results
}

func buildOutput(results []SearchResult) string {
	s := ""
	for _, result := range results {
		s += result.Title + " - <" + result.URL + ">\n"
	}

	return s
}

func Search(s *discordgo.Session, m *discordgo.MessageCreate, args []string) (error, string) {
	query := strings.Join(args, " ")
	results := searchGoogle(query, 5)

	appendHistory(m.Author.ID, query)

	output := buildOutput(results)
	_, err := s.ChannelMessageSend(m.ChannelID, output)
	if err != nil {
		return err, "Failure"
	}

	return err, "Success"
}