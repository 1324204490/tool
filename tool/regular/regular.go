/**
*@Author: Zzy
*@Date: 2019/5/10 10:44
*@Description:
 */
package regular

import (
	"regexp"
	"strings"
)

/**
IsPhone()
验证手机号格式是否正确
*/
func IsPhone(phone string) bool {
	b, _ := regexp.MatchString("^1([38][0-9]|4[579]|5[0-3,5-9]|6[6]|7[0135678]|9[89])\\d{8}$", phone)
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

/*
验证身份证号
*/
func IsIdCard(idCard string) bool {
	var coefficient = []int32{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var code = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
	if len(idCard) != 18 {
		return false
	}
	idByte := []byte(strings.ToUpper(idCard))
	sum := int32(0)
	for i := 0; i < 17; i++ {
		sum += int32(byte(idByte[i])-byte('0')) * coefficient[i]
	}
	return code[sum%11] == idByte[17]
}
