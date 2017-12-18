/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-08 17:48:30
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-09 18:50:41
***********************************************/
package controllers

// import (
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/astaxie/beego"
// 	"github.com/api-admin/models"
// )

type GoApiController struct {
	BaseController
}

// func (c *GoApiController) URLMapping() {
// 	c.Mapping("List", c.List)
// }

// // @router /goapi/list [post,get]
func (self *GoApiController) List() {
	self.display()
}
