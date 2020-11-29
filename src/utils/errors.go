package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/bwmarrin/discordgo"
)
type Webhooks struct {
	ErrorHook string `json:"ErrorHook"`

}
type XatuPicture struct{
	Url string `json:"url"`
}
type Embed struct {
	Title string `json:"title"`
	Timestamp string `json:"timestamp"`
	Description string `json:"description"`
	Color int `json:"color"`
	Thumbnail []XatuPicture 
}

type ErrorStruct struct {
Embeds                     []Embed `json:"embeds"`

}


func ContainerErrorHandler(s *discordgo.Session, m *discordgo.MessageCreate){
	jsonFile, err := os.Open("utils/discordwebhooks.json")
	if err != nil {
		fmt.Println(err)
	}
	data := ErrorStruct{
		Embeds: []Embed{
			Embed{
				Title: "Error report",
				Timestamp: time.Now().Format(time.RFC3339),
				Description: "```go\n" + string(debug.Stack()) + "```",
				Color: 9851206,
			},
		},
	}
 
	file, _ := json.MarshalIndent(data, "", " ")
 
	_ = ioutil.WriteFile("utils/errors.json", file, 0644)

	
	errorJson, errorJsonError := os.Open("utils/errors.json")
	if(errorJsonError != nil){

	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	errorValue, _ := ioutil.ReadAll(errorJson)
	var webhook Webhooks
	WebhookError := json.Unmarshal(byteValue, &webhook)
	if(WebhookError != nil){
		fmt.Println("We failed, seems the hook has no values!")
	}
	
	req, err := http.Post(webhook.ErrorHook, "application/json", bytes.NewBuffer(errorValue))
	if(err != nil){

	}
	fmt.Println((req))
}