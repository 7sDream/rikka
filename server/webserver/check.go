package webserver

import (
	"flag"
	"github.com/7sDream/rikka/common/util"

	pathutil "path/filepath"
)

var argWebServerRootDir = flag.String(
	"wsroot", "server/webserver",
	"Root dir of rikka web server, templates should be saved in wsroot/templates "+
		"and static files should be saved in wsroot/static",
)

var (
	staticDirName   = "static"
	templateDirName = "templates"

	webServerRootDirPath = "" // will get from cli argument
	staticDirPath        = "" // append webServerRootDirPath with staticDirName
	templateDirPath      = "" // append webServerRootDirPath with templateDirName

	homeTemplateFileName         = "index.html"
	viewTemplateFileName         = "view.html"
	finishedViewTemplateFileName = "viewFinish.html"

	homeTemplateFilePath         = ""
	viewTemplateFilePath         = ""
	finishedViewTemplateFilePath = ""

	staticFilesName = []string{
		"common.css", "index.css", "view.css",
		"copy.js", "getSrc.js", "onError.js", "checkForm.js",
		"rikka.png", "favicon.png",
	}
)

// updatePathVars update module level path var
func updatePathVars(root string) {
	webServerRootDirPath = root
	staticDirPath = pathutil.Join(webServerRootDirPath, staticDirName)
	templateDirPath = pathutil.Join(webServerRootDirPath, templateDirName)

	homeTemplateFilePath = pathutil.Join(templateDirPath, homeTemplateFileName)
	viewTemplateFilePath = pathutil.Join(templateDirPath, viewTemplateFileName)
	finishedViewTemplateFilePath = pathutil.Join(templateDirPath, finishedViewTemplateFileName)
}

// calcRequireFileList calc file list that web server require
func calcRequireFileList(root string) []string {

	updatePathVars(root)

	requireFiles := []string{
		webServerRootDirPath,

		templateDirPath,
		homeTemplateFilePath, viewTemplateFilePath, finishedViewTemplateFilePath,

		staticDirPath,
	}

	for _, filename := range staticFilesName {
		requireFiles = append(requireFiles, pathutil.Join(staticDirPath, filename))
	}

	return requireFiles
}

// Check needed files like html, css, js, logo, etc...
func checkFiles() {
	l.Info("Args wsroot = ", *argWebServerRootDir)

	absWebServerDir, err := pathutil.Abs(*argWebServerRootDir)
	if err != nil {
		l.Fatal("Proviede web server root dir", *argWebServerRootDir, "is a invalid path")
	}
	l.Debug("Change web server dir to absolute:", absWebServerDir)

	l.Info("Check needed files")

	for _, filepath := range calcRequireFileList(absWebServerDir) {
		if !util.CheckExist(filepath) {
			l.Fatal(filepath, "not exist, check failed")
		} else {
			l.Debug("File", filepath, "exist, check passed")
		}
	}

	l.Info("Check needed files successfully")
}
