package Monitor

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"ipsd_vsc/Utils"
)

const (
	MARKDOWNPAGE_METADATA_END = `[//]: # "METADATA_END_29b43fcf-5b71-4b15-a048-46765f5ef048"`
)

func ReadMarkdownPageProperties(filePath string) (*MarkdownProperties, bool, error) {
	if Utils.PathIsExist(filePath) == false {
		var errMsg = "Monitor.ReadMarkdownPageProperties: Markdown File not exist " + filePath
		Utils.Logger.Println(errMsg)
		return nil, false, errors.New(errMsg)
	}

	metadataFilePath, errMetaDataFilePath := Utils.GetMetaFilePathWithSameName(filePath)

	if errMetaDataFilePath != nil {
		var errMsg = "Monitor.ReadMarkdownPageProperties: cannot find the metadata file associated with  " + filePath
		Utils.Logger.Println(errMsg)
		return nil, false, errors.New(errMsg)
	}

	bFileContent, errReadFile := ioutil.ReadFile(metadataFilePath)

	if errReadFile != nil {
		var errMsg string
		errMsg = "Monitor.ReadMarkdownPageProperties: Cannot read Markdown file " + filePath
		Utils.Logger.Println(errMsg)

		return nil, false, errors.New(errMsg)
	}

	var mdProperties MarkdownProperties

	errUnmarshal := json.Unmarshal(bFileContent, &mdProperties)
	if errUnmarshal != nil {
		Utils.Logger.Println(errUnmarshal.Error())
		return nil, false, errUnmarshal
	}

	var mdP *MarkdownProperties
	mdP = &mdProperties

	return mdP, true, nil
}

type MarkdownProperties struct {
	Title       string
	Author      string
	Description string
	IsTop       bool
}

type MarkdownPage struct {
	FilePath     string
	ID           string
	LastModified string
}
