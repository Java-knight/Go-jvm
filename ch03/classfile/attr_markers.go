package classfile

// 已弃用（Deprecated）属性，Java5.0开始提供了@Deprecated注解, 表示该 属性/方法 弃用
type DeprecatedAttribute struct {
	MarkerAttribute
}

// Synthetic属性：标记源文件不存在, 由编译器生成的类成员, Synthetic主要为了支持嵌套类和嵌套接口
type SyntheticAttribute struct {
	MarkerAttribute
}

// 标记属性：只有标记作用, 没有任何数据
type MarkerAttribute struct {
}

func (this *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
