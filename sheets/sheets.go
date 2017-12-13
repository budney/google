package sheets

import (
	"github.com/budney/google/client"
	"github.com/budney/google/oauth/secret"

	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

// GetService() returns a Google sheets service receiver
func GetService(secretFile string, cacheFile string) (*sheets.Service, error) {
	// Empty context, no timeout or cancellation callback
	ctx := context.Background()

	// Client secret for completing authentication
	b, err := secret.Get(secretFile)
	if err != nil {
		log.Printf("Unable to read client secret file %s: %v", secretFile, err)
		return nil, err
	}

	// Request a read/write connection
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
        log.Printf("Unable to parse client secret file %s: %v", secretFile, err)
		return nil, err
	}

	// Connect and authenticate, getting an HTTP client
	httpClient := client.Create(ctx, config, cacheFile)

	// Turn the HTTP client into a Google sheets service receiver
	srv, err := sheets.New(httpClient)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets service receiver: %v", err)
		return nil, err
	}

	// Success!
	return srv, nil
}
