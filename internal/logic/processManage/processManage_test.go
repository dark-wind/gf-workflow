package processManage

import (
	"encoding/json"
	"github.com/gogf/gf/v2/util/gconv"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

//func Test_ProcessCreate(t *testing.T) {
//	processID := createProcess("新生开学报到流程", "v1", "")
//	var process entity.ProcessInfos
//	g.Model(entity.ProcessInfos{}).Where("id", processID).Scan(&process)
//
//	gtest.C(t, func(t *gtest.T) {
//		t.Assert(process.Id, 1)
//		t.Assert(process.ProcessName, "新生开学报到流程")
//	})
//}

type switchNode struct {
	Conditions string // 条件
	RoleID     string // 角色id
	NodeName   string
}

func Test_AddNode(t *testing.T) {

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
	nodeInfoJson, _ := json.Marshal(&structInfo)
	addNode(gconv.String(1), "开始", "start", nodeInfoJson)
	addNode(gconv.String(1), "入学资格审核", "normal", nodeInfoJson)
	addNode(gconv.String(1), "缴费情况审核", "countersign", nodeInfoJson)
	addNode(gconv.String(1), "学历审核", "normal", nodeInfoJson)
	//addNode(gconv.String(1), "档案接收", "switch", "[{\"Conditions\":\"党员\",\"RoleID\":\"1\",\"NodeName\":\"党支部审批\"},{\"Conditions\":\"非党员\",\"RoleID\":\"3\",\"NodeName\":\"团支部审批\"}]")
	addNode(gconv.String(1), "入住确认", "normal", nodeInfoJson)
}
