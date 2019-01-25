package internal

import (
	"strings"
	"text/template"
)

// TemplateFuncs ...
var TemplateFuncs = template.FuncMap{
	"escapenl": func(text string) string {
		return strings.Replace(template.HTMLEscapeString(text), "\n", "\\n", -1)
	},
}
