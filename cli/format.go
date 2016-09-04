package main

import (
	"bytes"
	"flag"
	"text/template"

	"github.com/7sDream/rikka/api"
)

var formatList = []*bool{
	flag.Bool("r", false, "reStructuredText format"),
	flag.Bool("b", false, "BBCode format"),
	flag.Bool("h", false, "HTML format"),
	flag.Bool("m", false, "markdown format"),
	flag.Bool("s", true, "src address format"),
}

var formatMap = []string{
	".. image:: {{ .URL }}",
	"[img]{{ .rURL }}[/img]",
	"<img src=\"{{ .URL }}\" >",
	"![]({{ .URL }})",
	"{{ .URL }}",
}

func format(url *api.URL) string {

	var templateStr string

	for i, v := range formatList {
		if *v {
			templateStr = formatMap[i]
			break
		}
	}

	template, err := template.New("_").Parse(templateStr)
	if err != nil {
		l.Fatal("Error happened when create template with string", templateStr, ":", err)
	}
	l.Debug("Create template with string", templateStr, "successfully")

	strWriter := bytes.NewBufferString("")

	if err = template.Execute(strWriter, url); err != nil {
		l.Fatal("Error happened when execute template :", err)
	}
	l.Debug("Execute template successfully")

	return strWriter.String()
}
