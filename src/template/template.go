package template

import (
	"os"
	"strings"
	"text/template"
)

var LinkMergeRequestTemplate *template.Template
var TagTemplate *template.Template
var TextTemplate *template.Template

func GetSrcPath() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	currentDir := strings.Split(wd, "\\")
	srcIndex := 0
	for index, value := range currentDir {
		if value == "src" {
			srcIndex = index
			break
		}
	}

	return strings.Join(currentDir[:srcIndex+1], "\\")
}

func init() {
	var err error
	workdir := GetSrcPath()
	workdir += "\\template"
	LinkMergeRequestTemplate, err = template.ParseFiles(workdir + "/link_merge_request.tpl")
	if err != nil {

	}
	//TagTemplate, err = template.ParseFiles(currentDir + "/link_tag_all.tpl")
	TagTemplate, err = template.ParseFiles(workdir + "/link_tag_all.tpl")

	if err != nil {

	}
	//TextTemplate, err = template.ParseFiles(currentDir + "/text_push.tpl")
	TextTemplate, err = template.ParseFiles(workdir + "/text_push.tpl")
	if err != nil {

	}
}
