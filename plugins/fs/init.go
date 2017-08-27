package fs

import (
	"os"
	pathUtil "path/filepath"

	"github.com/7sDream/rikka/common/util"
)

// Init is the plugin init function, will be called when plugin be load.
func (fsp fsPlugin) Init() {
	// where to store file
	l.Info("Start plugin fs")

	l.Info("Args dir =", *argFilesDir)
	l.Info("Args fsDebugSleep =", *argFsDebugSleep)

	absFilesDir, err := pathUtil.Abs(*argFilesDir)
	if err == nil {
		l.Debug("Abs path of image file dir:", absFilesDir)
		imageDir = absFilesDir
	} else {
		l.Fatal("A error happened when change image dir to absolute path:", err)
	}
	// if target dir not exist, create it
	if util.CheckExist(absFilesDir) {
		l.Debug("Image file dir already exist")
	} else {
		l.Debug("Image file dir not exist, try to create it")
		err = os.MkdirAll(absFilesDir, 0755)
		if err == nil {
			l.Debug("Create dir", absFilesDir, "successfully")
		} else {
			l.Fatal("A error happened when try to create image dir:", err)
		}
	}

	l.Info("Fs plugin start successfully")
}
