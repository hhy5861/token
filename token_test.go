package actoken

import (
	"testing"
)

var enCode = "wFFFjtv86PIe8FtSiKlHO8aqzK3dI4XS411xNIzsbb8="

func TestEncrypt(t *testing.T) {
	token := GenerateToken(8)
	Encrypt([]byte(token))
}

func BenchmarkEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		token := GenerateToken(8)
		Encrypt([]byte(token))
	}
}

func TestDecrypt(t *testing.T) {
	Decrypt(enCode)
}

func BenchmarkDecrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decrypt(enCode)
	}
}
