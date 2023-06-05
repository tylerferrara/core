package pkg

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type Question struct {
	text string
}

type Answer struct {
	text string
}

type Session struct {
	complete bool
	question Question
	answer   Answer
}

type ContinuousSession struct {
	sessions []Session
}

func GenContinuousSession(s []Session) ContinuousSession {
	return ContinuousSession{
		sessions: s,
	}
}

func Ask(question string) Session {
	return Session{
		complete: false,
		question: Question{text: question},
		answer:   Answer{text: ""},
	}
}

func Launch(c ContinuousSession) ([]Session, []openai.ChatCompletionMessage) {
	// new chatgpt client
	sessions := c.sessions

	token := os.Getenv("API_KEY")
	client := openai.NewClient(token)
	messages := make([]openai.ChatCompletionMessage, 0)
	fmt.Println("Conversation")
	fmt.Println("---------------------")

	for _, session := range sessions {
		question := session.question.text
		fmt.Print("\n--- Question -> " + question + "\n")
		// convert CRLF to LF
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: question,
		})

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    openai.GPT3Dot5Turbo,
				Messages: messages,
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}

		content := resp.Choices[0].Message.Content
		session.answer.text = content
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: content,
		})
		fmt.Println("ANSWER:\n " + content)
	}

	return sessions, messages
}
