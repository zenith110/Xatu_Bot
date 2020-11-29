package bot

// Code originally created by https://github.com/mgerb/discord-bot-tutorial
// Edited by Abrahan Nevarez for the FE book club
import (
	"fmt"
	"strings"

	"../commands"
	"../config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}
		split_arguments := strings.Split(m.Content, " ")
		InitalArgument := strings.ToLower(split_arguments[0])
		switch InitalArgument{
		case "!ping":
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
		case "!countdown":
			commands.Countdown(s, m)
		case "!pubsub":
			commands.Pubsub_Fetch(s, m)
		case "!fe":
			commands.FeData(s, m)
		case "!help":
			commands.Help(s, m)
		case "!dog":
			commands.Doggo_Runner(s, m)
		case "!role":
			commands.RoleCaller(s,m, BotID)
		case "!problem":
			commands.Problem(s,m)
		case "!animal":
			commands.AnimalRunner(s,m)
		default:
			nickname := m.Author.Mention()
			s.ChannelMessageSend(m.ChannelID, nickname + " Please use the help command to see our current commands!")
		
		}
	}
}