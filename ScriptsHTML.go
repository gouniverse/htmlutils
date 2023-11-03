package htmlutils

import (
	"log"
	"strings"
)

// ScriptsHTML the HTML from scripts string
func ScriptsHTML(str string) string {
	scripts := strings.Split(str, ",")

	scriptsHTML := ""

	for _, script := range scripts {

		if strings.HasPrefix(script, "http://") || strings.HasPrefix(script, "https://") {
			scriptsHTML += "<script src=\"" + script + "\"></script>"
			continue
		}

		path := "views/" + script

		if !fileExists(path) {
			println(path + " does not exists")
			continue
		}

		scriptText, getContentsError := fileGetContents(path)

		if getContentsError != nil {
			log.Println(path + " get contents error")
			continue
		}

		scriptText, err := MinScript(scriptText)
		if err != nil {
			log.Println(err)
		}
		scriptsHTML += "<script>" + scriptText + "</script>"
	}

	return scriptsHTML
}
