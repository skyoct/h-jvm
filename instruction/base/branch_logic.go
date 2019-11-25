package base

import "h-jvm/runtimedata"

// 当前的pc加上偏移量
func Branch(frame *runtimedata.Frame, offset int) {
	pc := frame.Thread().Pc()
	nextPc := pc + offset
	frame.SetNextPc(nextPc)
}
