package hymn

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/blackflagsoftware/agenda/config"
	"github.com/xuri/excelize/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/option"
)

type (
	GoogleSheet struct {
		SheetName string
	}
)

func (g *GoogleSheet) LookupSheet(dateStr string) (HymnSheet, error) {
	thisDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Printf("Unable to parse date string: %v", err)
		return HymnSheet{}, err
	}
	shortDate := thisDate.Format("01-02-06")
	ctx := context.Background()

	b, err := os.ReadFile(path.Join(config.CredPath, "credentials.json"))
	if err != nil {
		fmt.Printf("Unable to read client secret file: %v", err)
		return HymnSheet{}, err
	}

	config, err := google.ConfigFromJSON(b, drive.DriveReadonlyScope)
	if err != nil {
		fmt.Printf("Unable to parse client secret file to config: %v", err)
		return HymnSheet{}, err
	}
	client := getClient(config)

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		fmt.Printf("Unable to retrieve Sheets client: %v", err)
		return HymnSheet{}, err
	}

	spreadsheetId := g.SheetName

	resp, err := srv.Files.Get(spreadsheetId).Download()
	if err != nil {
		fmt.Printf("Unable to download file: %v", err)
		return HymnSheet{}, err
	}
	defer resp.Body.Close()

	f, err := excelize.OpenReader(resp.Body)
	if err != nil {
		fmt.Printf("Unable to open XLSX file from memory: %v", err)
		return HymnSheet{}, err
	}

	rows, err := f.GetRows(strconv.Itoa(thisDate.Year()))
	if err != nil {
		fmt.Printf("Unable to get rows from sheet: %v", err)
		return HymnSheet{}, err
	}

	hymnSheet := HymnSheet{}
	for i, row := range rows {
		if len(row) > 0 && row[0] == shortDate {
			hymnSheet.Opening = parseTitle(rows[i+1][2], false)
			hymnSheet.Sacrament = parseTitle(rows[i+1][4], false)
			hymnSheet.Intermediate = parseTitle(rows[i+1][6], true)
			hymnSheet.Closing = parseTitle(rows[i+1][8], false)
		}
	}
	return hymnSheet, nil
}

func parseTitle(title string, returnFull bool) string {
	numSplit := strings.SplitN(title, " ", 2)
	numStr := ""
	if len(numSplit) < 2 {
		numStr = numSplit[0]
	}
	numStr = numSplit[0]
	// go thorugh numStr and determine if it is a number
	newNum := []rune{}
	for _, char := range numStr {
		if char < '0' || char > '9' {
			continue
		}
		newNum = append(newNum, char)
	}
	if len(newNum) == 0 {
		if returnFull {
			if strings.ToLower(strings.TrimSpace(title)) == "testimonies" {
				return ""
			}
			return title
		}
		return ""
	}
	return string(newNum)
}

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

// getTokenFromWeb requests a token from the web and saves it to a file.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		fmt.Printf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		fmt.Printf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// saveToken saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Printf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func getClient(cfg *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := path.Join(config.CredPath, "token.json")
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(cfg)
		saveToken(tokFile, tok)
	}
	return cfg.Client(context.Background(), tok)
}
