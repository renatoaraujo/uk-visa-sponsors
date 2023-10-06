package cmd

import (
	"fmt"

	"renatoaraujo/uk-visa-sponsors/internal"
	"renatoaraujo/uk-visa-sponsors/pkg"

	"github.com/spf13/cobra"
)

var companyName string

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a specific company by it's name",
	Run: func(cmd *cobra.Command, args []string) {
		scraper := pkg.NewScraper("https://www.gov.uk/government/publications/register-of-licensed-sponsors-workers")
		service := internal.NewHandler(scraper)
		service.Find(companyName)
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
	findCmd.Flags().StringVarP(&companyName, "company", "c", "", "Company name")
	err := findCmd.MarkFlagRequired("company")
	if err != nil {
		panic(fmt.Errorf("failed to initialise; %w", err))
	}
}
