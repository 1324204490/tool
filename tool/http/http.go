/**
*@Author: Zzy
*@Date: 2019/4/23 10:50
*@Description:
 */
package http

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

/**
url 请求地址后面可追加请求参数
v 返回的实例对象 传指针
*/
func Get(url string, v interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return err
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	json.Unmarshal(b, &v)
	return nil
}

/**
url 请求地址
v 返回的实例对象
v 返回的实例对象 传指针
body 参数内容
info := make(map[string]interface{})
info["bindNbr"] = conf.Conf.BindNbr
info["calleeNbr"] = calleeNbr
info["playInfoList"] = []interface{}{playInfoList}
info["displayNbr"] = conf.Conf.DisplayNbr
//struct 到json str
infoString, err := json.Marshal(info)
if err != nil {
   log.Println(err)
}
body := bytes.NewBuffer(infoString)
*/
func Post(url string, contentType string, body io.Reader, v interface{}) error {
	r, err := http.Post(url, contentType, body)
	if err != nil {
		log.Println(err)
		return err
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	json.Unmarshal(b, &v)
	return nil
}
