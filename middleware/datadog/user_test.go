package middleware

import (
	"bytes"
	"fmt"
	"testing"

	datadog "github.com/zorkian/go-datadog-api"
)

func Test_PrintUserConfiguration(t *testing.T) {
	cases := []struct {
		in       *datadog.User
		expected string
	}{
		{
			in: &datadog.User{
				Handle:   setString("aaaa@example.com"),
				Disabled: nil,
				Email:    setString("aaaa@example.com"),
				Name:     nil,
				Role:     nil,
				IsAdmin:  nil,
			},
			expected: "\n# User Configuration for aaaa@example.com\nresource \"datadog_user\" \"user_aaaa@example.com\" {\n\temail  = \"aaaa@example.com\"\n\thandle = \"aaaa@example.com\"\n\tname   = \"<nil>\"\n}\n",
		}, {
			in: &datadog.User{
				Handle:   setString("bbb@example.com"),
				Disabled: nil,
				Email:    setString("bbb@example.com"),
				Name:     nil,
			},
			expected: "\n# User Configuration for bbb@example.com\nresource \"datadog_user\" \"user_bbb@example.com\" {\n\temail  = \"bbb@example.com\"\n\thandle = \"bbb@example.com\"\n\tname   = \"<nil>\"\n}\n",
		},
	}

	for _, c := range cases {
		buffer = &bytes.Buffer{}
		writer = buffer
		if err := PrintUserConfiguration(writer, c.in); err != nil {
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
