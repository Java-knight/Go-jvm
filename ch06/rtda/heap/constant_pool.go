package heap

import (
	"Go-jvm/ch06/classfile"
	"fmt"
)

/*
运行时常量池:
*/

// 常量（字面量和符号引用）
type Constant interface{}

// 运行时常量池
type ConstantPool struct {
	class    *Class
	constant []Constant
}

// 创建运行时常量池
func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	constants := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, constants}

	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			constants[i] = intInfo.Value() // int32
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			constants[i] = floatInfo.Value() // float32
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			constants[i] = longInfo.Value() // int64
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			constants[i] = doubleInfo.Value() // float64
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			constants[i] = stringInfo.String() // string
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			constants[i] = newClassRef(rtCp, classInfo) // 类符号引用
		case *classfile.ConstantFieldrefInfo:
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			constants[i] = newFieldRef(rtCp, fieldrefInfo) // 字段引用
		case *classfile.ConstantMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			constants[i] = newMethodRef(rtCp, methodrefInfo) // 普通方法引用
		case *classfile.ConstantInterfaceMethodrefInfo:
			interfaceMethodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			constants[i] = newInterfaceMethodRef(rtCp, interfaceMethodrefInfo) // 接口方法引用
		default:
			// TODO
		}
	}
	return rtCp
}

// getting constant
func (this *ConstantPool) GetConstant(index uint) Constant {
	if constant := this.constant[index]; constant != nil {
		return constant
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
