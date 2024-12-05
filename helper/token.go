package helper

import (
	"be-golang-chapter-44-45-implem/config"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

func GenerateToken(userID, email, username string, secretKey string) string {
	// Gabungkan data user
	data := fmt.Sprintf("%s:%s:%s:%d", userID, email, username, time.Now().Unix())

	// Buat hash HMAC menggunakan SHA-256
	h := hmac.New(sha256.New, []byte(config.Configs.Secret))
	h.Write([]byte(data))
	signature := base64.URLEncoding.EncodeToString(h.Sum(nil))

	// Encode data ke Base64
	tokenData := base64.URLEncoding.EncodeToString([]byte(data))

	// Gabungkan data dan tanda tangan
	return fmt.Sprintf("%s.%s", tokenData, signature)
}

func ValidasialidateToken(token string, secretKey string) (bool, string) {
	// Pisahkan token menjadi data dan signature
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return false, "Invalid token format"
	}

	tokenData, signature := parts[0], parts[1]

	// Decode data dari Base64
	data, err := base64.URLEncoding.DecodeString(tokenData)
	if err != nil {
		return false, "Invalid token data"
	}

	// Buat ulang signature untuk validasi
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write(data)
	expectedSignature := base64.URLEncoding.EncodeToString(h.Sum(nil))

	// Validasi signature
	if signature != expectedSignature {
		return false, "Invalid token signature"
	}

	return true, string(data)
}
