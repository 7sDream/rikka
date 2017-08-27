package main

import (
	"bytes"
	"flag"
	"text/template"

	"github.com/7sDream/rikka/api"
)

var (
	templateArgs = []*bool{
		flag.Bool("r", false, "reStructuredText format"),
		flag.Bool("b", false, "BBCode format"),
		flag.Bool("h", false, "HTML format"),
		flag.Bool("m", false, "Markdown format"),
		flag.Bool("s", true, "Src url format"),
	}

	templateStrings = []string{
		".. image:: {{ .URL }}",
		"[img]{{ .rURL }}[/img]",
		"<img src=\"{{ .URL }}\" >",
		"![]({{ .URL }})",
		"{{ .URL }}",
	}
)

func format(url *api.URL) string {

	var templateStr string

	for i, v := range templateArgs {
		if *v {
			templateStr = templateStrings[i]
			break
		}
	}

	htmlTemplate, err := template.New("_").Parse(templateStr)
	if err != nil {
		l.Fatal("Error happened when create htmlTemplate with string", templateStr, ":", err)
	}
	l.Debug("Create htmlTemplate with string", templateStr, "successfully")

	strWriter := bytes.NewBufferString("")

	if err = htmlTemplate.Execute(strWriter, url); err != nil {
		l.Fatal("Error happened when execute htmlTemplate :", err)
	}
	l.Debug("Execute htmlTemplate successfully")

	return strWriter.String()
}
