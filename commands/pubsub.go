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

func jsonParser(textMatch string, text string) string {
	removeData := strings.Replace(textMatch, text, "", -1)
	removeData = strings.Replace(removeData, `"`, "", -1)
	removeData = removeData[7:]
	removeData = strings.Replace(removeData, ",", "", -1)
	return removeData
}

func Pubsub(s *discordgo.Session, m *discordgo.MessageCreate)int{
			// If we have a secondary argument continue down this conditional
			if(len(m.Content) > 7) {
			// Grabs the name from the user who inputted it
			secondaryArgs := m.Content[8:]
			// Lowercases it for ease of use into the database
			secondaryArgsLower := strings.ToLower(secondaryArgs)

			// Removes all spaces in the name
			finalArg := strings.Replace(secondaryArgsLower, " ", "-", -1)
			fetchurl := "https://pubsub-api.dev/subs/?name=" + finalArg

			// Sends a post request to the url above
			req, err := http.Get(fetchurl)
			// Will always be NIL, ignore
			if err == nil{
			}
			
			// // Puts the body text bytes to be read into a variable so we can check length
			bodyData, err := ioutil.ReadAll(req.Body)
			bodyText := string(bodyData)
			// Returns an error whenever we get something that's not in the database
			if len(bodyData) < 300{
				s.ChannelMessageSend(m.ChannelID, "It seems that " + secondaryArgsLower + " is currently not in the database!\nPlease try again later")
				return 0
			}

			// // Parses the text for the subname
			subname_pattern := regexp.MustCompile(`.*"sub_name": .*`)
			subname_match := subname_pattern.FindStringSubmatch(string(bodyText))
			sub_name := jsonParser(subname_match[0], "sub_name")
			sub_name = strings.Replace(sub_name, `-`, " ", -1)

			last_sale_pattern := regexp.MustCompile(`.*"last_sale": .*`)
			last_sale_match := last_sale_pattern.FindStringSubmatch(string(bodyText))
			last_sale := jsonParser(last_sale_match[0], "last_sale")

			status_pattern := regexp.MustCompile(`.*"status": .*`)
			status_match := status_pattern.FindStringSubmatch(string(bodyText))
			status_sale := jsonParser(status_match[0], "status")

			price_pattern := regexp.MustCompile(`.*"price": .*`)
			price_match := price_pattern.FindStringSubmatch(string(bodyText))
			price := jsonParser(price_match[0], "price")

			image_pattern := regexp.MustCompile(`.*"image": .*`)
			image_match := image_pattern.FindStringSubmatch(string(bodyText))
			image := jsonParser(image_match[0], "image")
			
			notOnSaleEmbed := &discordgo.MessageEmbed{
				Color: 0x00ff00, // Green
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   sub_name,
						Value:  "Last seen on sale from " + last_sale + " for " + price,
						Inline: true,
					},
				},

				Image: &discordgo.MessageEmbedImage{
					URL: image,
				},

				Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
			}

			onSaleEmbed := &discordgo.MessageEmbed{
				Color: 0x00ff00, // Green
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   sub_name,
						Value:  "Currently on sale from " + last_sale + " for " + price,
						Inline: true,
					},
				},

				Image: &discordgo.MessageEmbedImage{
					URL: image,
				},

				Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
			}
			// Checks if currently on sale or not
			if(status_sale == " True"){
				// Sends message
				s.ChannelMessageSendEmbed(m.ChannelID, onSaleEmbed)
			}else{
				s.ChannelMessageSendEmbed(m.ChannelID, notOnSaleEmbed)
			}
	}else{
		fetchurl := "https://pubsub-api.dev/subs/?name="

			// Sends a post request to the url above
			req, err := http.Get(fetchurl)
			// Will always be NIL, ignore
			if err == nil{
			}
			
			// // Puts the body text bytes to be read into a variable so we can check length
			bodyData, err := ioutil.ReadAll(req.Body)
			bodyText := string(bodyData)

			// // Parses the text for the subname
			subname_pattern := regexp.MustCompile(`.*"sub_name": .*`)
			subname_match := subname_pattern.FindStringSubmatch(string(bodyText))
			sub_name := jsonParser(subname_match[0], "sub_name")
			sub_name = strings.Replace(sub_name, `-`, " ", -1)
			last_sale_pattern := regexp.MustCompile(`.*"last_sale": .*`)
			last_sale_match := last_sale_pattern.FindStringSubmatch(string(bodyText))
			last_sale := jsonParser(last_sale_match[0], "last_sale")

			status_pattern := regexp.MustCompile(`.*"status": .*`)
			status_match := status_pattern.FindStringSubmatch(string(bodyText))
			status_sale := jsonParser(status_match[0], "status")

			price_pattern := regexp.MustCompile(`.*"price": .*`)
			price_match := price_pattern.FindStringSubmatch(string(bodyText))
			price := jsonParser(price_match[0], "price")

			image_pattern := regexp.MustCompile(`.*"image": .*`)
			image_match := image_pattern.FindStringSubmatch(string(bodyText))
			image := jsonParser(image_match[0], "image")
			fmt.Println(status_sale)
			notOnSaleEmbed := &discordgo.MessageEmbed{
				Color: 0x00ff00, // Green
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   sub_name,
						Value:  "Last seen on sale from " + last_sale + " for " + price,
						Inline: true,
					},
				},

				Image: &discordgo.MessageEmbedImage{
					URL: image,
				},

				Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
			}

			onSaleEmbed := &discordgo.MessageEmbed{
				Color: 0x00ff00, // Green
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   sub_name,
						Value:  "Currently on sale from " + last_sale + " for " + price,
						Inline: true,
					},
				},

				Image: &discordgo.MessageEmbedImage{
					URL: image,
				},

				Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
			}
			if(status_sale == " True"){
				// Sends message
				s.ChannelMessageSendEmbed(m.ChannelID, onSaleEmbed)
			}else{
				s.ChannelMessageSendEmbed(m.ChannelID, notOnSaleEmbed)
			}
			
	}
	return 0
}