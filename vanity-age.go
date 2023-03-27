package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"filippo.io/age"

)

// Probably doesn't need to be any more than the number of cores you have
const THREADS = 8 

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`ERROR: No regular expression provided. Please specify a valid regular expression in the case of: regexp.MatchString("[your_regexp]", "age1*********************************************************"). age keys are 58 characters long, in all lowercase, excluding 'age1' prefix. Exceedingly notable to the author is that the charactes "b", "i", "o", and "1" after the first "1" in "age1"  are not present in any age public key.`)
		return
	}

	m, err := regexp.MatchString("[1bio]", os.Args[1])
	if err != nil {
		fmt.Println("ERROR: Unable to process this regexp") 
		fmt.Println(err.Error())
		return
	}

	if (m) {
		fmt.Println(`Query string contains one of: [1bio]`)
		return
	}

	query := "^age1(" + strings.ToLower(os.Args[1]) + ")$"

	keyChan := make(chan *age.X25519Identity)
	
	for i := 0; i < THREADS; i++ {
		go generate(query, keyChan)
	}

	key := <-keyChan

	fmt.Println("Created: ", time.Now().Format(time.RFC3339))
	fmt.Println("Public key: ", key.Recipient())
	fmt.Println(key)
}

func generate(query string, keyChan chan *age.X25519Identity) {
	for {
		k, _ := age.GenerateX25519Identity()
		m, err := regexp.MatchString(query, k.Recipient().String())	
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if (m) {
			keyChan <- k
		}
	}
}
