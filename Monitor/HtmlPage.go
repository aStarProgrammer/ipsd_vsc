package Monitor

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"ipsd_vsc/Utils"
)

type HtmlPage struct {
	FilePath     string
	ID           string
	LastModified string
}

func ReadHtmlProperties(filePath string) (*HtmlProperties, bool, error) {
	if Utils.PathIsExist(filePath) == false {
		var errMsg = "Monitor.ReadHtmlPageProperties: Html File not exist " + filePath
		Utils.Logger.Println(errMsg)
		return nil, false, errors.New(errMsg)
	}

	metadataFilePath, errMetaDataFilePath := Utils.GetMetaFilePathWithSameName(filePath)

	if errMetaDataFilePath != nil {
		var errMsg = "Monitor.ReadHtmlProperties: cannot find the metadata file associated with  " + filePath
		Utils.Logger.Println(errMsg)
		return nil, false, errors.New(errMsg)
	}

	bFileContent, errReadFile := ioutil.ReadFile(metadataFilePath)

	if errReadFile != nil {
		var errMsg string
		errMsg = "Monitor.ReadHtmlPageProperties: Cannot read Html file " + filePath
		Utils.Logger.Println(errMsg)

		return nil, false, errors.New(errMsg)
	}

	var htmProperties HtmlProperties
	errUnmarshal := json.Unmarshal(bFileContent, &htmProperties)
	if errUnmarshal != nil {
		Utils.Logger.Println(errUnmarshal.Error())
		return nil, false, errUnmarshal
	}
	var htmP *HtmlProperties
	htmP = &htmProperties

	return htmP, true, nil
}

type HtmlProperties struct {
	Title       string
	Author      string
	Description string
	IsTop       bool
}
