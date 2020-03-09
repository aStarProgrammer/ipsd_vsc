package Monitor

import (
	"errors"
	"io/ioutil"
	"ipsd_vsc/Utils"
	"strings"
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

	fileContent := string(bFileContent)

	metadataInfo := fileContent

	var tvalues = strings.Split(metadataInfo, ":#")

	var values []string

	for _, v := range tvalues {
		var t = strings.TrimSpace(v)
		if t != "" {
			values = append(values, t)
		}
	}

	if len(values) != 4 {
		var errMsg string
		errMsg = "Monitor.ReadMarkdownPageProperties: Cannot read metadata properties of Markdown file " + filePath
		Utils.Logger.Println(errMsg)

		return nil, false, errors.New(errMsg)
	}

	var title, author, description, vIsTop string
	var isTop bool
	title = strings.TrimSpace(values[0])
	title = values[0][strings.Index(title, ":")+1:]
	title = title[:len(title)-1]
	title = strings.TrimSpace(title)

	author = strings.TrimSpace(values[1])
	author = values[1][strings.Index(author, ":")+1:]
	author = author[:len(author)-1]
	author = strings.TrimSpace(author)

	description = strings.TrimSpace(values[2])
	description = values[2][strings.Index(description, ":")+1:]
	description = description[:len(description)-1]
	description = strings.TrimSpace(description)

	vIsTop = strings.TrimSpace(values[3])
	vIsTop = values[3][strings.Index(vIsTop, ":")+1:]
	vIsTop = vIsTop[:len(vIsTop)-1]
	vIsTop = strings.TrimSpace(vIsTop)
	if strings.Contains(vIsTop, "true") {
		isTop = true
	} else {
		isTop = false
	}

	var mdProperties MarkdownProperties
	var mdP *MarkdownProperties
	mdP = &mdProperties

	mdP.Title = title
	mdP.Author = author
	mdP.Description = description
	mdP.IsTop = isTop

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
