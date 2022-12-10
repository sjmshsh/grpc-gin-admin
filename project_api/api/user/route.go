package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sjmshsh/grpc-gin-admin/project_api/router"
	"log"
)

func init() {
	log.Println("init user router")
	router.Register(&RouterUser{})
}

type RouterUser struct {
}

func (*RouterUser) Router(r *gin.Engine) {
	// 初始化grpc客户端连接
	InitRpcUserClient()
	h := New()
	r.POST("/login", h.getCaptcha)
}
