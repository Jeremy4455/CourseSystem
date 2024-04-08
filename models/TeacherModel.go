package models

// Teacher 老师模型
type Teacher struct {
	Id        int       `orm:"auto;pk"`
	Name      string    `orm:"size(100)"`
	TeacherId string    `orm:"unique;size(20)"`
	Contact   string    `orm:"size(50)"`
	Courses   []*Course `orm:"reverse(many)"` // 一个老师可以教授多门课程
}
