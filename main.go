package dotdotdot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
)

func main() {
	// so that repl won't exit after the page is closed
	go keepalive()

	bot, err := discordgo.New("Bot " + os.Getenv("token"))

	if err != nil {
		panic(err)
	}

	// register events
	bot.AddHandler(ready)
	bot.AddHandler(messageCreate)

	err = bot.Open()

	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.")
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	err := s.UpdateStatus(0, "FUCK PYTHON I AM NOW A GO CODER")
	if err != nil {
		fmt.Println("Error updating status: ", err)
	}
	fmt.Println("Logged in as user " + s.State.User.ID)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	prefix := "'"
	if strings.HasPrefix(m.Content, prefix) {
		if strings.HasPrefix(m.Content, prefix+"ping") {
			_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
			if err != nil {
				fmt.Println("Error sending Pong! message: ", err)
			}
		}
	}
}
