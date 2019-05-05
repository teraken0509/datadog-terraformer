package middleware

import (
	"bytes"
	"fmt"
	"testing"

	datadog "github.com/zorkian/go-datadog-api"
)

func Test_PrintTimeBoardConfiguration(t *testing.T) {
	cases := []struct {
		in       *datadog.Dashboard
		expected string
	}{
		{
			in: &datadog.Dashboard{
				Id:          setInt(1000),
				Description: setString("Test Dashboard 1"),
				Title:       setString("Test Dashboard 1"),
			},
			expected: "\nresource \"datadog_timeboard\" \"timeboard_1000\" {\n\ttitle       = \"Test Dashboard 1\"\n\tdescription = \"Test Dashboard 1\"\n\tread_only   = <nil>\n  \n\t\n}\n",
		}, {
			in: &datadog.Dashboard{
				Id:          setInt(1001),
				Description: setString("Test Dashboard 2"),
				Title:       setString("Test Dashboard 2"),
			},
			expected: "\nresource \"datadog_timeboard\" \"timeboard_1001\" {\n\ttitle       = \"Test Dashboard 2\"\n\tdescription = \"Test Dashboard 2\"\n\tread_only   = <nil>\n  \n\t\n}\n",
		},
	}

	for _, c := range cases {
		buffer = &bytes.Buffer{}
		writer = buffer
		if err := PrintTimeBoardConfiguration(writer, c.in); err != nil {
			t.Errorf("PrintTimeBoardConfiguration: is return err: %s", err)
		}
		fmt.Println("===========================")
		fmt.Println(buffer.String())
		fmt.Println("===========================")
		if c.expected != buffer.String() {
			t.Errorf("got %q, want %q", buffer.String(), c.expected)
		}
	}
}
