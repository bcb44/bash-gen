package cmd

import (
	"github.com/bcb44/bash-gen/cmd/gen"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bg",
		Short: "bg is a tool for generating bash scripts",
		Long:  "bg is a tool for generating bash scripts. It uses the openai api to generate bash scripts based on a prompt.",
	}
	cmd.AddCommand(gen.Cmd())
	return cmd
}
