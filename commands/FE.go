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
	AB3 string
}
func fe_value(fe FE) string{
	return ("\nAverage score for section I: " + fe.average_score_section_I + 
	"\nAverage score for section II: " + fe.average_score_section_II + 
	"\nAverage score total: " + fe.average_score_total +
	"\nPassing line: " + fe.passing_line +
	"\nNumber of passing: " + fe.number_of_passing +
	"\nNumber of students "  + fe.number_of_students + 
	"\nPass rate: " + fe.pass_rate +
	"\nData structures A question 1 average: " + fe.DS_A1 +
	"\nData structures A question 2 average: " + fe.DS_A2 +
	"\nData structures A question 3 average: " + fe.DS_A3 + 
	"\n\nData structures B question 1 average: " + fe.DS_B1 +
	"\nData structures B question 2 average: " + fe.DS_B2 +
	"\nData structures B question 3 average: " + fe.DS_B3 +
	"\n\nAlgorithms A question 1 average: " + fe.AA1 +
	"\nAlgorithms A question 2 average: " + fe.AA2 + 
	"\nAlgorithms A question 3 average: " + fe.AA3 +
	"\n\nAlgorithms B question 1 average: " + fe.AB1 +
	"\nAlgorithms B question 2 average: " + fe.AB2 +
	"\nAlgorithms B question 3 average: " + fe.AB3 +
	"\n\nExam: " + fe.fe_exam  +  
	"\nSolutions: " + fe.fe_solutions)
	
 }
func jsonRipper(textMatch string, text string)string{
	removeData := strings.Replace(textMatch, text,  "", -1)
	removeData = strings.Replace(removeData, `"`, "", -1)
	removeData = strings.Replace(removeData, ",", "", -1)
	removeData = strings.Replace(removeData, ":", "", -1)
	removeData = strings.Replace(removeData, " ", "", -1)
	return removeData
}
func finder(what_to_find string, bodyText string) string{
	data_pattern := regexp.MustCompile(fmt.Sprintf(`.*"%s": .*`, what_to_find))
	data_match := data_pattern.FindStringSubmatch(bodyText)
	fmt.Println(data_match)
	fmt.Println(data_match[0] + " is the match!")
	data := jsonRipper(data_match[0], what_to_find)
	fmt.Println(data + " the current data")
	return data
}
// Manipulates bytes from the json to reassemble into the struct
func assignVals(byteValue string) FE{
	fmt.Println("In assign vals")
	var fe FE
	fe.fe_term = finder("fe_term", byteValue)
	fe.fe_exam = finder("fe_exam", byteValue)
	fe.fe_solutions = finder("fe_exam_solutions", byteValue)
	fe.average_score_section_I = finder("average_score_section_I", byteValue)
	fe.average_score_section_II = finder("average_score_section_II", byteValue)
	fe.average_score_total = finder("average_score_total", byteValue)
	fe.passing_line = finder("passing_line", byteValue)
	fe.number_of_passing = finder("number_of_passing", byteValue)
	fe.number_of_students = finder("number_of_students", byteValue)
	fe.pass_rate = finder("pass_rate", byteValue)
	fe.DS_A1 = finder("DS_A1", byteValue)
	fe.DS_A2 = finder("DS_A2", byteValue)
	fe.DS_A3 = finder("DS_A3", byteValue)
	fe.DS_B1 = finder("DS_B1", byteValue)
	fe.DS_B2 = finder("DS_B2", byteValue)
	fe.DS_B3 = finder("DS_B3", byteValue)
	
	fe.AA1 = finder("AA1", byteValue)
	fe.AA2 = finder("AA2", byteValue)
	fe.AA3 = finder("AA3", byteValue)
	fe.AB1 = finder("AB1", byteValue)
	fe.AB2 = finder("AB2", byteValue)
	fe.AB3 = finder("AB3", byteValue)
	return fe
}
func web_request(name string, s *discordgo.Session, m *discordgo.MessageCreate) FE{
	fetchurl := "http://ucf-cs-fe-api.tk/exam/?name=" + name
	var f FE
	// Sends a post request to the url above
	req, err := http.Get(fetchurl)
	// Will always be NIL, ignore
	if err == nil{
	}
	if req.StatusCode == 500{
		s.ChannelMessageSend(m.ChannelID, "Unfortunately, we do not have " + name + " in our system, try again later!")
		return f
	}else{
		fmt.Println("hi")
		bodyData, err := ioutil.ReadAll(req.Body)
		if err != nil{
		}
		bodyText := string(bodyData)
		assign_values := assignVals(bodyText)
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

func Individual_Exams(s *discordgo.Session, m *discordgo.MessageCreate, fe_exam_term string){
	fe_exam_term = strings.Replace(fe_exam_term, " ", "-", -1)
	fe := web_request(fe_exam_term, s, m)
	values := fe_value(fe)
	Embed(s, m, values, fe)
}
func FeData(s *discordgo.Session, m *discordgo.MessageCreate) int {
	if(len(m.Content) > 3){
			fe_exam_term := strings.ToUpper(string(m.Content[4])) + strings.ToLower(m.Content[5:])
			Individual_Exams(s, m, fe_exam_term)
	}else{
		s.ChannelMessageSend(m.ChannelID, "No term provided, randomizing!")
		fe := web_request("", s, m)
		values := fe_value(fe)
		Embed(s, m, values, fe)
	}
	return 0	
}
