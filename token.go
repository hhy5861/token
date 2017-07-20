package token

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/hhy5861/aescrypt"
)

var (
	key          = []byte("w32%y7893Ui62nvb")
	eReplacement = map[string]string{
		"+": "-",
		"/": "_",
	}

	dReplacement = map[string]string{
		"-": "+",
		"_": "/",
	}
)

func GenerateToken(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func Encrypt(data []byte) string {
	encrypted, err := aescrypt.AesEncrypt(data, key)
	if err == nil {
		baseCode := base64.StdEncoding.EncodeToString(encrypted)
		strCode := Strtr(baseCode, eReplacement)
		return strCode
	}

	return ""
}

func Decrypt(data string) (string, error) {
	strCode := Strtr(data, dReplacement)
	baseCode, err := base64.StdEncoding.DecodeString(strCode)
	if err == nil {
		metaData, err := aescrypt.AesDecrypt(baseCode, key)
		return string(metaData), err
	}

	return "", err
}

func Strtr(data string, Replace map[string]string) string {
	for k, v := range Replace {
		data = strings.Replace(data, k, v, -1)
	}

	return data
}
