// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProcessInfos is the golang structure for table process_infos.
type ProcessInfos struct {
	Id          uint64      `json:"id"          ` //
	ProcessName string      `json:"processName" ` // 流程名称
	Version     string      `json:"version"     ` // 流程版本
	Commnent    string      `json:"commnent"    ` // 流程备注
	DeletedAt   *gtime.Time `json:"deletedAt"   ` //
	CreatedAt   *gtime.Time `json:"createdAt"   ` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` //
}
