package models

// Course 课程模型
type Course struct {
	Id         int        `orm:"auto;pk"`
	Name       string     `orm:"size(100)"`
	CourseCode string     `orm:"unique;size(20)"`
	Credit     int        // 学分
	StartTime  string     // 上课时间
	Location   string     // 上课地点
	Teacher    *Teacher   `orm:"rel(fk)"`       // 一门课程由一个老师教授
	Students   []*Student `orm:"reverse(many)"` // 多个学生选修一门课程
}
