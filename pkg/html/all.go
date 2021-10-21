package html

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/StiviiK/go-modules-http-proxy/cmd"
)

const defaultRedirect = "https://pkg.go.dev%s"

var htmlTemplate *template.Template

func ParseTemplate(fs embed.FS) {
	htmlTemplate = template.Must(template.ParseFS(fs, "assets/template.html"))
}

func All(config *cmd.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// try to get module from url
		module := get_module(config, fmt.Sprintf("%s/%s", r.Host, strings.Split(r.URL.Path, "/")[1]))
		if module == nil {
			http.Error(w, "404 page not found", http.StatusNotFound)
			return
		}

		// explicitly copy module to avoid further changes
		module_copy := *module

		// Check if a redirect is set
		if module_copy.Redirect == "" {
			module_copy.Redirect = fmt.Sprintf(defaultRedirect, r.URL.Path)
		}

		// Render the template
		var buf bytes.Buffer
		if err := htmlTemplate.Execute(&buf, module_copy); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Only write the response on success
		w.Write(buf.Bytes())
	}
}

func get_module(config *cmd.Config, name string) *cmd.Module {
	for _, _module := range config.Modules {
		if _module.Package == name {
			return _module
		}
	}

	return nil
}
