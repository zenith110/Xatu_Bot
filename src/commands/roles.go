package commands

import
(
	"github.com/bwmarrin/discordgo"
	"strings"
	"fmt"
	"strconv"
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
func Role_Color(color string)int{
	// Casts to an int64
	n, err := strconv.ParseInt(color, 16, 64)  
    if err != nil {
        panic(err)
    }
    // Turns into a string
	final_val_string := strconv.FormatInt(n, 10)
	// Turns into an int
	final_val, final_err := strconv.Atoi(final_val_string)
	if(final_err != nil){

	}
    return final_val
}
// Handles the functionality of assigning roles 
// Includes assigning custom channels
func Role_Creator(split_arguments []string, s *discordgo.Session, m *discordgo.MessageCreate){
	nickname := m.Author.Mention()
	role := split_arguments[2]
	guild, guild_err := s.State.Channel(m.ChannelID)
	if(guild_err != nil){
			fmt.Print("test")
	}
	new_role, role_err := s.GuildRoleCreate(guild.GuildID)
	fmt.Println(new_role)
	if(role_err != nil){
		fmt.Println("Failed to make role!")
	}
	guild_roles, err := s.GuildRoles(guild.GuildID)
	if(err != nil){

	}
	// Gets us the specified role id
	color := split_arguments[3]
	final_color := Role_Color(color)
	role_id := Role_ID_Finder(guild_roles, "new role")
	final_role, final_err := s.GuildRoleEdit(guild.GuildID, string(role_id), role, final_color, false, 1, true)
	if(final_err != nil){

	}
	fmt.Println(final_role)
			
	s.ChannelMessageSend(m.ChannelID,  nickname + ", the role " + role + " has been created!")
}
func Role_Joiner(split_arguments []string, s *discordgo.Session, m *discordgo.MessageCreate){
	nickname := m.Author.Mention()
	role := split_arguments[2]
	// Loop through the guild struct of the server to find the role data
	guild, guild_err := s.State.Channel(m.ChannelID)

	// Assigns the roles using the GuildID
	guild_roles, err := s.GuildRoles(guild.GuildID)

	if(err != nil){
	}

	if(guild_err != nil){
	}

	// Gets us the specified role id
	role_id := Role_ID_Finder(guild_roles, role)
			
	// Sanity check for a role cannot be assigned, send a message
	edit := s.GuildMemberRoleAdd(guild.GuildID, m.Author.ID, string(role_id))

	if(edit != nil){
		s.ChannelMessageSend(m.ChannelID, "Sorry " + nickname + ", unable to assign " + role + ", try another one!")	
	}else{
		s.ChannelMessageSend(m.ChannelID,  nickname + ", the role " + role + " has been assigned!")
	}
}

func Role_Help(s *discordgo.Session, m *discordgo.MessageCreate){
	nickname := m.Author.Mention()
	guild, guild_err := s.State.Channel(m.ChannelID)
	guild_roles, err := s.GuildRoles(guild.GuildID)
	if(guild_err != nil){
		fmt.Print("test")
	}

	if(err != nil){

	}
	roles := []string{}
	for i := 0; i < len(guild_roles); i++{
		roles = append(roles, guild_roles[i].Name)
	}
	message := strings.Join(roles, ",")
	// Remove the reserved users from the list
	message = strings.Replace(message, "Admin,", "", -1)
	message = strings.Replace(message, "@everyone,", "", -1)
	message = strings.Replace(message, "doggo,", "", -1)
	message = strings.Replace(message, "Tomato,", "", -1)
	message = strings.Replace(message, "memer,", "", -1)
	message = strings.Replace(message, "deceased,", "", -1)
	message = strings.Replace(message, "Resident Plague Doctor,", "", -1)
	message = strings.Replace(message, "Cicada,", "", -1)
	message = strings.Replace(message, "organizer,", "", -1)
	message = strings.Replace(message, "spammer,", "", -1)
	message = strings.Replace(message, "Xatu,", "", -1)
	s.ChannelMessageSend(m.ChannelID,  nickname + ", the current roles available are: " + message + "\nTo get a role use role join <name-of-role>\nTo create a role, use role create <name-of-role> <color>(Keep in mind that colors are in hex!)")
}
func Role_Caller(s *discordgo.Session, m *discordgo.MessageCreate, BotID string){
		if(len(m.Content) > 5){
		split_arguments := strings.Split(m.Content, " ")
		sub_argument := strings.ToLower(split_arguments[1])
		if(sub_argument == "help"){
			Role_Help(s, m)
			// s.ChannelMessageSend(m.ChannelID,  nickname + ", the current roles available are: Leetcode, and Among us!\nTo get a role use role join <name-of-role>\nTo create a role, use role create <name-of-role> <color>(Keep in mind that colors are in hex!)")
		}else if(sub_argument == "join"){
			Role_Joiner(split_arguments, s, m)
		}else if(sub_argument == "create"){
			Role_Creator(split_arguments, s, m)
		}else{
			fmt.Println("Hi, nothing is provided!")
		}
	}else{
		nickname := m.Author.Mention()
		s.ChannelMessageSend(m.ChannelID,  nickname + ", Please input a command!")
	}
}