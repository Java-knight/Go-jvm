package base

import "Go-jvm/ch05/rtda"

// 设置线程的 next程序计数器
func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
