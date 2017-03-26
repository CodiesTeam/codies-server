package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func MD5Sum(args ...string) string {
	h := md5.New()
	for _, arg := range args {
		io.WriteString(h, arg)
	}
	return hex.EncodeToString(h.Sum(nil))
}
