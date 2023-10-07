package commands

import (
	"fmt"
	"log"
	"strings"

	"renatoaraujo/uk-visa-sponsors/internal/sponsors"
	"renatoaraujo/uk-visa-sponsors/pkg/data"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var companyName string

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find the company by it's name and the type of the Visa they are licensed to provide.",
	Long: `Easily verify if a company is a UK Visa sponsor and identify the types of Visas they offer. For deeper insights, leverage information powered by OpenAI GPT. 

Note: Some companies are registered under different names with the government. For instance, you might find 'Facebook' instead of 'Meta'. 

For AI-enhanced details, ensure your search result is unique. Future updates aim to use OpenAI to better match company names in the list.
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		p := data.NewCSVProcessor()
		// TODO: enable configuration, but since this is likely to keep for a while I will keep hardcoded
		f := data.NewCSVScraper("https://www.gov.uk/government/publications/register-of-licensed-sponsors-workers")

		handler, err := sponsors.NewHandler(f, p, true)
		if err != nil {
			return fmt.Errorf("failed to initialise the handler; %w", err)
		}

		orgs := handler.Organisations.SearchOrganisationsByName(companyName)

		for _, org := range orgs {
			quotedVisaTypes := make([]string, len(org.VisaType))
			for i, v := range org.VisaType {
				quotedVisaTypes[i] = fmt.Sprintf("\"%s\"", v)
			}
			color.Green("%s is authorized to sponsor the following visa types: %s.", org.Name, strings.Join(quotedVisaTypes, ", "))

		}

		fmt.Println()
		color.Yellow("Keep in mind, just because a company is licensed to sponsor a Visa doesn't mean they necessarily will for your role.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
	findCmd.Flags().StringVarP(&companyName, "company", "c", "", "Company name.")
	err := findCmd.MarkFlagRequired("company")
	if err != nil {
		log.Fatalf("failed to initialise; %s", err)
	}
}
