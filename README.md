# Go-jvm

## 命令行工具开发
进行测试：
```
# user是一个用户名
/Users/user/GolandProjects/Go-jvm/ch01/cmd -version
version 0.0.1

/Users/user/GolandProjects/Go-jvm/ch01/cmd -cp foo/bar/ MyApp arg1 arg2
classpath:foo/bar/ class:MyApp args:[arg1 arg2]

/Users/user/GolandProjects/Go-jvm/ch01/cmd -help
Usage: /Users/lidong.han/GolandProjects/Go-jvm/ch01/cmd [-options] class [args...]
```

## 搜索class文件
Java在写一个Hello World程序时，需要涉及到类加载，HotSpot VM是按照classpath去加载类的，加载顺序

（1）启动类路径（bootstrap classpath）（jre\lib\rt.jar）

（2）扩展类路径（extension classpath）(jre\lib\ext\)

（3）用户类路径（user classpath）(当前目录，可以通过classpath指定)

但是，这里我采用命令行指定类加载路径.

**实现类路径：**
类路径就是一个大的整体，它由启动类路径、扩展类路径和用户类路径三个小路径组成。三个小路径又分别由更小的路径构成。这就是组合模式（composite pattern）

这里是有4中Entry实体
* 第一种就是WildcardEntry: 用来解决用户需要找的class文件路径是一个模糊参数，/usr/load/java/src/*
* 第二种就是CompositeEntry: 用一个Entry切片 将是以".jar\.JAR"结尾的文件存放起来
* 第三种就是ZipEntry: 专门用来解析以".jar\.JAR\.zip\.ZIP"结尾的Entry
* 第四种就是DirEntry: 专门用来解析以路径+className.class的文件读取

可以说，前两种都是为了search class, 后两种都为了parse class.

### 测试
```
# 编译代码
go build    
# 执行程序, 没有输入jre的配置, 回去获取JAVA_HOME中配置的jre                                                    
/Users/user/GolandProjects/Go-jvm/ch02/ch02 java.lang.Object
# 执行程序   
/Users/user/GolandProjects/Go-jvm/ch02/ch02 -Xjre "jre的位置" java.lang.Object

```


