package token

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
)

// Get() requests a token, either from a cache file
// or, failing that, from the web.
func Get(cacheFile string, config *oauth2.Config) (*oauth2.Token, error) {
	tok, err := getFromFile(cacheFile)
	if err != nil {
		tok, err = getFromWeb(config)
		if err != nil {
			err = saveToken(cacheFile, tok)
			if err != nil {
				// The error itself was already logged.
				log.Print("Proceeding without saving token")
                err = nil
			}
		}
	}

	return tok, err
}

// getFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n> ", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Printf("Unable to read authorization code %v", err)
		return nil, err
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Printf("Unable to retrieve token from web %v", err)
		return nil, err
	}
	return tok, nil
}

// getFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func getFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) error {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Printf("Unable to cache oauth token: %v", err)
		return err
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(token)
	return err
}
