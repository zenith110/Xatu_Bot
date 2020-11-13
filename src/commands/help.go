package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)
type Commands struct {
	Commands []Command `json:"commands"`
}

type Command struct {
	CommandPrefix string   `json:"command-prefix"`        
	Name          string   `json:"name"`                  
	Description   string   `json:"description"`           
	Example       string   `json:"example"`                             
	SubCommands   []string `json:"sub-commands,omitempty"`
}

func Values_For_Help() string{
	jsonFile, err := os.Open("commands.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	var inputs Commands
	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonErr := json.Unmarshal(byteValue, &inputs)
	if(jsonErr != nil){
		fmt.Println(jsonErr)
	}
	fmt.Println(len(inputs.Commands))
	data := []string{}
	for i := 0; i < len(inputs.Commands); i++{
		data = append(data, fmt.Sprintf("%s\n%s  \n%s\n Sub-commands: %s\n", inputs.Commands[i].Name, inputs.Commands[i].Example, inputs.Commands[i].Description, inputs.Commands[i].SubCommands))
	}
	fmt.Println(data)
	message := strings.Join(data, "\n")
	message = message + "\nhttps://www.cs.ucf.edu/registration/new/ to register for the FE exam!"
	return message
}
func Help(s *discordgo.Session, m *discordgo.MessageCreate){
		jsonFile, err := os.Open("commands.json")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		var inputs Commands
		byteValue, _ := ioutil.ReadAll(jsonFile)
		jsonErr := json.Unmarshal(byteValue, &inputs)
		if(jsonErr != nil){

		}
		
		Help_Embed := &discordgo.MessageEmbed{
			Color: 0x965146, // Green
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Commands",
					Value: 	Values_For_Help(),
					Inline: true,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Find me here: https://github.com/zenith110/Xatu_Bot",
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://img.pokemondb.net/sprites/sword-shield/normal/xatu.png",
			},
	}
	s.ChannelMessageSendEmbed(m.ChannelID, Help_Embed)	
}

