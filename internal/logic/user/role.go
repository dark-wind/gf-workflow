package user

import (
	"fmt"
	"gf-workflow/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

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
