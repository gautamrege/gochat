package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/gautamrege/gochat/api"
	"github.com/gautamrege/gochat/mocks"
	"github.com/stretchr/testify/assert"
)

func TestChat(t *testing.T) {
	req := &api.ChatRequest{
		From: &api.Handle{
			Name: "name",
			Host: "host",
			Port: 12345,
		},
		To: &api.Handle{
			Name: "name",
			Host: "host",
			Port: 12345,
		},
		Message: "WTF",
		Source:  "ws",
	}

	mockChatter := &mocks.Chatter{}

	mockChatServer := chatServer{
		chatter: mockChatter,
	}

	message, err := json.Marshal(req)
	assert.Nil(t, err)

	mockChatter.On("Render", string(message)).Return(nil)
	mockChatter.On("Moderate", *req).Return(nil)

	_, err = mockChatServer.Chat(context.TODO(), req)
	assert.Nil(t, err)

	handle, ok := USERS.Get("name")
	assert.Equal(t, ok, true)
	assert.Equal(t, handle.Name, req.From.Name)

	mockChatter.AssertExpectations(t)
}
