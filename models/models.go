package models

import (
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

// Teacher 老师模型
type Teacher struct {
	Id        int       `orm:"auto;pk"`
	Name      string    `orm:"size(100)"`
	TeacherId string    `orm:"unique;size(20)"`
	Contact   string    `orm:"size(50)"`
	Courses   []*Course `orm:"reverse(many)"` // 一个老师可以教授多门课程
}

// Student 学生模型
type Student struct {
	Id        int       `orm:"auto;pk"`
	Name      string    `orm:"size(100)"`
	StudentId string    `orm:"unique;size(20)"`
	Class     string    `orm:"size(50)"`
	Contact   string    `orm:"size(50)"`
	Courses   []*Course `orm:"rel(m2m)"` // 学生和课程是多对多关系
}

// Course 课程模型
type Course struct {
	Id         int        `orm:"auto;pk"`
	Name       string     `orm:"size(100)"`
	CourseCode string     `orm:"unique;size(20)"`
	Credit     int        // 学分
	StartTime  time.Time  // 上课时间
	Location   string     // 上课地点
	Teacher    *Teacher   `orm:"rel(fk)"`       // 一门课程由一个老师教授
	Students   []*Student `orm:"reverse(many)"` // 多个学生选修一门课程
}

func RegisterDB() {
	// 注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	mysqluser, _ := beego.AppConfig.String("mysql::userName")
	mysqlpass, _ := beego.AppConfig.String("mysql::password")
	mysqlurls, _ := beego.AppConfig.String("mysql::urls")
	mysqldb, _ := beego.AppConfig.String("mysql::databaseName")

	// 设置数据库连接
	orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+")/"+mysqldb+"?charset=utf8")
}
func init() {
	orm.RegisterModel(new(Teacher), new(Student), new(Course))
	RegisterDB()
}
