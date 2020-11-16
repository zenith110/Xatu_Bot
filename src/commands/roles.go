package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	// "runtime/debug"
)

func RoleIDFinder(guildroles []*discordgo.Role, role string) string{
	for i := 0; i < len(guildroles); i++{
		if(guildroles[i].Name == role){
			return string(guildroles[i].ID)
		}
	}
	i := ""
	return i
}
func RoleColor(color string)int{
	// Casts to an int64
	n, err := strconv.ParseInt(color, 16, 64)  
    if err != nil {
        panic(err)
    }
    // Turns into a string
	finalvalstring := strconv.FormatInt(n, 10)
	// Turns into an int
	finalval, finalerr := strconv.Atoi(finalvalstring)
	if(finalerr != nil){

	}
    return finalval
}
// Handles the functionality of assigning roles 
// Includes assigning custom channels
func RoleCreator(SplitArguments []string, s *discordgo.Session, m *discordgo.MessageCreate){
	nickname := m.Author.Mention()
	role := SplitArguments[2]
	guild, guilderr := s.State.Channel(m.ChannelID)
	if(guilderr != nil){
	}
	NewRole, RoleErr := s.GuildRoleCreate(guild.GuildID)
	if(RoleErr != nil){
		fmt.Print(NewRole)
	}
	guildroles, err := s.GuildRoles(guild.GuildID)
	if(err != nil){

	}
	// Gets us the specified role id
	color := SplitArguments[3]

	color = strings.Replace(color, "#", "", -1)
	FinalColor := RoleColor(color)
	RoleID := RoleIDFinder(guildroles, "new role")
	
	FinalRole, FinalErr := s.GuildRoleEdit(guild.GuildID, string(RoleID), role, FinalColor, false, 1, true)
	if(FinalErr != nil){
		fmt.Println(FinalRole)
	}
			
	s.ChannelMessageSend(m.ChannelID,  nickname + ", the role " + role + " has been created!")
}
func RoleJoiner(splitarguments []string, s *discordgo.Session, m *discordgo.MessageCreate){
	nickname := m.Author.Mention()
	role := splitarguments[2]
	// Loop through the guild struct of the server to find the role data
	guild, GuildErr := s.State.Channel(m.ChannelID)

	// Assigns the roles using the GuildID
	GuildRoles, err := s.GuildRoles(guild.GuildID)

	if(err != nil){
	}

	if(GuildErr != nil){
	}

	// Gets us the specified role id
	RoleID := RoleIDFinder(GuildRoles, role)
			
	// If user has no roles and wishes to add, add the role but check if it exist
	if(len(m.Member.Roles) == 0){
		edit := s.GuildMemberRoleAdd(guild.GuildID, m.Author.ID, string(RoleID))
		if(edit != nil){
			s.ChannelMessageSend(m.ChannelID, "Sorry " + nickname + ", unable to assign " + role + ", try another one!")	
		}else{
			s.ChannelMessageSend(m.ChannelID,  nickname + ", the role " + role + " has been assigned!")
		}
	}else if(len(m.Member.Roles) >= 1){
		// Counter for keeping track of the tag
		tag := 0
		for i := 0; i < len(m.Member.Roles); i++{
			// If we already have it, let's delete it!
			if(m.Member.Roles[i] == string(RoleID)){
				tag = 0
				break
			}else{
				tag += 1
				continue
			}
		}

		if(tag > 0 ){
			edit := s.GuildMemberRoleAdd(guild.GuildID, m.Author.ID, string(RoleID))
			if(edit != nil){
				s.ChannelMessageSend(m.ChannelID, "Sorry " + nickname + ", unable to assign " + role + ", try another one!")	
			}else{
				s.ChannelMessageSend(m.ChannelID,  nickname + ", you've been assigned the " + role + " role!")	
			}
		}else if(tag == 0){
			s.GuildMemberRoleRemove(guild.GuildID, m.Author.ID, string(RoleID))
			s.ChannelMessageSend(m.ChannelID,  nickname + ", you've been removed from " + role )
		}

	}
}

func RoleHelp(s *discordgo.Session, m *discordgo.MessageCreate){
	NickName := m.Author.Mention()
	guild, GuildErr := s.State.Channel(m.ChannelID)
	GuildRoles, err := s.GuildRoles(guild.GuildID)
	if(GuildErr != nil){
		
	}

	if(err != nil){

	}
	roles := []string{}
	for i := 0; i < len(GuildRoles); i++{
		roles = append(roles, GuildRoles[i].Name)
	}
	message := strings.Join(roles, ",")
	// Remove the reserved users from the list
	message = strings.Replace(message, "admin,", "", -1)
	message = strings.Replace(message, "@everyone,", "", -1)
	message = strings.Replace(message, "doggo,", "", -1)
	message = strings.Replace(message, "Tomato,", "", -1)
	message = strings.Replace(message, "memer,", "", -1)
	message = strings.Replace(message, "Deceased,", "", -1)
	message = strings.Replace(message, "Resident Plague Doctor,", "", -1)
	message = strings.Replace(message, "Cicada,", "", -1)
	message = strings.Replace(message, "organizer,", "", -1)
	message = strings.Replace(message, "spammer,", "", -1)
	message = strings.Replace(message, "Xatu,", "", -1)
	message = strings.Replace(message, "FE Study Bot,", "", -1)
	s.ChannelMessageSend(m.ChannelID,  NickName + ", \nthe current roles available are: \n" + message + "\nTo get a role use !role join <name-of-role>\nTo create a role, use !role create <name-of-role> <color>(Keep in mind that colors are in hex!)")
}
func RoleCaller(s *discordgo.Session, m *discordgo.MessageCreate, BotID string){
		if(len(m.Content) > 5){
		splitarguments := strings.Split(m.Content, " ")
		subargument := strings.ToLower(splitarguments[1])
		if(subargument == "help"){
			RoleHelp(s, m)
		}else if(subargument == "join"){
			RoleJoiner(splitarguments, s, m)
		}else if(subargument == "create"){
			RoleCreator(splitarguments, s, m)
		}else{
			NickName := m.Author.Mention()
			s.ChannelMessageSend(m.ChannelID,  NickName + ", Please input provide a secondary argument!")
		}
	}else{
		NickName := m.Author.Mention()
		s.ChannelMessageSend(m.ChannelID,  NickName + ", Please input a command!")
	}
}