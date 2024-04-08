package models

// Student 学生模型
type Student struct {
	Id        int       `orm:"auto;pk"`
	Name      string    `orm:"size(100)"`
	StudentId string    `orm:"unique;size(20)"`
	Class     string    `orm:"size(50)"`
	Contact   string    `orm:"size(50)"`
	Courses   []*Course `orm:"rel(m2m)"` // 学生和课程是多对多关系
}
