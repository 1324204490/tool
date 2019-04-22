/**
*@Author: Zzy
*@Date: 2019/4/22 15:20
*@Description:
 */
package transcoding

import (
	"crypto/md5"
	"encoding/hex"
)

func StrToMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}
