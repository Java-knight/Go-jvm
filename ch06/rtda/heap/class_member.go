package heap

import "Go-jvm/ch06/classfile"

/*
类成员：字段和方法都属于类的成员，它们有一些相同都信息（访问标志、名字、描述符）
*/

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class // 通过字段和方法访问到所属的类
}

// 从class文件中复制数据
func (this *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	this.accessFlags = memberInfo.AccessFlags()
	this.name = memberInfo.Name()
	this.descriptor = memberInfo.Descriptor()
}

// check vars(public\private\protected\static\final\synthetic)
func (this *ClassMember) IsPublic() bool {
	return 0 != this.accessFlags&ACC_PUBLIC
}
func (this *ClassMember) IsPrivate() bool {
	return 0 != this.accessFlags&ACC_PRIVATE
}
func (this *ClassMember) IsProtected() bool {
	return 0 != this.accessFlags&ACC_PROTECTED
}
func (this *ClassMember) IsStatic() bool {
	return 0 != this.accessFlags&ACC_STATIC
}
func (this *ClassMember) IsFinal() bool {
	return 0 != this.accessFlags&ACC_FINAL
}

func (this *ClassMember) IsSynthetic() bool {
	return 0 != this.accessFlags&ACC_SYNTHETIC
}

// getter name\descriptor\class
func (this *ClassMember) Name() string {
	return this.name
}
func (this *ClassMember) Descriptor() string {
	return this.descriptor
}
func (this *ClassMember) Class() *Class {
	return this.class
}

// 权限访问限定的check jvms 5.4.4
// 传入一个 类d 是都可以访问本类c
func (this *ClassMember) isAccessibleTo(d *Class) bool {
	if this.IsPublic() { // public
		return true
	}
	c := this.class
	if this.IsProtected() { // protected
		return d == c || d.isSubClassOf(c) || c.getPackageName() == d.getPackageName()
	}
	if !this.IsPrivate() { // 没有写访问限定修饰符，同一个文件的类，但是不是public
		return c.getPackageName() == d.getPackageName()
	}
	return d == c // private
}
