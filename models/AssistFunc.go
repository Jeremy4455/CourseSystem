package models

func CheckTime(busytime, classtime string) bool {

	return true
}

func StudentTimeConflict(s *Student, c *Class) bool {
	if s == nil || c == nil {
		return false
	}
	classes := s.Classes
	for _, class := range classes {
		if CheckTime(class.Class.ClassTime, c.ClassTime) == false {
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
		if CheckTime(class.ClassTime, c.ClassTime) == false {
			return false
		}
	}
	return true
}
