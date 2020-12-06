package commands

import (
	"fmt"
	"strings"

	"../utils"
	"./problem_utalities"
	"github.com/bwmarrin/discordgo"
)
type DSN struct {
	ProblemName      string `json:"Problem_Name"`     
	SeenOn           string `json:"Seen_On"`          
	ProblemStatement string `json:"Problem_Statement"`
	ProblemCode      string `json:"Problem_Code"`
	Solution 		 string `json:"Solution"`
	Status string 
}
type Stack struct {
	ProblemName      string `json:"Problem_Name"`     
	SeenOn           string `json:"Seen_On"`          
	ProblemStatement string `json:"Problem_Statement"`
	ProblemCode      string `json:"Problem_Code"`
	Solution 		 string `json:"Solution"`
	Status string 
}

// Looks for secondary argument 
func Problem(s *discordgo.Session, m *discordgo.MessageCreate){
		split_arguments := strings.Split(m.Content, " ")
		material_subject := strings.ToLower(split_arguments[1])
		sub_argument := strings.ToLower(split_arguments[2])
		if(material_subject == "help"){
			s.ChannelMessageSend(m.ChannelID, "```Hello, the current problems subjects available are: dsn, and stacks!\n Use !problem <subject> random to get a random problem!\nUse !problem <subject> term <term-year> in order to get a specific ``")
		}else if(material_subject == "dsn"){
			problem_utalities.DSNLogic(sub_argument, split_arguments, s, m)
		}else if(material_subject == "stack"){
			problem_utalities.StackLogic(sub_argument, split_arguments, s, m)
		}else{
			fmt.Println("We screwed upppp")
			utils.ContainerErrorHandler(s, m)	
		}
	}
