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
	image string

}
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
func jsonParser(textMatch string, text string) string {
	removeData := strings.Replace(textMatch, text, "", -1)
	removeData = strings.Replace(removeData, `"`, "", -1)
	removeData = removeData[6:]
	removeData = strings.Replace(removeData, ",", "", -1)
	return removeData
}
func parser(what_to_find string, text string) string{
	// // Parses the text for the subname
	data_pattern := regexp.MustCompile(fmt.Sprintf(`.*"%s": .*`, what_to_find))
	data_match := data_pattern.FindStringSubmatch(text)
	data := jsonParser(data_match[0], what_to_find)
	data = strings.Replace(data, `-`, " ", -1)
	return data
}
func dates(what_to_find, text string) string{
	// // Parses the text for the subname
	data_pattern := regexp.MustCompile(fmt.Sprintf(`.*"%s": .*`, what_to_find))
	data_match := data_pattern.FindStringSubmatch(text)
	data := jsonParser(data_match[0], what_to_find)
	return data
}
func assign_vals(text string) Pubsub{
	var sub Pubsub
	sub.sub_name = parser("sub_name", text)
	sub.last_sale = dates("last_sale", text)
	sub.status = parser("status", text)
	sub.price = parser("price", text)
	sub.image = parser("image", text)
	return sub
}
func web_hit(text string, s * discordgo.Session, m *discordgo.MessageCreate) Pubsub{
	fetchurl := "https://pubsub-api.dev/subs/?name=" + text
	// Sends a post request to the url above
	req, err := http.Get(fetchurl)
	// Will always be NIL, ignore
	if err == nil{
		s.ChannelMessage(m.ChannelID, "Sorry, " + text + " isn't currently in our database, try again soon!")
	}
		// // Puts the body text bytes to be read into a variable so we can check length
		bodyData, err := ioutil.ReadAll(req.Body)
		if err == nil{
			s.ChannelMessage(m.ChannelID, "Sorry, " + text + " isn't currently in our database, try again soon!")
		}
		bodyText := string(bodyData)
		assign_values := assign_vals(bodyText)
		return assign_values
}

func Pubsub_fetch(s *discordgo.Session, m *discordgo.MessageCreate){
			// If we have a secondary argument continue down this conditional
			if(len(m.Content) > 7) {
			// Grabs the name from the user who inputted it
			secondary_args := m.Content[8:]
			// Lowercases it for ease of use into the database
			secondary_args = strings.ToLower(secondary_args)
			
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