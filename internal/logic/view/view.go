package view

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type View struct{}

func (View) Task(r *ghttp.Request) {
	err := r.Response.WriteTpl("task.tpl", g.Map{
		"id":   123,
		"name": "john",
	})
	if err != nil {
		fmt.Println(err)
	}
}
