package sponsors

import (
	"strings"
)

type OrganisationList struct {
	Organisations []Organisation
}

type Organisation struct {
	Name     string `csv:"Organisation Name"`
	VisaType string `csv:"Route"`
}

func (o *OrganisationList) SearchOrganisationsByName(name string) []Organisation {
	var found []Organisation

	for _, org := range o.Organisations {
		if strings.Contains(strings.ToLower(org.Name), strings.ToLower(name)) {
			found = append(found, org)
		}
	}

	return found
}
