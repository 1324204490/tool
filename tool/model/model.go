/**
*@Author: Zzy
*@Date: 2019/4/22 15:33
*@Description:
 */
package model

import "tool/tool/timeformat"

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt timeformat.JSONTime
	UpdatedAt timeformat.JSONTime
	DeletedAt *timeformat.JSONTime `sql:"index"`
}
