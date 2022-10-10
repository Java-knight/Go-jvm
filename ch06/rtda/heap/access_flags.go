package heap

/*
具体的访问标识符，Java虚拟机规范 规定好的
*/

const (
	ACC_PUBLIC       = 0x0001 // class field method, 公共标识符号
	ACC_PRIVATE      = 0x0002 // ----- field method, 私有标识符号
	ACC_PROTECTED    = 0x0004 // ----- field method, 保护标识符号
	ACC_STATIC       = 0x0008 // ----- field method, 静态标识符号
	ACC_FINAL        = 0x0010 // class field method, 常量标识符号
	ACC_SUPER        = 0x0020 // class ----- ------, 继承extends
	ACC_SYNCHRONIZED = 0x0020 // ----- ----- method, 同步synchronized（可见性、原子性、有序性）
	ACC_VOLATILE     = 0x0040 // ----- field ------, volatile关键字（可见性、有序性、不能保证原子性）
	ACC_BRIDGE       = 0x0040 // ----- ----- method, 方法的桥接（比如子类重写父类的方法）
	ACC_TRANSIENT    = 0x0080 // ----- field ------,
	ACC_VARARGS      = 0x0080 // ----- ----- method, 变长参数（方法中的入参...）
	ACC_NATIVE       = 0x0100 // ----- ----- method, 标识native method
	ACC_INTERFACE    = 0x0200 // class ----- ------, 表示接口（特殊的类）
	ACC_ABSTRACT     = 0x0400 // class ----- ------, 抽象类
	ACC_STRICT       = 0x0800 // ----- ----- method
	ACC_SYNTHETIC    = 0x1000 // class field method, 标识（一般出现在内部类）
	ACC_ANNOTATION   = 0x2000 // class ----- ------, 注解
	ACC_ENUM         = 0x4000 // class field ------, 枚举类型

)
