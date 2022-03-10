package bot 
import "testing"

func TestStart(t *testing.T) {
	if Start() != "Bot is running!" {
		t.Error("Expected success message from bot")
	}
}