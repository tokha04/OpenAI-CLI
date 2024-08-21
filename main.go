package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
	"strconv"
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	ctx := context.Background()
	client := openai.NewClient(apiKey)
	for {
		fmt.Print("\n> ")

		reader := bufio.NewReader(os.Stdin)

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("read line: %s-\n", line)
		complete(ctx, client, line)
	}
}

func makeRequest(question string) openai.ChatCompletionRequest {
	maxToken, _ := strconv.Atoi(os.Getenv("MAX_TOKEN"))
	temperature, _ := strconv.ParseFloat(os.Getenv("TEMPERATURE"), 32)

	chatMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: question,
	}
	questions := []openai.ChatCompletionMessage{chatMessage}

	return openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo,
		Messages:    questions,
		MaxTokens:   maxToken,
		Temperature: float32(temperature),
	}
}

func complete(ctx context.Context, client *openai.Client, question string) {
	request := makeRequest(question)

	resp, err := client.CreateChatCompletion(ctx, request)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Message.Content)
}
