package view

import (
	"fmt"
	"gf-workflow/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type View struct{}

func (View) Task(r *ghttp.Request) {
	var task entity.Tasks
	g.Model(entity.Tasks{}).Where("process_name", "新生开学报到流程").Scan(&task)

	var processList []entity.ProcessInfos
	g.Model(entity.ProcessInfos{}).Scan(&processList)

	list := map[uint64]string{}

	for _, infos := range processList {
		list[infos.Id] = infos.ProcessName
	}

	data := g.Map{}
	if task.AssigneeRoleName == "招生办" {
		data = g.Map{
			"zsbProcessName":   task.ProcessName,
			"zsbStartUserName": task.StartUserName,
		}
	}
	if task.AssigneeRoleName == "辅导员,财务处" {
		data = g.Map{
			"fdyProcessName":   task.ProcessName,
			"fdyStartUserName": task.StartUserName,
			"cwcProcessName":   task.ProcessName,
			"cwcStartUserName": task.StartUserName,
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

	data["taskId"] = gconv.String(task.Id)
	data["list"] = list

	err := r.Response.WriteTpl("task.tpl", data)
	if err != nil {
		fmt.Println(err)
	}
}
