/**
*@Author: Zzy
*@Date: 2019/5/15 17:32
*@Description:
 */
package file

import (
	"compress/flate"
	"fmt"
	"github.com/mholt/archiver"
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

//创建一个ZIP压缩包
//filePath 为需要压缩到ZIP里的文件路径
//zipPath 为创建ZIP压缩包的路径
func CreateZip(filePath []string, zipPath string) bool {
	z := archiver.Zip{
		CompressionLevel:       flate.DefaultCompression,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ContinueOnError:        false,
		OverwriteExisting:      false,
		ImplicitTopLevelFolder: false,
	}
	//判断文件是否存在
	if Exists(zipPath) {
		//存在就先删除文件
		if err := os.Remove(zipPath); err != nil {
			log.Println(err)
			return false
		}
	}
	err := z.Archive(filePath, zipPath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
