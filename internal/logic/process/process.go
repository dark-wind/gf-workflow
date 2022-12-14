package process

import (
	"context"
	"encoding/json"
	"fmt"
	"gf-workflow/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

type StartReq struct {
	g.Meta     `path:"/start" method:"post" summary:"发起流程" tags:"流程控制"`
	ProcessID  string `v:"required" dc:"流程ID"`
	UserID     string `v:"required" dc:"发起用户ID"`
	Conditions string `v:"required" dc:"用于分支判断的值"`
}
type StartRes struct {
	Reply string `dc:"Reply content"`
}

type CompleteReq struct {
	g.Meta `path:"/complete" method:"post" summary:"完成当前任务" tags:"流程控制"`
	TaskID string `v:"required" dc:"任务ID"`
	UserID string `v:"required" dc:"完成当前任务的用户ID"`
}
type CompleteRes struct {
	Reply string `dc:"Reply content"`
}

type ListReq struct {
	g.Meta `path:"/list" method:"get" summary:"工作列表" tags:"流程控制"`
	UserID string `v:"required" dc:"当前用户ID"`
}
type ListRes struct {
	Reply string         `dc:"Reply content"`
	Data  []entity.Tasks `json:"data"`
}

type RejectReq struct {
	g.Meta `path:"/reject" method:"post" summary:"驳回" tags:"流程控制"`
	TaskID string `v:"required" dc:"任务ID"`
}
type RejectRes struct {
	Reply string `dc:"Reply content"`
}

type DispatchReq struct {
	g.Meta `path:"/dispatch" method:"post" summary:"重新指派候选人角色" tags:"流程控制"`
	TaskID string `v:"required" dc:"任务ID"`
	RoleID string `v:"required" dc:"角色ID"`
}
type DispatchRes struct {
	Reply string `dc:"Reply content"`
}

type UpdateReq struct {
	g.Meta     `path:"/update" method:"post" summary:"修改任务的相关数据" tags:"流程控制"`
	TaskID     string `v:"required" dc:"任务ID"`
	Conditions string `v:"required" dc:"相关数据"`
}
type UpdateRes struct {
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
		Conditions:    req.Conditions,
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
	complete(gconv.String(taskId), gconv.String(user.Id))

	res = &StartRes{}
	res.Reply = gconv.String(taskId)
	return res, err
}

func (Process) List(ctx context.Context, req *ListReq) (res *ListRes, err error) {
	//查用户
	user := entity.Users{}
	err = g.Model(entity.Users{}).Where("id", req.UserID).Scan(&user)
	if err != nil {
		return nil, err
	}

	res = &ListRes{}
	//根据user的role_id查task
	//err = g.Model(entity.Tasks{}).Where("assignee_role_id", user.RoleId).Where("status", "run").Scan(&res.Data)
	g.Model(entity.Tasks{}).WhereLike("assignee_role_name", "%"+user.RoleName+"%").Where("status", "run").Scan(&res.Data)
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.WriteJson(res)
	return res, err
}

func (Process) Complete(ctx context.Context, req *CompleteReq) (res *CompleteRes, err error) {
	complete(req.TaskID, req.UserID)
	return nil, err
}

func (Process) Reject(ctx context.Context, req *RejectReq) (res *RejectRes, err error) {
	// 查任务
	task := entity.Tasks{}
	err = g.Model(entity.Tasks{}).Where("id", req.TaskID).Scan(&task)
	if err != nil {
		return
	}
	task.Status = "fail"
	task.AssigneeRoleName = ""
	task.AssigneeRoleId = ""
	task.NodeId = ""
	task.NodeName = "被驳回，流程结束"
	result, err := g.Model(entity.Tasks{}).Save(&task)
	if err != nil {
		fmt.Println(result, err)
		return nil, err
	}

	res = &RejectRes{}
	g.RequestFromCtx(ctx).Response.WriteJson(res)
	return res, err
}
func (Process) Dispatch(ctx context.Context, req *DispatchReq) (res *DispatchRes, err error) {
	task := entity.Tasks{}
	err = g.Model(entity.Tasks{}).Where("id", req.TaskID).Scan(&task)
	if err != nil {
		return
	}

	role := entity.Roles{}
	err = g.Model(entity.Roles{}).Where("id", req.RoleID).Scan(&role)

	task.AssigneeRoleId = gconv.String(role.Id)
	task.AssigneeRoleName = role.Name

	result, err := g.Model(entity.Tasks{}).Save(&task)
	if err != nil {
		fmt.Println(result, err)
		return nil, err
	}

	res = &DispatchRes{}
	g.RequestFromCtx(ctx).Response.WriteJson(res)
	return res, err
}

func (Process) UpdateTask(ctx context.Context, req *UpdateReq) (res *UpdateRes, err error) {
	task := entity.Tasks{}
	err = g.Model(entity.Tasks{}).Where("id", req.TaskID).Scan(&task)
	if err != nil {
		return
	}

	task.Conditions = req.Conditions
	result, err := g.Model(entity.Tasks{}).Save(&task)
	if err != nil {
		fmt.Println(result, err)
		return nil, err
	}
	res = &UpdateRes{}
	g.RequestFromCtx(ctx).Response.WriteJson(res)
	return res, err
}

func complete(taskId string, userId string) {
	// 查任务
	task := entity.Tasks{}
	err := g.Model(entity.Tasks{}).Where("id", taskId).Scan(&task)
	if err != nil {
		return
	}

	user := entity.Users{}
	err = g.Model(entity.Users{}).Where("id", userId).Scan(&user)
	if err != nil {
		return
	}
	// 查流程当前节点
	currentNode := entity.ProcessDefines{}
	err = g.Model(entity.ProcessDefines{}).Where("id", task.NodeId).Scan(&currentNode)
	if currentNode.Type == "countersign" {
		// 先审批，后减少，最后往下
		roleIds := strings.Split(task.AssigneeRoleId, ",")
		roleNames := strings.Split(task.AssigneeRoleName, ",")
		var newRoleIds []string
		for _, roleId := range roleIds {
			if user.RoleId == roleId {
				continue
			}
			newRoleIds = append(newRoleIds, roleId)
		}
		var newRoleNames []string
		for _, roleName := range roleNames {
			if user.RoleName == roleName {
				continue
			}
			newRoleNames = append(newRoleNames, roleName)
		}
		task.AssigneeRoleId = strings.Join(newRoleIds, ",")
		task.AssigneeRoleName = strings.Join(newRoleNames, ",")
		// 如果所有角色都审批了，可以推下一步，否则返回
		if len(newRoleNames) != 0 {
			g.Model(entity.Tasks{}).Save(&task)
			return
		}
	}

	// 查流程下一个节点
	nextNode := entity.ProcessDefines{}
	err = g.Model(entity.ProcessDefines{}).Where("id", currentNode.NextId).Scan(&nextNode)

	// 更新task
	if nextNode.NextId == "" {
		//最后一个节点
		task.Status = "finish"
		task.AssigneeRoleId = ""
		task.AssigneeRoleName = ""
		// 角色名称
		task.NodeName = "结束"
		g.Model(entity.Tasks{}).Save(&task)
	} else {
		//正常节点
		if nextNode.Type == "normal" {
			NormalMove(task, nextNode.NodeInfo, nextNode)
		}
		if nextNode.Type == "countersign" {
			CountersignMove(task, nextNode.NodeInfo, nextNode)
		}
		if nextNode.Type == "switch" {
			switchMove(task, nextNode.NodeInfo, nextNode)
		}
	}

}

type normalNode struct {
	RoleID   string // 角色id
	RoleName string // 角色名
}

type Countersign struct {
	RoleID      string // 角色id
	RoleName    string // 角色名
	IsCompleted string
}

// 处理普通类型审批节点
func NormalMove(task entity.Tasks, nodeInfoJson string, nextNode entity.ProcessDefines) entity.Tasks {

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
func CountersignMove(task entity.Tasks, nodeInfoJson string, nextNode entity.ProcessDefines) entity.Tasks {
	var CountersignList []normalNode
	err := json.Unmarshal(gconv.Bytes(nodeInfoJson), &CountersignList)
	if err != nil {
		fmt.Println("json解析错误: ", err)
	}

	var roleIds []string
	var roleNames []string
	count := 0
	for _, node := range CountersignList {
		roleIds = append(roleIds, node.RoleID)
		roleNames = append(roleNames, node.RoleName)
		count++
	}
	task.AssigneeRoleId = strings.Join(roleIds, ",")
	task.AssigneeRoleName = strings.Join(roleNames, ",")

	task.NodeId = gconv.String(nextNode.Id)
	task.NodeName = nextNode.NodeName
	g.Model(entity.Tasks{}).Save(&task)

	return task
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
			task.AssigneeRoleName = node.NodeName
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
