package process

import (
	"encoding/json"
	"fmt"
	"gf-workflow/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func Test_SwitchMove(t *testing.T) {
	// mock task nodeInfo nextNode
	task := entity.Tasks{}
	err := g.Model(entity.Tasks{}).Where("id", 1).Scan(&task)
	if err != nil {
		fmt.Println(err)
	}
	nextNode := entity.ProcessDefines{
		NodeName: "",
		Id:       5,
	}
	structInfo := []switchNode{
		{
			Conditions: "党员",
			RoleID:     "1",
			NodeName:   "党支部审批",
		},
		{
			Conditions: "非党员",
			RoleID:     "3",
			NodeName:   "团支部审批",
		},
	}
	nodeInfoJson, err := json.Marshal(&structInfo)
	if err != nil {
		return
	}
	taskRs := switchMove(task, gconv.String(nodeInfoJson), nextNode)
	gtest.C(t, func(t *gtest.T) {
		t.Assert(taskRs.AssigneeRoleId, "3")
		t.Assert(taskRs.NodeName, "团支部审批")
	})
}
