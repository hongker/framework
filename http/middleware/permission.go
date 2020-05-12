package middleware

import (
	"framework/app"
	"framework/http/response"
	"framework/util/strings"
	"github.com/gin-gonic/gin"
	"net/http"
)

const(
	// 默认角色
	defaultRole = "anonymous"
)

// Permission 权限校验中间件
func Permission(ctx *gin.Context)  {
	// 获取角色
	// TODO 获取方式待定
	role := strings.Default(ctx.Query("role"), defaultRole)

	// 根据角色,路由，请求方法校验权限
	hasPermission, err := app.PermissionManager().EnforceSafe(role, ctx.Request.URL.Path, ctx.Request.Method)

	// 程序异常
	if err != nil {
		// TODO 记录日志
		response.Wrap(ctx).Error(500, "system error")
		ctx.Abort()
	}

	// 没有权限
	if !hasPermission {
		ctx.String(http.StatusUnauthorized, "StatusUnauthorized")
		ctx.Abort()
	}

	ctx.Next()
}
