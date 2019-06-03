/**
*@Author: Zzy
*@Date: 2019/5/9 14:48
*@Description: 金蝶K3Cloud接口调用
 */
package kingdee

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/**
url 请求地址
setOfBooks 账套ID
m 请求参数
m["userName"] 用户名
m["pwd"]   密码
m["isKickOff"] 是否踢人 url的后半截必须为Kingdee.BOS.WebApi.ServicesStub.AuthService.ValidateUser2.common.kdsvc
m["lang"] 语言
后续请求接口都必须带cookie 以及请求头必须带token 响应头需把ContextType 改为text/plain;charset=UTF-8
c.Ctx.ResponseWriter().Header().Set("Content-Type", "text/plain;charset=UTF-8")
*/
func Login(url, setOfBooks string, m map[string]interface{}) map[string]string {
	var slice []interface{}
	slice = append(slice, setOfBooks)
	slice = append(slice, m["userName"])
	slice = append(slice, m["pwd"])
	slice = append(slice, m["isKickOff"])
	slice = append(slice, m["lang"])
	s, err := json.Marshal(slice)
	if err != nil {
		log.Println(err)
	}
	var m2 = make(map[string]interface{})
	m2["format"] = 1
	m2["useragent"] = "ApiClient"
	m2["parameters"] = string(s)
	infoString, err := json.Marshal(m2)
	if err != nil {
		log.Println(err)
	}
	body := strings.NewReader(string(infoString))
	r, err := http.Post(url, "application/json", body)
	if err != nil {
		log.Println(err)
	}
	cookie := r.Header.Get("Set-Cookie")
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var mp = make(map[string]string)
	mp["data"] = string(b)
	mp["cookie"] = cookie
	return mp
}

/**
用于请求金蝶的方法
url 请求地址
data 经过格式化的请求数据
需要设置登录获取的Cookie
*/
func Request(url, data, cookie string) string {
	if cookie == "" {
		return "cookie没传"
	}
	client := &http.Client{}
	req, _ := http.NewRequest("Post", url, strings.NewReader(data))
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	return string(b)
}

/**
将请求的数据格式化
*/
func Biography(m map[string]string) string {
	delete(m, "cookie")
	var slice []interface{}
	for _, v := range m {
		slice = append(slice, v)
	}
	s, err := json.Marshal(slice)
	if err != nil {
		log.Println(err)
	}
	var m2 = make(map[string]interface{})
	m2["format"] = 1
	m2["useragent"] = "ApiClient"
	m2["rid"] = -123054561
	m2["parameters"] = string(s)
	m2["v"] = "1.0"
	infoString, err := json.Marshal(m2)
	if err != nil {
		log.Println(err)
	}
	return string(infoString)
}
