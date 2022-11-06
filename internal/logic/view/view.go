package view

import (
	"fmt"
	"gf-workflow/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type View struct{}

func (View) Task(r *ghttp.Request) {
	var task entity.Tasks
	g.Model(entity.Tasks{}).Where("id", 2).Scan(&task)
	err := r.Response.WriteTpl("task.tpl", g.Map{
		"id":   123,
		"name": "john",
	})
	if err != nil {
		fmt.Println(err)
	}
}
