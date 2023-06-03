package mw

//
//import (
//	"context"
//	"errors"
//	"github.com/bytedance/gopkg/util/logger"
//	"github.com/lyleshaw/open-plugin/biz/dal/mysql"
//	model "github.com/lyleshaw/open-plugin/biz/model/orm_gen"
//	"github.com/lyleshaw/open-plugin/biz/model/service"
//	"github.com/lyleshaw/open-plugin/pkg/constants"
//	utils2 "github.com/lyleshaw/open-plugin/pkg/utils"
//	"net/http"
//	"time"
//
//	"github.com/cloudwego/hertz/pkg/app"
//	"github.com/cloudwego/hertz/pkg/common/hlog"
//	"github.com/cloudwego/hertz/pkg/common/utils"
//	"github.com/hertz-contrib/jwt"
//)
//
//var (
//	JwtMiddleware *jwt.HertzJWTMiddleware
//	IdentityKey   = "identity"
//)
//
//func InitJwt() {
//	var err error
//	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
//		Realm:         "test zone",
//		Key:           []byte("secret key"),
//		Timeout:       2400 * time.Hour,
//		MaxRefresh:    time.Hour,
//		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
//		TokenHeadName: "Bearer",
//		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
//			resp := service.LoginResp{
//				Code:    int32(code),
//				Message: "success",
//				Data: &service.LoginData{
//					Token:  token,
//					Expire: expire.Format(constants.TimeFormat),
//				},
//			}
//			c.JSON(http.StatusOK, resp)
//		},
//		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
//			var req service.LoginReq
//			err = c.BindAndValidate(&req)
//			if err != nil {
//				c.String(400, err.Error())
//				return nil, err
//			}
//
//			user, err := mysql.GetUserByEmail(ctx, req.Email)
//			if err != nil {
//				return nil, err
//			}
//			if user == nil {
//				return nil, errors.New("user not found")
//			}
//			if !(utils2.VerifyPassword(req.Password, user.Password)) {
//				return nil, errors.New("wrong password")
//			}
//			return user, nil
//		},
//		IdentityKey: IdentityKey,
//		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
//			claims := jwt.ExtractClaims(ctx, c)
//			logger.Infof("claims = %+v", claims)
//			user, _ := mysql.GetUserByEmail(ctx, claims[IdentityKey].(string))
//			return user
//		},
//
//		PayloadFunc: func(data interface{}) jwt.MapClaims {
//			if v, ok := data.(*model.User); ok {
//				return jwt.MapClaims{
//					IdentityKey: v.Email,
//				}
//			}
//			return jwt.MapClaims{}
//		},
//		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
//			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
//			return e.Error()
//		},
//		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
//			c.JSON(http.StatusOK, utils.H{
//				"code":    code,
//				"message": message,
//			})
//		},
//	})
//	if err != nil {
//		panic(err)
//	}
//}
