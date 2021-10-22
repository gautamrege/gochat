package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gautamrege/gochat/api"
)

var Moderation chan api.ChatRequest
var AbuseWords = []string{"shit", "damn"}

func moderateChat(quit chan bool) {
	// Bufferred channel
	Moderation = make(chan api.ChatRequest, 50)

	// Set word boundaries for all abuses
	for i, word := range AbuseWords {
		AbuseWords[i] = fmt.Sprintf("\\b%s\\b", word)
	}
	abuses := strings.Join(AbuseWords, "|")

	//fmt.Printf("%+v", abuses)

	for {
		select {
		case <-quit:
			return
		case req := <-Moderation:
			flagAbuse, err := regexp.MatchString(abuses, strings.ToLower(req.Message))
			if err != nil {
				fmt.Println("Moderation failure: ", err)
				return
			}

			if flagAbuse && req.Source == "ws" {
				// TODO: Handle abuse rendering
				WS.Render(fmt.Sprintf("Abuse: %s", req.Message))
			}

		}
	}
}
