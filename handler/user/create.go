package user

import (
	"api/handler"
	"api/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	admin2 := c.Param("username")
	log.Infof("URL username : %s", admin2)

	desc := c.Param("desc")
	log.Infof("URL key param desc : %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("HeaderContent-Type : %s", contentType)

	log.Debugf("username is: [%s],password is [%s]", r.Username, r.Password)

	if r.Username == "" {
		handler.SendResponse(
			c,
			errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db:xx.xx.xx.xx")),
			nil,
		)
		return
	}

	if r.Password == "" {
		handler.SendResponse(c, fmt.Errorf("password is empty"), nil)
	}

	rsp := CreateResponse{}
	handler.SendResponse(c, nil, rsp)
}
