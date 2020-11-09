package bot

// Code originally created by https://github.com/mgerb/discord-bot-tutorial
// Edited by Abrahan Nevarez for the FE book club
import (
	"fmt"
	"../config"
	"github.com/bwmarrin/discordgo"
	"strings"
	"../commands"
	// "reflect"
	// "errors"
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
		inital_argument := strings.ToLower(split_arguments[0])
		
		if m.Content == "!ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
		}else if inital_argument == "!countdown"{
			commands.Countdown(s, m)
		}else if strings.Contains(m.Content,"!pubsub"){
			commands.Pubsub_Fetch(s, m)
		}else if strings.Contains(m.Content,"!fe"){
			commands.FeData(s, m)
		}else if inital_argument == "!help"{
			commands.Help(s, m)
		}
		}else if strings.Contains(m.Content, "!dog"){
			commands.Doggo_Runner(s, m)
		}else if strings.Contains(m.Content, "!role"){
			commands.Role_Caller(s,m, BotID)
		}
	}
	
	
