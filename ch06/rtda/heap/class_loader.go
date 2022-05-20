package heap

import (
	"Go-jvm/ch06/classfile"
	"Go-jvm/ch06/classpath"
	"fmt"
)

/*
类加载器

LoadClass()逻辑：
*/

type ClassLoader struct {
	cp       *classpath.Classpath // 搜索和读取class文件
	classMap map[string]*Class    // loaded classes，classMap字段记录已经加载的类数据，key是类的完全限定名。它就充当了方法区的概念
}

// 创建classloader
func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

// 将类数据（xxx.class—>解析—>方法区）加载到方法区
func (this *ClassLoader) LoadClass(name string) *Class {
	if class, ok := this.classMap[name]; ok {
		return class // 类已经加载
	}
	return this.loadNonArrayClass(name) // 加载非数组的class
}

// 加载非数组的class
// 加载 —> 链接 —> 初始化
func (this *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := this.readClass(name)
	class := this.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

// 通过类名从classpath按照固定的解析格式读取到class的byte数组和entry类型
func (this *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := this.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

// 将读取到的data数据转化为固定的class类型（parse）, 并存入classMap（方法区）【每个class类型加载都需要做的】
// resolveSuperClass 和 resolveInterface 字段存放超类名和直接接口名列表，这些类名其实都是符号引用
func (this *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = this
	resolveSuperClass(class)
	resolveInterface(class)
	this.classMap[class.name] = class
	return class
}

// 解析class，将对应的 byte数组 获得classfile的（对应的class格式）
func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

// 每个类的父类都是java.lang.Object。如果当前类名不是Object类就继续调用LoadClass，递归的去加载其父类，直到加载到当前类是Object
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

// 对类的全部接口进行加载
func resolveInterface(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

// 链接(验证、准备、解析)，这里只实现类验证和准备阶段
func link(class *Class) {
	verify(class)
	prepare(class)
}

// 链接的验证阶段
func verify(class *Class) {
	// todo 这里的实现可以参考jvm4.10节
}

// 链接的准备阶段
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

// 给实例变量字段设置编号（计算实例变量个数）
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil { // 递归的去调用，直到 Object 上（顶级父类）
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() { // 非static
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() { // long\double 占用两个slot
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

// 给静态变量字段设置编号（计算静态变量个数）
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

// 给类变量分配内存空间，并赋初值
// （1）go语言会保证新创建的 slot 结构体有默认值（num字段为0, ref字段为nil），而浮点数0 编码之后和整数0相同，
// 所以不用做任何操作就可以保证静态变量有默认初始值（数字类型是0，引用类型是null）。
// （2）如果静态比那里属于基本类型 或 String类型，有final修饰符，且它的值在编译期已知，则该值存储在 class 文件常量池中。
// （3）initStaticFinalVar函数从常量池中加载常量值，然后给静态变量赋值
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

// 从常量池中加载常量值，然后给静态变量赋值
// map.put(slotId, constantPool->val)
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String": // TODO 后续实现
			panic("todo: string后面实现init")
		}
	}
}
