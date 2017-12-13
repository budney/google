package secret

import (
    "io/ioutil"
    "log"
)

func Get(fileName string) ([]byte, error) {
    // Client secret for completing authentication
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("Unable to read client secret file %s: %v", fileName, err)
        return nil, err
	}

    return b, nil
}

