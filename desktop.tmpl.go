package desktopfiles

import (
	"strings"
	"text/template"
)

var tmpl = template.Must(template.New(".desktop").Funcs(template.FuncMap{
	"List": func(s []string) string {
		return strings.Join(s, ";") + ";"
	},
}).Parse(`[Desktop Entry]
Version=1.0
Name={{.Name}}
{{- if .LocalizedNames}}{{range $l, $n := .LocalizedNames}}
Name[{{$l}}]={{$n}}
{{end}}{{end}}

{{- if .GenericName}}GenericName={{.GenericName}}{{end}}
{{- if .LocalizedGenericNames}}{{range $l, $g := .LocalizedGenericNames}}
GenericName[{{$l}}]={{$g}}
{{end}}{{end}}

Comment={{.Comment}}
{{- if .LocalizedComments}}{{range $l, $c := .LocalizedComments}}
Comment[{{$l}}]={{$c}}
{{end}}{{end}}
Exec={{.Exec}}
Icon={{.Icon}}
Terminal={{.Terminal}}
Type=Application
{{- if .MimeTypes}}
MimeType={{List .MimeTypes}}
{{end}}
StartupNotify=true
{{- if .Categories}}
Categories={{List .Categories}}
{{end}}
{{- if .Keywords}}
Keywords={{List .Keywords}}
{{end}}
{{- if .Actions}}
Actions={{List .Actions}}

{{range .Actions}}
[Desktop Action {{.}}]
Name={{.Name}}
{{- range $l, $n := .LocalizedNames}}
Name[{{$l}}]={{$n}}
{{end -}}
Exec={{.Exec}}
{{end}}
{{end}}
`))
