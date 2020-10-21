package commands

import (
	"time"
	"github.com/bwmarrin/discordgo"
	"fmt"
	"math"
)

func Date(year, month, day int) time.Time {
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func Countdown(s *discordgo.Session, m *discordgo.MessageCreate) {
	currentTime := time.Now()
	nextFe := Date(2021, 01, 16)
	daysBetween := fmt.Sprintf("%.0f", math.Round(nextFe.Sub(currentTime).Hours()/ 24))

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0x00ff00, // Green
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Days till FE",
				Value:  "There are " + daysBetween + " days before the next FE!", 
				Inline: true,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
