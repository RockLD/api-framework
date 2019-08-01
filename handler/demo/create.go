package demo

import (
	"apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/util"
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

	if err := r.checkParam(); err != nil {
		handler.SendResponse(c, err, "2")
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
		handler.SendResponse(c, errno.ErrDatabase, err)
		return
	}

	rsp := CreateResponse{Username: r.Username}
	handler.SendResponse(c, nil, rsp)
}

func (r *CreateRequest) checkParam() error {
	if r.Username == "" {
		return errno.New(errno.ErrValidation, nil).Add("username is empty.")
	}

	if r.Password == "" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty.")
	}
	return nil
}
