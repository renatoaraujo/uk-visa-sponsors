package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sponsors",
	Short: "Find out if your dream company are able to sponsor your visa to work in the UK",
	Long: `You want to apply for a company but they are quite vague if they are visa sponsors?
Discover if your dream company are able to sponsor your visa to work in the UK by doing a very simple search by the company name`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from the root command!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
