package problem_utalities
import(
	"github.com/bwmarrin/discordgo"
	"time"
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

func Stack_Embed(stack Stack, s *discordgo.Session, m *discordgo.MessageCreate){
	Stack_Embed := &discordgo.MessageEmbed{
		Color: 0x00ff00, // Green
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   stack.ProblemName + " - " + stack.SeenOn,
				Value:  "Problem statement: " + stack.ProblemStatement + "\nProblem code: ```c\n" + stack.ProblemCode + "```" + "\nSolution: " + "\n", 
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: stack.Solution,
		},

		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
	}
	s.ChannelMessageSendEmbed(m.ChannelID, Stack_Embed)
}
func Random_Stack(s *discordgo.Session, m *discordgo.MessageCreate)Stack{
	fetchurl := "https://fetchit.dev/FE/questions/stack/?name="
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
//Using the criteria set by the user, fetch the required problem
func Individual_Stack(name string, s *discordgo.Session, m *discordgo.MessageCreate) Stack{
	stack := Stack{}
	fetchurl := "https://fetchit.dev/FE/questions/stack/?name=" + name
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
