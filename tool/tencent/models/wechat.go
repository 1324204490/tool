/**
*@Author: Zzy
*@Date: 2019/7/5 14:32
*@Description: 用于解析微信返回的json信息
 */
package models

type XmlEncrypt struct {
	AppId   string `xml:"AppId"`
	Encrypt string `xml:"Encrypt"`
}
type XmlNotice struct {
	ToUserName string `xml:"ToUserName"`
	Encrypt    string `xml:"Encrypt"`
}
type ReturnXml struct {
	AppId                 string `xml:"AppId"`
	CreateTime            string `xml:"CreateTime"`
	InfoType              string `xml:"InfoType"`
	ComponentVerifyTicket string `xml:"ComponentVerifyTicket"`
}
type Component struct {
	AppId                string `json:"component_appid"`
	Appsecret            string `json:"component_appsecret"`
	VerifyTicket         string `json:"component_verify_ticket"`
	ComponentAccessToken string `json:"component_access_token"`
	PreAuthCode          string `json:"pre_auth_code"`
}

/*
解析模板
*/
type TemplateAuto struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	TemplateList []struct {
		CreateTime             int    `json:"create_time"`
		UserVersion            string `json:"user_version"`
		UserDesc               string `json:"user_desc"`
		TemplateID             int    `json:"template_id"`
		SourceMiniprogramAppid string `json:"source_miniprogram_appid"`
		SourceMiniprogram      string `json:"source_miniprogram"`
		Developer              string `json:"developer"`
	} `json:"template_list"`
}
type Template struct {
	TemplateID  int    `json:"template_id"`
	ExtJSON     string `json:"ext_json"`
	UserVersion string `json:"user_version"`
	UserDesc    string `json:"user_desc"`
}
type Message struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Auditid string `json:"auditid"`
}

//用于刷新token
type AuthorizerAccessToken struct {
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

//解析小程序设置的类目
type Category struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	CategoryList []struct {
		FirstClass  string `json:"first_class"`
		SecondClass string `json:"second_class"`
		FirstID     int    `json:"first_id"`
		SecondID    int    `json:"second_id"`
		ThirdClass  string `json:"third_class,omitempty"`
		ThirdID     int    `json:"third_id,omitempty"`
	} `json:"category_list"`
}

//解析小程序页面路径
type Page struct {
	Errcode  int      `json:"errcode"`
	Errmsg   string   `json:"errmsg"`
	PageList []string `json:"page_list"`
}

//用于审核
type Examine struct {
	ItemList []ItemList `json:"item_list"`
}
type ItemList struct {
	Address     string `json:"address"`
	Tag         string `json:"tag"`
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
	FirstID     int    `json:"first_id"`
	SecondID    int    `json:"second_id"`
	Title       string `json:"title"`
	ThirdClass  string `json:"third_class,omitempty"`
	ThirdID     int    `json:"third_id,omitempty"`
}

//查询审核状态
type AuditStatus struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	Auditid    int    `json:"auditid"`
	Status     int    `json:"status"`
	Reason     string `json:"reason"`
	ScreenShot string `json:"ScreenShot"`
}

//客服消息
type Xml struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        string `xml:"MsgId"`
}
