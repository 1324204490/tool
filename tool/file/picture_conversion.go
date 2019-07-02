/**
*@Author: Zzy
*@Date: 2019/4/22 17:03
*@Description:
 */
package file

import (
	"bufio"
	"github.com/prometheus/common/log"
	"io"
	"os"
)

/**
图片互相转换
fileNameAddr 需要转换的图片地址
toFileNameAddr 转换后的图片地址
*/
func PictureConversion(fileNameAddr, toFileNameAddr string) bool {
	file, err := os.Open(fileNameAddr)
	reader := bufio.NewReader(file)
	//创建文件
	fCreate, err := os.Create(toFileNameAddr)
	if err != nil {
		log.Error(err)
		return false
	}
	defer fCreate.Close()
	_, err = io.Copy(fCreate, reader)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func VerificationImg(suffix string) bool {
	if suffix == ".bmp" || suffix == ".jpg" ||
		suffix == ".png" || suffix == ".jpeg" ||
		suffix == ".webp" || suffix == ".pcx" ||
		suffix == ".tif" || suffix == ".tga" ||
		suffix == ".fpx" || suffix == ".psd" ||
		suffix == ".pcd" {
	} else {
		return false
	}
	return true
}
