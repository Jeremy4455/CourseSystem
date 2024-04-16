package models

import "github.com/astaxie/beego/orm"

// Teacher 老师模型
type Teacher struct {
	TeacherId string   `orm:"unique;size(20);pk"`
	Name      string   `orm:"size(100)"`
	Mobile    string   `orm:"size(100)"`
	Email     string   `orm:"size(100)"`
	Classes   []*Class `orm:"reverse(many)"`
}

func GetTeacher(teacherId string) (*Teacher, error) {
	o := orm.NewOrm()

	teacher := &Teacher{TeacherId: teacherId}

	err := o.Read(teacher)
	if err != nil {
		// 处理错误
		return nil, err
	}

	return teacher, nil
}

func AddTeacher(teacherId, name, mobile, email string) error {
	teacher, err := GetTeacher(teacherId)
	if teacher != nil {
		return err
	}

	o := orm.NewOrm()
	teacher = &Teacher{TeacherId: teacherId, Name: name, Mobile: mobile, Email: email}
	_, err = o.Insert(teacher)
	if err != nil {
		return err
	}

	AddUser(teacherId, name, "teacher123", "teacher")
	return nil
}

func DeleteTeacher(teacherId string) error {
	o := orm.NewOrm()

	teacher := &Teacher{TeacherId: teacherId}
	err := o.Read(teacher)
	if err != nil {
		return err
	}
	_, err = o.Delete(teacher)
	if err != nil {
		return err
	}

	DeleteUser(teacherId)
	return nil
}

func ReviseTeacher() {}
