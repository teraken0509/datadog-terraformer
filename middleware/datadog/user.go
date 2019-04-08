package middleware

import (
	"os"
	"text/template"

	"github.com/kterada0509/datadog-terraformer/internal"

	datadog "github.com/zorkian/go-datadog-api"
)

// GetUser ...
func (c *Credential) GetUser(id string) (*datadog.User, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}

	user, err := c.Client.GetUser(id)
	if err != nil {
		return nil, err
	}
	return &user, err
}

// PrintUserConfiguration ...
func PrintUserConfiguration(user *datadog.User) error {
	tmpl := template.Must(template.New("userTemplate").Funcs(internal.TemplateFuncs).Parse(userTemplate))

	if err := tmpl.Execute(os.Stdout, *user); err != nil {
		return err
	}
	return nil
}

const userTemplate = `
# User Configuration for {{ .Handle }}
resource "datadog_user" "user_{{ .Handle }}" {
	{{- if .Disabled }}
	disabled = {{ .Disabled }}
	{{- end }}
	email  = "{{ .Email }}"
	handle = "{{ .Handle }}"
	name   = "{{ .Name }}"
	{{- if .Role }}
	//Deprecated
	role = "{{ .Role }}"
	{{- end }}
	{{- if .IsAdmin }}
	//Deprecated
	is_admin = {{ .IsAdmin }}
	{{- end }}
}
`
