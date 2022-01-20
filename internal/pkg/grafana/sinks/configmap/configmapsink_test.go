package configmap

import (
	"fmt"
	"testing"
)

func TestZZZZ(t *testing.T) {
	s := DashboardSink{}
	op, _ := s.makePatchOp("zz", "whatever", `"zzz": yyy""`)
	fmt.Println(string(op))
}
