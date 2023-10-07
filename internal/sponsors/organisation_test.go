package sponsors_test

import (
	"testing"

	"renatoaraujo/uk-visa-sponsors/internal/sponsors"

	"github.com/stretchr/testify/require"
)

func TestOrganisations_List(t *testing.T) {
	orgs := &sponsors.Organisations{}
	orgs.AddOrUpdateVisaType("Org1", "Type1")
	orgs.AddOrUpdateVisaType("Org2", "Type2")

	result := orgs.List()
	require.Len(t, result, 2)
	require.Equal(t, "Org1", result[0].Name)
	require.Equal(t, "Org2", result[1].Name)
}

func TestOrganisations_SearchOrganisationsByName(t *testing.T) {
	orgs := &sponsors.Organisations{}
	orgs.AddOrUpdateVisaType("AwesomeOrg", "TypeA")
	orgs.AddOrUpdateVisaType("GreatOrg", "TypeB")
	orgs.AddOrUpdateVisaType("OrgAwesome", "TypeC")

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

func TestOrganisations_AddOrUpdateVisaType(t *testing.T) {
	orgs := &sponsors.Organisations{}

	orgs.AddOrUpdateVisaType("Org1", "Type1")
	require.Equal(t, "Org1", orgs.List()[0].Name)
	require.Contains(t, orgs.List()[0].VisaType, "Type1")

	orgs.AddOrUpdateVisaType("Org1", "Type2")
	require.Contains(t, orgs.List()[0].VisaType, "Type2")

	orgs.AddOrUpdateVisaType("Org2", "TypeA")
	require.Len(t, orgs.List(), 2)
	require.Contains(t, orgs.List()[1].VisaType, "TypeA")
}
