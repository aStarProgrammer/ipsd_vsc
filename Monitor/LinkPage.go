package Monitor

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"ipsd_vsc/Utils"
	"path/filepath"
	"strings"
)

type LinkMeta struct {
	Url    string
	Title  string
	Author string
	IsTop  bool
}

type LinkPage struct {
	Url          string
	ID           string
	LastModified string
	Title        string
}

func ReadLinkMetadatasFromFolder(folderPath string) ([]LinkMeta, error) {
	if Utils.PathIsExist(folderPath) == false {
		return nil, errors.New("ReadLinkMetadatasFromFolder: FolderPath not exist " + folderPath)
	}

	files, errReadDir := ioutil.ReadDir(folderPath)

	if errReadDir != nil {
		return nil, errReadDir
	}

	var linkMetas []LinkMeta
	var fileNames []string
	var all, success int
	all = 0
	success = 0
	for _, file := range files {
		var fName = file.Name()
		var fPath = filepath.Join(folderPath, fName)

		if strings.HasSuffix(fPath, ".lik.json") {
			all = all + 1
			linkMeta, errReadLink := ReadLinkMetaFromFile(fPath)

			if errReadLink == nil {
				success = success + 1
				linkMetas = append(linkMetas, *linkMeta)
			} else {
				fileNames = append(fileNames, fPath)
			}
		}
	}

	if all != success {
		var errMsg = "ReadLinkMetasFromFolder: Didn't read all the liks"
		for _, fileName := range fileNames {
			errMsg = errMsg + fileName
		}
		Utils.Logger.Println(errMsg)
	}

	return linkMetas, nil
}

func ReadLinkMetaFromFile(filePath string) (*LinkMeta, error) {
	if Utils.PathIsExist(filePath) == false {
		return nil, errors.New("ReadLinkMetaFromFile: FilePath not exist " + filePath)
	}

	bFileContent, errReadFile := ioutil.ReadFile(filePath)

	if errReadFile != nil {
		var errMsg string
		errMsg = "ReadLinkMetaFromFile: : Cannot read Link Meta file " + filePath
		Utils.Logger.Println(errMsg)

		return nil, errors.New(errMsg)
	}

	var link *LinkMeta
	errUnmarshal := json.Unmarshal(bFileContent, &link)
	if errUnmarshal != nil {
		Utils.Logger.Println(errUnmarshal.Error())
		return nil, errUnmarshal
	}

	return link, nil
}
