package processManage

import (
	"context"
	"fmt"
	"gf-workflow/internal/model/entity"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type ProcessManage struct{}

type ListReq struct {
	g.Meta `path:"/manage/list" method:"get" summary:"流程列表" tags:"模型管理"`
}
type ListRes struct {
	Reply string                `dc:"Reply content"`
	Data  []entity.ProcessInfos `json:"data"`
}

type CreateReq struct {
	g.Meta      `path:"/manage/create" method:"post" summary:"新建流程" tags:"模型管理"`
	ProcessName string `v:"required" dc:"流程名称"`
	Version     string `v:"required" dc:"版本号"`
	Comment     string ` dc:"备注"`
}
type CreateRes struct {
	Reply string                `dc:"Reply content"`
	Data  []entity.ProcessInfos `json:"data"`
}

type AddNodeReq struct {
	g.Meta    `path:"/manage/add-node" method:"post" summary:"添加节点" tags:"模型管理"`
	ProcessId string `v:"required" dc:"所属流程ID"`
	NodeName  string `v:"required" dc:"节点名称"`
	NodeInfo  []byte `v:"required" dc:"节点信息"`
	NodeType  string `v:"required" dc:"节点类型"`
}
type AddNodeRes struct {
	Reply string                `dc:"Reply content"`
	Data  []entity.ProcessInfos `json:"data"`
}

func (ProcessManage) List(ctx context.Context, req *ListReq) (res *ListRes, err error) {
	res = &ListRes{}
	err = g.Model(entity.ProcessInfos{}).Scan(&res.Data)
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.WriteJson(res)
	return res, err
}

func (ProcessManage) Create(ctx context.Context, req *CreateReq) (res *CreateRes, err error) {
	processId := createProcess(req.ProcessName, req.Version, req.Comment)

	res = &CreateRes{}
	g.RequestFromCtx(ctx).Response.WriteJson(processId)
	return res, err
}

func (ProcessManage) AddNode(ctx context.Context, req *AddNodeReq) (res *AddNodeRes, err error) {
	addNode(req.ProcessId, req.NodeName, req.NodeType, req.NodeInfo)
	//res = &AddNodeReq{}
	//g.RequestFromCtx(ctx).Response.WriteJson(processId)
	return res, err
}

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

func addNode(processId string, nodeName string, nodeType string, nodeInfo []byte) {
	var preNode entity.ProcessDefines
	g.Model(entity.ProcessDefines{}).Where("process_id", processId).Where("next_id", "").Scan(&preNode)
	// 第一个节点
	if preNode.Id == 0 {
		node := entity.ProcessDefines{
			ProcessId: processId,
			NodeName:  "开始",
			Type:      "start",
			NodeInfo:  gconv.String(nodeInfo),
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
			NodeInfo:  gconv.String(nodeInfo),
		}
		newNodeId, err := g.Model(entity.ProcessDefines{}).InsertAndGetId(&newNode)
		if err != nil {
			fmt.Println(err)
		}
		preNode.NextName = nodeName
		preNode.NextId = gconv.String(newNodeId)
		g.Model(entity.ProcessDefines{}).Where("id", preNode.Id).Save(&preNode)
	}
}
