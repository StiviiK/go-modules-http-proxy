package html

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/StiviiK/go-modules-http-proxy/cmd/serve/config"
	"html/template"
	"net/http"
	"strings"
)

const defaultRedirect = "https://pkg.go.dev%s"

var htmlTemplate *template.Template

func ParseTemplate(fs embed.FS) {
	htmlTemplate = template.Must(template.ParseFS(fs, "assets/template.html"))
}

func All(config *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// try to get module from url
		module := getModule(config, fmt.Sprintf("%s/%s", r.Host, strings.Split(r.URL.Path, "/")[1]))
		if module == nil {
			http.Error(w, "404 page not found", http.StatusNotFound)
			return
		}

		// explicitly copy module to avoid further changes
		moduleCopy := *module

		// Check if a redirect is set
		if moduleCopy.Redirect == "" {
			moduleCopy.Redirect = fmt.Sprintf(defaultRedirect, r.URL.Path)
		}

		// Render the template
		var buf bytes.Buffer
		if err := htmlTemplate.Execute(&buf, moduleCopy); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Only write the response on success
		if _, err := w.Write(buf.Bytes()); err != nil {
			// Todo: Log error or else
			return
		}
	}
}

func getModule(config *config.Config, name string) *config.Module {
	for _, _module := range config.Modules {
		if _module.Package == name {
			return _module
		}
	}

	return nil
}
