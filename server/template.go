package server

import "html/template"

var homeTemplate = template.Must(template.New("home").Parse(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>diiig</title>
    <style>
      .entry form {
          display: inline-block;
      }
    </style>
</head>
<body>
<h1>Create New Topic</h1>
<form action="/topics" method="post">
    <input type="text" name="topic">
    <button type="submit">Create</button>
</form>

<h1>Top 20 Topics</h1>
<ul>
{{range .Topics}}
<li class="entry">
    <form action="/vote" method="post">
        <input type="hidden" name="topic" value="{{.Name}}">
        <input type="hidden" name="score" value="1">
        <button type="submit">👍</button>
    </form>
    {{.Score}}
    <form action="/vote" method="post">
        <input type="hidden" name="topic" value="{{.Name}}">
        <input type="hidden" name="score" value="-1">
        <button type="submit">👎</button>
    </form>
    <strong>{{.Name}}</strong>
</li>
{{end}}
</ul>
</body>
</html>
`))
