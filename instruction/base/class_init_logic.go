package base

import (
	"fmt"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

func InitClass(thread *runtimedata.Thread, class *metaspace.Class) {
	class.StartInit() // 设置标记为true
	exeClinit(thread, class)  //执行clinit方法
	initSuperClass(thread, class)
}

// 执行client
func exeClinit(thread *runtimedata.Thread, class *metaspace.Class){
	clinit := class.GetClinitMethod() // 获取clinit方法
	if clinit != nil {
		fmt.Println("[invoke clinit: " + class.Name() + "]")
		newFrame := thread.NewFrame(clinit) // 创建一个帧来执行clinit
		thread.PushFrame(newFrame)
	}
}

// 执行super class init
func initSuperClass(thread *runtimedata.Thread, class *metaspace.Class){
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}