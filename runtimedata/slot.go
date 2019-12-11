package runtimedata

// 局部变量表的容量以变量槽（Slot）为最小单位，32位虚拟机中一个Slot可以存放32位（4 字节）以内的数据类型
type Slot struct {
	val int32      // 存放int float （double和lang需要两个）
	ref *Reference // 存放引用
}
