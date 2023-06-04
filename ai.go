package main

import (
	"log"
	"net/http"
	"os"
	"time"

	chatgpt "github.com/chatgp/chatgpt-go"
)

func launch() {
	// new chatgpt client
	token := `copy-from-cookies`
	cfValue := "copy-from-cookies"

	cookies := []*http.Cookie{
		{
			Name:  os.Getenv("GOLANG_SECURE_NEXT_AUTH_SESSION_TOKEN_CHAT_GPT"),
			Value: token,
		},
		{
			Name:  os.Getenv("GOLANG_CF_CLEARANCE_CHAT_GPT"),
			Value: cfValue,
		},
	}

	cli := chatgpt.NewClient(
		chatgpt.WithDebug(true),
		chatgpt.WithTimeout(60*time.Second),
		chatgpt.WithCookies(cookies),
	)

	// chat in continuous conversation

	// first message
	message := "say hello to me"
	text, err := cli.GetChatText(message)
	if err != nil {
		log.Fatalf("get chat text failed: %v", err)
	}

	log.Printf("q: %s, a: %s\n", message, text.Content)

	// continue conversation with new message
	conversationID := text.ConversationID
	parentMessage := text.MessageID
	newMessage := "again"

	newText, err := cli.GetChatText(newMessage, conversationID, parentMessage)
	if err != nil {
		log.Fatalf("get chat text failed: %v", err)
	}

	log.Printf("q: %s, a: %s\n", newMessage, newText.Content)
}
