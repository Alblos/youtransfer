package batch

import "github.com/spf13/cobra"

func UploadBatch(path string) {}

var UploadBatchCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a batch to a given destination",
}
