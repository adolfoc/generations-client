package handlers

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

var templateDirs = []string{"views/partials"}

func ExecuteView(viewName string, viewParams interface{}, w http.ResponseWriter) error {
	log := common.StartLog("templates", "ExecuteView")

	htmlView := fmt.Sprintf("%s.html", viewName)
	templates, err := GetTemplates(htmlView)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		log.FailedReturn()
		return err
	}

	err = templates.ExecuteTemplate(w, viewName, viewParams)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		log.FailedReturn()
		return err
	}

	log.NormalReturn()
	return nil
}

func GetTemplates(mainTemplate string) (templates *template.Template, err error) {
	var allFiles []string
	for _, dir := range templateDirs {
		files2, _ := ioutil.ReadDir(dir)
		for _, file := range files2 {
			filename := file.Name()
			if strings.HasSuffix(filename, ".html") {
				filePath := filepath.Join(dir, filename)
				allFiles = append(allFiles, filePath)
			}
		}
	}

	mainTemplatePath := fmt.Sprintf("views/%s", mainTemplate)
	allFiles = append(allFiles, mainTemplatePath)

	templates, err = template.New("").ParseFiles(allFiles...)
	return
}


