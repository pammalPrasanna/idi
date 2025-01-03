package templates

import (
	"strings"
	"text/template"
)

func capitalize(txt string) string {
	t := strings.Split(txt, "")
	return strings.ToUpper(t[0]) + strings.Join(t[1:], "")
}

var templateFunctions = template.FuncMap{
	"capitalize": capitalize,
}
