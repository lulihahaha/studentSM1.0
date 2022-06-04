package main

import (
	"fmt"
	"studentSM1.0/modle"
)

func main() {
	var (
		choose int
	)
	studentDB := make([]modle.Student, 2, 100)
	studentDB[0] = modle.Student{
		Id:    1,
		Name:  "肥狗",
		Class: "土木1班",
		Age:   22,
		Major: "土木工程",
	}
	studentDB[1] = modle.Student{
		Id:    2,
		Name:  "PJ",
		Class: "大数据2班",
		Age:   18,
		Major: "数据科学",
	}
	fmt.Println("欢迎使用学生管理系统1.0")
	for {
		fmt.Println("++++++++++++++++++++++")
		fmt.Printf("%10s\n", "1.查看学生信息")
		fmt.Printf("%10s\n", "2.新增学生信息")
		fmt.Printf("%10s\n", "3.删除学生信息")
		fmt.Printf("%10s\n", "4.修改学生信息")
		fmt.Printf("%8s\n", "5.退出系统")
		fmt.Println("++++++++++++++++++++++")
		fmt.Print("请选择对应的操作:")
		fmt.Scan(&choose)
		switch choose {
		case 1:
			{
				//fmt.Printf("查询函数查询的studentDB的地址：%p\n", studentDB)
				fmt.Printf("开始执行%s操作...\n", "查看学生信息")
				fmt.Printf("本次查询共查询到%d人\n", len(studentDB))
				fmt.Println("........................")
				for _, v := range studentDB {
					v.ShowStudentInfo()
					fmt.Println("-----------------------------------")
				}
			}
		case 2:
			{
				fmt.Printf("开始执行%s操作...\n", "新增学生信息")
				fmt.Print("请依次输入新增学生的学号、姓名、班级、年龄和专业：")
				var (
					id    int64
					name  string
					class string
					age   int
					major string
				)
				fmt.Scan(&id, &name, &class, &age, &major)
				modle.AddStudentInfo(id, name, class, age, major, &studentDB)
				fmt.Printf("新增学生信息如下：%v\n", studentDB[len(studentDB)-1])
			}
		case 3:
			{
				fmt.Printf("开始执行%s操作...\n", "删除学生信息")
				fmt.Print("请输入要删除的学生的学号：")
				var id int64
				fmt.Scan(&id)
				modle.DeleteStudentInfo(id, &studentDB)
			}
		case 4:
			{
				fmt.Printf("开始执行%s操作...\n", "修改学生信息")
				var id int64
				fmt.Print("请输入需要进行修改操作的学生的学号：")
				fmt.Scan(&id)
				//fmt.Printf("即将进行修改函数的studentDB的地址：%p\n", studentDB)
				modle.ModifyStudentInfo(id, &studentDB)
				//fmt.Printf("修改函数出来后的studentDB的地址：%p\n", studentDB)
			}
		case 5:
			{
				goto breakTag
			}
		default:
			{
				fmt.Println("无效输入，请重新输入！")
			}
		}
	}
breakTag:
	fmt.Println("+++++++++++++++++++++++++++++++")
	fmt.Println("感谢使用学生管理系统1.0!!!")
}
