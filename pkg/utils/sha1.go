package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

// SHA1 return the sha1(string)
func SHA1(content string) string {
	h:=sha1.New()
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
}

// Password2Secret convert the password(plaintext) to salted, use double SHA1 and saltA saltB, database store its output
func Password2Secret(password,saltA,saltB string) string {
	return SHA1(saltB+SHA1(saltA+password))
}