package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var companyName string

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a specific company by it's name",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("yes! but not sure, logic is not ready :)")
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
