package user

import (
	"api/handler"
	"api/model"
	"api/pkg/errno"
	"api/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

/**
创建用户
1、从HTTP消息体获取参数
2、参数校验
3、加密密码
4、在数据库中添加数据记录
5、返回结果
*/
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{Username: r.Username}
	handler.SendResponse(c, nil, rsp)
}

func (r *CreateRequest) checkParam() error {
	return nil
}
