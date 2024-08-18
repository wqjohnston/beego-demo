package controllers

import (
	"fyoukuApi/models"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//获取频道地区列表
// @router /channel/region [*]
func (c *BaseController) ChannelRegion() {
	channelId, _ := c.GetInt("channelId")
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
	}
	num, regions, err := models.GetChannelRegion(channelId)
	if err == nil {
		c.Data["json"] = ReturnSuccess(0, "success", regions, num)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	}
}

//获取频道类型列表
// @router /channel/type [*]
func (c *BaseController) ChannelType() {
	channelId, _ := c.GetInt("channelId")
	if channelId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定频道")
		c.ServeJSON()
	}

	num, types, err := models.GetChannelType(channelId)
	if err == nil {
		c.Data["json"] = ReturnSuccess(0, "success", types, num)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
	}
}
