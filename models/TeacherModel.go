package models

// Teacher 老师模型
type Teacher struct {
	Name      string    `orm:"size(100)"`
	TeacherId string    `orm:"unique;size(20);pk"`
	Mobile    string    `orm:"size(100)"`
	Email     string    `orm:"size(100)"`
	Courses   []*Course `orm:"reverse(many)"` // 一个老师可以教授多门课程
}
