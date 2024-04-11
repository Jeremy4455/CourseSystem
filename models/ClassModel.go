package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
)

type Class struct {
	Id        int      `orm:"pk"`
	Course    *Course  `orm:"rel(fk);index"` // Course作为外键
	Teacher   *Teacher `orm:"rel(fk);index"` // Teacher作为外键
	Semester  string   `orm:"index"`
	ClassTime string
	Capacity  int
	Location  string
	Students  []*Student `orm:"rel(m2m)"`
}

// 多字段唯一键
func (c *Class) TableUnique() [][]string {
	return [][]string{
		{"Course", "Teacher", "Semester"},
	}
}

func ExistClass(c *Course, t *Teacher, s string) bool {
	o := orm.NewOrm()
	q := o.QueryTable("class")
	exist := q.Filter("Course", c).Filter("Teacher", t).Filter("Semester", s).Exist()
	return exist
}

func SearchClasses(courseCode, courseName, courseTeacherId, courseSemester, courseTime, classroom string) (*[]Class, error) {
	if courseSemester == "" {
		return nil, nil
	}

	o := orm.NewOrm()
	q := o.QueryTable("class").Filter("Semester", courseSemester)

	if courseCode != "" {
		course, err := GetCourse(courseCode, courseName)
		if err != nil {
			return nil, nil
		}
		q = q.Filter("Course", course)
	}

	if courseTeacherId != "" {
		teacher, err := GetTeacher(courseTeacherId)
		if err != nil {
			return nil, nil
		}
		q = q.Filter("Teacher", teacher)
	}

	if courseTime != "" {
		q = q.Filter("ClassTime__contains", courseTime)
	}

	if classroom != "" {
		q = q.Filter("Location__contains", classroom)
	}

	var classes []Class
	_, err := q.All(&classes)

	if err != nil {
		// 处理错误
		fmt.Println("Error querying courses from database:", err)
	}
	return &classes, nil
}
