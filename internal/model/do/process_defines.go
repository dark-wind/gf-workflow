// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ProcessDefines is the golang structure of table process_defines for DAO operations like Where/Data.
type ProcessDefines struct {
	g.Meta     `orm:"table:process_defines, do:true"`
	Id         interface{} //
	ProcessId  interface{} // 流程的id
	NodeName   interface{} // 节点名称
	NextId     interface{} // 下节点id
	NextName   interface{} // 下一节点名称
	RejectId   interface{} // 驳回节点id
	RejectName interface{} // 驳回节点名称
	Type       interface{} // 节点类型
	NodeInfo   interface{} // 节点信息
}
