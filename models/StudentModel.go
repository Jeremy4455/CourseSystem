package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

// Student 学生模型
type Student struct {
	StudentId string          `orm:"unique;size(20);pk"`
	Name      string          `orm:"size(100)"`
	Class     string          `orm:"size(50)"`
	Grade     float64         `orm:"size(20)"`
	Classes   []*ClassStudent `orm:"rel(m2m);on_delete(cascade)"`
}

func GetStudent(studentId string) (*Student, error) {
	o := orm.NewOrm()

	student := &Student{StudentId: studentId}

	err := o.Read(student)
	if err != nil {
		// 处理错误
		return nil, err
	}

	return student, nil
}

func CreateStudent(studentId, name, class, grade string) error {
	student, err := GetStudent(studentId)
	if student != nil {
		return err
	}

	g, err := strconv.ParseFloat(grade, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	student = &Student{StudentId: studentId, Name: name, Class: class, Grade: g}
	_, err = o.Insert(student)
	if err != nil {
		return err
	}

	err = AddUser(studentId, name, "123456", "student")
	if err != nil {
		return err
	}
	return nil
}

func DeleteStudent(studentId string) error {
	o := orm.NewOrm()

	student := &Student{StudentId: studentId}
	err := o.Read(student)
	if err != nil {
		return err
	}
	_, err = o.Delete(student)
	if err != nil {
		return err
	}

	err = DeleteUser(studentId)
	if err != nil {
		return err
	}
	return nil
}

func GetAllStudents() ([]*Student, error) {
	o := orm.NewOrm()
	q := o.QueryTable("student")
	var students []*Student
	_, err := q.All(&students)
	return students, err
}

func GetStudents(studentId, name, class string) ([]*Student, error) {
	if studentId == "" && name == "" && class == "" {
		return nil, nil
	}
	// 获取 ORM 对象
	o := orm.NewOrm()

	// 创建查询对象
	var students []*Student

	// 构建查询条件
	qs := o.QueryTable("student")
	if studentId != "" {
		qs = qs.Filter("StudentId", studentId)
	}
	if name != "" {
		qs = qs.Filter("Name__contains", name)
	}
	if class != "" {
		qs = qs.Filter("Class__contains", class)
	}
	// 执行查询
	_, err := qs.All(&students)
	return students, err
}

func ReviseStudent(studentId, name, class string) error {
	students, err := GetStudents(studentId, "", "")
	if err != nil {
		return err
	}
	s := students[0]
	if name != "" {
		s.Name = name
	}
	if class != "" {
		s.Class = class
	}
	_, err = orm.NewOrm().Update(s)
	if err != nil {
		return err
	}
	return nil
}
