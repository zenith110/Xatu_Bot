# Xatu
## Table of Contents
1. [Goals](#goals)
2. [Setup](#setup)
### Goals
    To serve the FE Book Club and the UCF Community with resources to pass and excel at the CS Foundation Exam.
    To encourage the continuation of one's learning through developing plugins for the bot
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




