package processManage

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

type switchNode struct {
	Conditions string // 条件
	RoleID     string // 角色id
	NodeName   string
}

type countersignNode struct {
	RoleID   string
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

	switchNodeInfo := []switchNode{
		{
			Conditions: "党员",
			RoleID:     "5",
			NodeName:   "党支部审批",
		},
		{
			Conditions: "非党员",
			RoleID:     "6",
			NodeName:   "团支部审批",
		},
	}
	switchNodeJson, _ := json.Marshal(&switchNodeInfo)

	conterSignNodeInfo := []countersignNode{
		{
			RoleID:   "2",
			RoleName: "辅导员",
		},
		{
			RoleID:   "3",
			RoleName: "财务处",
		},
	}
	conterSignJson, _ := json.Marshal(&conterSignNodeInfo)
	fmt.Println(conterSignJson)

	addNode(gconv.String(1), "开始", "start", startJson)
	addNode(gconv.String(1), "入学资格审核", "normal", startJson)
	//addNode(gconv.String(1), "缴费情况审核", "countersign", conterSignJson)
	addNode(gconv.String(1), "学历审核", "normal", xueliJson)
	addNode(gconv.String(1), "档案接收", "switch", switchNodeJson)
	addNode(gconv.String(1), "入住确认", "normal", ruzhuJson)

	var lastNode entity.ProcessDefines
	g.Model(entity.ProcessDefines{}).Where("process_id", "1").Where("next_id", "").Scan(&lastNode)

	gtest.C(t, func(t *gtest.T) {
		t.Assert(lastNode.NodeName, "入住确认")
	})
}
