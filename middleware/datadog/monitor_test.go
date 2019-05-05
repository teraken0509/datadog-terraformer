package middleware

import (
	"bytes"
	"fmt"
	"testing"

	datadog "github.com/zorkian/go-datadog-api"
)

func Test_PrintMonitorConfiguration(t *testing.T) {
	cases := []struct {
		in       *datadog.Monitor
		expected string
	}{
		{
			in: &datadog.Monitor{
				Id:      setInt(10000),
				Type:    setString("query_alert"),
				Name:    setString("Test Monitor 1"),
				Message: setString("test message 1"),
				Options: &datadog.Options{},
				Tags: []string{
					"test",
					"aaaa",
				},
			},
			expected: "\n# Monitor Configuration for 10000\nresource \"datadog_monitor\" \"monitor_10000\" {\n\tname               = \"Test Monitor 1\"\n\ttype               = \"query_alert\"\n\tmessage            = \"test message 1\"\n  \n\tquery = \"<nil>\"\n  \n\tnotify_no_data    = <nil>\n\tnotify_audit = <nil>\n\tinclude_tags = <nil>\n\trequire_full_window = <nil>\n\tlocked = <nil>\n  \n\ttags = [\n\t  \"test\",\n\t  \"aaaa\",\n\t]\n}\n",
		}, {
			in: &datadog.Monitor{
				Id:      setInt(10001),
				Type:    setString("query_alert"),
				Name:    setString("Test Monitor 2"),
				Message: setString("test message 2"),
				Options: &datadog.Options{},
				Tags: []string{
					"test",
					"bbbb",
				},
			},
			expected: "\n# Monitor Configuration for 10001\nresource \"datadog_monitor\" \"monitor_10001\" {\n\tname               = \"Test Monitor 2\"\n\ttype               = \"query_alert\"\n\tmessage            = \"test message 2\"\n  \n\tquery = \"<nil>\"\n  \n\tnotify_no_data    = <nil>\n\tnotify_audit = <nil>\n\tinclude_tags = <nil>\n\trequire_full_window = <nil>\n\tlocked = <nil>\n  \n\ttags = [\n\t  \"test\",\n\t  \"bbbb\",\n\t]\n}\n",
		},
	}

	for _, c := range cases {
		buffer = &bytes.Buffer{}
		writer = buffer
		if err := PrintMonitorConfiguration(writer, c.in); err != nil {
			t.Errorf("PrintUserConfiguration: is return err: %s", err)
		}
		fmt.Println("===========================")
		fmt.Println(buffer.String())
		fmt.Println("===========================")
		if c.expected != buffer.String() {
			t.Errorf("got %q, want %q", buffer.String(), c.expected)
		}
	}
}
