package webserver

import "github.com/7sDream/rikka/common/util"

// Check needed files like html, css, js, logo, etc...
func checkFiles() {
	l.Info("Check needed files")

	requireFiles := []string{
		"templates",
		"templates/index.html", "templates/view.html", "templates/viewFinish.html",
		"static",
		"static/common.css", "static/index.css", "static/view.css",
		"static/copy.js", "static/getSrc.js", "static/onError.js", "static/checkForm.js",
		"static/rikka.png", "static/favicon.png",
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
