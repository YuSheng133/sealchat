package main

import "text/template"

const tmpl = `// Code generated by go generate; DO NOT EDIT.
package gen

// {{.MapName}} 提供权限字符串到描述的映射
var {{.MapName}} = map[string]string{
{{- range .Perms}}
    "{{.Key}}": "{{.Desc}}",
{{- end}}
}


// {{.ArrayName}} 提供权限字符串到描述的映射
var {{.ArrayName}} = []map[string]string{
{{- range .Perms}}
	{"key": "{{.Key}}", "desc": "{{.Desc}}"},
{{- end}}
}
`

var tmplGo = template.Must(template.New("perm").Parse(tmpl))
