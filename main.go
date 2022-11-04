package main

import (
	"context"
	"fmt"
	"gf-workflow/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

const (
	swaggerUIPageContent = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <meta name="description" content="SwaggerUI"/>
  <title>SwaggerUI</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@latest/swagger-ui.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="https://unpkg.com/swagger-ui-dist@latest/swagger-ui-bundle.js" crossorigin></script>
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

type StartReq struct {
	g.Meta    `path:"/start" method:"post"`
	ProcessID string `v:"required" dc:"流程ID"`
	UserID    string `v:"required" dc:"发起用户ID"`
}
type StartRes struct {
	Reply string `dc:"Reply content"`
}

type ListReq struct {
	g.Meta `path:"/list" method:"get"`
	Name   string `v:"required" dc:"Your name"`
}
type ListRes struct {
	Reply string `dc:"Reply content"`
}

type Process struct{}

func (Process) Start(ctx context.Context, req *StartReq) (res *StartRes, err error) {
	// 查流程
	process :=entity.ProcessInfos{}
	err= g.Model(entity.ProcessInfos{}).Where("id", req.ProcessID).Scan(&process)
	if err != nil {
		return nil, err
	}

	// 查流程第一个节点
	node :=entity.ProcessDefines{}
	err= g.Model(entity.ProcessDefines{}).Where("process_id", req.ProcessID).Where("")Scan(&node)
	if err != nil {
		return nil, err
	}

	user := entity.Users{}
	// 查用户
	err = g.Model(entity.Users{}).Where("id", req.UserID).Scan(&user)
	if err != nil {
		return nil, err
	}

	fmt.Println(user.Name, ctx)




	//// 创建task
	//g.Model(entity.Tasks{}).Insert(&entity.Tasks{
	//	StartUserId: user,
	//})
	return
}

func (Process) List(ctx context.Context, req *ListReq) (res *ListRes, err error) {

	return
}

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/try", func(r *ghttp.Request) {
			r.Response.Write(swaggerUIPageContent)
		})
		group.Bind(
			new(Process),
		)
	})

	s.Run()
}
