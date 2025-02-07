package templates

import (
	"os"
	"strings"
	"text/template"
)

func capitalize(txt string) string {
	t := strings.Split(txt, "")
	return strings.ToUpper(t[0]) + strings.Join(t[1:], "")
}

func trimS(txt string) string {
	sIndex := strings.LastIndex(txt, "s")
	lenText := len(txt)
	lastIndex := lenText - 1
	if sIndex == lastIndex {
		return string(txt[:lastIndex])
	}
	return txt
}

func pathSep() string {
	return string(os.PathSeparator)
}

var templateFunctions = template.FuncMap{
	"capitalize": capitalize,
	"trimS":      trimS,
	"pathSep": pathSep,
}
