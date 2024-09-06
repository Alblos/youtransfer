package batch

import "github.com/spf13/cobra"

func ShareBatch(path string) {}

var ShareBatchCmd = &cobra.Command{
	Use:   "share",
	Short: "Share a batch with a given url",
}
