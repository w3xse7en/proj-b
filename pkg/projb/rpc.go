package projb

import (
	"fmt"
	c "github.com/w3xse7en/third-proj-c"
)

func RpcB(i int) string {
	return fmt.Sprintf("rpc b input %v", c.ToolCItoa(i))
}
