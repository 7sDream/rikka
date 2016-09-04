package webserver

import "github.com/7sDream/rikka/common/util"

// Check needed files like html, css, js, logo, etc...
func checkFiles() {
	l.Info("Check needed files")

	requireFiles := []string{
		"templates", "templates/index.html", "templates/view.html", "templates/viewFinish.html",
		"static", "static/main.css", "static/index.css", "static/view.css", "static/rikka.png",
		"static/copy.js", "static/getSrc.js",
	}

	for _, filepath := range requireFiles {
		if !util.CheckExist(filepath) {
			l.Fatal(filepath, "not exist, check failed")
		} else {
			l.Debug("File", filepath, "exist, check passed")
		}
	}

	l.Info("Check needed files successfully")
}
