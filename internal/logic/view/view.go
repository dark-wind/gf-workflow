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
	g.Model(entity.Tasks{}).Where("id", 3).Scan(&task)
	var data g.Map
	if task.AssigneeRoleName == "招生办" {
		data = g.Map{
			"zsbProcessName":   task.ProcessName,
			"zsbStartUserName": task.StartUserName,
		}
	}
	if task.AssigneeRoleName == "辅导员" {
		data = g.Map{
			"fdyProcessName":   task.ProcessName,
			"fdyStartUserName": task.StartUserName,
		}
	}
	if task.AssigneeRoleName == "财务处" {
		data = g.Map{
			"cwcProcessName":   task.ProcessName,
			"cwcStartUserName": task.StartUserName,
		}
	}
	if task.AssigneeRoleName == "导师" {
		data = g.Map{
			"dsProcessName":   task.ProcessName,
			"dsStartUserName": task.StartUserName,
		}
	}
	if task.AssigneeRoleName == "党支部" {
		data = g.Map{
			"dzbProcessName":   task.ProcessName,
			"dzbStartUserName": task.StartUserName,
		}
	}
	if task.AssigneeRoleName == "团支部" {
		data = g.Map{
			"tzbProcessName":   task.ProcessName,
			"tzbStartUserName": task.StartUserName,
		}
	}
	if task.AssigneeRoleName == "宿管" {
		data = g.Map{
			"sgProcessName":   task.ProcessName,
			"sgStartUserName": task.StartUserName,
		}
	}
	err := r.Response.WriteTpl("task.tpl", data)
	if err != nil {
		fmt.Println(err)
	}
}
