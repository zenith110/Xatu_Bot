package commands

import (
	"time"
	"github.com/bwmarrin/discordgo"
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
	"net/http"
	"encoding/json"
	"math/rand"
)

type DogBreeds struct {
	Message map[string][]string `json:"message"`
	Status  string              `json:"status"`
}

type Dog struct{
	name,
	image,
	status string
}
// Exports the data for when we parse the json file from the server
func dog_embed(dog Dog) *discordgo.MessageEmbed{
	dog_Embed := &discordgo.MessageEmbed{
		Color: 0x00ff00, // Green
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   dog.name,
				Value:  "Boop boop, I am " + dog.name + " feel free to boop my nose!",
				Inline: true,
			},
		},

		Image: &discordgo.MessageEmbedImage{
			URL: dog.image,
		},

		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
	}
	return dog_Embed
}
// Strips of data we text matched with
func Dog_Json_Parser(textMatch string, text string, extra string) string {
	removeData := strings.Replace(textMatch, text, "", -1)
	removeData = strings.Replace(removeData, `"`, "", -1)
	removeData = strings.Replace(removeData, extra, "", -1)
	removeData = strings.Replace(removeData, "}", "", -1)
	removeData = removeData[2:]
	removeData = strings.Replace(removeData, ",", "", -1)
	removeData = strings.Replace(removeData, `\`, ``, -1)
	return removeData
}

func Dog_Filter(what_to_find string, text string) string{
	data_pattern := regexp.MustCompile(fmt.Sprintf(`.*"%s":.*`, what_to_find))
	data_match := data_pattern.FindStringSubmatch(text)
	data := Dog_Json_Parser(data_match[0], what_to_find, "status:success")
	return data
}
func Dog_Values(text string, Dog_Name string) Dog{
	var dog Dog 
	dog.name = Dog_Name
	dog.image = Dog_Filter("message", text)
	dog.status = "OK"
	return dog
}
func Dog_Web(Dog_Name string) Dog{
	Dog_Url := "https://dog.ceo/api/breed/"+ Dog_Name + "/images/random"
	var dog Dog
	// Sends a post request to the url above
	req, err := http.Get(Dog_Url)
	// Will always be NIL, ignore
	if err == nil{
	}
	// If we cannot find the dog, simply change the status code
	if req.StatusCode == 500{
		dog.status = "500"
		return dog
	}else{
		bodyData, err := ioutil.ReadAll(req.Body)
		if err != nil{
		}

		bodyText := string(bodyData)
		Dog_Data := Dog_Values(bodyText, Dog_Name)
		return Dog_Data
	}
}
// Still needs some tweaking
func Random_Dog(s *discordgo.Session, m *discordgo.MessageCreate) Dog{
	var dog DogBreeds
	Dog_Url := "https://dog.ceo/api/breeds/list/all"
	req, err := http.Get(Dog_Url)
	if err == nil{
	}
	data, err := ioutil.ReadAll(req.Body)
	json.Unmarshal(data, &dog)
	var Dog_Names []string
	for k := range dog.Message{
		Dog_Names = append(Dog_Names, k)
	}

	n := rand.Intn(len(Dog_Names))
	Dog := Dog_Web(Dog_Names[n])
	return Dog
}

func Doggo_Runner(s *discordgo.Session, m *discordgo.MessageCreate){
	if (len(m.Content) > 4){
		Dog_Name := strings.ToLower(string(m.Content[5:]))

		Doggo := Dog_Web(Dog_Name)

		Dog_Message := dog_embed(Doggo)

		if (Doggo.status == "OK"){
			s.ChannelMessageSendEmbed(m.ChannelID, Dog_Message)
		}else if (Doggo.status == "500"){

		}
	}else{
		Random_Doggo := Random_Dog(s, m)
		Random_Dog_Message := dog_embed(Random_Doggo)
		s.ChannelMessageSendEmbed(m.ChannelID, Random_Dog_Message)
	}
}