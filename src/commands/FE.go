package commands

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"../utils"
	"github.com/bwmarrin/discordgo"
)
type FE struct {
	FeTerm                string `json:"fe_term"`                 
	FeExam                string `json:"fe_exam"`                 
	FeExamSolutions       string `json:"fe_exam_solutions"`       
	AverageScoreSectionI  string `json:"average_score_section_I"` 
	AverageScoreSectionII string `json:"average_score_section_II"`
	AverageScoreTotal     string `json:"average_score_total"`     
	PassingLine           string `json:"passing_line"`            
	NumberOfPassing       string `json:"number_of_passing"`       
	NumberOfStudents      string `json:"number_of_students"`      
	PassRate              string `json:"pass_rate"`               
	DsA1                  string `json:"DS_A1"`                   
	DsA2                  string `json:"DS_A2"`                   
	DsA3                  string `json:"DS_A3"`                   
	DsB1                  string `json:"DS_B1"`                   
	DsB2                  string `json:"DS_B2"`                   
	DsB3                  string `json:"DS_B3"`                   
	Aa1                   string `json:"AA1"`                     
	Aa2                   string `json:"AA2"`                     
	Aa3                   string `json:"AA3"`                     
	Ab1                   string `json:"AB1"`                     
	Ab2                   string `json:"AB2"`                     
	Ab3                   string `json:"AB3"` 
	status string                    
}
type FE_Command struct{
	alias []string
	description string
	input_example string
	sub_commands []string
	argugument_runner string
}
func FE_Command_Creator() *FE_Command{
	var fe FE_Command
	fe.alias = append(fe.alias, "fe", "Fe")
	fe.description = "Recieve the stats behind a specific FE exam!"
	fe.sub_commands = append(fe.sub_commands, "random", "stats")
	fe.input_example = "!fe <[" + strings.Join(fe.sub_commands, ",") + "]>"
	fe.argugument_runner = "commands.FeData(s, m)"
	return &fe
}

func fe_value(fe FE) string{
	return ("\nAverage score for section I: " + fe.AverageScoreSectionI + "%" + 
	"\nAverage score for section II: " + fe.AverageScoreSectionII + "%" +
	"\nAverage score total: " + fe.AverageScoreTotal  + "%" +
	"\nPassing line: " + fe.PassingLine + "%" +
	"\nNumber of passing: " + fe.NumberOfPassing  + 
	"\nNumber of students "  + fe.NumberOfStudents + 
	"\nPass rate: " + fe.PassRate + "%" +
	"\nData structures A question 1 average: " + fe.DsA1 + "%" +
	"\nData structures A question 2 average: " + fe.DsA2 + "%" +
	"\nData structures A question 3 average: " + fe.DsA3 + "%" +
	"\n\nData structures B question 1 average: " + fe.DsB1 + "%" +
	"\nData structures B question 2 average: " + fe.DsB2 + "%" +
	"\nData structures B question 3 average: " + fe.DsB3 + "%" +
	"\n\nAlgorithms A question 1 average: " + fe.Aa1 + "%" +
	"\nAlgorithms A question 2 average: " + fe.Aa2 + "%" +
	"\nAlgorithms A question 3 average: " + fe.Aa3 + "%" +
	"\n\nAlgorithms B question 1 average: " + fe.Ab1 + "%" +
	"\nAlgorithms B question 2 average: " + fe.Ab2 + "%" +
	"\nAlgorithms B question 3 average: " + fe.Ab3 + "%" +
	"\n\nExam: " + fe.FeExam  +  
	"\nSolutions: " + fe.FeExamSolutions) 
	
 }

func Web_Request(name string, s *discordgo.Session, m *discordgo.MessageCreate) FE{
	fetchurl := "https://fetchit.dev/FE/exam/?name=" + name
	var f FE
	// Sends a post request to the url above
	req, err := http.Get(fetchurl)
	// Will always be NIL, ignore
	if err == nil{
	}
	// If we cannot find the FE exam, simply exit 
	if req.StatusCode == 404{
		f.status = "404"
		s.ChannelMessageSend(m.ChannelID, "```Seems we were not able to fetch the " + name + " exam data, please try again later```")
		utils.ContainerErrorHandler(s, m)
		return f
	}else{
		bodyData, err := ioutil.ReadAll(req.Body)
		if err != nil{
		}

		jsonErr := json.Unmarshal(bodyData, &f)
		if(jsonErr != nil){
			
		}
		return f
	}
}
func Embed(s *discordgo.Session, m *discordgo.MessageCreate, values string, fe FE){
	fe_data := &discordgo.MessageEmbed{
		Color: 0x965146, // Green
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   fe.FeTerm + " Data",
				Value:  values,
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Services provided by: https://fetchit.dev/",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://img.pokemondb.net/sprites/sword-shield/normal/xatu.png",
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
	}
	s.ChannelMessageSendEmbed(m.ChannelID, fe_data)
}
func Help_Exam_Info(state *discordgo.Session, m *discordgo.MessageCreate){
	fetchurl := "https://fetchit.dev/FE/exams/allexams"
	// Sends a post request to the url above
	req, err := http.Get(fetchurl)
	// Will always be NIL, ignore
	if err != nil{
	}
	// Cannot process request
	if req.StatusCode == 404{
		state.ChannelMessageSend(m.ChannelID, "```Seems we were not able to fetch the current exam list, please try again later```")
		utils.ContainerErrorHandler(state, m)
	}else{
		bodyData, err := ioutil.ReadAll(req.Body)
		if err == nil{
		}
		// Converts the bytes recieved into a readable form
		bodyText := string(bodyData)

		// Splits the string into a string[]
		s := strings.Split(bodyText, ",")

		// replace the first and last part of the string
		s[0] = strings.Replace(s[0], `[{"Exam":"`, "", -1)
		s[0] = strings.Replace(s[0], `"}`, "", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], `{"Exam":"`, " and ", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], `"}]`, "", -1)
		s[len(s) - 1] = strings.Replace(s[len(s) - 1], `\n`, "", -1)
		
		// Loops through the remaining items and removes the json encoding from them
		for i := 1; i < len(s) - 1; i++{
			s[i] = strings.Replace(s[i], `{"Exam":"`, "", -1)
			s[i] = strings.Replace(s[i], `"}`, "", -1)
		}
		// Joins back into one giant string
		exam_list := strings.Join(s, ", ")
		// Sends out the exam list
		state.ChannelMessageSend(m.ChannelID, "```Hello, these are the current exams available!\n" + exam_list + "```")
	}
}

func Individual_Exams(s *discordgo.Session, m *discordgo.MessageCreate, fe_exam_term string){
	fe_exam_term = strings.Replace(fe_exam_term, " ", "-", -1)
	fe := Web_Request(fe_exam_term, s, m)
	if(fe.status == "404"){
	}else if(fe.status == "200"){
		values := fe_value(fe)
		Embed(s, m, values, fe)
	}
}
func FeData(s *discordgo.Session, m *discordgo.MessageCreate) int {
	if(len(m.Content) > 3){
			fe_exam_term := strings.ToUpper(string(m.Content[4])) + strings.ToLower(m.Content[5:])
			if strings.Contains(m.Content[4:], "help"){
				Help_Exam_Info(s, m)
			}else{
				Individual_Exams(s, m, fe_exam_term)
			}
	}else{
		s.ChannelMessageSend(m.ChannelID, "No term provided, randomizing!")
		fe := Web_Request("", s, m)
		values := fe_value(fe)
		Embed(s, m, values, fe)
	}
	return 0	
}