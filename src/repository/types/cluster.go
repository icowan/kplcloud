/**
 * @Time : 6/29/21 9:40 AM
 * @Author : solacowa@gmail.com
 * @File : cluster
 * @Software: GoLand
 */

package types

import "time"

type Cluster struct {
	Id        int64      `gorm:"column:id;primary_key" json:"id"`
	Name      string     `gorm:"column:name;index;unique;comment:'集群标识'" json:"name"` // 标识
	Alias     string     `gorm:"column:alias;comment:'别名'" json:"alias"`              // 别名
	Remark    string     `gorm:"column:remark;comment:'备注'" json:"remark"`            // 备注
	Label     string     `gorm:"column:remark;comment:'标签'" json:"label"`             // 标签
	Status    string     `gorm:"column:status;comment:'状态:0未启用,1:正常'" json:"status"`  // 状态
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`                 // 创建时间
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`                 // 更新时间
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (e *Cluster) TableName() string {
	return "cluster"
}
