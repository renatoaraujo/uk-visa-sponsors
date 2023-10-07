package commands

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sponsors",
	Short: "Find out if your dream company are able to sponsor your visa to work in the UK",
	Long: `You want to apply for a company but they are quite vague if they are visa sponsors?
Discover if your dream company are able to sponsor your visa to work in the UK by doing a very simple search by the company name.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red("failed to execute; %s", err)
		os.Exit(1)
	}
}
