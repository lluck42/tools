package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func RsaEncryptWithPrivateKey(privateKeyStr, secretText string) (string, error) {

	// 解析PEM格式的私钥
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block containing the key123")
	}

	var privKey *rsa.PrivateKey

	// 尝试PKCS8格式
	if key, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		privKey = key.(*rsa.PrivateKey)
	} else {
		// 尝试PKCS1格式
		privKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return "", fmt.Errorf("failed to parse private key: %v", err)
		}
	}
	////////////////////////////////////

	// 2. 使用私钥加密
	ciphertext, err := rsa.SignPKCS1v15(rand.Reader, privKey, 0, []byte(secretText))
	if err != nil {
		return "", fmt.Errorf("加密字符串[%s]时遇到异常: %v", secretText, err)
	}

	// 3. 返回Base64编码结果
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func RsaDecryptWithPrivateKey(privateKeyStr, encryptedText string) (string, error) {
	// fmt.Println("privateKeyStr, encryptedText:", privateKeyStr, encryptedText)
	// Base64解码
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", fmt.Errorf("base64 decode failed: %v", err)
	}

	// 解析PEM格式的私钥
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block containing the key")
	}

	var privKey *rsa.PrivateKey

	// 尝试PKCS8格式
	if key, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		privKey = key.(*rsa.PrivateKey)
	} else {
		// 尝试PKCS1格式
		privKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return "", fmt.Errorf("failed to parse private key: %v", err)
		}
	}

	// 使用私钥解密
	decryptedBytes, err := rsa.DecryptPKCS1v15(nil, privKey, encryptedBytes)
	if err != nil {
		return "", fmt.Errorf("RSA decryption failed: %v", err)
	}
	fmt.Println("解密结果:", string(decryptedBytes))
	return string(decryptedBytes), nil
}
