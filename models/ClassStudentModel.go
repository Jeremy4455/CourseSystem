package models

type ClassStudent struct {
	Id          string   `orm:"pk"`
	Class       *Class   `orm:"rel(fk);on_delete(cascade)"`
	Student     *Student `orm:"rel(fk)"`
	Performance float64
	Score       float64
	Grade       string
}

func PickClass()   {}
func DropClass()   {}
func UpdateClass() {}
