package commands

import (
	"time"
	"github.com/bwmarrin/discordgo"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"regexp"
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
	"\n\nExam: " + fe.fe_exam + 
	"\nSolutions: " + fe.fe_solutions)
	
 }
func jsonRipper(textMatch string, text string)string{
	removeData := strings.Replace(textMatch, text,  "", -1)
	removeData = strings.Replace(removeData, `"`, "", -1)
	removeData = strings.Replace(removeData, ",", "", -1)
	removeData = strings.Replace(removeData, ":", "", -1)
	return removeData
}
func finder(what_to_find string, bodyText string) string{
	data_pattern := regexp.MustCompile(fmt.Sprintf(`.*"%s" : .*`, what_to_find))
	data_match := data_pattern.FindStringSubmatch(bodyText)
	fmt.Println(data_match[0] + " is the match!")
	data := jsonRipper(data_match[0], what_to_find)
	fmt.Println(data + " the current data")
	return data
}
// Manipulates bytes from the json to reassemble into the struct
func assignVals(byteValue []byte) FE{
	var fe FE
	fe.fe_term = finder("fe_term", string(byteValue))
	fe.fe_exam = finder("fe_exam", string(byteValue))
	fe.fe_solutions = finder("fe_exam_solutions", string(byteValue))
	fe.average_score_section_I = finder("average_score_section_I", string(byteValue))
	fe.average_score_section_II = finder("average_score_section_II", string(byteValue))
	fe.average_score_total = finder("average_score_total", string(byteValue))
	fe.passing_line = finder("passing_line", string(byteValue))
	fe.number_of_passing = finder("number_of_passing", string(byteValue))
	fe.number_of_students = finder("number_of_students", string(byteValue))
	fe.pass_rate = finder("pass_rate", string(byteValue))
	fe.DS_A1 = finder("DS_A1", string(byteValue))
	fe.DS_A2 = finder("DS_A2", string(byteValue))
	fe.DS_A3 = finder("DS_A3", string(byteValue))
	fe.DS_B1 = finder("DS_B1", string(byteValue))
	fe.DS_B2 = finder("DS_B2", string(byteValue))
	fe.DS_B3 = finder("DS_B3", string(byteValue))
	
	fe.AA1 = finder("AA1", string(byteValue))
	fe.AA2 = finder("AA2", string(byteValue))
	fe.AA3 = finder("AA3", string(byteValue))
	fe.AB1 = finder("AB1", string(byteValue))
	fe.AB2 = finder("AB2", string(byteValue))
	fe.AB3 = finder("AB3", string(byteValue))
	return fe
}
func openJsonFile(s *discordgo.Session, m *discordgo.MessageCreate, file string, fe_term string){
	json_file, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(json_file)
	fe := assignVals(byteValue)
	values := fe_value(fe)
	s.ChannelMessageSend(m.ChannelID, "As predicted, here is the " + fe_term + " exam!")
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0x00ff00, // Green
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   fe.fe_term + " Stats",
				Value:  values,
				Inline: true,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
func FeData(s *discordgo.Session, m *discordgo.MessageCreate) {
	if(len(m.Content) > 3){
	fe_exam_term := strings.ToUpper(string(m.Content[4])) + strings.ToLower(m.Content[5:])
	if _, err := os.Stat("FE/" + fe_exam_term + ".json"); err == nil {
		fmt.Println("File exist, now opening")
		openJsonFile(s, m, "FE/" + fe_exam_term + ".json", fe_exam_term) 
	  } else {
		s.ChannelMessageSend(m.ChannelID, "Sorry, but " + fe_exam_term + " is currently not availabe") 
	  }
	}else{
		fmt.Println("No term provided, randomizing!")
	}		
}
