package user

import (
	"fmt"
	"gf-workflow/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

func createUser(name string, pwd string, roleId string, roleName string) int64 {
	user := entity.Users{
		Name:     name,
		Password: pwd,
		RoleId:   roleId,
		RoleName: roleName,
	}
	userId, err := g.Model(entity.Users{}).InsertAndGetId(&user)
	if err != nil {
		fmt.Println(err)
	}
	return userId
}
