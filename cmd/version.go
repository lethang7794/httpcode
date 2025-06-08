package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	appVersion = "dev"
	appCommit  = "unknown"
	appDate    = "unknown"
)

// SetVersionInfo sets the version information from main
func SetVersionInfo(version, commit, date string) {
	appVersion = version
	appCommit = commit
	appDate = date
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  `Display version, commit, and build date information for httpcode.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("httpcode version %s\n", appVersion)
		fmt.Printf("Commit: %s\n", appCommit)
		fmt.Printf("Built: %s\n", appDate)
		fmt.Printf("Source: https://github.com/lethang7794/httpcode\n")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
