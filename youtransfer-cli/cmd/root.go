package cmd

import (
	"fmt"
	"os"
	"youtransfer/cli/cmd/batch"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ytransfer",
	Short: "YouTransfer is a simple but elegant file transfer & sharing service",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(batch.BatchRootCmd)
	rootCmd.AddCommand(ShareCmd)
	rootCmd.AddCommand(SendCmd)
	rootCmd.AddCommand(UploadCmd)
	ShareCmd.PersistentFlags().IntP("duration", "d", 0, "Set the duration for the file to be available, in minutes")
	ShareCmd.PersistentFlags().StringP("password", "p", "", "Set a password for the file")
	ShareCmd.PersistentFlags().StringP("email", "e", "", "Set the email to send the download link to")
	ShareCmd.Args = cobra.PositionalArgs(
		func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("requires a path to the file")
			}
			return nil
		},
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
