package webserver

import (
	"flag"
	"github.com/7sDream/rikka/common/util"

	pathUtil "path/filepath"
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
		"css",
		"css/common.css", "css/index.css", "css/view.css",
		"js",
		"js/copy.js", "js/getSrc.js", "js/onError.js", "js/checkForm.js",
		"image",
		"image/rikka.png", "image/favicon.png",
	}
)

// updatePathVars update module level path var
func updatePathVars(root string) {
	webServerRootDirPath = root
	staticDirPath = pathUtil.Join(webServerRootDirPath, staticDirName)
	templateDirPath = pathUtil.Join(webServerRootDirPath, templateDirName)

	homeTemplateFilePath = pathUtil.Join(templateDirPath, homeTemplateFileName)
	viewTemplateFilePath = pathUtil.Join(templateDirPath, viewTemplateFileName)
	finishedViewTemplateFilePath = pathUtil.Join(templateDirPath, finishedViewTemplateFileName)
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
		requireFiles = append(requireFiles, pathUtil.Join(staticDirPath, filename))
	}

	return requireFiles
}

// Check needed files like html, css, js, logo, etc...
func checkFiles() {
	l.Info("Args wsroot = ", *argWebServerRootDir)

	absWebServerDir, err := pathUtil.Abs(*argWebServerRootDir)
	if err != nil {
		l.Fatal("Provided web server root dir", *argWebServerRootDir, "is a invalid path")
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
