package middleware

import (
	"os"
	"text/template"

	"github.com/kterada0509/datadog-terraformer/internal"
	datadog "github.com/zorkian/go-datadog-api"
)

// GetTimeboard ...
func (c *Credential) GetTimeboard(id int) (*datadog.Dashboard, error) {
	board, err := c.Client.GetDashboard(id)
	if err != nil {
		return nil, err
	}
	return board, err
}

// PrintTimeBoardConfiguration ...
func PrintTimeBoardConfiguration(board *datadog.Dashboard) error {
	tmpl := template.Must(template.New("timeboardTemplate").Funcs(internal.TemplateFuncs).Parse(timeboardTemplate))

	if err := tmpl.Execute(os.Stdout, *board); err != nil {
		return err
	}
	return nil
}

const timeboardTemplate = `
resource "datadog_timeboard" "timeboard_{{ .Id }}" {
	title       = "{{ .Title }}"
	description = "{{ .Description }}"
	read_only   = {{ .ReadOnly }}
  
	{{ range .Graphs }}
	graph {
		title = "{{ .Title }}"
		viz   = "{{ .Definition.Viz }}"
		
		{{- range .Definition.Requests }}
		request {
		  	q    = "{{ .Query }}"
		  	type = "{{ .Type }}"
			{{- if .Aggregator }}
			aggregator = "{{ .Aggregator }}"
			{{- end }}
			{{- if .Stacked }}
			stacked = {{ .Stacked }}
			{{- end }}
		  	{{- if or .Style.Palette .Style.Width .Style.Type }}
			style {
				{{- if .Style.Palette }}
				palette = "{{ .Style.Palette }}"
				{{- end }}
				{{- if .Style.Width }}
				width = "{{ .Style.Width }}"
				{{- end }}
				{{- if .Style.Type }}
				type = "{{ .Style.Type }}"
				{{- end }}
			}
			{{- end }}
			{{ range .ConditionalFormats }}
			conditional_format {
				{{- if .Palette }}
				palette = "{{ .Palette }}"
				{{- end }}
				{{- if .Comparator }}
				comparator = "{{ .Comparator }}"
				{{- end }}
				{{- if .Value }}
				value = "{{ .Value }}"
				{{- end }}
				{{- if .CustomFgColor }}
				custom_fg_color = "{{ .CustomFgColor }}"
				{{- end }}
				{{- if .CustomBgColor }}
				custom_bg_color = "{{ .CustomBgColor }}"
				{{- end }}
			}
			{{- end }}  
		}
		{{- end }}

		{{- if .Definition.Events }}
		events = [
		{{- range .Definition.Events }}
			"{{ .Query }}"
		{{- end }}
		]
		{{- end }}
		{{- if .Definition.Autoscale }}
		autoscale = {{ .Definition.Autoscale }}
		{{- end }}
		{{- if .Definition.Precision }}
		precision = "{{ .Definition.Precision }}"
		{{- end }}
		{{- if .Definition.CustomUnit }}
		custom_unit = "{{ .Definition.CustomUnit }}"
		{{- end }}
		{{- if .Definition.TextAlign }}
		text_align = "{{ .Definition.TextAlign }}"
		{{- end }}
		{{- if .Definition.Style }}
		style {
			{{- if .Definition.Style.Palette }}
			palette = "{{ .Definition.Style.Palette }}"
			{{- end }}
			{{- if .Definition.Style.PaletteFlip }}
			palette_flip = {{ .Definition.Style.PaletteFlip }}
			{{- end }}
			{{- if .Definition.Style.FillMin }}
			fill_min = "{{ .Definition.Style.FillMin }}"
			{{- end }}
			{{- if .Definition.Style.FillMax }}
			fill_max = "{{ .Definition.Style.FillMax }}"
			{{- end }}
		}
		{{- end }}
		{{- if .Definition.Groups }}
		group = [
		{{ range .Definition.Groups }}
			"{{ . }}"
		{{- end }}
		]
		{{- end }}
		{{- if .Definition.IncludeNoMetricHosts }}
		include_no_metric_hosts = {{ .Definition.IncludeNoMetricHosts }}
		{{- end }}
		{{- if .Definition.IncludeUngroupedHosts }}
		include_ungrouped_hosts = {{ .Definition.IncludeUngroupedHosts }}
		{{- end }}
		{{- if .Definition.NodeType }}
		node_type = "{{ .Definition.NodeType }}"
		{{- end }}
		{{ range .Definition.Scopes }}
		scope = [
				"{{ .}}"
		]
		{{- end }}

		{{- if or .Definition.Yaxis.Min .Definition.Yaxis.Min .Definition.Yaxis.Scale }}
		yaxis {
			{{- if  .Definition.Yaxis.Min }}
			min = "{{ .Definition.Yaxis.Min }}"
			{{- end }}
			{{- if  .Definition.Yaxis.Max }}
			max = "{{ .Definition.Yaxis.Max }}"
			{{- end }}
			{{- if  .Definition.Yaxis.Scale }}
			scale = "{{ .Definition.Yaxis.Scale }}"
			{{- end }}
		}
		{{- end }}

		{{- range .Definition.Markers }}
		markers {
		  type = "{{ .Type }}"
		  value = "{{ .Value }}"
		  {{- if .Label }}
		  label = "{{ .Label }}"
		  {{- end }}
		}
		{{- end }}

	}
	{{ end }}

	{{- range .TemplateVariables }}
	template_variable {
	  name = "{{ .Name }}"
	  prefix = "{{ .Prefix }}"
	  default = "{{ .Default }}"
	}
	{{- end }}
}
`
