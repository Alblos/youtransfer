package cmd

import "github.com/spf13/cobra"

func SendFile(path string) {}

var SendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a file",
}
