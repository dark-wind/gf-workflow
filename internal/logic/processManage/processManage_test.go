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

type normalNodeInfo struct {
	RoleId   string
	RoleName string
}

func Test_AddNode(t *testing.T) {

	startNodeInfo := &normalNodeInfo{
		RoleId:   "1",
		RoleName: "招生办",
	}

	startJson, _ := json.Marshal(startNodeInfo)

	xueliNodeInfo := &normalNodeInfo{
		RoleId:   "4",
		RoleName: "导师",
	}
	xueliJson, _ := json.Marshal(xueliNodeInfo)

	ruzhuNodeInfo := &normalNodeInfo{
		RoleId:   "7",
		RoleName: "宿管",
	}
	ruzhuJson, _ := json.Marshal(ruzhuNodeInfo)

	addNode(gconv.String(1), "开始", "start", startJson)
	addNode(gconv.String(1), "入学资格审核", "normal", startJson)
	//addNode(gconv.String(1), "缴费情况审核", "countersign", startJson)
	addNode(gconv.String(1), "学历审核", "normal", xueliJson)
	//addNode(gconv.String(1), "档案接收", "switch", "[{\"Conditions\":\"党员\",\"RoleID\":\"1\",\"NodeName\":\"党支部审批\"},{\"Conditions\":\"非党员\",\"RoleID\":\"3\",\"NodeName\":\"团支部审批\"}]")
	addNode(gconv.String(1), "入住确认", "normal", ruzhuJson)
}
