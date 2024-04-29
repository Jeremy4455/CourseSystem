package AdminTransactionControllers

import (
	"CourseSystem/controllers"
	"CourseSystem/models"
)

type AdminTransactionControllerCommit struct {
	controllers.BaseController
}

func (c *AdminTransactionControllerCommit) Get() {
	c.TplName = ""
}

// 做几个按钮
// | 开始选课 | 结束选课 | 敲定课程 | 成绩归档 |
// 开始选课返回begin
// 结束选课返回end
// 敲定课程返回done
// 成绩归档返回archive

func (c *AdminTransactionControllerCommit) Post() {
	trans := c.GetString("Trans")

	switch trans {
	case "begin":
		models.ChangeLevel(models.PER_CLASS, models.STUDENT_PICK_DROP_CLASS)
	case "end":
		models.UpgradeFromLevel(models.STUDENT_PICK_DROP_CLASS)
	case "done":
		models.UpgradeFromLevel(models.ADMIN_PICK_DROP_CLASS)
	case "archive":
		models.UpgradeFromLevel(models.ADMIN_UPDATE_GRADE)
	}
}