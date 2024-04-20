package models

import "github.com/astaxie/beego/orm"

// Teacher 老师模型
type Teacher struct {
	TeacherId string   `orm:"unique;size(20);pk"`
	Name      string   `orm:"size(100)"`
	Mobile    string   `orm:"size(100)"`
	Email     string   `orm:"size(100)"`
	Classes   []*Class `orm:"reverse(many);on_delete(cascade)"`
}

func GetTeacher(teacherId, name, mobile, email string) ([]*Teacher, error) {
	o := orm.NewOrm()
	q := o.QueryTable("teacher")

	var teacher []*Teacher
	if teacherId != "" {
		_, err := q.Filter("TeacherId", teacherId).All(&teacher)
		if err != nil {
			return nil, err
		}
	}
	if name != "" {
		_, err := q.Filter("Name", name).All(&teacher)
		if err != nil {
			return nil, err
		}
	}
	if mobile != "" {
		_, err := q.Filter("Mobile", mobile).All(&teacher)
		if err != nil {
			return nil, err
		}
	}
	if email != "" {
		_, err := q.Filter("Email", email).All(&teacher)
		if err != nil {
			return nil, err
		}
	}

	return teacher, nil
}

func GetAllTeachers() ([]*Teacher, error) {
	o := orm.NewOrm()
	q := o.QueryTable("teacher")
	var teachers []*Teacher
	_, err := q.All(&teachers)
	return teachers, err
}

func GetTeachers(teacherId, name, mobile, email string) ([]*Teacher, error) {
	if teacherId == "" && name == "" && mobile == "" && email == "" {
		return nil, nil
	}
	// 获取 ORM 对象
	o := orm.NewOrm()

	// 创建查询对象
	var teachers []*Teacher

	// 构建查询条件
	qs := o.QueryTable("teacher")
	if teacherId != "" {
		qs = qs.Filter("TeacherId", teacherId)
	}
	if name != "" {
		qs = qs.Filter("Name__contains", name)
	}
	if mobile != "" {
		qs = qs.Filter("Mobile", mobile)
	}
	if email != "" {
		qs = qs.Filter("Email", email)
	}
	// 执行查询
	_, err := qs.All(&teachers)
	return teachers, err
}

func AddTeacher(teacherId, name, mobile, email string) error {
	teachers, err := GetTeacher(teacherId, "", "", "")
	if len(teachers) != 0 {
		return err
	}

	o := orm.NewOrm()
	teacher := &Teacher{TeacherId: teacherId, Name: name, Mobile: mobile, Email: email}
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

func ReviseTeacher(t *Teacher, name, mobile, email string) bool {
	if t == nil {
		return false
	}
	if name != "" {
		t.Name = name
	}
	if mobile != "" {
		t.Mobile = mobile
	}
	if email != "" {
		t.Email = email
	}
	_, err := orm.NewOrm().Update(t)
	if err != nil {
		return false
	}
	return true
}
