package middleware

import (
	"os"
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

func PrintMonitorConfiguration(monitor *datadog.Monitor) error {
	tmpl := template.Must(template.New("monitor.tmpl").Funcs(internal.TemplateFuncs).ParseFiles("templates/datadog/monitor.tmpl"))

	if err := tmpl.Execute(os.Stdout, *monitor); err != nil {
		return err
	}
	return nil
}
