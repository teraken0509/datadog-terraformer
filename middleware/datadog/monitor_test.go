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
				Query:   setString("test query 1"),
				Options: &datadog.Options{
					NotifyAudit:       setBool(true),
					NotifyNoData:      setBool(true),
					RenotifyInterval:  setInt(10),
					NewHostDelay:      setInt(5),
					EvaluationDelay:   setInt(300),
					IncludeTags:       setBool(true),
					RequireFullWindow: setBool(true),
					Locked:            setBool(true),
				},
				Tags: []string{
					"test",
					"aaaa",
				},
			},
			expected: "\n# Monitor Configuration for 10000\nresource \"datadog_monitor\" \"monitor_10000\" {\n\tname               = \"Test Monitor 1\"\n\ttype               = \"query_alert\"\n\tmessage            = \"test message 1\"\n  \n\tquery = \"test query 1\"\n\n\tnotify_no_data    = true\n\tnew_host_delay = 5\n\tevaluation_delay = 300\n\trenotify_interval = 10\n\tnotify_audit = true\n\tinclude_tags = true\n\trequire_full_window = true\n\tlocked = true\n  \n\ttags = [\n\t  \"test\",\n\t  \"aaaa\",\n\t]\n}\n",
		}, {
			in: &datadog.Monitor{
				Id:      setInt(10001),
				Type:    setString("query_alert"),
				Name:    setString("Test Monitor 2"),
				Message: setString("test message 2"),
				Query:   setString("test query 2"),
				Options: &datadog.Options{
					NotifyAudit:       setBool(false),
					NotifyNoData:      setBool(false),
					Locked:            setBool(false),
					IncludeTags:       setBool(false),
					RequireFullWindow: setBool(false),
				},
				Tags: []string{
					"test",
					"bbbb",
				},
			},
			expected: "\n# Monitor Configuration for 10001\nresource \"datadog_monitor\" \"monitor_10001\" {\n\tname               = \"Test Monitor 2\"\n\ttype               = \"query_alert\"\n\tmessage            = \"test message 2\"\n  \n\tquery = \"test query 2\"\n\n\tnotify_no_data    = false\n\tnotify_audit = false\n\tinclude_tags = false\n\trequire_full_window = false\n\tlocked = false\n  \n\ttags = [\n\t  \"test\",\n\t  \"bbbb\",\n\t]\n}\n",
		}, {
			in: &datadog.Monitor{
				Id:      setInt(10001),
				Type:    setString("query_alert"),
				Name:    setString("Test Monitor 3"),
				Message: setString("test message 3"),
				Query:   setString("test query 3"),
				Options: &datadog.Options{
					NotifyAudit:       setBool(false),
					NotifyNoData:      setBool(false),
					IncludeTags:       setBool(false),
					RequireFullWindow: setBool(false),
					Locked:            setBool(false),
					ThresholdWindows: &datadog.ThresholdWindows{
						RecoveryWindow: setString("last_15m"),
						TriggerWindow:  setString("last_15m"),
					},
				},
				Tags: []string{
					"test",
					"bbbb",
				},
			},
			expected: "\n# Monitor Configuration for 10001\nresource \"datadog_monitor\" \"monitor_10001\" {\n\tname               = \"Test Monitor 3\"\n\ttype               = \"query_alert\"\n\tmessage            = \"test message 3\"\n  \n\tquery = \"test query 3\"\n\tthreshold_windows {\n\t\trecovery_window\t= \"last_15m\"\n\t\ttrigger_window\t= \"last_15m\"\n\t}\n\n\tnotify_no_data    = false\n\tnotify_audit = false\n\tinclude_tags = false\n\trequire_full_window = false\n\tlocked = false\n  \n\ttags = [\n\t  \"test\",\n\t  \"bbbb\",\n\t]\n}\n",
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
