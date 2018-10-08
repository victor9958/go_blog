package controllers

type WorkersController struct {
	BaseController
}

func (c *WorkersController)Get(){
	c.TplName = "ceshi.html"

}
