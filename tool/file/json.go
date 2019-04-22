/**
*@Author: Zzy
*@Date: 2019/4/22 17:17
*@Description:
 */
package file

import (
	"encoding/json"
	"io/ioutil"
)

/**
解析Json配置文件
fileName 文件名
v   结构体
*/
func JsonLoad(fileNameAddr string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(fileNameAddr)
	if err != nil {
		return
	}
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
