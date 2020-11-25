package projb

import (
	"fmt"
	"github.com/w3xse7en/proj-a/pkg/proja"
	c "github.com/w3xse7en/third-proj-c"
)

func RpcB(i int) string {
	return fmt.Sprintf("rpc b input %v", c.ToolCItoa(i))
}

func AccessProjA(i int) string {
	return proja.RpcA(i)
}
