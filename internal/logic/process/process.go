package process

import (
	"context"
	"encoding/json"
	"fmt"
	"gf-workflow/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"strconv"
)

type StartReq struct {
	g.Meta    `path:"/start" method:"post"`
	ProcessID string `v:"required" dc:"流程ID"`
	UserID    string `v:"required" dc:"发起用户ID"`
}
type StartRes struct {
	Reply string `dc:"Reply content"`
}

type CompleteReq struct {
	g.Meta `path:"/complete" method:"post"`
	TaskID string `v:"required" dc:"任务ID"`
	UserID string `v:"required" dc:"完成当前任务的用户ID"`
}
type CompleteRes struct {
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
	process := entity.ProcessInfos{}
	err = g.Model(entity.ProcessInfos{}).Where("id", req.ProcessID).Scan(&process)
	if err != nil {
		return nil, err
	}

	// 查流程第一个节点
	node := entity.ProcessDefines{}
	err = g.Model(entity.ProcessDefines{}).Where("process_id", req.ProcessID).Where("type", "start").Scan(&node)
	if err != nil {
		return nil, err
	}

	user := entity.Users{}
	// 查用户
	err = g.Model(entity.Users{}).Where("id", req.UserID).Scan(&user)
	if err != nil {
		return nil, err
	}

	g.Model("task").Insert(&entity.Tasks{
		Id:                1,
		StartUserId:       strconv.FormatUint(user.Id, 10),
		StartUserName:     user.Name,
		NodeId:            strconv.FormatUint(node.Id, 10),
		NodeName:          node.NodeName,
		ProcessId:         strconv.FormatUint(process.Id, 10),
		ProcessName:       process.ProcessName,
		AssigneeRoleId:    node.NextId,
		AssigneeRoleName:  node.NextName,
		AssigneeRoleCount: 1,
		Status:            "run",
	})
	fmt.Println(user.Name, ctx)
	//// 创建task
	//g.Model(entity.Tasks{}).Insert(&entity.Tasks{
	//	StartUserId: user,
	//})
	return
}

func (Process) List(ctx context.Context, req *ListReq) (res *ListRes, err error) {
	//查用户
	user := entity.Users{}
	err = g.Model(entity.Users{}).Where("name", req.Name).Scan(&user)
	if err != nil {
		return nil, err
	}

	//根据user的role_id查task
	var tasks []entity.Tasks
	err = g.Model(entity.Users{}).Where("assignee_role_id", user.RoleId).Scan(&tasks)
	if err != nil {
		return nil, err
	}

	fmt.Println(tasks)
	return
}

func (Process) Complete(ctx context.Context, req *CompleteReq) (res *CompleteRes, err error) {
	// 查任务
	task := entity.Tasks{}
	err = g.Model(entity.Tasks{}).Where("id", req.TaskID).Scan(&task)
	// 查流程下一个节点
	currentNode := entity.ProcessDefines{}
	err = g.Model(entity.ProcessDefines{}).Where("id", task.NodeId).Scan(&currentNode)

	nextNode := entity.ProcessDefines{}
	err = g.Model(entity.ProcessDefines{}).Where("id", currentNode.NextId).Scan(&nextNode)
	// 更新task
	if nextNode.Type == "normal" {
		NormalMove()
	}
	if nextNode.Type == "countersign" {
		CountersignMove()
	}

	if nextNode.Type == "switch" {
		switchMove(task, nextNode.NodeInfo, nextNode)
	}
	return nil, err
}

func NormalMove() {

}
func CountersignMove() {

}

type switchNode struct {
	Conditions string // 条件
	RoleID     string // 角色id
}

func switchMove(task entity.Tasks, nodeInfoJson string, nextNode entity.ProcessDefines) entity.Tasks {
	var switchList []switchNode
	err := json.Unmarshal(gconv.Bytes(nodeInfoJson), &switchList)
	if err != nil {
		fmt.Println("json解析错误: ", err)
	}

	ifMatch := 0
	for _, node := range switchList {
		if task.Conditions == node.Conditions {
			task.AssigneeRoleId = node.RoleID
			// 角色名称
			task.NodeName = node.NodeName
			ifMatch = 1
			break
		}
	}
	if ifMatch == 0 {
		fmt.Println("流程选择条件错误，没有匹配到对应值: ", err)
	}
	task.NodeId = gconv.String(nextNode.Id)
	g.Model(entity.Tasks{}).Save(&task)

	return task
}
