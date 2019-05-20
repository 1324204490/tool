/**
*@Author: Zzy
*@Date: 2019/4/23 11:08
*@Description:
 */
package model

type Result struct {
	Status bool
	Msg    string
	Data   interface{}
}

/*
d为通用返回数据，e为可选参数，e不为空时返回失败result
用于微服务
*/
func SetResult(result *Result, d interface{}, e ...string) {
	if len(e) > 0 {
		result.Status = false
		result.Msg = e[0]
		result.Data = nil
	} else {
		result.Status = true
		result.Msg = "success"
		result.Data = d
	}
}

/*
d为通用返回数据，e为可选参数，e不为空时返回失败result
用于普通项目
*/
func NewResult(d interface{}, e ...string) *Result {
	if len(e) > 0 {
		return &Result{false, e[0], nil}
	}
	return &Result{true, "success", d}
}
