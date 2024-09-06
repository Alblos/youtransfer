package batch

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var BatchRootCmd = &cobra.Command{
	Use:   "batch",
	Short: "YouTransfer is a simple but elegant file transfer & sharing service",
}

func init() {
	BatchRootCmd.AddCommand(ShareBatchCmd)
	BatchRootCmd.AddCommand(UploadBatchCmd)
	BatchRootCmd.AddCommand(SendBatchCmd)
}

func Execute() {
	if err := BatchRootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
