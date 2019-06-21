package tencent

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/spf13/cast"
	"time"
)

/*
获取腾讯云点播的签名
secretId
secretKey
d 多久过期
random 随机数
*/
func GetCloudOnDemandSignature(secretId, secretKey string, d time.Duration, random string) string {
	secret := "secretId=" + secretId + "&currentTimeStamp=" + cast.ToString(time.Now().Unix()) + "&expireTime=" + cast.ToString(time.Now().Add(d).Unix()) + "&random=" + random
	key := []byte(secretKey)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(secret))
	SignatureTmp := string(mac.Sum(nil))
	SignatureTmp += secret
	return base64.StdEncoding.EncodeToString([]byte(SignatureTmp))
}
