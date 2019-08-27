/**
*@Author: Zzy
*@Date: 2019/4/23 10:50
*@Description:
 */
package http

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//body = strings.NewReader(string(by)
func NewRequest(method, url string, body io.Reader) (by []byte, err error) {
	method = strings.ToUpper(method)
	if method == "POST" {
		client := &http.Client{}
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		defer req.Body.Close()
		if err != nil {
			log.Println(err)
			return nil, err
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	} else if method == "GET" {
		req, err := http.Get(url)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		defer req.Body.Close()
		return ioutil.ReadAll(req.Body)
	}
	return nil, nil
}
