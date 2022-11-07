package main

import (
	"gf-workflow/internal/logic/processManage"
	"gf-workflow/internal/logic/view"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"gf-workflow/internal/logic/process"
)

const (
	swaggerUIPageContent = `

<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <meta name="description" content="SwaggerUI"/>
  <title>SwaggerUI</title>
  <link rel="stylesheet" href="public/swagger-ui.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="public/swagger-ui-bundle.js" crossorigin></script>
<script>
	window.onload = () => {
		window.ui = SwaggerUIBundle({
			url:    '/api.json',
			dom_id: '#swagger-ui',
		});
	};
</script>
</body>
</html>
`
)

func main() {
	s := g.Server()
	// 工作流引擎
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/try", func(r *ghttp.Request) {
			r.Response.Write(swaggerUIPageContent)
		})
		group.Bind(
			new(process.Process),
			new(processManage.ProcessManage),
		)
	})

	// 视图
	s.BindObject("/view", view.View{})
	// 静态文件服务
	s.SetServerRoot(".")
	s.Run()
}
