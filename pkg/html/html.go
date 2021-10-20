package html

import (
	"embed"
	"net/http"
	"strings"
	"text/template"
)

type GitHubTemplateData struct {
	Hostname        string
	GitHubNamespace string
	BasePackage     string
	FullPackage     string
}

var (
	githubTemplate *template.Template
)

func ParseTemplates(fs embed.FS) {
	githubTemplate = template.Must(template.ParseFS(fs, "assets/github.template.html"))
}

func GitHubHandler(w http.ResponseWriter, r *http.Request) {
	data := GitHubTemplateData{
		Hostname:        "go.uber.org",
		GitHubNamespace: "uber-go",
		BasePackage:     strings.Split(r.URL.Path, "/")[1],
		FullPackage:     r.URL.Path[1:],
	}

	if err := githubTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
