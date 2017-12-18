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

type PythonApiController struct {
	BaseController
}

func (self *PythonApiController) List() {
	self.display()
}
