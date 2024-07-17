package models

import (
	"gorm.io/gorm"
)

type SubmitBasic struct {
	gorm.Model
	Identity        string        `gorm:"column:identity;type:varchar(36);" json:"identity"`                     // 问题表的唯一标识
	ProblemIdentity string        `gorm:"column:problem_identity;type:varchar(36);" json:"problem_identity"`     // 问题表的唯一标识
	ProblemBasic    *ProblemBasic `gorm:"foreignKey:problem_identity;references:identity;" json:"problem_basic"` // 关联问题基础表
	UserIdentity    string        `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`           // 用户表的唯一标识
	UserBasic       *UserBasic    `gorm:"foreignKey:user_identity;references:identity;" json:"user_basic"`       // 关联用户基础表
	Path            string        `gorm:"column:path;type:varchar(255);" json:"path"`                            //代码存放路径
	Status          int           `gorm:"column:status;type:tinyint(1);" json:"status"`
}

func (table *SubmitBasic) TableName() string {
	return "submit_basic"
}

func GetSubmitList(problemIdentity string, userIdentity string, status int) *gorm.DB {
	tx := DB.Model(new(SubmitBasic)).Preload("ProblemBasic").Preload("UserBasic")
	if problemIdentity != "" {
		tx.Where("problem_identity=?", problemIdentity)
	}
	if userIdentity != "" {
		tx.Where("user_identity=?", userIdentity)
	}
	if status != 0 {
		tx.Where("status=?", status)
	}
	return tx
}
