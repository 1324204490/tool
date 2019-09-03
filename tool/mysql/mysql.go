/**
*@Author: Zzy
*@Date: 2019/9/3 17:58
*@Description:
 */
package mysql

//查询本月代码
//time 时间字段 使用方法为 db.where(QueryThisMonth(created_at))
func QueryThisMonth(time string) string {
	return "date_format(" + time + ", '%Y-%m') = DATE_FORMAT(now(), '%Y-%m')"
}

//查询当天代码
//time 时间字段
func QueryThisDay(time string) string {
	return "date_format(" + time + ", '%Y-%m') = DATE_FORMAT(now(), '%Y-%d')"
}

//查询本周代码
//time 时间字段
func QueryThisWeek(time string) string {
	return "YEARWEEK(date_format(" + time + ",'%Y-%m-%d') - INTERVAL 1 DAY) = YEARWEEK(now() - INTERVAL 1 DAY)"
}
