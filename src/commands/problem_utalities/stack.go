package problem_utalities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Stack struct {
	ProblemName      string `json:"Problem_Name"`     
	SeenOn           string `json:"Seen_On"`          
	ProblemStatement string `json:"Problem_Statement"`
	ProblemCode      string `json:"Problem_Code"`
	Solution 		 string `json:"Solution"`
	Status string 
}

func Stack_Embed(stack Stack, s *discordgo.Session, m *discordgo.MessageCreate){
	Stack_Embed := &discordgo.MessageEmbed{
		Color: 0x00ff00, // Green
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   stack.ProblemName + " - " + stack.SeenOn,
				Value:  "Solution: ||" + stack.Solution + "||", 
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Problem statement: " + stack.ProblemStatement + "\nProblem code: \n" + stack.ProblemCode,
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
	}
	s.ChannelMessageSendEmbed(m.ChannelID, Stack_Embed)
}
func Random_Stack(s *discordgo.Session, m *discordgo.MessageCreate)Stack{
	fetchurl := "https://fetchit.dev/FE/questions/stack/?name="
	// fetchurl := "http://127.0.0.1:5000/FE/questions/stack/?name="
	stack := Stack{}
	// Sends a post request to the url above
	req, err := http.Get(fetchurl)
	// Will always be NIL, ignore
	if err == nil{
	}
	// If we cannot find the FE exam, simply exit 
	if req.StatusCode == 500{
		stack.Status = "500"
		return stack
	}else{
		bodyData, err := ioutil.ReadAll(req.Body)
		if err != nil{
		}
		json.Unmarshal(bodyData, &stack)
		Stack_Embed(stack, s, m)
	}
	return stack
}

func StackProblemHelp(state *discordgo.Session, m *discordgo.MessageCreate, Problem string, Formated_String string)int{
	fetchurl := "https://fetchit.dev/FE/questions/all" + Problem + "/"
	// Sends a post request to the url above
	req, err := http.Get(fetchurl)
	// Will always be NIL, ignore
	if err != nil{
	}
	// Cannot process request
	if req.StatusCode == 404{
		state.ChannelMessageSend(m.ChannelID, "Seems we were not able to fetch the current " + Problem + " data, please try again later")
		return 404
	}else{
		bodyData, err := ioutil.ReadAll(req.Body)
		if err == nil{
		}
		// Converts the bytes recieved into a readable form
		bodyText := string(bodyData)

		// Splits the string into a string[]
		s := strings.Split(bodyText, ",")

		// replace the first and last part of the string
		s[0] = strings.Replace(s[0], fmt.Sprintf(`[{"%s":"`, Formated_String), "", -1)
		s[0] = strings.Replace(s[0], `"}`, "", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], fmt.Sprintf(`{"%s":"`, Formated_String), "and ", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], `"}]`, "", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], `\n`, "", -1)
		
		// Loops through the remaining items and removes the json encoding from them
		for i := 1; i < len(s) - 1; i++{
			s[i] = strings.Replace(s[i], `{"%s":"`, "", -1)
			s[i] = strings.Replace(s[i], `"}`, "", -1)
		}
		// Joins back into one giant string
		problem_list := strings.Join(s, ", ")
		state.ChannelMessageSend(m.ChannelID, "```Hello, these are the current " + Problem + " problems available!\n" + problem_list + "```")
	}
	return 200
}
func StackLogic(sub_argument string, split_arguments []string, s *discordgo.Session, m *discordgo.MessageCreate){
	if(sub_argument == "help"){
		StackProblemHelp(s, m, "stacks", "Stack")
	}else if(sub_argument == "random"){
		Random_Stack(s, m)
	}else if(sub_argument == "term"){
		term_name := strings.ToLower(split_arguments[3])
		Stack := Individual_Stack(term_name, s, m)
		if(Stack.Status == "404"){
			s.ChannelMessageSend(m.ChannelID, "```Unfortunately, we do not have the item you are looking for in our system,  use the help sub argument  to find the list of available stack problems!```")
		}else if(Stack.Status == "200"){
			Stack_Embed(Stack, s, m)
		}
	}else{
		StackProblemHelp(s, m, "stacks", "Stack")
	}
}
//Using the criteria set by the user, fetch the required problem
func Individual_Stack(name string, s *discordgo.Session, m *discordgo.MessageCreate) Stack{
	stack := Stack{}
	fetchurl := "https://fetchit.dev/FE/questions/stack/?name=" + name
	// fetchurl := "http://127.0.0.1:5000/FE/questions/stack/?name=" + name
	// Sends a post request to the url above
	req, err := http.Get(fetchurl)
	// Will always be NIL, ignore
	if err == nil{
	}
	// If we cannot find the FE exam, simply exit 
	if req.StatusCode == 404{
		stack.Status = "404"
		s.ChannelMessageSend(m.ChannelID, "```Unfortunately, we do not have " + name + " in our system, use the help sub argument  to find the list of available stack problems!```")
		return stack
	}else{
		bodyData, err := ioutil.ReadAll(req.Body)
		if err != nil{
		}
		json.Unmarshal(bodyData, &stack)
		stack.Status = "200"
		return stack
	}
}
