package main

/**
// （Functional Options）函数式选项模式
// 下面的代码示例，采用函数的实现形式来处理 Student 结构体内部属性的修改
//
// 适用场景：路由处理中间件、系统(各种)风险判定、风险流水的processors、完整流水的processors等流水线式的处理
//
**/

// 我们定义学生（一个聚类根）的基本属性
type Student struct {
	name	string
	sex 	int
	address	string
	number	int32
}

// 需求：如果需要更改这个学生的属性，常规的设计方式是提供一个公共的方法，比如：
//func (student *Student) setName(name string)  {
//	student.name = name
//}

// 但是这样无法适配不同情况的处理，应：

// 定义专门修改 Student 属性的类型
type modifyStudentOption func(*Student)

func setStudentName(name string) modifyStudentOption {
	return func(student *Student) {
		student.name = name
	}
}

func setStudentNameWithPrefix(name string) modifyStudentOption {
	return func(student *Student) {
		student.name = "[A]" + name
	}
}

func FillMessage(student *Student, deals ...modifyStudentOption)  {
	for _, deal := range deals{
		deal(student)
	}
}

// 函数式编程
func main()  {
	// 实例化一个 Student
	student := new(Student)

	FillMessage(student,
		setStudentName("张三"),
		setStudentNameWithPrefix("张三"),
	)
}