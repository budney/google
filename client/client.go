package client

import (
	"github.com/budney/google/oauth/token"

	"encoding/json"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

// Create() uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func Create(ctx context.Context, config *oauth2.Config, cacheFile string) *http.Client {
    tok, err := token.Get(cacheFile, config)
	if err != nil {
		log.Fatalf("Unable to get OAUTH token: %v", err)
	}
	return config.Client(ctx, tok)
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
