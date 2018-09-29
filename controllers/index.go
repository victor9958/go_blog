package controllers

type IndexController struct {
	BaseController
}


func (c *IndexController) Get() {
	//str:=strconv.Itoa(66666)
	//c.Ctx.WriteString(str)
	//fmt.Sprintf("任务[%d]上一次执行尚未结束，本次被忽略。", j.id

	///c.Ctx.WriteString(fmt.Sprintf("这是字符串中加[%s]","sfdsfsd"))

	c.TplName = "index.html"
}