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
func ContainerErrorHandler(s *discordgo.Session, m *discordgo.MessageCreate){
	jsonFile, err := os.Open("utils/discordwebhooks.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var webhook Webhooks
	WebhookError := json.Unmarshal(byteValue, &webhook)
	if(WebhookError != nil){
		fmt.Println("We failed, seems the hook has no values!")
	}
	EmbedData, err := json.Marshal(map[string]string{
		"content": "```go\nError reported on " + time.Now().Format(time.RFC3339) + "\n"+ string(debug.Stack()) + "```",
	})
	req, err := http.Post(webhook.ErrorHook, "application/json", bytes.NewBuffer(EmbedData))
	if(err != nil){

	}
	fmt.Println((req))
}