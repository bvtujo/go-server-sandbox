package template

import (
	"html/template"

	"github.com/bvtujo/go-server-sandbox/internal/pkg/points"
)

const indexTpl = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>&#x2699 The System</title>
</head>
<body>
	<h1> &#x2699 The System &#x1F423</h1>
	<table>
	  <tr>
		<th>Rank</th>
		<th>User</th>
		<th>Points</th>
	  </tr>
	  {{ range $ind, $val := .Users }}
	  <tr>
		<td>{{ $ind }}</td>
		<td>{{ .Username }}</td>
		<td>{{ .Points }}</td>
	  </tr>
	  {{end}}
	</table>
</body>
</html>
`

type indexPageData struct {
	Users []points.User
}

// GetIndex renders a template according to a list of input structs
func GetIndex() *template.Template {
	tmpl := template.Must(template.ParseGlob(indexTpl))
	return tmpl
}
