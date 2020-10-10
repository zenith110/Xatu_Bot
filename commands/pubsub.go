package commands
import(
	"strings"
	"net/http"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"regexp"
	"time"
	"fmt"
)
type Pubsub struct{
	sub_name,
	last_sale,
	status,
	price,
	image,
	status_code string
}

// Exports the data for when we parse the json file from the server
func embed(sub Pubsub, value_string string) *discordgo.MessageEmbed{
	Embed := &discordgo.MessageEmbed{
		Color: 0x00ff00, // Green
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   sub.sub_name,
				Value:  value_string,
				Inline: true,
			},
		},

		Image: &discordgo.MessageEmbedImage{
			URL: sub.image,
		},

		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
	}
	return Embed
}
// Strips of data we text matched with
func jsonParser(textMatch string, text string) string {
	removeData := strings.Replace(textMatch, text, "", -1)
	removeData = strings.Replace(removeData, `"`, "", -1)
	removeData = removeData[6:]
	removeData = strings.Replace(removeData, ",", "", -1)
	return removeData
}
// Parses the data using regex to find the specific instance of that string
func parser(what_to_find string, text string) string{
	// // Parses the text for the subname
	data_pattern := regexp.MustCompile(fmt.Sprintf(`.*"%s": .*`, what_to_find))
	data_match := data_pattern.FindStringSubmatch(text)
	data := jsonParser(data_match[0], what_to_find)
	data = strings.Replace(data, `-`, " ", -1)
	return data
}

func assign_vals(text string) Pubsub{
	var sub Pubsub
	sub.sub_name = parser("sub_name", text)
	sub.last_sale = parser("last_sale", text)
	sub.status = parser("status", text)
	sub.price = parser("price", text)
	sub.image = parser("image", text)
	return sub
}
func web_hit(text string, s *discordgo.Session, m *discordgo.MessageCreate) Pubsub{
	var pub Pubsub
	fetchurl := "https://pubsub-api.dev/subs/?name=" + text
	// Sends a post request to the url above
	req, err := http.Get(fetchurl)
	// Will always be NIL, ignore
	if err != nil{
	}
	if req.StatusCode == 500{
		s.ChannelMessageSend(m.ChannelID, "Unfortunately, we do not have " + text + " in our system, try again later!")
		return pub
	}else{
		bodyData, err := ioutil.ReadAll(req.Body)
		if err == nil{
		}
		bodyText := string(bodyData)
		assign_values := assign_vals(bodyText)
		return assign_values
	}
}

func Help_info(state *discordgo.Session, m *discordgo.MessageCreate){
	fetchurl := "https://pubsub-api.dev/allsubs/"
	// Sends a post request to the url above
	req, err := http.Get(fetchurl)
	// Will always be NIL, ignore
	if err != nil{
	}
	// Cannot process request
	if req.StatusCode == 500{
		state.ChannelMessageSend(m.ChannelID, "Seems we were not able to fetch the current pubsub list, please try again later")
	}else{
		bodyData, err := ioutil.ReadAll(req.Body)
		if err == nil{
		}
		// Converts the bytes recieved into a readable form
		bodyText := string(bodyData)

		// Splits the string into a string[]
		s := strings.Split(bodyText, ",")

		// replace the first and last part of the string
		s[0] = strings.Replace(s[0], `[{"name":"`, "", -1)
		s[0] = strings.Replace(s[0], `"}`, "", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], `{"name":"`, " and ", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], `"}]`, "", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], `\n`, "", -1)
		
		// Loops through the remaining items and removes the json encoding from them
		for i := 1; i < len(s) - 1; i++{
			s[i] = strings.Replace(s[i], `{"name":"`, "", -1)
			s[i] = strings.Replace(s[i], `"}`, "", -1)
		}
		// Joins back into one giant string
		sub_list := strings.Join(s, ", ")
		// Sends out the pub list
		state.ChannelMessageSend(m.ChannelID, "Hello, these are the current subs available!\n" + sub_list)
	}
}

func Individual_Subs(s *discordgo.Session, m *discordgo.MessageCreate, secondary_args string){
	// Removes all spaces in the name
	if !strings.Contains(secondary_args ," "){
		fmt.Println("Let's not modify it!")
	}else{
		fmt.Println("Found a space, let's strip it!")
		secondary_args = strings.Replace(secondary_args, " ", "-", -1)
	
	}
	sub := web_hit(secondary_args, s, m)

	onSaleEmbed := embed(sub, "Currently on sale from " + sub.last_sale + " for " + sub.price)
	notOnSaleEmbed := embed(sub, "Last seen on sale from " + sub.last_sale + " for " + sub.price)
	// Checks if currently on sale or not
	if(sub.status == " True"){
		// Sends message
		s.ChannelMessageSendEmbed(m.ChannelID, onSaleEmbed)
		}else{
			s.ChannelMessageSendEmbed(m.ChannelID, notOnSaleEmbed)
		}
}

func Pubsub_fetch(s *discordgo.Session, m *discordgo.MessageCreate){
			// If we have a secondary argument continue down this conditional
			if(len(m.Content) > 7) {
			// Grabs the name from the user who inputted it
			secondary_args := m.Content[8:]
			// Lowercases it for ease of use into the database
			secondary_args = strings.ToLower(secondary_args)
			// If help is found, run the help sub command, otherwise run the pubsub check
			if strings.Contains(secondary_args, "help"){
				Help_info(s, m)
			}else{
				Individual_Subs(s, m, secondary_args)
			}

	// Randomize the pubsub if we have just !pubsub
	}else if(len(m.Content) <= 7){
			sub := web_hit("", s, m)
			onSaleEmbed := embed(sub, "Currently on sale from " + sub.last_sale + " for " + sub.price)
			notOnSaleEmbed := embed(sub, "Last seen on sale from " + sub.last_sale + " for " + sub.price)

			if(sub.status == "True"){
				// Sends message
				s.ChannelMessageSendEmbed(m.ChannelID, onSaleEmbed)
			}else{
				s.ChannelMessageSendEmbed(m.ChannelID, notOnSaleEmbed)
			}
			
	}
}