package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sjmshsh/grpc-gin-admin/project_api/api"
	"github.com/sjmshsh/grpc-gin-admin/project_api/config"
	"github.com/sjmshsh/grpc-gin-admin/project_api/router"
	"github.com/sjmshsh/grpc-gin-admin/project_common"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	config.InitConfig()
	project_common.Run(r, config.C.SC.Name, config.C.SC.Addr, nil)
}
