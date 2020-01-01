package base

import (
	"fmt"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

func InvokeMethod(invokeFrame *runtimedata.Frame, method *metaspace.Method) {
	thread := invokeFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argSlotCount := int(method.ArgsSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invokeFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
	// 原生方法不支持
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.Class().Name(), method.Name(), method.Descriptor()))
		}
	}
}
