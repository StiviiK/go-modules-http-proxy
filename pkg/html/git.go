package html

import (
	"net/http"
	"strings"
	"text/template"
)

type GitTemplateData struct {
	Hostname    string
	GitInstance string
	Namespace   string
	BasePackage string
	FullPackage string
}

var (
	gitTemplate *template.Template
)

func Git(hostname string, gitinstance string, namespace string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := GitTemplateData{
			Hostname:    hostname,
			GitInstance: gitinstance,
			Namespace:   namespace,
			BasePackage: strings.Split(r.URL.Path, "/")[1],
			FullPackage: r.URL.Path[1:],
		}

		if err := gitTemplate.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
