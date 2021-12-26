package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"../utils"
	"github.com/bwmarrin/discordgo"
)

type AnimalData struct {
	Animal []Animal `json:"animal"`
}

type Animal struct {
	Image string `json:"Image"`
	Status string
}

type AnimalNames struct {
	Name string `json:"name"`
	Status string
}
// Exports the data for when we parse the json file from the server
func AnimalHelp(AnimalNames string) *discordgo.MessageEmbed{
	AnimalHelpEmbed := &discordgo.MessageEmbed{
		Color: 0x965146, // Green
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Animal list",
				Value:  AnimalNames,
				Inline: true,
			},
		},

		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://img.pokemondb.net/sprites/sword-shield/normal/xatu.png",
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
	}
	return AnimalHelpEmbed
}
// Exports the data for when we parse the json file from the server
func AnimalEmbed(animal Animal, SpeciesName string) *discordgo.MessageEmbed{
	AnimalEmbed := &discordgo.MessageEmbed{
		Color: 0x965146, // Green
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   SpeciesName,
				Value:  "A wild " + SpeciesName + " has appeared!",
				Inline: true,
			},
		},

		Image: &discordgo.MessageEmbedImage{
			URL: animal.Image,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "And just like that, they're gone!",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://img.pokemondb.net/sprites/sword-shield/normal/xatu.png",
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
	}
	return AnimalEmbed
}

func AnimalFetch(AnimalName string) Animal{
	AnimalUrl := "https://fetchit.dev/species/?name="+ AnimalName
	var AnimalHolder AnimalData
	var animal Animal
	// Sends a post request to the url above
	req, err := http.Get(AnimalUrl)
	// Will always be NIL, ignore
	if err == nil{
	}
	// If we cannot find the dog, simply change the status code
	if req.StatusCode == 404{
		fmt.Println("I messed up")
		animal.Status = "500"
		return animal
	}else{
		bodyData, err := ioutil.ReadAll(req.Body)
		if err != nil{
		}
		json.Unmarshal(bodyData, &AnimalHolder)
		return AnimalHolder.Animal[0]
	}
}

func AllAnimals(FormattedString string) string{
	AnimalUrl := "https://fetchit.dev/species/allspecies/"
	var animal AnimalNames
	// Sends a post request to the url above
	req, err := http.Get(AnimalUrl)
	// Will always be NIL, ignore
	if err == nil{
	}
	// If we cannot find the dog, simply change the status code
	if req.StatusCode == 404{
		animal.Status = "404"
		return "404"
	}else{
		bodyData, err := ioutil.ReadAll(req.Body)
		if err != nil{
		}
		// Converts the bytes recieved into a readable form
		bodyText := string(bodyData)

		// Splits the string into a string[]
		s := strings.Split(bodyText, ",")
		// replace the first and last part of the string
		s[0] = strings.Replace(s[0], fmt.Sprintf(`[{"%s":"`, FormattedString), "", -1)
		s[0] = strings.Replace(s[0], `"}`, "", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], fmt.Sprintf(`{"%s":"`, FormattedString), " ", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], `"}]`, "", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], `\n`, "", -1)
		
		// Loops through the remaining items and removes the json encoding from them
		for i := 1; i < len(s) - 1; i++{
			s[i] = strings.Replace(s[i], fmt.Sprintf(`{"%s":"`, FormattedString),  "", -1)
			s[i] = strings.Replace(s[i], `"}`, "", -1)
		}
		// Joins back into one giant string
		SpeciesList := strings.Join(s, "\n")
		return SpeciesList
	}
}

func AnimalRunner(s *discordgo.Session, m *discordgo.MessageCreate){
	StringSplit := strings.Split(m.Content, " ")
	SubArgument := strings.ToLower(StringSplit[1])
	if(SubArgument == "fetch"){
		species := strings.ToLower(StringSplit[2])
		AnimalInfo := AnimalFetch(species)
		if(AnimalInfo.Status == "500"){
			utils.ContainerErrorHandler(s, m)
			s.ChannelMessageSend(m.ChannelID, "```That species is not available, please check the help sub argument!```")
		}else{
		AnimalData := AnimalEmbed(AnimalInfo, species)
		s.ChannelMessageSendEmbed(m.ChannelID, AnimalData)
		}
	}else if(SubArgument == "help"){
		AnimalNames := AllAnimals("name")
		AnimalHelpInfo := AnimalHelp(AnimalNames)
		s.ChannelMessageSendEmbed(m.ChannelID, AnimalHelpInfo)
	}
}
