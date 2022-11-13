package user

import (
	"context"
	"fmt"
	"gf-workflow/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type CreateRoleReq struct {
	g.Meta   `path:"/role/create" method:"post" summary:"添加角色" tags:"用户管理"`
	RoleName string `v:"required" dc:"角色名称"`
}
type CreateRoleRes struct {
	Reply string                `dc:"Reply content"`
	Data  []entity.ProcessInfos `json:"data"`
}

type Role struct{}

func (Role) CreateRole(ctx context.Context, req *CreateRoleReq) (res *CreateRoleRes, err error) {
	roleID := createRole(req.RoleName)
	fmt.Println(roleID)
	return res, err
}

func createRole(roleName string) int64 {
	roleId, err := g.Model(entity.Roles{}).InsertAndGetId(&entity.Roles{
		Name: roleName,
	})
	if err != nil {
		fmt.Println(err)
	}
	return roleId
}

func findRoleById(id int) entity.Roles {
	var role entity.Roles
	err := g.Model(entity.Roles{}).Where("id", id).Scan(&role)
	if err != nil {
		fmt.Println(err)
	}
	return role
}
