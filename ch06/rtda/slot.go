package rtda

import "Go-jvm/ch06/rtda/heap"

// ｜-----｜-----｜-----｜-----｜
// ｜slot1｜     ｜     ｜     ｜
// ｜--- -｜-----｜-----｜-----｜
// ｜slot2｜     ｜     ｜     ｜
// ｜-----｜-----｜-----｜-----｜

// 局部变量表中的Slot槽位, 因为局部变量表需要用数组来表示（使用slot就可以看作数组的一个元素，方便操作）
type Slot struct {
	num int32        // 基本数据类型（byte\short\char\int\long\float\double\bool）
	ref *heap.Object // 引用(接口类型、类类型、数组类型、null类型)
}
