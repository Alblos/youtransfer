package cmd

import (
	u "youtransfer/cli/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

func ShareFile(path string, duration int) {
	u.UploadFile(path)
	color.Yellow("URL copied to clipboard (http://localhost:8080/file/123456)")

	// Copy URL to clipboard
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	clipboard.Write(clipboard.FmtText, []byte("http://localhost:8080/file/123456"))
}

var ShareCmd = &cobra.Command{
	Use:   "share [path]",
	Short: "Share a file",
	Run: func(cmd *cobra.Command, args []string) {
		ShareFile(args[0], 0)
	},
}
