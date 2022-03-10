package bot 
import (
	"testing"
)

func TestSearch(t *testing.T) {
	var tests = []struct {
			input    string
			expectedTitle string
			expectedURL string
	}{
			{"How many dogs are there", "How Many Dogs Are In The World? - Pawsome Advice", "https://pawsomeadvice.com/dog/how-many-dogs-are-in-the-world/"},
			{"What is the United States population", "United States Population (LIVE) - Worldometer", "https://www.worldometers.info/world-population/us-population/"},
			{"What is a good golang project idea", "What are some good project ideas for a Go beginner? : r/golang", "https://www.reddit.com/r/golang/comments/ojjgum/what_are_some_good_project_ideas_for_a_go_beginner/"},
	}


	for _, test := range tests {
		results :=searchGoogle(test.input, 1)
		if results[0].Title != test.expectedTitle {
			t.Error("Title mismatch")
		}
		if results[0].URL != test.expectedURL {
			t.Error("URL mismatch")
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
			searchGoogle("I wonder how long this takes", 5)
	}
}

func TestOutput(t *testing.T) {
	var tests = []struct {
		input []SearchResult
		output string
	}{
		{[]SearchResult{ SearchResult{ Title:"A brilliant title", URL:"https://wow.com" } }, "A brilliant title - <https://wow.com>\n"},
		{[]SearchResult{ SearchResult{ Title:"Home of Make School", URL:"https://makeschool.com" } }, "Home of Make School - <https://makeschool.com>\n"},
	}

	for _, test := range tests {
		if output := buildOutput(test.input); output != test.output {
			t.Error("Incorrectly built output")
		}
	}
}