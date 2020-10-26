package commands

import
(
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Role_ID_Finder(guild_roles []*discordgo.Role, role string) string{
	for i := 0; i < len(guild_roles); i++{
		if(guild_roles[i].Name == role){
			return string(guild_roles[i].ID)
		}
	}
	i := ""
	return i
}
// Handles the functionality of assigning roles 
// Includes assigning custom channels
func Role_Caller(s *discordgo.Session, m *discordgo.MessageCreate, BotID string){
	if(len(m.Content) > 5){
		nickname := m.Author.Mention()
		role := m.Content[6:]
		if(strings.Contains(role, "help")){
			s.ChannelMessageSend(m.ChannelID,  nickname + ", the current roles available are: Leetcode, and Among us")
		}else{
			// Loop through the guild struct of the server to find the role data
			guild, guild_err := s.State.Channel(m.ChannelID)
			// Assigns the roles using the GuildID
			guild_roles, err := s.GuildRoles(guild.GuildID)
			if(err != nil){
			}
			if(guild_err != nil){
			}
			role_id := Role_ID_Finder(guild_roles, role)
			// Sanity check for a role cannot be assigned, send a message
			edit := s.GuildMemberRoleAdd(guild.GuildID, m.Author.ID, string(role_id))

			if(edit != nil){
				s.ChannelMessageSend(m.ChannelID, "Sorry " + nickname + ", unable to assign " + role + ", try another one!")	
			}else{
			s.ChannelMessageSend(m.ChannelID,  nickname + ", the role " + role + " has been assigned!")
			}
		}
	}
}