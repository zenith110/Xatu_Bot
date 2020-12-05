# Xatu
## Table of Contents
1. [Goals](#goals)
2. [Setup](#setup)
    1. [Config file](#config-file)
    2. [Installing dependecies](#installing-dependecies)
    3. [Running the bot itself](#running-the-bot-itself)
3. [Creating your first plugin](#creating-your-first-plugin)
	1. [Utilities](#utilities)
		1. [Errors](#errors)
### Goals
    To serve the FE Book Club and the UCF Community with resources to pass and excel at the CS Foundation Exam.
    To encourage the continuation of one's learning through developing plugins for the bot.
### Setup
#### Config file
config.json would look like this:
```
{
    "Token": "BOT_TOKEN",
    "BotPrefix": "!"
}
```
#### Installing dependecies
Utalizing the below command will install all dependencies required to run Xatu:
```go
go get .
```
#### Running the bot itself
Once that has been installed, migrate to the src directory and run the following command to get Xatu up and running:
```go
go run main.go
```
### Creating your first plugin
You will need to first fork the repo, then clone it to any directory that you wish to use.
Once cloned, we will go into the commands folder to create the plugin.
Create a new file with the .go extension, naming it FirstPlugin.go for this example.
#### The fiirst plugin
The basic structe of a plugin written for Xatu is composed of:
```go
package commands
import (
"github.com/bwmarrin/discordgo"
)
```
Every folder is a package within go. Think of it as importing files similar to python, but in a directory scope.
Let's create a function that will simply respond with pong!
```go
func PingPong(s *discordgo.Session, m *discordgo.MessageCreate){
    s.ChannelMessageSend(m.ChannelID, "pong!")
}
```
S in this case is state, allowing us to access everything involving the server, and m is for various components of the message system. ChannelID will allow us to send our command in the current channel that we are in. With that, the plugin for ping is complete.

Now, let's head over to bot.go
```go
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}
		split_arguments := strings.Split(m.Content, " ")
		InitalArgument := strings.ToLower(split_arguments[0])
		switch InitalArgument{
		case "!ping":
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
		case "!countdown":
			commands.Countdown(s, m)
		case "!pubsub":
			commands.Pubsub_Fetch(s, m)
		case "!fe":
			commands.FeData(s, m)
		case "!help":
			commands.Help(s, m)
		case "!dog":
			commands.Doggo_Runner(s, m)
		case "!role":
			commands.RoleCaller(s,m, BotID)
		case "!problem":
			commands.Problem(s,m)
		case "!animal":
			commands.AnimalRunner(s,m)
		default:
			nickname := m.Author.Mention()
			s.ChannelMessageSend(m.ChannelID, nickname + " Please use the help command to see our current commands!")
		
		}
	}
}
```
You will want to replace _, _ = s.ChannelMessageSend(m.ChannelID, "pong"), with commands.PingPong(s, m) to get the command to properly work.

And with that, congrulations! You've made your first plugin. Now head on over to pull requests, and make a pull request comparing your fork to the repo, and submit it to begin contributing!

#### Utilities
To ensure debugging within a live enviroment, we have some utilities that will assist with the process.

##### Errors
ContainerErrorHandler from the package utils will be used to logging all errors, and will assist in debugging.



