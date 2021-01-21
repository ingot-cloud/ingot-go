package digest

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

// MD5 MD5摘要
func MD5(b []byte) string {
	h := md5.New()
	_, _ = h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// MD5String MD5摘要
func MD5String(s string) string {
	return MD5([]byte(s))
}

// SHA1 SHA1摘要
func SHA1(b []byte) string {
	h := sha1.New()
	_, _ = h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA1String SHA1摘要
func SHA1String(s string) string {
	return SHA1([]byte(s))
}
