package modle

import "fmt"

// Student 学生类型
type Student struct {
	Id    int64
	Name  string
	Class string
	Age   int
	Major string
}
// newStudent 学生类型的构造函数
func newStudent(id int64, name string, class string, age int, major string) *Student {
	return &Student{
		Id:    id,
		Name:  name,
		Class: class,
		Age:   age,
		Major: major,
	}
}

// ShowStudentInfo 查看学生信息
func (s *Student) ShowStudentInfo() {
	fmt.Printf("学生学号：%d\n", s.Id)
	fmt.Printf("学生姓名：%s\n", s.Name)
	fmt.Printf("学生班级：%s\n", s.Class)
	fmt.Printf("学生年龄：%d\n", s.Age)
	fmt.Printf("学生专业：%s\n", s.Major)
}

// AddStudentInfo 新增学生信息
func AddStudentInfo(id int64, name string, class string, age int, major string, studentDB *[]Student) {
	new := newStudent(id, name, class, age, major)
	*studentDB = append(*studentDB, *new)
}

// DeleteStudentInfo 删除学生信息
func DeleteStudentInfo(id int64, studentDB *[]Student) {
	s1 := *studentDB
	for i, v := range s1 {
		if id == v.Id {
			s1 = append(s1[:i], s1[i+1:]...)
			break
		}
	}
}

// ModifyStudentInfo 修改学生信息
func ModifyStudentInfo(id int64, studentDB *[]Student) {
	//fmt.Printf("修改函数接受到的studentDB的地址：%p\n", *studentDB)
	s1 := *studentDB
	fmt.Println("请选择对应要修改的字段值：")
	fmt.Printf("%10s\n", "n->修改姓名")
	fmt.Printf("%10s\n", "c->修改班级")
	fmt.Printf("%10s\n", "a->修改年龄")
	fmt.Printf("%10s\n", "m->修改专业")
	var modify string
	fmt.Print("请选择要修改的字段：")
	fmt.Scan(&modify)
	for i, v := range s1 {
		if v.Id == id {
			switch modify {
			case "n":
				{
					var name string
					fmt.Print("请输入修改后的姓名：")
					fmt.Scan(&name)
					s1[i].Name = name
				}
			case "c":
				{
					var class string
					fmt.Print("请输入修改后的班级：")
					fmt.Scan(&class)
					s1[i].Class = class
				}
			case "a":
				{
					var age int
					fmt.Print("请输入修改后的年龄：")
					fmt.Scan(&age)
					s1[i].Age = age
				}
			case "m":
				{
					var major string
					fmt.Print("请输入修改后的专业：")
					fmt.Scan(&major)
					s1[i].Major = major
				}
			default:
				{
					fmt.Println("无效输入，请重新输入！")
				}
			}
			s1[i].ShowStudentInfo()
			break
		}
	}
	//fmt.Printf("修改函数修改完成后的studentDB的地址：%p\n", s1)
}