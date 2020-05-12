package rbac

import (
	"github.com/casbin/casbin"
)


// NewManager 返回一个rbac权限管理器
func NewManager(params ...interface{}) (*casbin.Enforcer, error) {
	authEnforcer, err := casbin.NewEnforcerSafe(params)
	if err != nil {
		return nil, err
	}
	return authEnforcer, nil
}
