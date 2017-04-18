package common

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

var (
	idChars    = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	idCharsLen = int64(len(idChars))

	idChan = func() <-chan string {
		c := make(chan string)
		offset := time.Date(2017, time.April, 18, 0, 0, 0, 0, time.UTC)
		go func() {
			seq := int64(0)
			previousId := int64(0)
			for {
				inc, err := rand.Int(rand.Reader, big.NewInt(499))
				if err != nil {
					panic(err)
				}
				if seq > 9500 {
					seq = 0
				}
				seq += inc.Int64()
				// eg: 150706370050000 + seq, and seq < 10000, will failed 200years later
				id := int64(time.Now().Sub(offset).Seconds()*float64(time.Microsecond)*10)*10000 + seq
				if id < previousId {
					id += 1
				}
				previousId = id
				c <- base62encode(id)
			}
		}()
		return c
	}()
)

// base62encode 10进制转换为62进制
func base62encode(n int64) string {
	if n < 0 {
		panic(fmt.Sprintf("supposed n > 0 but %v", n))
	}
	if n < idCharsLen {
		return string(idChars[n])
	}
	ret := ""
	for n != 0 {
		ret = string(idChars[n%idCharsLen]) + ret
		n = n / idCharsLen
	}
	return ret
}

// NewUniqueID get unique ID
func NewUniqueID() string {
	return <-idChan
}
