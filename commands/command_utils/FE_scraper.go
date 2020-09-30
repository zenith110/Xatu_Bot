package command_utils

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "os"
        "golang.org/x/net/context"
        "golang.org/x/oauth2"
        "golang.org/x/oauth2/google"
        "google.golang.org/api/sheets/v4"
        
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
// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
        // The file token.json stores the user's access and refresh tokens, and is
        // created automatically when the authorization flow completes for the first
        // time.
        tokFile := "token.json"
        tok, err := tokenFromFile(tokFile)
        if err != nil {
                tok = getTokenFromWeb(config)
                saveToken(tokFile, tok)
        }
        return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
        authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
        fmt.Printf("Go to the following link in your browser then type the "+
                "authorization code: \n%v\n", authURL)

        var authCode string
        if _, err := fmt.Scan(&authCode); err != nil {
                log.Fatalf("Unable to read authorization code: %v", err)
        }

        tok, err := config.Exchange(context.TODO(), authCode)
        if err != nil {
                log.Fatalf("Unable to retrieve token from web: %v", err)
        }
        return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
        f, err := os.Open(file)
        if err != nil {
                return nil, err
        }
        defer f.Close()
        tok := &oauth2.Token{}
        err = json.NewDecoder(f).Decode(tok)
        return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
        fmt.Printf("Saving credential file to: %s\n", path)
        f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
        if err != nil {
                log.Fatalf("Unable to cache oauth token: %v", err)
        }
        defer f.Close()
        json.NewEncoder(f).Encode(token)
}
func FE_write(data FE) []byte{
        json_data :=   "{\n\t" + fmt.Sprintf(`"fe_term" : "%s",`, data.fe_term) +  
         "\n\t" + fmt.Sprintf(`"fe_exam" : "%s",`, data.fe_exam) +
         "\n\t" + fmt.Sprintf(`"fe_exam_solutions" : "%s",`, data.fe_solutions) +
         "\n\t" + fmt.Sprintf(`"average_score_section_I" : "%s",`, data.average_score_section_I) +
         "\n\t" + fmt.Sprintf(`"average_score_section_II" : "%s",`, data.average_score_section_II) +
         "\n\t" + fmt.Sprintf(`"average_score_total" : "%s",`, data.average_score_total) +
         "\n\t" + fmt.Sprintf(`"passing_line" : "%s",`, data.passing_line) +
         "\n\t" + fmt.Sprintf(`"number_of_passing" : "%s",`, data.number_of_passing) +
         "\n\t" + fmt.Sprintf(`"number_of_students" : "%s",`, data.number_of_students) +
         "\n\t" + fmt.Sprintf(`"pass_rate" : "%s",`, data.pass_rate) +
         "\n\t" + fmt.Sprintf(`"DS_A1" : "%s",`, data.DS_A1) + 
         "\n\t" + fmt.Sprintf(`"DS_A2" : "%s",`, data.DS_A2) +
         "\n\t" + fmt.Sprintf(`"DS_A3" : "%s",`, data.DS_A3) +
         "\n\t" + fmt.Sprintf(`"DS_B1" : "%s",`, data.DS_B1) +
         "\n\t" + fmt.Sprintf(`"DS_B2" : "%s",`, data.DS_B2) +
         "\n\t" + fmt.Sprintf(`"DS_B3" : "%s",`, data.DS_B3) +
         "\n\t" + fmt.Sprintf(`"AA1" : "%s",`, data.AA1) +
         "\n\t" + fmt.Sprintf(`"AA2" : "%s",`, data.AA2)  +
         "\n\t" + fmt.Sprintf(`"AA3" : "%s",`, data.AA3) +
         "\n\t" + fmt.Sprintf(`"AB1" : "%s",`, data.AB1) +
         "\n\t" + fmt.Sprintf(`"AB2" : "%s",`, data.AB2) +
         "\n\t" + fmt.Sprintf(`"AB3" : "%s"`, data.AB3) +
         "\n}"
        return []byte(json_data)
}
func FE_data_grabber(value string) string{
        var data FE
        b, err := ioutil.ReadFile("api_keys/credentials.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }
        client := getClient(config)

        srv, err := sheets.New(client)
        if err != nil {
                log.Fatalf("Unable to retrieve Sheets client: %v", err)
        }

        spreadsheetId := "1ePlnBEed-PA9Ni8xWwc67MnHc8CqQMWd_ugFZafaPeo"
        fmt.Println(string(value))
        readRange := "FE!A" + string(value) + ":V"
        resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
        if err != nil {
                log.Fatalf("Unable to retrieve data from sheet: %v", err)
        }

        if len(resp.Values) == 0 {
                fmt.Println("No data found.")
        } else {
                for _, row := range resp.Values {
                                                // Export the data to struct data
                                                data.fe_term = fmt.Sprintf("%v", row[0])
                                                data.fe_exam = fmt.Sprintf("%v", row[1]) 
                                                data.fe_solutions = fmt.Sprintf("%v", row[2])
                                                data.average_score_section_I = fmt.Sprintf("%v", row[3])
                                                data.average_score_section_II = fmt.Sprintf("%v", row[4]) 
                                                data.average_score_total = fmt.Sprintf("%v", row[5])
                                                data.passing_line = fmt.Sprintf("%v", row[6])
                                                data.number_of_passing = fmt.Sprintf("%v", row[7]) 
                                                data.number_of_students = fmt.Sprintf("%v", row[8]) 
                                                data.pass_rate = fmt.Sprintf("%v", row[9])
                                                data.DS_A1 = fmt.Sprintf("%v", row[10])
                                                data.DS_A2 = fmt.Sprintf("%v", row[11]) 
                                                data.DS_A3 = fmt.Sprintf("%v", row[12]) 
                                                data.DS_B1 = fmt.Sprintf("%v", row[13]) 
                                                data.DS_B2 = fmt.Sprintf("%v", row[14]) 
                                                data.DS_B3 = fmt.Sprintf("%v", row[15]) 
                                                data.AA1 = fmt.Sprintf("%v", row[16]) 
                                                data.AA2 = fmt.Sprintf("%v", row[17]) 
                                                data.AA3 = fmt.Sprintf("%v", row[18]) 
                                                data.AB1 = fmt.Sprintf("%v", row[19]) 
                                                data.AB2 = fmt.Sprintf("%v", row[20])
                                                data.AB3 = fmt.Sprintf("%v", row[21])
                                        }
                                
                                        file := FE_write(data)
                                        _ = ioutil.WriteFile("FE/" + data.fe_term + ".json", file, 0644)
                                        return string(file)
                                }
        return "hi"
}