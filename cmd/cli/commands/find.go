package commands

import (
	"fmt"
	"log"
	"strings"

	"renatoaraujo/uk-visa-sponsors/internal/sponsors"
	"renatoaraujo/uk-visa-sponsors/pkg/data"
	"renatoaraujo/uk-visa-sponsors/pkg/httputils"
	"renatoaraujo/uk-visa-sponsors/pkg/scraper"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const govUKBaseURL = "https://www.gov.uk"

var (
	companyName string
	dataSource  string
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find the company by its name and the type of Visa they are licensed to provide.",
	Long: `Easily verify if a company is a UK Visa sponsor and identify the types of Visas they offer. 

If you want to use a different list, use the --datasource with the CSV file url. Keep in mind that for now the format should be the same as the official from gov.uk website.

Note: Some companies are registered under different names with the government. For instance, you might find 'Facebook' instead of 'Meta'.
`,
	RunE: executeFindCommand,
}

func executeFindCommand(cmd *cobra.Command, args []string) error {
	handler, err := setupHandler()
	if err != nil {
		return err
	}

	orgs := handler.OrganisationList.SearchOrganisationsByName(companyName)
	printOrganisationInfo(orgs)

	fmt.Println()
	color.Yellow("Keep in mind, just because a company is licensed to sponsor a Visa doesn't mean they necessarily will for your role.")

	return nil
}

func setupHandler() (*sponsors.Handler, error) {
	p := data.NewCSVProcessor()
	govUkHttpClient, err := httputils.NewClient(govUKBaseURL, 60)
	if err != nil {
		return nil, fmt.Errorf("failed to setup HTTP client; %w", err)
	}

	s := scraper.NewWebScraper(govUkHttpClient)
	handler := sponsors.NewHandler(s, p)
	if err = handler.Load(dataSource); err != nil {
		return nil, fmt.Errorf("failed to initialise the handler; %w", err)
	}

	return &handler, nil
}

func printOrganisationInfo(orgs []sponsors.Organisation) {

	m := make(map[string][]string)
	for _, org := range orgs {
		m[org.Name] = append(m[org.Name], fmt.Sprintf("\"%s\"", org.VisaType))
	}

	for company, visaTypes := range m {
		color.Green("%s is authorized to sponsor the following visa types: %s.", company, strings.Join(visaTypes, ", "))
	}
}

func init() {
	rootCmd.AddCommand(findCmd)
	findCmd.Flags().StringVarP(&companyName, "company", "c", "", "Company name.")
	findCmd.Flags().StringVarP(&dataSource, "datasource", "d", "", "Custom datasource (for example an old list of sponsors).")
	err := findCmd.MarkFlagRequired("company")
	if err != nil {
		log.Fatalf("failed to initialise; %s", err)
	}
}
