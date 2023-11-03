package htmlutils

import (
	"regexp"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
)

// MinScript minifies the script
func MinScript(scriptString string) (string, error) {
	m := minify.New()
	// m.AddFunc("text/css", css.Minify)
	// m.AddFunc("text/html", html.Minify)
	// m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	// m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	// m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
	return m.String("text/javascript", scriptString)
}
