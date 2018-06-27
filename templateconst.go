package main

const pagestart = `

<html>
<body>
`
const pageend = `
</body>
</html>
`

func showimage(current string, next string) string {
	return "<a href='" + next + " '><img src='" + current + "'></a>"
}
