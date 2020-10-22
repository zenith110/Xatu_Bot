package commands
import(
	"strings"
	"github.com/bwmarrin/discordgo"
	"time"
	// "fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)
type Stack struct {
	ProblemName      string `json:"Problem_Name"`     
	SeenOn           string `json:"Seen_On"`          
	ProblemStatement string `json:"Problem_Statement"`
	ProblemCode      string `json:"Problem_Code"`
	Solution 		 string `json:"Solution"`
	Status string 
}
type StackHelp struct {
	Stack string `json:"Stack"`
}
func Stack_Embed(stack Stack) *discordgo.MessageEmbed{
	Stack_Embed := &discordgo.MessageEmbed{
		Color: 0x00ff00, // Green
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   stack.ProblemName + " - " + stack.SeenOn,
				Value:  "Problem statement: " + stack.ProblemStatement + "\nProblem code: ```" + stack.ProblemCode + "```" + "\nSolution: " + "\n", 
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: stack.Solution,
		},

		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
	}
	return Stack_Embed
}
func Random_Stack() Stack{
	fetchurl := "https://ucf-cs-fe-api.tk/stack/?name="
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
		return stack
	}
}
func Individual_Stack(name string, s *discordgo.Session, m *discordgo.MessageCreate) Stack{
	fetchurl := "https://ucf-cs-fe-api.tk/stack/?name=" + name
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
		return stack
	}
}
func Help_Stack_Info(state *discordgo.Session, m *discordgo.MessageCreate){
	fetchurl := "https://ucf-cs-fe-api.tk/allstacks/"
	// Sends a post request to the url above
	req, err := http.Get(fetchurl)
	// Will always be NIL, ignore
	if err != nil{
	}
	// Cannot process request
	if req.StatusCode == 500{
		state.ChannelMessageSend(m.ChannelID, "Seems we were not able to fetch the current exam list, please try again later")
	}else{
		bodyData, err := ioutil.ReadAll(req.Body)
		if err == nil{
		}
		// Converts the bytes recieved into a readable form
		bodyText := string(bodyData)

		// Splits the string into a string[]
		s := strings.Split(bodyText, ",")

		// replace the first and last part of the string
		s[0] = strings.Replace(s[0], `[{"Stack":"`, "", -1)
		s[0] = strings.Replace(s[0], `"}`, "", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], `{"Stack":"`, "and ", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], `"}]`, "", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], `\n`, "", -1)
		
		// Loops through the remaining items and removes the json encoding from them
		for i := 1; i < len(s) - 1; i++{
			s[i] = strings.Replace(s[i], `{"Stack":"`, "", -1)
			s[i] = strings.Replace(s[i], `"}`, "", -1)
		}
		// Joins back into one giant string
		stack_list := strings.Join(s, ", ")
		state.ChannelMessageSend(m.ChannelID, "```Hello, these are the current stack problems available!\n" + stack_list + "```")
	}
}
func Stack_Runner(s *discordgo.Session, m *discordgo.MessageCreate){
	if len(m.Content) > 6{
		name := strings.ToUpper(string(m.Content[7])) + strings.ToLower(string(m.Content[8:]))
		solo_stack := Individual_Stack(name, s, m)
		if strings.Contains(m.Content[7:], "help"){
			Help_Stack_Info(s, m)
		}else{
			if solo_stack.Status == "500"{
				s.ChannelMessageSend(m.ChannelID, "```Unfortunately, we do not have " + name + " in our system, try again later!```")
			}else{
			Stack_Message := Stack_Embed(solo_stack)
			s.ChannelMessageSendEmbed(m.ChannelID, Stack_Message)
			}
		}
	}else{
		stack := Random_Stack()
		Stack_Message := Stack_Embed(stack)
		s.ChannelMessageSendEmbed(m.ChannelID, Stack_Message)
	}
}