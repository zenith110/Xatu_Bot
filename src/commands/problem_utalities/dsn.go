package problem_utalities
import(
	"github.com/bwmarrin/discordgo"
	"time"
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
func DSN_Embed(dsn DSN, s *discordgo.Session, m *discordgo.MessageCreate){
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
	s.ChannelMessageSendEmbed(m.ChannelID, DSN_Embed)
}
func Random_DSN(s *discordgo.Session, m *discordgo.MessageCreate) DSN{
	fetchurl := "https://fetchit.dev/FE/questions/dsn/?name="
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
		DSN_Embed(dsn, s, m)
		return dsn
	}
}

//Using the criteria set by the user, fetch the required problem
func Individual_DSN(name string, s *discordgo.Session, m *discordgo.MessageCreate) DSN{
	dsn := DSN{}
	fetchurl := "https://fetchit.dev/FE/questions/dsn/?name=" + name
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
		DSN_Embed(dsn, s, m)
		return dsn
	}
}