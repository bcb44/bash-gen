package main

import (
	"context"
	"fmt"
	"os"
	"os/user"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("user.Current error: %v\n", err)
		return
	}
	envPath := usr.HomeDir + "/.bg/.env"
	err = godotenv.Load(envPath)
	fmt.Println(envPath)
	fmt.Println(os.Getenv("openai_key"))
	if err != nil {
		fmt.Printf("godotenv.Load error: %v\n", err)
		return
	}

	client := openai.NewClient(os.Getenv("openai_key"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Generate a bash script to port forward to a kubernetes pod in namespace arcgis based on a parameter pod_name",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
