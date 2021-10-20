package html

import (
	"bytes"
	"net/http"
	"strings"
	"text/template"
)

type GitTemplateData struct {
	Hostname    string
	GitInstance string
	BasePackage string
	FullPackage string
}

var (
	gitTemplate *template.Template
)

func Git(gitinstance string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := GitTemplateData{
			Hostname:    r.Host,
			GitInstance: gitinstance,
			BasePackage: strings.Split(r.URL.Path, "/")[1],
			FullPackage: r.URL.Path[1:],
		}

		var buf bytes.Buffer
		if err := gitTemplate.Execute(&buf, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Only write the response on success
		w.Write(buf.Bytes())
	}
}
