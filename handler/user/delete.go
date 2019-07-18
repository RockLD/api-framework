package user

import (
	"api/handler"
	"api/model"
	"api/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
