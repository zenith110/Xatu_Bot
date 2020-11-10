package commands
import(
	"strings"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"fmt"
	"./problem_utalities"
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

func Problem_Help(state *discordgo.Session, m *discordgo.MessageCreate, Problem string, Formated_String string){
	fetchurl := "https://fetchit.dev/FE/questions/all" + Problem + "/"
	fmt.Println(fetchurl)
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
}
// Looks for secondary argument 
func Problem(s *discordgo.Session, m *discordgo.MessageCreate){
		split_arguments := strings.Split(m.Content, " ")
		material_subject := strings.ToLower(split_arguments[1])
		sub_argument := strings.ToLower(split_arguments[2])
		if(material_subject == "dsn"){
			if(sub_argument == "help"){
				Problem_Help(s, m, "dsn", "DSN")
			}else if(sub_argument == "random"){
				problem_utalities.Random_DSN(s, m)
			}
		}else if(material_subject == "stack"){
			if(sub_argument == "help"){
				Problem_Help(s, m, "stacks", "Stack")
			}else if(sub_argument == "random"){
				problem_utalities.Random_Stack(s, m)
			}
		}
	}