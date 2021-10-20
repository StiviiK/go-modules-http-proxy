package html

import (
	"embed"
	"text/template"
)

func ParseTemplates(fs embed.FS) {
	gitTemplate = template.Must(template.ParseFS(fs, "assets/templates/git.html"))
}
