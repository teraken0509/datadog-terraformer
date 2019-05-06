package middleware

import (
	"bytes"
	"fmt"
	"testing"

	datadog "github.com/zorkian/go-datadog-api"
)

func Test_PrintDowntimeConfiguration(t *testing.T) {
	cases := []struct {
		in       *datadog.Downtime
		expected string
	}{
		{
			in: &datadog.Downtime{
				Id: setInt(1000),
				Scope: []string{
					"*",
				},
				Start: setInt(1457121738),
				End:   setInt(1557121738),
				Recurrence: &datadog.Recurrence{
					Type:   setString("days"),
					Period: setInt(3),
				},
				Active:   setBool(false),
				Disabled: setBool(true),
			},
			expected: "\nresource \"datadog_downtime\" \"downtime_1000\" {\n\tactive = false\n\tdisabled = true\n\tscope = [\n\t\t\"*\",\n\t]\n\tstart = 1457121738\n\tend = 1557121738\n\trecurrence {\n\t\tperiod = 3\n\t\ttype   = \"days\"\n\t}\n}\n",
		}, {
			in: &datadog.Downtime{
				Id: setInt(1001),
				Scope: []string{
					"aaaaa",
					"bbbbb",
				},
				Start: setInt(1457121738),
				End:   setInt(1557121738),
				Recurrence: &datadog.Recurrence{
					Type:   setString("weeks"),
					Period: setInt(3),
					WeekDays: []string{
						"Mon",
						"Fri",
					},
				},
				Active:   setBool(true),
				Disabled: setBool(false),
			},
			expected: "\nresource \"datadog_downtime\" \"downtime_1001\" {\n\tactive = true\n\tdisabled = false\n\tscope = [\n\t\t\"aaaaa\",\n\t\t\"bbbbb\",\n\t]\n\tstart = 1457121738\n\tend = 1557121738\n\trecurrence {\n\t\tperiod = 3\n\t\ttype   = \"weeks\"\n\t\tweek_days = [\n\t\t\t\"Mon\",\n\t\t\t\"Fri\",\n\t\t]\n\t}\n}\n",
		},
	}

	for _, c := range cases {
		buffer = &bytes.Buffer{}
		writer = buffer
		if err := PrintDowntimeConfiguration(writer, c.in); err != nil {
			t.Errorf("PrintDowntimeConfiguration: is return err: %s", err)
		}
		fmt.Println("===========================")
		fmt.Println(buffer.String())
		fmt.Println("===========================")
		if c.expected != buffer.String() {
			t.Errorf("got %q, want %q", buffer.String(), c.expected)
		}
	}
}
