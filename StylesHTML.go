package htmlutils

import (
	"log"
	"strings"
)

// StylesHTML the styles
func StylesHTML(str string) string {
	styles := strings.Split(str, ",")

	htmlStyles := ""

	for _, style := range styles {

		if strings.HasPrefix(style, "http://") || strings.HasPrefix(style, "https://") {
			htmlStyles += "<link href=\"" + style + "\"  rel=\"stylesheet\" />"
			continue
		}

		path := "views/" + style

		if !fileExists(path) {
			println(path + " does not exists")
			continue
		}

		scriptText, getContentsError := fileGetContents(path)

		if getContentsError != nil {
			log.Println(path + " get contents error")
			continue
		}

		styleText, err := MinScript(scriptText)
		if err != nil {
			log.Println(err)
		}

		htmlStyles += "<style>" + styleText + "</style>"
	}

	return htmlStyles
}
