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

func cleanName(txt string) string {
	cleaner := func(t []string) string {
		f := ""
		for i, s := range t {
			if i == 0 {
				f += s
				continue
			}
			f += capitalize(s)
		}
		return f
	}

	if strings.Contains(txt, "-") {
		return cleaner(strings.Split(txt, "-"))
	} else if strings.Contains(txt, "_") {
		return cleaner(strings.Split(txt, "_"))
	}
	return txt
}

var templateFunctions = template.FuncMap{
	"capitalize": capitalize,
	"trimS":      trimS,
	"pathSep":    pathSep,
	"cleanName":  cleanName,
}
