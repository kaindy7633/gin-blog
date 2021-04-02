package util

import (
	"crypto/md5"
	"encoding/hex"
	"gin-blog/pkg/logging"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	_, err := m.Write([]byte(value))
	if err != nil {
		logging.Info(err)
	}

	return hex.EncodeToString(m.Sum(nil))
}
