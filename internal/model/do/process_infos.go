// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProcessInfos is the golang structure of table process_infos for DAO operations like Where/Data.
type ProcessInfos struct {
	g.Meta      `orm:"table:process_infos, do:true"`
	Id          interface{} //
	ProcessName interface{} // 流程名称
	Version     interface{} // 流程版本
	Commnent    interface{} // 流程备注
	DeletedAt   *gtime.Time //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
