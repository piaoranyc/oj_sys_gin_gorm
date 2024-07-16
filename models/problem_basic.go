package models

import (
	"gorm.io/gorm"
)

type ProblemBasic struct {
	gorm.Model
	Identity   string `gorm:"column:identity;type:varchar(36);" json:"identity"` // 问题表的唯一标识
	Title      string `gorm:"column:title;type:varchar(255);" json:"title"`      // 文章标题
	Content    string `gorm:"column:content;type:text;" json:"content"`          // 文章正文
	MaxRuntime int    `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"`
	MaxMemory  int    `gorm:"column:max_mem;type:int(11);" json:"max_memory"`
	CategoryId string `gorm:"column:category_id;type:varchar(255);" json:"category_id"`
}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(keyword string) *gorm.DB {
	return DB.Model(new(ProblemBasic)).
		Where("title like ? OR content like ?", "%"+keyword+"%", "%"+keyword+"%")

}
