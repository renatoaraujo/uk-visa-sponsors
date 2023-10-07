package sponsors

import (
	"strings"
)

type GenerativeAI interface {
	RefineSearch([]map[string]string) ([]map[string]string, error)
	GenerateContent([]string) string
}

type Organisations struct {
	list []Organisation
	ai   *GenerativeAI
}

type Organisation struct {
	Name        string
	VisaType    []string
	Description string
}

func (o *Organisations) List() []Organisation {
	return o.list
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

func (o *Organisations) AddGenerativeAI(ai GenerativeAI) {
	o.ai = &ai
}
