package gen

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var outfile string
var run bool

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen",
		Short: "Generate bash script",
		RunE:  runE,
	}
	cmd.Flags().StringVarP(&outfile, "outfile", "o", "", "The file to write the generated bash script to")
	cmd.Flags().BoolVarP(&run, "run", "r", false, "Run the generated bash script")
	return cmd
}

func runE(cmd *cobra.Command, args []string) error {
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
		fmt.Printf("Error making openai api request %v\n", err)
		return nil
	}
	if outfile != "" {
		//write resp to file outfile
		os.WriteFile(outfile, []byte(resp.Choices[0].Message.Content), 0644)
	}
	if run {
		if outfile == "" {
			os.WriteFile("temp.sh", []byte(resp.Choices[0].Message.Content), 0644)
			outfile = "temp.sh"
		}
		err := exec.Command("/bin/sh", outfile).Run()
		if err != nil {
			fmt.Printf("Error running bash script: %v\n", err)
			return nil
		}
	}
	return nil
}
