package processManage

import (
	"fmt"
	"gf-workflow/internal/model/entity"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func createProcess(processName string, version string, comment string) int64 {
	process := entity.ProcessInfos{
		ProcessName: processName,
		Version:     version,
		Commnent:    comment,
	}
	processID, err := g.Model(entity.ProcessInfos{}).InsertAndGetId(&process)
	if err != nil {
		fmt.Println(err)
	}
	return processID
}

func addNode(processId string, nodeName string, nodeType string, nodeInfo string) {
	var preNode entity.ProcessDefines
	g.Model(entity.ProcessDefines{}).Where("process_id", processId).Where("next_id", "").Scan(&preNode)
	// 第一个节点
	if preNode.Id == 0 {
		node := entity.ProcessDefines{
			ProcessId: processId,
			NodeName:  "开始",
			Type:      "start",
		}
		insert, err := g.Model(entity.ProcessDefines{}).Insert(&node)
		if err != nil {
			fmt.Println(err, insert)
		}
	} else {
		// 不是第一个节点
		newNode := entity.ProcessDefines{
			ProcessId: processId,
			NodeName:  nodeName,
			Type:      nodeType,
			NodeInfo:  nodeInfo,
		}
		newNodeId, err := g.Model(entity.ProcessInfos{}).InsertAndGetId(&newNode)
		if err != nil {
			fmt.Println(err)
		}
		preNode.NextName = nodeName
		preNode.NextId = gconv.String(newNodeId)
		g.Model(entity.ProcessDefines{}).Where("id", preNode.Id).Save(&preNode)
	}
}
