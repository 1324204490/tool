/**
*@Author: Zzy
*@Date: 2019/4/23 11:29
*@Description:
 */
package transcoding

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

/**
MD5加密
*/
func StrToMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

/**
SHA256加密
*/
func StrToSHA256(data string) string {
	t := sha256.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
