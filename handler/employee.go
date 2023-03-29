package handler

import (
	"sesi_8/helper"
	"sesi_8/model"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) CreateEmployee(c *gin.Context) {
	in := model.Employee{}

	err := c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	//call service
	res, err := h.app.CreateEmployee(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}
