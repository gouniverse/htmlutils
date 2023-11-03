package htmlutils

import (
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
)

// MinCSS minifies the CSS
func MinCSS(cssString string) (string, error) {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	return m.String("text/css", cssString)
}
