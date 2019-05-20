/**
*@Author: Zzy
*@Date: 2019/5/15 17:32
*@Description:
 */
package file

import (
	"log"
	"os"
)

/**
用于判断文件夹是否已经存在
*/
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/**
创建文件夹
*/
func CreateFolder(path string) bool {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
