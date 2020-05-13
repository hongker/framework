// rbac 权限管理组件
// 基于https://github.com/casbin/casbin集成的rbac组件，做到基于角色的权限控制
package rbac

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
)

// rbacModel 定义基于角色的rbac模型
const rbacModel =`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`

// NewManagerWithAdapter 基于自定义适配器的权限管理器
func NewManagerWithAdapter(adapter interface{}) (*casbin.Enforcer, error) {
	m, err := model.NewModelFromString(rbacModel)
	if err != nil {
		return nil, err
	}
	authEnforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		return nil, err
	}
	return authEnforcer, nil
}
