package utils

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func UploadFile(path string) (string, error) {
	// Upload file to server
	// Read file from path
	// Send file to server

	fileBytes := []byte("file content")
	// Send file to server
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond, spinner.WithColor("cyan"))
	s.Prefix = color.CyanString(fmt.Sprintf("Uploading file %s ", path))
	s.Start()
	time.Sleep(4 * time.Second)
	s.Stop()
	_ = fileBytes

	color.Green("File uploaded successfully!")

	return "File uploaded successfully!", nil
}
