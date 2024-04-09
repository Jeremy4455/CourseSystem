package models

// Student 学生模型
type Student struct {
	Name      string    `orm:"size(100)"`
	StudentId string    `orm:"unique;size(20);pk"`
	Class     string    `orm:"size(50)"`
	Grade     float64   `orm:"size(20)"`
	Courses   []*Course `orm:"rel(m2m)"` // 学生和课程是多对多关系
}
