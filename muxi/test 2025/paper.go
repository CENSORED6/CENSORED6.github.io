package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/sha256"
    "encoding/base64"
    "fmt"
    "strings"
)

func main() {
	var 

}

func DecodeEncryptedBase64(paper string, password string) (string, error) {
    raw, err := base64.StdEncoding.DecodeString(strings.TrimSpace(paper))
    if err != nil {
       return "", fmt.Errorf("base64 解码失败: %w", err)
    }
    key := sha256.Sum256([]byte(password))
    block, err := aes.NewCipher(key[:])
    if err != nil {
       return "", err
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
       return "", err
    }
    nonceSize := gcm.NonceSize()
    if len(raw) < nonceSize {
       return "", fmt.Errorf("数据长度不足: %d < %d", len(raw), nonceSize)
    }
    nonce := raw[:nonceSize]
    ciphertext := raw[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
       return "", fmt.Errorf("GCM 解密失败: %w", err)
    }
    return string(plaintext), nil
}