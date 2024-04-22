package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

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

func StudentTimeConflict(s *Student, c *Class) bool {
	if s == nil || c == nil {
		return false
	}
	classes := s.Classes
	for _, class := range classes {
		if !CheckTime(class.Class.ClassTime, c.ClassTime) {
			return false
		}
	}
	return true
}

func TeacherTimeConflict(t *Teacher, c *Class) bool {
	if t == nil || c == nil {
		return false
	}
	classes := t.Classes
	for _, class := range classes {
		if !CheckTime(class.ClassTime, c.ClassTime) {
			return false
		}
	}
	return true
}

func ClassConflict(c1, c2 *Class) bool {
	if c1.Location != c2.Location {
		return true
	}
	return CheckTime(c1.ClassTime, c2.ClassTime)
}

func Syncronize() {
	o := orm.NewOrm()

	var students []*Student
	_, err := o.QueryTable("student").All(&students)
	if err != nil {
		return
	}

	for _, student := range students {
		err := o.Read(&User{Id: student.StudentId})
		if err == nil {
			continue
		}
		AddUser(student.StudentId, student.Name, "123456", "student")
	}

	var teachers []*Teacher
	_, err = o.QueryTable("teacher").All(&teachers)
	if err == nil {
		return
	}

	for _, teacher := range teachers {
		err := o.Read(&User{Id: teacher.TeacherId})
		if err != nil {
			continue
		}
		AddUser(teacher.TeacherId, teacher.Name, "teacher123", "teacher")
	}
}
func GetId(table_name string) (int, error) {
	var idv int
	o := orm.NewOrm()
	count, err := o.QueryTable(table_name).Count()
	if err != nil {
		return -1, err
	}
	if count == 0 {
		return 0, nil
	} else {
		var results []orm.Params

		_, err = o.QueryTable(table_name).OrderBy("-id").Limit(1).Values(&results)
		if err != nil {
			return -1, err
		}
		if len(results) > 0 {
			idv = results[0]["Id"].(int)
		}
	}
	return idv, nil
}
