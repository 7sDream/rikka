package main

import "github.com/7sDream/rikka/client"

type taskRes struct {
	Index    int
	Filepath string
	Res      string
	Err      error
}

func (res taskRes) String() string {
	if res.Err != nil {
		return res.Filepath + ": " + res.Err.Error()
	}
	return res.Filepath + ": " + res.Res
}

func (res taskRes) StringWithoutFilepath() string {
	if res.Err != nil {
		return "Error:" + res.Err.Error()
	}
	return res.Res
}

func buildErrorRes(index int, filepath string, err error) *taskRes {
	return &taskRes{
		Index:    index,
		Filepath: filepath,
		Err:      err,
	}
}

func worker(host string, filepath string, index int, out chan *taskRes) {
	absFilepath, fileContent, err := readFile(filepath)
	if err != nil {
		out <- buildErrorRes(index, filepath, err)
		return
	}
	l.Info("Read file", absFilepath, "successfully")

	taskID, err := client.Upload(host, absFilepath, fileContent, getPassword())
	if err != nil {
		out <- buildErrorRes(index, filepath, err)
		return
	}
	l.Info("Upload successfully, get taskID:", taskID)

	err = client.WaitFinish(host, taskID)
	if err != nil {
		out <- buildErrorRes(index, filepath, err)
		return
	}
	l.Info("Task state comes to finished")

	pURL, err := client.GetURL(host, taskID)
	if err != nil {
		out <- buildErrorRes(index, filepath, err)
		return
	}
	l.Info("Url gotten:", *pURL)

	formatted := format(pURL)
	l.Info("Make final formatted text successfully:", formatted)

	out <- &taskRes{
		Index:    index,
		Filepath: filepath,
		Res:      formatted,
		Err:      nil,
	}
}
