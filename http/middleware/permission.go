package middleware

import (
	"github.com/hongker/framework/app"
	"github.com/hongker/framework/http/response"
	"github.com/hongker/framework/util/strings"
	"github.com/gin-gonic/gin"
	"net/http"
)

const(
	// 默认角色
	defaultRole = "1"
)

// Permission 权限校验中间件
func Permission(ctx *gin.Context)  {
	// 获取UID
	// TODO 获取方式待定
	id := strings.Default(ctx.Query("id"), defaultRole)

	// 根据用户ID,路由，请求方法校验权限
	hasPermission, err := app.PermissionManager().Enforce(id, ctx.Request.URL.Path, ctx.Request.Method)

	// 程序异常
	if err != nil {
		// TODO 记录日志
		response.Wrap(ctx).Error(500, "system error")
		ctx.Abort()
	}

	// 没有权限
	if !hasPermission {
		response.Wrap(ctx).Error(http.StatusUnauthorized, "StatusUnauthorized")
		ctx.Abort()
	}

	ctx.Next()
}

