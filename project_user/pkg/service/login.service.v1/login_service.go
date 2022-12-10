package login_service_v1

import (
	"context"
	"github.com/sjmshsh/grpc-gin-admin/project_common"
	"github.com/sjmshsh/grpc-gin-admin/project_common/logs"
	"github.com/sjmshsh/grpc-gin-admin/project_user/pkg/dao"
	model2 "github.com/sjmshsh/grpc-gin-admin/project_user/pkg/model"
	"github.com/sjmshsh/grpc-gin-admin/project_user/pkg/repo"
	"go.uber.org/zap"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

func New() *LoginService {
	return &LoginService{
		cache: dao.Rc,
	}
}

func (ls *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {
	// 1. 获取参数
	mobile := msg.Mobile
	// 2. 校验参数
	if !project_common.VerifyMobile(mobile) {
		return nil, model2.NoLegalMobile
	}
	// 3. 生成验证码(随机4位1000-9999或者6位1000000-99999)
	codeInt := rand.New(rand.NewSource(10)).Int()
	code := strconv.Itoa(codeInt)
	// 4. 调用短信平台(三方 放入go协程中执行 接口可以快速响应)
	go func() {
		time.Sleep(2 * time.Second)
		zap.L().Info("短信平台调用成功, 发送短信 INFO")
		logs.LG.Debug("短信平台调用成功, 发送短信 DEBUG")
		logs.LG.Debug("短信平台调用成功, 发送短信 ERROR")
		// redis 假设后续缓存可能存在MySQL当作，也可能存在mongo当作，也可能是memcache当中
		// 5. 存储验证码 redis 当中 15分钟
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := ls.cache.Put(c, "REGISTER"+mobile, code, 15*time.Minute)
		if err != nil {
			log.Printf("验证码存入redis出错, couse by: %v", err)
		}
		log.Printf("将手机号和验证码存入redis成功: REGISTER_%S : %s", mobile, code)
	}()
	return &CaptchaResponse{Code: code}, nil
}

func (LoginService) mustEmbedUnimplementedLoginServiceServer() {}
