package commands

import (
	// "time"
	"github.com/bwmarrin/discordgo"
	// "fmt"
	// "./command_utils"
	// "os"
	// "io/ioutil"
	// "strings"
	// "regexp"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := "```Greetings, I am Xatu. I am can both predict the future, and guide the present. Below are what I offer." + 
	"\nPubsubs - You crave these sandwiches, and I shall predict when they are on sale!" +
	"\nFE - I see you are studying for the FE, very well. Using this command, you will be allowed to get data on whichever FE you choose! I offer stats, as well as links to the exams and their solutions for your convience." + 
	"\nYou may find my source code here: ```"
	s.ChannelMessageSend(m.ChannelID, message)		
}
