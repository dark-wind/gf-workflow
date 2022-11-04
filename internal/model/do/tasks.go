// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Tasks is the golang structure of table tasks for DAO operations like Where/Data.
type Tasks struct {
	g.Meta            `orm:"table:tasks, do:true"`
	Id                interface{} //
	StartUserId       interface{} // 发起用户id
	StartUserName     interface{} // 发起用户名称
	NodeId            interface{} // 当前节点ID
	NodeName          interface{} // 当前节点名称
	ProcessId         interface{} // 流程ID
	ProcessName       interface{} // 流程名称
	AssigneeRoleId    interface{} // 应当处理该任务的角色
	AssigneeRoleName  interface{} // 应当处理该任务的角色
	AssigneeRoleCount interface{} // 表示当前任务需要多少种角色审批之后才能结束
	ActType           interface{} // 表示任务类型 "or"表示或签，即一个人通过或者驳回就结束，"and"表示会签，要所有人通过就流             转到下一步，如果有一个人驳回那么就跳转到上一步
	Conditions        interface{} // 当前任务的条件值，由外部写入，用于流程的判断分支
	Status            interface{} // 任务的状态：run、pause、finish
	DeletedAt         *gtime.Time //
	CreatedAt         *gtime.Time //
	UpdatedAt         *gtime.Time //
}
