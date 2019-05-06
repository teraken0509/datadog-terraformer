package middleware

import (
	"io"
	"text/template"

	"github.com/kterada0509/datadog-terraformer/internal"
	datadog "github.com/zorkian/go-datadog-api"
)

// GetDowntime ...
func (c *Credential) GetDowntime(id int) (*datadog.Downtime, error) {
	downtime, err := c.Client.GetDowntime(id)
	if err != nil {
		return nil, err
	}
	return downtime, err
}

// PrintDowntimeConfiguration ...
func PrintDowntimeConfiguration(w io.Writer, downtime *datadog.Downtime) error {
	tmpl := template.Must(template.New("downtimeTemplate").Funcs(internal.TemplateFuncs).Parse(downtimeTemplate))

	if err := tmpl.Execute(w, *downtime); err != nil {
		return err
	}
	return nil
}

const downtimeTemplate = `
resource "datadog_downtime" "downtime_{{ .Id }}" {
	{{- if .Active }}
	active = {{ .Active }}
	{{- end }}
	{{- if .Disabled }}
	disabled = {{ .Disabled }}
	{{- end }}
	{{- if .Scope }}
	scope = [
	{{- range .Scope }}
		"{{ . }}",
	{{- end }}
	]
	{{- end }}
	{{- if .Start }}
	start = {{ .Start }}
	{{- end }}
	{{- if .End }}
	end = {{ .End }}
	{{- end }}

	{{- if .Message }}
	message = "{{ .Message }}"
	{{- end }}
	{{- if .MonitorId }}
	monitor_id   = {{ .MonitorId }}
	{{- end }}

	{{- if .Recurrence }}
	recurrence {
		period = {{ .Recurrence.Period }}
		type   = "{{ .Recurrence.Type }}"
		{{- if .Recurrence.UntilDate }}
		until_date   = {{ .Recurrence.UntilDate }}
		{{- end }}
		{{- if .Recurrence.UntilOccurrences }}
		until_occurrences   = {{ .Recurrence.UntilOccurrences }}
		{{- end }}
		
		{{- if .Recurrence.WeekDays }}
		week_days = [
		{{- range .Recurrence.WeekDays }}
			"{{ . }}",
		{{- end }}
		]
		{{- end }}
	}
	{{- end }}
}
`
