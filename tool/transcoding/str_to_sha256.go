/**
*@Author: Zzy
*@Date: 2019/4/22 15:21
*@Description:
 */
package transcoding

import (
	"crypto/sha256"
	"fmt"
	"io"
)

func StrToSHA256(data string) string {
	t := sha256.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
