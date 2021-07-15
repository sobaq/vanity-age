package main

import (
	"fmt"
	"os"
	"time"

	"filippo.io/age"
	"github.com/danwakefield/fnmatch"
)

const THREADS = 50

func main() {
	if len(os.Args) == 1 {
		fmt.Println(
			`no query
rules:
  *     - match anything
  ?     - match any single character
  [seq] - match any character in seq
 [!seq] - match any character not in seq
age keys are 58 characters long, excluding 'age1'`)
		return
	}

	query := "age1" + os.Args[1]

	keyChan := make(chan *age.X25519Identity)
	for i := 0; i < THREADS; i++ {
		go generate(query, keyChan)
	}

	key := <-keyChan

	fmt.Println("# created:", time.Now().Format(time.RFC3339))
	fmt.Println("# public key:", key.Recipient())
	fmt.Println(key)
}

func generate(query string, keyChan chan *age.X25519Identity) {
	for {
		k, _ := age.GenerateX25519Identity()

		if fnmatch.Match(query, k.Recipient().String(), fnmatch.FNM_IGNORECASE) {
			keyChan <- k
		}
	}
}
