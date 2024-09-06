package batch

import "github.com/spf13/cobra"

func SendBatchFile(path string) {}

var SendBatchCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a batch to a given destination",
}
