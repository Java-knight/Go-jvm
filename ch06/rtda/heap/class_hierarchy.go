package heap

/*
class.go文件的副本，为了防止class文件过长，放到这个文件
*/

// 在三种情况下，S 类型的引用值可以赋值给 T 类型：
// （1）S 和 T 是同一类型；
// （2）T 是类且 S 是 T 的子类；
// （3）或者 T 是接口且 S 实现了 T 接口
func (this *Class) isAssignableFrom(other *Class) bool {
	s, t := other, this
	if s == t {
		return true
	}
	if !t.IsInterface() {
		return s.isSubClassOf(t)
	} else {
		return s.isImplements(t)
	}
}

// 类的继承判断
// 判断 S 是否是 T 的子类，实际上也就是判断 T 是否是 S 的（直接或间接）超类。
func (this *Class) isSubClassOf(other *Class) bool {
	for c := this.superClass; c != nil; c = c.superClass { // 递归
		if c == other {
			return true
		}
	}
	return false
}

// 类实现接口判断
// 判断 S 是否实现了 T 接口，就看 S 或 S的（直接或间接）超类是否实现类某个接口 T'，T' 要么是 T，要么是 T 的子接口。
func (this *Class) isImplements(iface *Class) bool {
	for c := this; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// 接口的继承判断
func (this *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range this.interfaces { // 递归
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}
