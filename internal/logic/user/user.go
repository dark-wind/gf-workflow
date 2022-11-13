package user

import (
	"context"
	"fmt"
	"gf-workflow/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type CreateUserReq struct {
	g.Meta   `path:"/user/create" method:"post" summary:"添加用户" tags:"用户管理"`
	Name     string `v:"required" dc:"角色名称"`
	Pwd      string `v:"required" dc:"角色名称"`
	RoleId   string `v:"required" dc:"角色名称"`
	RoleName string `v:"required" dc:"角色名称"`
}
type CreateUserRes struct {
	Reply string                `dc:"Reply content"`
	Data  []entity.ProcessInfos `json:"data"`
}

type User struct{}

func (User) CreateUser(ctx context.Context, req *CreateUserReq) (res *CreateUserRes, err error) {
	userId := createUser(req.Name, req.Pwd, req.RoleId, req.RoleName)
	fmt.Println(userId)
	return res, err
}

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
