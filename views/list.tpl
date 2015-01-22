<html>
<head>

</head>
<body>
<div>
Message:{{.msg}}
</div>
<div>
{{range $key, $val := .jas}}
<div>
{{$key}}
{{$val.Title}}
{{$val.Text}}
{{$val.IsPublished}}
</div>
{{end}} 
</div>
</body>
</html>