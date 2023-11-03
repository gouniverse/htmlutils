package htmlutils

import (
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

// MinHTML minifies the CSS
func MinHTML(htmlString string) (string, error) {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	return m.String("text/html", htmlString)
}
