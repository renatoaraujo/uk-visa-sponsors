package cmd

import (
	"fmt"
	"log"
	"renatoaraujo/uk-visa-sponsors/internal/sponsors"
	"renatoaraujo/uk-visa-sponsors/pkg/data"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var companyName string

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find the company by it's name and the type of the Visa they are licensed to provide.",
	RunE: func(cmd *cobra.Command, args []string) error {

		p := data.NewCSVProcessor()
		f := data.NewCSVScraper("https://www.gov.uk/government/publications/register-of-licensed-sponsors-workers")

		handler, err := sponsors.NewHandler(f, p, true)
		if err != nil {
			return fmt.Errorf("failed to initialise the handler; %w", err)
		}

		orgs := handler.Organisations.SearchOrganisationsByName(companyName)
		for _, org := range orgs {
			color.Green("company %s is licensed to provide the %s Visa", org.Name, org.VisaType)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
	findCmd.Flags().StringVarP(&companyName, "company", "c", "", "Company name")
	err := findCmd.MarkFlagRequired("company")
	if err != nil {
		log.Fatalf("failed to initialise; %w", err)
	}
}
