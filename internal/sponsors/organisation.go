package sponsors

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

type Organisations struct {
	list []Organisation
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

// TODO: decouple and improve configuration for the command
func (o *Organisation) AddDescription(role string) error {

	message := fmt.Sprintf(`Provide a concise summary (up to 1024 characters) about the company %s based in the UK, focusing on its business nature, core activities, and market reputation, especially in the context of %s. If no details are available, respond with "N/A". 
- For insights on the company's culture and its reputation as a desirable workplace for %s, include a section after the main content, formatted as: \nCulture:\n{information}\n.
- If you encounter details about their salary range for %s or if they're recognized for offering competitive salaries in this domain, add a section formatted as: \nSalaries:\n{information}. If no specific salary information is available and the company is a subsidiary, consider information about the parent company's salary reputation for %s roles.
- Provide insights on the hiring process's complexity for %s positions, whether it's challenging to secure such a role, and any preparation time typically required for interviews or technical interviews, formatted as: \nHiring Process:\n{information}.
`, o.Name, o.Name, role, role, role, role)

	client := openai.NewClient(os.Getenv("OPEN_AI_TOKEN"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Temperature: 0.0001,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		return fmt.Errorf("gpt generation error; %w", err)
	}

	o.Description = resp.Choices[0].Message.Content
	return nil
}
