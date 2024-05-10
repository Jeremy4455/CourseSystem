package models

import (
	"errors"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
)

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func CheckTime(busytime, classtime string) bool {
	var split = func(str string) ([]string, [][2]int) {
		var day []string
		var time [][2]int
		for i, n := 0, len(str); i < n; i++ {
			start := i
			for i < n && (str[i] < '0' || str[i] > '9') {
				i++
			}
			day = append(day, str[start:i])

			var t [2]int
			start = i
			for i < n && str[i] != '-' {
				i++
			}
			s := str[start:i]
			num, _ := strconv.Atoi(s)

			t[0] = num
			i++

			start = i
			for i < n && '0' <= str[i] && str[i] <= '9' {
				i++
			}
			s = str[start:i]
			num, _ = strconv.Atoi(s)
			t[1] = num
			time = append(time, t)

			for i < n && str[i] != ' ' {
				i++
			}
		}
		return day, time
	}
	d1, t1 := split(busytime)
	d2, t2 := split(classtime)
	for i, x := range d1 {
		for j, y := range d2 {
			if x != y {
				continue
			}
			if t2[j][1] < t1[i][0] || t1[i][1] < t2[j][0] {
				continue
			}
			return false
		}
	}
	return true
}

func StudentTimeConflict(classes []*Class, c *Class) bool {
	if classes == nil || c == nil {
		return false
	}
	for _, class := range classes {
		if !CheckTime(class.ClassTime, c.ClassTime) {
			return false
		}
	}
	return true
}

func TeacherTimeConflict(classes []*Class, c *Class) bool {
	if classes == nil || c == nil {
		return false
	}

	for _, class := range classes {
		if !CheckTime(class.ClassTime, c.ClassTime) {
			return false
		}
	}
	return true
}

func ClassConflict(classes []*Class, c *Class) bool {
	if classes == nil || c == nil {
		return false
	}
	for _, class := range classes {
		if class.Id == c.Id {
			continue
		}
		if !CheckTime(class.ClassTime, c.ClassTime) {
			return false
		}
	}
	return true
}

func UpgradeLevel(c *Class) error {
	if c == nil {
		return errors.New("不存在该课程")
	}
	o := orm.NewOrm()
	c.Level++
	if _, err := o.Update(c); err != nil {
		return err
	}
	return nil
}

func UpgradeFromLevel(level int) error {
	var classes []*Class
	o := orm.NewOrm()
	if _, err := o.QueryTable("Class").Filter("Level", level).All(&classes); err != nil {
		return err
	}
	for _, class := range classes {
		class.Level++
		if _, err := o.Update(class); err != nil {
			return err
		}
	}
	return nil
}
func ChangeLevel(start, end int) error {
	var classes []*Class
	o := orm.NewOrm()
	if _, err := o.QueryTable("Class").Filter("Level", start).All(&classes); err != nil {
		return err
	}
	for _, class := range classes {
		class.Level = end
		if _, err := o.Update(class); err != nil {
			return err
		}
	}
	return nil
}

func Upgrade2ReadOnly() error {
	var classes []*Class
	var cs []*ClassStudent
	o := orm.NewOrm()
	if _, err := o.QueryTable("Class").Filter("Level", ADMIN_UPDATE_GRADE).All(&classes); err != nil {
		return err
	}
	if _, err := o.QueryTable("ClassStudent").Filter("Class__in", classes).All(&cs); err != nil {
		return err
	}
	for _, classStudent := range cs {
		classStudent.Student.Grade = (classStudent.Student.Grade/MAX_GRADE*float64(classStudent.Student.Credit) + classStudent.Fscore/100*float64(classStudent.Class.Course.Credit)) / (float64(classStudent.Student.Credit + classStudent.Class.Course.Credit))
		classStudent.Student.Credit += classStudent.Class.Course.Credit
		if _, err := o.Update(classStudent); err != nil {
			return err
		}
	}
	for _, class := range classes {
		class.Level++
		if _, err := o.Update(class); err != nil {
			return err
		}
	}
	return nil
}

func Syncronize() {
	o := orm.NewOrm()

	var students []*Student
	if _, err := o.QueryTable("student").All(&students); err != nil {
		return
	}

	for _, student := range students {
		if err := o.Read(&User{Id: student.StudentId}); err != nil {
			continue
		}
		AddUser(student.StudentId, student.Name, "123456", "student")
	}

	var teachers []*Teacher
	if _, err := o.QueryTable("teacher").All(&teachers); err != nil {
		return
	}

	for _, teacher := range teachers {
		if err := o.Read(&User{Id: teacher.TeacherId}); err != nil {
			continue
		}
		AddUser(teacher.TeacherId, teacher.Name, "teacher123", "teacher")
	}
}

func ExistClass(c *Course, t *Teacher, s string) bool {
	o := orm.NewOrm()
	q := o.QueryTable("class")
	exist := q.Filter("Course", c).Filter("Teacher", t).Filter("Semester", s).Exist()
	return exist
}

func GetPickedCount(c *Class) (int, error) {
	o := orm.NewOrm()
	q := o.QueryTable("ClassStudent").Filter("Class", c)
	count, err := q.Count()
	if err != nil {
		return -1, err
	}
	return int(count), nil
}
func CalculateGrade(grade float64) float64 {
	if grade < 60 {
		return 0.0
	}
	if grade < 63.9 {
		return 1.0
	}
	if grade < 65.9 {
		return 1.5
	}
	if grade < 67.9 {
		return 1.7
	}
	if grade < 71.9 {
		return 2.0
	}
	if grade < 74.9 {
		return 2.3
	}
	if grade < 77.9 {
		return 2.7
	}
	if grade < 81.9 {
		return 3.0
	}
	if grade < 84.9 {
		return 3.3
	}
	if grade < 89.9 {
		return 3.7
	}
	return 4.0
}

func NewSemester(s string) error {
	for _, se := range Semesters {
		if s == se {
			return errors.New("学期已存在")
		}
	}
	Semesters = append(Semesters, s)
	return nil
}
