/**
 * @Time: 2020/4/23 14:12
 * @Author: solacowa@gmail.com
 * @File: repository
 * @Software: GoLand
 */

package repository

import (
	"github.com/go-kit/kit/log"
	redisclient "github.com/icowan/redis-client"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"

	"github.com/kplcloud/kplcloud/src/repository/sysnamespace"
	"github.com/kplcloud/kplcloud/src/repository/syspermission"
	"github.com/kplcloud/kplcloud/src/repository/sysrole"
	"github.com/kplcloud/kplcloud/src/repository/syssetting"
	"github.com/kplcloud/kplcloud/src/repository/sysuser"
)

type Repository interface {
	SysSetting() syssetting.Service
	SysUser() sysuser.Service
	SysNamespace() sysnamespace.Service
	SysRole() sysrole.Service
	SysPermission() syspermission.Service
}

type repository struct {
	sysSetting    syssetting.Service
	sysUser       sysuser.Service
	sysNamespace  sysnamespace.Service
	sysRole       sysrole.Service
	sysPermission syspermission.Service
}

func (r *repository) SysPermission() syspermission.Service {
	return r.sysPermission
}

func (r *repository) SysRole() sysrole.Service {
	return r.sysRole
}

func (r *repository) SysNamespace() sysnamespace.Service {
	return r.sysNamespace
}

func (r *repository) SysUser() sysuser.Service {
	return r.sysUser
}

func (r *repository) SysSetting() syssetting.Service {
	return r.sysSetting
}

func NewRepository(db *gorm.DB, logger log.Logger, traceId string, tracer opentracing.Tracer, redis redisclient.RedisClient) Repository {
	// 平台系统相关仓库
	sysSetting := syssetting.New(db)
	sysSetting = syssetting.NewLogging(logger, traceId)(sysSetting)

	sysNamespace := sysnamespace.New(db)
	sysNamespace = sysnamespace.NewLogging(logger, traceId)(sysNamespace)

	sysUser := sysuser.New(db)
	sysUser = sysuser.NewLogging(logger, traceId)(sysUser)

	sysRole := sysrole.New(db)
	sysRole = sysrole.NewLogging(logger, traceId)(sysRole)

	sysPermission := syspermission.New(db)
	sysPermission = syspermission.NewLogging(logger, traceId)(sysPermission)

	if tracer != nil {
		sysSetting = syssetting.NewTracing(tracer)(sysSetting)
		sysUser = sysuser.NewTracing(tracer)(sysUser)
		sysNamespace = sysnamespace.NewTracing(tracer)(sysNamespace)
		sysRole = sysrole.NewTracing(tracer)(sysRole)
		//sysPermission = sysrole.NewTracing(tracer)(sysPermission)
	}

	if redis != nil {
	}

	return &repository{
		sysSetting:    sysSetting,
		sysUser:       sysUser,
		sysNamespace:  sysNamespace,
		sysRole:       sysRole,
		sysPermission: sysPermission,
	}
}
