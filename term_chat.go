package main

import (
	"fmt"
	"strings"

	"github.com/gautamrege/gochat/api"
)

type TermChat struct {
}

func (t *TermChat) Input() (string, error) {
	textInput, _ := stdReader.ReadString('\n')

	// convert CRLF to LF
	textInput = strings.Replace(textInput, "\n", "", -1)
	return textInput, nil
}

func (t *TermChat) Moderate(req api.ChatRequest) {
	return
}

func (t *TermChat) Render(message string) error {
	fmt.Println(message)
	return nil
}
