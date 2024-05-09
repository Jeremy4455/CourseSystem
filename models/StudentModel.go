package models

import (
	"errors"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
)

// Student 学生模型
type Student struct {
	StudentId string  `orm:"unique;size(20);pk"`
	Name      string  `orm:"size(100)"`
	Class     string  `orm:"size(50)"`
	Grade     float64 `orm:"size(20)"`
	Credit    int
	Classes   []*ClassStudent `orm:"reverse(many);on_delete(cascade)"`
}

func GetStudent(studentId string) (*Student, error) {
	o := orm.NewOrm()
	sql := "CALL GetStudentById(?)"
	var s []*Student

	_, err := o.Raw(sql, studentId).QueryRows(&s)
	if err != nil {
		return nil, err
	}
	if len(s) == 0 {
		return nil, errors.New("没有该学生")
	}
	return s[0], nil
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
	student = &Student{
		StudentId: studentId,
		Name:      name,
		Class:     class,
		Grade:     g,
		Credit:    0,
	}
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
	//if studentId == "" && name == "" && class == "" {
	//	return nil, nil
	//}
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
