package commands

import (
	"github.com/bwmarrin/discordgo"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := "```Greetings, I am Xatu. I am can both predict the future, and guide the present. Below are what I offer." + 
	"\n\nUse the help command with the various commands to see correct spelling/data available!" + 
	"\n\nMisc commands:" + 
	"\n!pubsub - You crave these sandwiches, and I shall predict when they are on sale!" +
	"\n!dog to view good boys/good girls!" + 
	"\n\nFE specific commands:" + 
	"\n!fe - I see you are studying for the FE, very well. Using this command, you will be allowed to get data on whichever FE you choose! I offer stats, as well as links to the exams and their solutions for your convience." + 
	"\n!countdown - To see how much time is left till the next FE exam is, work dilligently." +
	"\n!stack to view some stack problems" +
	"\n!dsn - to view some memory management problems" +
	"\n\nhttp://www.cs.ucf.edu/registration/exm/ to register for the FE exam!" + 
	"\n\nYou may find my source code here: https://github.com/zenith110/Xatu_Bot```"

	s.ChannelMessageSend(m.ChannelID, message)		
}
