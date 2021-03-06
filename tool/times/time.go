/**
*@Author: Zzy
*@Date: 2019/4/22 16:03
*@Description:
 */
package times

import (
	"strconv"
	"time"
)

//获取系统当前时间,用于time类型的字段赋值
func GetTimeNow() *time.Time {
	currentTimeData := time.Date(time.Now().Year(), time.Now().Month(),
		time.Now().Day(), time.Now().Hour(), time.Now().Minute(),
		time.Now().Second(), time.Now().Nanosecond(), time.Local)
	return &currentTimeData
}

//获取当前时间戳
func GetTimeUnix() string {
	timeUnix := time.Now().UnixNano()
	i := int64(timeUnix)
	fileName := strconv.FormatInt(i, 10)
	return fileName
}

//time格式转string
//例子  获取当前时间  TimeFormatString(time.Now())
func TimeFormatString(time time.Time) string {
	baseFormat := "2006-01-02 15:04:05"
	strTime := time.Format(baseFormat)
	return strTime
}

//获取当前时间 string 没有小尾巴
func GetTimeStringNowNoTail() string {
	t := TimeFormatString(*GetTimeNow())
	return t[:len(t)-9]
}

//获取当前时间 2019-07-30 14:12:20
func GetTimeString() string {
	return TimeFormatString(time.Now())
}
