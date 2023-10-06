package sponsors

import (
	"strings"
)

type Organisations struct {
	list []Organisation
}

type Organisation struct {
	Name     string
	VisaType []string
}

func (o *Organisations) SearchOrganisationsByName(name string) []Organisation {
	var found []Organisation

	for _, org := range o.list {
		if strings.Contains(strings.ToLower(org.Name), strings.ToLower(name)) {
			found = append(found, org)
		}
	}

	return found
}

func (o *Organisations) AddOrUpdateVisaType(name string, visaType string) {
	for i, org := range o.list {
		if org.Name == name {
			o.list[i].VisaType = append(o.list[i].VisaType, visaType)
			return
		}
	}

	newOrg := Organisation{
		Name:     name,
		VisaType: []string{visaType},
	}
	o.list = append(o.list, newOrg)
}
