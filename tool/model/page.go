/**
*@Author: Zzy
*@Date: 2019/4/23 11:08
*@Description:
 */
package model

type Page struct {
	CurrentPage int //当前页码
	PageSize    int //页大小
	PageCount   int //总数
}

/**
用于gorm  MYSQL 分页查询
示例如下
db.Limit(page.PageSize).Offset(model.PageIndex(page))
*/
func PageIndex(page *Page) int {
	if page.CurrentPage == 0 {
		page.CurrentPage = 1
	}
	return (page.CurrentPage - 1) * page.PageSize
}
