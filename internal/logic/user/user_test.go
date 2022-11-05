package user

import (
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func Test_UserCreate(t *testing.T) {
	//createUser("蔡徐坤","123456","0","练习生")

	roles := []string{
		"招生办",
		"辅导员",
		"财务处",
		"导师",
		"党支部",
		"团支部",
		"宿管",
	}
	usernames := []string{
		"路飞",
		"冯宝宝",
		"范闲",
		"阿贝尔",
		"五条悟",
		"川建国",
		"栗子姨",
	}
	for i, role := range roles {
		createRole(role)
		createUser(usernames[i], "123456", gconv.String(i+1), role)
	}

	gtest.C(t, func(t *gtest.T) {
		//t.Assert(taskRs.AssigneeRoleId, "3")
		//t.Assert(taskRs.NodeName, "团支部审批")
	})
}
