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
type DSN struct {
	ProblemName      string `json:"Problem_Name"`     
	SeenOn           string `json:"Seen_On"`          
	ProblemStatement string `json:"Problem_Statement"`
	ProblemCode      string `json:"Problem_Code"`
	Solution 		 string `json:"Solution"`
	Status string 
}

func DSN_Embed(dsn DSN) *discordgo.MessageEmbed{
	DSN_Embed := &discordgo.MessageEmbed{
		Color: 0x00ff00, // Green
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   dsn.ProblemName + " - " + dsn.SeenOn,
				Value:  "Problem statement: " + dsn.ProblemStatement + "\nProblem code: ```" + dsn.ProblemCode + "```" + "\nSolution: " + "\n", 
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: dsn.Solution,
		},

		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
	}
	return DSN_Embed
}
func Random_DSN() DSN{
	fetchurl := "https://fetchit.dev/FE/questions/stack/?name="
	dsn := DSN{}
	// Sends a post request to the url above
	req, err := http.Get(fetchurl)
	// Will always be NIL, ignore
	if err == nil{
	}
	// If we cannot find the FE exam, simply exit 
	if req.StatusCode == 500{
		dsn.Status = "500"
		return dsn
	}else{
		bodyData, err := ioutil.ReadAll(req.Body)
		if err != nil{
		}
		json.Unmarshal(bodyData, &dsn)
		return dsn
	}
}
func Individual_DSN(name string, s *discordgo.Session, m *discordgo.MessageCreate) DSN{
	fetchurl := "https://fetchit.dev/FE/questions/stack/?name=" + name
	dsn := DSN{}
	// Sends a post request to the url above
	req, err := http.Get(fetchurl)
	// Will always be NIL, ignore
	if err == nil{
	}
	// If we cannot find the FE exam, simply exit 
	if req.StatusCode == 500{
		dsn.Status = "500"
		return dsn
	}else{
		bodyData, err := ioutil.ReadAll(req.Body)
		if err != nil{
		}
		json.Unmarshal(bodyData, &dsn)
		return dsn
	}
}
func Help_DSN_Info(state *discordgo.Session, m *discordgo.MessageCreate){
	fetchurl := "https://fetchit.dev/FE/questions/allstacks"
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
		s[0] = strings.Replace(s[0], `[{"DSN":"`, "", -1)
		s[0] = strings.Replace(s[0], `"}`, "", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], `{"DSN":"`, "and ", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], `"}]`, "", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], `\n`, "", -1)
		
		// Loops through the remaining items and removes the json encoding from them
		for i := 1; i < len(s) - 1; i++{
			s[i] = strings.Replace(s[i], `{"DSN":"`, "", -1)
			s[i] = strings.Replace(s[i], `"}`, "", -1)
		}
		// Joins back into one giant string
		dsn_list := strings.Join(s, ", ")
		state.ChannelMessageSend(m.ChannelID, "```Hello, these are the current dns problems available!\n" + dsn_list + "```")
	}
}
func DSN_Runner(s *discordgo.Session, m *discordgo.MessageCreate){
	if len(m.Content) > 4{
		name := strings.ToUpper(string(m.Content[5])) + strings.ToLower(string(m.Content[6:]))
		solo_dsn := Individual_DSN(name, s, m)
		if strings.Contains(m.Content[5:], "help"){
			Help_DSN_Info(s, m)
		}else{
			if solo_dsn.Status == "500"{
				s.ChannelMessageSend(m.ChannelID, "```Unfortunately, we do not have " + name + " in our system, try again later!```")
			}else{
			DSN_Message := DSN_Embed(solo_dsn)
			s.ChannelMessageSendEmbed(m.ChannelID, DSN_Message)
			}
		}
	}else{
		dsn := Random_DSN()
		DSN_Message := DSN_Embed(dsn)
		s.ChannelMessageSendEmbed(m.ChannelID, DSN_Message)
	}
}