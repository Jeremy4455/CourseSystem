package models

import "github.com/beego/beego/orm"

// Course 课程模型
type Course struct {
	CourseCode string   `orm:"unique;size(20);pk"`
	Name       string   `orm:"size(100)"`
	College    string   `orm:"size(100)"`
	Credit     int      // 学分
	Classes    []*Class `orm:"reverse(many)"`
}

func GetCourse(courseCode, name string) (*Course, error) {
	if courseCode == "" && name == "" {
		return nil, nil
	}

	o := orm.NewOrm()
	var course *Course
	if courseCode == "" {
		course = &Course{Name: name}
	} else if name == "" {
		course = &Course{CourseCode: courseCode}
	} else {
		course = &Course{CourseCode: courseCode, Name: name}
	}

	err := o.Read(course)
	if err != nil {
		// 处理错误
		return nil, err
	}
	return course, nil
}
func AddCourse(courseCode, name, college string, credit int) error {
	course, err := GetCourse(courseCode, name)
	if course != nil {
		return err
	}

	o := orm.NewOrm()
	course = &Course{CourseCode: courseCode, Name: name, College: college, Credit: credit}
	o.Insert(course)
	return nil
}

func DeleteCourse(courseCode string) error {
	o := orm.NewOrm()

	course := &Course{CourseCode: courseCode}
	err := o.Read(course)
	if err != nil {
		return err
	}
	o.Delete(course)
	return nil
}
