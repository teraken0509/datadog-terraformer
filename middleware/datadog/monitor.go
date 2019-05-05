package middleware

import (
	"io"
	"text/template"

	"github.com/kterada0509/datadog-terraformer/internal"

	datadog "github.com/zorkian/go-datadog-api"
)

// GetMonitor ...
func (c *Credential) GetMonitor(id int) (*datadog.Monitor, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}

	monitor, err := c.Client.GetMonitor(id)
	if err != nil {
		return nil, err
	}
	return monitor, err
}

// PrintMonitorConfiguration ...
func PrintMonitorConfiguration(w io.Writer, monitor *datadog.Monitor) error {
	tmpl := template.Must(template.New("monitorTemplate").Funcs(internal.TemplateFuncs).Parse(monitorTemplate))

	if err := tmpl.Execute(w, *monitor); err != nil {
		return err
	}
	return nil
}

const monitorTemplate = `
# Monitor Configuration for {{ .Id }}
resource "datadog_monitor" "monitor_{{ .Id }}" {
	name               = "{{ .Name }}"
	type               = "{{ .Type }}"
	message            = "{{ .Message | escapenl }}"
	{{- if .Options.EscalationMessage }}
	escalation_message = "{{ .Options.EscalationMessage }}"
	{{- end }}
  
	query = "{{ .Query }}"
  
	{{- if .Options.Thresholds }}
	thresholds {
	  {{- if .Options.Thresholds.Ok }}
	  ok                = {{ .Options.Thresholds.Ok }}
	  {{- end }}
	  {{- if .Options.Thresholds.Warning }}
	  warning           = {{ .Options.Thresholds.Warning }}
	  {{- end }}
	  {{- if .Options.Thresholds.WarningRecovery }}
	  warning_recovery  = {{ .Options.Thresholds.WarningRecovery }}
	  {{- end }}
	  {{- if .Options.Thresholds.Critical }}
	  critical          = {{ .Options.Thresholds.Critical }}
	  {{- end }}
	  {{- if .Options.Thresholds.CriticalRecovery }}
	  critical_recovery = {{ .Options.Thresholds.CriticalRecovery }}
	  {{- end }}
	}
	{{- end }}
  
	notify_no_data    = {{ .Options.NotifyNoData }}
	{{- if .Options.NewHostDelay }}
	new_host_delay = {{ .Options.NewHostDelay }}
	{{- end }}
	{{- if .Options.EvaluationDelay }}
	evaluation_delay = {{ .Options.EvaluationDelay }}
	{{- end }}
	{{- if .Options.NoDataTimeframe }}
	no_data_timeframe = {{ .Options.NoDataTimeframe }}
	{{- end }}
	{{- if .Options.RenotifyInterval }}
	renotify_interval = {{ .Options.RenotifyInterval }}
	{{- end }}
	notify_audit = {{ .Options.NotifyAudit }}
	{{- if .Options.TimeoutH }}
	timeout_h    = {{ .Options.TimeoutH }}
	{{- end }}
	include_tags = {{ .Options.IncludeTags }}
	require_full_window = {{ .Options.RequireFullWindow }}
	locked = {{ .Options.Locked }}
  
	{{- if .Options.Silenced }}
	silenced {
	{{- range $key, $val := .Options.Silenced }}
	  "{{ $key }}" = {{ $val }}
	{{- end }}
	}
	{{- end }}
  
	tags = [
	{{- range .Tags }}
	  "{{ . }}",
	{{- end }}
	]
}
`
