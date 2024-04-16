package models

type ClassStudent struct {
	Id          string   `orm:"pk"`
	Class       *Class   `orm:"rel(fk);on_delete(cascade)"`
	Student     *Student `orm:"rel(fk)"`
	Performance float64
	Score       float64
	Grade       string
}

func PickClass(s *Student, c *Class) bool {
	return true
}
func DropClass(s *Student, c *Class) bool {
	return true
}
func UpdateClass() {}
