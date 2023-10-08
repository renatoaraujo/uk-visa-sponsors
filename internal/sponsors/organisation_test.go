package sponsors_test

import (
	"testing"

	"renatoaraujo/uk-visa-sponsors/internal/sponsors"

	"github.com/stretchr/testify/require"
)

func TestOrganisations_SearchOrganisationsByName(t *testing.T) {
	orgs := &sponsors.OrganisationList{
		Organisations: []sponsors.Organisation{
			{
				Name:     "AwesomeOrg",
				VisaType: "TypeA",
			},
			{
				Name:     "GreatOrg",
				VisaType: "TypeB",
			},
			{
				Name:     "OrgAwesome",
				VisaType: "TypeA",
			},
		},
	}

	tests := []struct {
		name     string
		search   string
		expected int
	}{
		{name: "Search by exact name", search: "AwesomeOrg", expected: 1},
		{name: "Search by partial name", search: "Awesome", expected: 2},
		{name: "Search by non-existent name", search: "NonExistent", expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := orgs.SearchOrganisationsByName(tt.search)
			require.Len(t, result, tt.expected)
		})
	}
}
