/**
 * @Time: 2020/3/27 17:34
 * @Author: solacowa@gmail.com
 * @File: responsestatus
 * @Software: GoLand
 */

package encode

import (
	"github.com/pkg/errors"
)

type ResStatus string

var ResponseMessage = map[ResStatus]int{
	Invalid:        400,
	InvalidParams:  400,
	ErrParamsPhone: 401,
	ErrBadRoute:    401,
	ErrSystem:      500,
	ErrNotfound:    404,
	ErrLimiter:     429,

	ErrAccountNotFound:    404,
	ErrAccountLogin:       1002,
	ErrAccountLoginIsNull: 1003,
	ErrAccountNotLogin:    501,
	ErrAccountASD:         1004,
	ErrAccountLocked:      1005,
	ErrAuthNotLogin:       501,

	// 系统API
	ErrSysRoleNotfound:     2001,
	ErrSysRoleSave:         2002,
	ErrSysRoleUserNotfound: 2003,
	ErrSysRoleUser:         2004,
	ErrSysRoleUserLen:      2005,
	ErrSysRoleUserDelete:   2006,
	ErrSysUserNotfound:     2007,
}

const (
	// 公共错误信息
	Invalid        ResStatus = "invalid"
	InvalidParams  ResStatus = "请求参数错误"
	ErrNotfound    ResStatus = "不存在"
	ErrBadRoute    ResStatus = "请求路由错误"
	ErrParamsPhone ResStatus = "手机格式不正确"
	ErrLimiter     ResStatus = "太快了,等我一会儿..."

	// 中间件错误信息
	ErrSystem             ResStatus = "系统错误"
	ErrAccountNotLogin    ResStatus = "用户没登录"
	ErrAuthNotLogin       ResStatus = "请先登录"
	ErrAccountLoginIsNull ResStatus = "用户名和密码不能为空"
	ErrAccountLogin       ResStatus = "用户名或密码错误"
	ErrAccountNotFound    ResStatus = "账号不存在"
	ErrAccountASD         ResStatus = "权限验证失败"
	ErrAccountLocked      ResStatus = "用户已被锁定"

	// 系统API
	ErrSysRoleNotfound     ResStatus = "角色不存在"
	ErrSysRoleSave         ResStatus = "角色保证错误"
	ErrSysRoleUserNotfound ResStatus = "角色用户不存在"
	ErrSysRoleUser         ResStatus = "用户角色配置失败"
	ErrSysRoleUserLen      ResStatus = "请选择用户"
	ErrSysRoleUserDelete   ResStatus = "角色删除失败"
	ErrSysUserNotfound     ResStatus = "用户不存在"
)

func (c ResStatus) String() string {
	return string(c)
}

func (c ResStatus) Error() error {
	return errors.New(string(c))
}

func (c ResStatus) Wrap(err error) error {
	return errors.Wrap(err, string(c))
}
