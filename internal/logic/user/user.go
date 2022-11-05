package user

import (
	"fmt"
	"gf-workflow/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

func createUser(name string, pwd string, roleId string, roleName string) entity.Users {
	user := entity.Users{
		Name:     name,
		Password: pwd,
		RoleId:   roleId,
		RoleName: roleName,
	}
	insert, err := g.Model(entity.Users{}).Insert(&user)
	if err != nil {
		fmt.Println(insert)
	}
	return user
}
