package memory_test

import (
	"gostudy/memory"
	"testing"
)

func TestCalls(t *testing.T) {
	in := [3]uint16{0x2a, 0x1eaf, 0x0}
	//println(in[0], in[1], in[2])
	res := memory.Inc(in)
	//res = memory.Inc2(res)
	//memory.Inc3(1)
	//println(res[0], res[1], res[2])
	memory.IncPtr(&res)
	//println(res[0], res[1], res[2])
}
