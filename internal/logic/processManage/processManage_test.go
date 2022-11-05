package processManage

import (
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func Test_ProcessCreate(t *testing.T) {
	process := createProcess("新生开学报到流程", "v1", "")

	gtest.C(t, func(t *gtest.T) {
		t.Assert(process.Id, 1)
		t.Assert(process.ProcessName, "新生开学报到流程")
	})
}

func Test_AddNode(t *testing.T) {
	node := addNode()
}
