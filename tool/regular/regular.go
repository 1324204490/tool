/**
*@Author: Zzy
*@Date: 2019/5/10 10:44
*@Description:
 */
package regular

import "regexp"

/**
IsPhone()
验证手机号格式是否正确
*/
func IsPhone(phone string) bool {
	b, _ := regexp.MatchString("^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\\d{8}$", phone)
	return b
}

/**
IsEmail()
验证邮箱
*/
func IsEmail(email string) bool {
	b, _ := regexp.MatchString("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$", email)
	return b
}

/**
IsPwd()
密码(以字母开头，长度在6~18之间，只能包含字母、数字和下划线)
*/
func IsPwd(pwd string) bool {
	b, _ := regexp.MatchString("^[a-zA-Z]\\w{5,17}$", pwd)
	return b
}

/**
验证车牌号格式
*/
func IsVehicleNumber(vehicleNumber string) bool {
	b, _ := regexp.MatchString("^[京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领A-Z]{1}[A-Z]{1}[A-Z0-9]{4}[A-Z0-9挂学警港澳]{1}$", vehicleNumber)
	return b
}
