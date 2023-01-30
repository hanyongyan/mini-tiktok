package mw

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	jwtutil "mini_tiktok/pkg/utils"
	"net/http"
	"time"
)

// JwtMiddleware jwt中间件
func JwtMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		//从请求头中获取token
		tokenStr := c.Query("token") //从url参数获取
		fmt.Println(tokenStr)
		//用户不存在
		if tokenStr == "" {
			tokenStr = c.PostForm("token") //body获取
			if tokenStr == "" {
				c.JSON(http.StatusOK, utils.H{"code": 0, "msg": "用户不存在"})
				c.Abort() //阻止执行
				return
			}
		}
		//token格式错误
		//tokenSlice := strings.SplitN(tokenStr, " ", 2)
		//if len(tokenSlice) != 2 && tokenSlice[0] != "Bearer" {
		//	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "token格式错误"})
		//	c.Abort() //阻止执行
		//	return
		//}
		//验证token
		tokenStruck, ok := jwtutil.CheckToken(tokenStr)
		if !ok {
			c.JSON(http.StatusOK, utils.H{"code": 0, "msg": "token不正确"})
			c.Abort() //阻止执行
			return
		}
		//token超时
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			c.JSON(http.StatusOK, utils.H{"code": 0, "msg": "token过期"})
			c.Abort() //阻止执行
			return
		}
		fmt.Println("jwt校验正确，允许通过----------")
		c.Next(ctx)
	}
}
