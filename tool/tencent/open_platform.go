/**
*@Author: Zzy
*@Date: 2019/8/16 11:37
*@Description: 开放平台流程
 */
package tencent

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"education/models/wechat"
	"education/tool"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"tool/tool/tencent/models"
)

type Platform struct {
	EncodingAESKey string
	AppSecret      string
	Aes            []byte
}

func (p *Platform) NewPlatform(AppID, AppSecret, RedirectUri, EncodingAESKey string) *Platform {
	p.AppSecret = AppSecret
	p.EncodingAESKey = EncodingAESKey
	p.Aes = encodingAESKey2AESKey(EncodingAESKey)
	return p
}

//接收消息
func (p *Platform) Notice(request *http.Request) (err error, appId, accessToken string) {
	defer request.Body.Close()
	var encrypt models.XmlEncrypt
	e, _ := ioutil.ReadAll(request.Body)
	err = xml.Unmarshal(e, &encrypt)
	if err != nil {
		log.Println("解析xml错误", err)
		return err, "", ""
	}
	if encrypt.Encrypt == "" {
		log.Println("xml的Encrypt为空")
		return errors.New("xml的Encrypt为空"), "", ""
	}
	cipherData, err := base64.StdEncoding.DecodeString(encrypt.Encrypt)
	if err != nil {
		log.Println(err)
		return err, "", ""
	}
	plainData, err := aesDecrypt(cipherData, p.Aes)
	if err != nil {
		log.Println(err)
		return err, "", ""
	}
	accessToken = parseEncryptTextRequestBody(plainData)
	if accessToken == "" {
		return errors.New("获取accessToken失败"), "", ""
	}
	return nil, encrypt.AppId, accessToken
}
func encodingAESKey2AESKey(encodingKey string) []byte {
	data, _ := base64.StdEncoding.DecodeString(encodingKey + "=")
	return data
}
func aesDecrypt(cipherData []byte, aesKey []byte) ([]byte, error) {
	k := len(aesKey) //PKCS#7
	if len(cipherData)%k != 0 {
		return nil, errors.New("crypto/cipher:cipherText size is not multiple of aes key length")
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainData := make([]byte, len(cipherData))
	blockMode.CryptBlocks(plainData, cipherData)
	return plainData, nil
}
func parseEncryptTextRequestBody(plainData []byte) (accessToken string) {
	buf := bytes.NewReader(plainData[16:20])
	var length int32
	binary.Read(buf, binary.BigEndian, &length)
	var returnXml wechat.ReturnXml
	xml.Unmarshal(plainData[20:20+length], &returnXml)
	if returnXml.ComponentVerifyTicket == "" {
		return ""
	}
	var component models.Component
	component.AppId = returnXml.AppId
	component.Appsecret = tool.ConfJson.AppSecret
	component.VerifyTicket = returnXml.ComponentVerifyTicket
	componentByte, err := json.Marshal(component)
	if err != nil {
		log.Println(err)
	}
	body, err := tool.NewRequest("Post", "https://api.weixin.qq.com/cgi-bin/component/api_component_token", strings.NewReader(string(componentByte)))
	var component2 models.Component
	err = json.Unmarshal(body, &component2)
	if err != nil {
		log.Println(err)
	}
	return component2.ComponentAccessToken
}
