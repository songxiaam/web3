package test

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

var PushHtml = template.Must(template.New("https").Parse(`
<html>
<head>
<title>Https Test</title>
<meta charset="utf-8">
</head>
<body>
	<h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`,
))

func TestPush(c *gin.Context) {
	if pusher := c.Writer.Pusher(); pusher != nil {
		if err := pusher.Push("/assets/app.js", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}

	c.HTML(http.StatusOK, "http", gin.H{
		"status": "success",
	})

}
