package commands

import (
	"time"
	"github.com/bwmarrin/discordgo"
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
	"net/http"
)
type FE struct{
	fe_term,
	fe_exam,
	fe_solutions,
	average_score_section_I,
	average_score_section_II,
	average_score_total, 
	passing_line,
	number_of_passing,
	number_of_students,
	pass_rate,
	DS_A1,
	DS_A2,
	DS_A3,
	DS_B1,
	DS_B2,
	DS_B3,
	AA1,
	AA2,
	AA3,
	AB1,
	AB2,
	AB3,
	status string
}
func fe_value(fe FE) string{
	return ("\nAverage score for section I: " + fe.average_score_section_I + "%" + 
	"\nAverage score for section II: " + fe.average_score_section_II + "%" +
	"\nAverage score total: " + fe.average_score_total + "%" +
	"\nPassing line: " + fe.passing_line + "%" +
	"\nNumber of passing: " + fe.number_of_passing + 
	"\nNumber of students "  + fe.number_of_students + 
	"\nPass rate: " + fe.pass_rate + "%" +
	"\nData structures A question 1 average: " + fe.DS_A1 + "%" +
	"\nData structures A question 2 average: " + fe.DS_A2 + "%" +
	"\nData structures A question 3 average: " + fe.DS_A3 + "%" +
	"\n\nData structures B question 1 average: " + fe.DS_B1 + "%" +
	"\nData structures B question 2 average: " + fe.DS_B2 + "%" +
	"\nData structures B question 3 average: " + fe.DS_B3 + "%" +
	"\n\nAlgorithms A question 1 average: " + fe.AA1 + "%" +
	"\nAlgorithms A question 2 average: " + fe.AA2 + "%" +
	"\nAlgorithms A question 3 average: " + fe.AA3 + "%" +
	"\n\nAlgorithms B question 1 average: " + fe.AB1 + "%" +
	"\nAlgorithms B question 2 average: " + fe.AB2 + "%" +
	"\nAlgorithms B question 3 average: " + fe.AB3 + "%" +
	"\n\nExam: " + fe.fe_exam  +  
	"\nSolutions: " + fe.fe_solutions)
	
 }
func Json_Ripper(textMatch string, text string)string{
	removeData := strings.Replace(textMatch, text,  "", -1)
	removeData = strings.Replace(removeData, `"`, "", -1)
	removeData = strings.Replace(removeData, ",", "", -1)
	removeData = strings.Replace(removeData, ":", "", -1)
	removeData = strings.Replace(removeData, " ", "", -1)
	return removeData
}
func Finder(what_to_find string, bodyText string) string{
	data_pattern := regexp.MustCompile(fmt.Sprintf(`.*"%s": .*`, what_to_find))
	data_match := data_pattern.FindStringSubmatch(bodyText)
	data := Json_Ripper(data_match[0], what_to_find)
	return data
}
// Manipulates bytes from the json to reassemble into the struct
func Assign_Vals(byteValue string) FE{
	var fe FE
	fe.fe_term = Finder("fe_term", byteValue)
	fe.fe_exam = Finder("fe_exam", byteValue)
	fe.fe_solutions = Finder("fe_exam_solutions", byteValue)

	fe.average_score_section_I = Finder("average_score_section_I", byteValue)
	fe.average_score_section_II = Finder("average_score_section_II", byteValue)
	fe.average_score_total = Finder("average_score_total", byteValue)

	fe.passing_line = Finder("passing_line", byteValue)

	fe.number_of_passing = Finder("number_of_passing", byteValue)
	fe.number_of_students = Finder("number_of_students", byteValue)
	fe.pass_rate = Finder("pass_rate", byteValue)

	fe.DS_A1 = Finder("DS_A1", byteValue)
	fe.DS_A2 = Finder("DS_A2", byteValue)
	fe.DS_A3 = Finder("DS_A3", byteValue)
	fe.DS_B1 = Finder("DS_B1", byteValue)
	fe.DS_B2 = Finder("DS_B2", byteValue)
	fe.DS_B3 = Finder("DS_B3", byteValue)
	
	fe.AA1 = Finder("AA1", byteValue)
	fe.AA2 = Finder("AA2", byteValue)
	fe.AA3 = Finder("AA3", byteValue)
	fe.AB1 = Finder("AB1", byteValue)
	fe.AB2 = Finder("AB2", byteValue)
	fe.AB3 = Finder("AB3", byteValue)
	return fe
}

func Web_Request(name string, s *discordgo.Session, m *discordgo.MessageCreate) FE{
	fetchurl := "https://ucf-cs-fe-api.tk/exam/?name=" + name
	var f FE
	// Sends a post request to the url above
	req, err := http.Get(fetchurl)
	// Will always be NIL, ignore
	if err == nil{
	}
	// If we cannot find the FE exam, simply exit 
	if req.StatusCode == 500{
		f.status = "500"
		return f
	}else{
		bodyData, err := ioutil.ReadAll(req.Body)
		if err != nil{
		}
		bodyText := string(bodyData)
		assign_values := Assign_Vals(bodyText)
		return assign_values
	}
}
func Embed(s *discordgo.Session, m *discordgo.MessageCreate, values string, fe FE){
	fe_data := &discordgo.MessageEmbed{
		Color: 0x00ff00, // Green
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   fe.fe_term + " Data",
				Value:  values,
				Inline: true,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
	}
	s.ChannelMessageSendEmbed(m.ChannelID, fe_data)
}
func Help_Exam_Info(state *discordgo.Session, m *discordgo.MessageCreate){
	fetchurl := "https://ucf-cs-fe-api.tk/allexams/"
	// Sends a post request to the url above
	req, err := http.Get(fetchurl)
	// Will always be NIL, ignore
	if err != nil{
	}
	// Cannot process request
	if req.StatusCode == 500{
		state.ChannelMessageSend(m.ChannelID, "Seems we were not able to fetch the current exam list, please try again later")
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
	if(fe.status == "500"){
		s.ChannelMessageSend(m.ChannelID, "```Unfortunately, we do not have " + fe_exam_term + " in our system, try again later!```")
	}else{
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