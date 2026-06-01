package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("reading .env is failed")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API KEY is not found")
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
		option.WithBaseURL("https://api.moonshot.cn/v1"),
	)

	prompt1 := "write a fast sort"
	prompt2 := "use python"

	res, err := client.Chat.Completions.New(
		context.Background(), openai.ChatCompletionNewParams{
			Model: "kimi-k2.6",
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(prompt1),
				openai.UserMessage(prompt2),
			},
		},
	)

	if err != nil {
		log.Fatalf("error:%v", err)
	}

	fmt.Println(res.Choices[0].Message.Content)

}
