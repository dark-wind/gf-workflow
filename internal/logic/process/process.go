package process

import (
	"context"
	"encoding/json"
	"fmt"
	"gf-workflow/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
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

	// 查用户
	user := entity.Users{}
	err = g.Model(entity.Users{}).Where("id", req.UserID).Scan(&user)
	if err != nil {
		return nil, err
	}

	// 创建task
	taskId, err := g.Model(entity.Tasks{}).InsertAndGetId(&entity.Tasks{
		StartUserId:   gconv.String(user.Id),
		StartUserName: user.Name,
		NodeId:        gconv.String(node.Id),
		NodeName:      node.NodeName,
		ProcessId:     gconv.String(process.Id),
		ProcessName:   process.ProcessName,
		//开始节点不需要审批，后面调用一次complete
		//AssigneeRoleId:    "",
		//AssigneeRoleName:  "",
		//AssigneeRoleCount: 1,
		Status: "run",
	})
	if err != nil {
		return nil, err
	}

	//调用一次complete
	complete(gconv.String(taskId))

	res = &StartRes{}
	res.Reply = gconv.String(taskId)
	return res, err
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
	complete(req.TaskID)
	return nil, err
}

func complete(taskId string) {
	// 查任务
	task := entity.Tasks{}
	err := g.Model(entity.Tasks{}).Where("id", taskId).Scan(&task)
	if err != nil {
		return
	}
	// 查流程下一个节点
	currentNode := entity.ProcessDefines{}
	err = g.Model(entity.ProcessDefines{}).Where("id", task.NodeId).Scan(&currentNode)

	nextNode := entity.ProcessDefines{}
	err = g.Model(entity.ProcessDefines{}).Where("id", currentNode.NextId).Scan(&nextNode)
	// 更新task
	if nextNode.Type == "normal" {
		NormalMove(task, nextNode.NodeInfo, nextNode)
	}
	if nextNode.Type == "countersign" {
		CountersignMove()
	}

	if nextNode.Type == "switch" {
		switchMove(task, nextNode.NodeInfo, nextNode)
	}
}

type normalNode struct {
	RoleID     string // 角色id
	RoleName   string // 角色名
}

// 处理普通类型审批节点
func NormalMove(task entity.Tasks, nodeInfoJson string, nextNode entity.ProcessDefines) entity.Tasks{

	task.NodeId = gconv.String(nextNode.Id)
	task.NodeName = nextNode.NodeName

	var normal normalNode
	err := json.Unmarshal(gconv.Bytes(nodeInfoJson), &normal)
	if err != nil {
		fmt.Println("json解析错误: ", err)
	}
	task.AssigneeRoleId = normal.RoleID
	task.AssigneeRoleName = normal.RoleName

	g.Model(entity.Tasks{}).Save(&task)
	return task
}

// 处理会签审批节点
func CountersignMove() {

}

type switchNode struct {
	Conditions string // 条件
	RoleID     string // 角色id
	NodeName   string
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
