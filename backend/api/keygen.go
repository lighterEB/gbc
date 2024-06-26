package api

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"gbc/backend/bridge"
	"gbc/backend/entity"
	"gbc/backend/response"
)

func KeyGenRequest(info string) interface{} {
	var license entity.License
	infoBytes := []byte(info)
	json.Unmarshal(infoBytes, &license)
	license.CreateLis()
	key, err := keyGen(&license)
	if err != nil {
		return response.NewErrorResponse[string](500, "failed to generation a key")
	}
	return response.NewSuccessResponse(key)

}

func keyGen(license *entity.License) (string, error) {
	caPath := "resources/ca.crt"
	// 获取证书
	ca, _ := bridge.Fs.ReadFile(caPath)
	certBase64 := base64.StdEncoding.EncodeToString(ca)

	// 签名数据
	// licensePart := `{"licenseId":"7GC5726T07","licenseeName":"gurgles tumbles","assigneeName":"","assigneeEmail":"","licenseRestriction":"","checkConcurrentUse":false,"products":[{"code":"PCWMP","fallbackDate":"2026-09-14","paidUpTo":"2026-09-14","extended":true},{"code":"GO","fallbackDate":"2026-09-14","paidUpTo":"2026-09-14","extended":false},{"code":"PSI","fallbackDate":"2026-09-14","paidUpTo":"2026-09-14","extended":true}],"metadata":"0120230914PSAX000005","hash":"TRIAL:1805249793","gracePeriodDays":7,"autoProlongated":false,"isAutoProlongated":false}`
	// licensePartBytes := []byte(licensePart)
	licensePartBytes, err := json.Marshal(license)
	if err != nil {
		return "", err
	}
	licensePartBase64 := base64.StdEncoding.EncodeToString(licensePartBytes)
	licenseId := license.LicenseId
	// 私钥
	pk, err := parsePrivateKey()
	if err != nil {
		return "", err
	}

	// fmt.Println("certBase64: ", certBase64)
	// fmt.Println("licensePartBase64: ", licensePartBase64)
	// fmt.Println("licenseId: ", licenseId)
	// fmt.Println("parsePrivateKey: ", pk)

	// 使用SHA1withRSA算法创建签名器
	hashed := sha1.Sum(licensePartBytes)
	signature, err := rsa.SignPKCS1v15(rand.Reader, pk.(*rsa.PrivateKey), crypto.SHA1, hashed[:])
	if err != nil {
		return "", err
	}
	signBytes := []byte(signature)
	signBase64 := base64.StdEncoding.EncodeToString(signBytes)
	lis := licenseId
	lis = fmt.Sprintf("%s-%s-%s-%s", lis, licensePartBase64, signBase64, certBase64)
	return lis, nil
}

func parsePrivateKey() (interface{}, error) {
	// 读取PEM文件
	caPath := "resources/ca.key"
	pemData, err := bridge.Fs.ReadFile(caPath)
	if err != nil {
		return nil, err
	}

	// 解析PEM数据块
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, errors.New("failed to parse PEM block")
	}

	// 根据不同的加密算法进行特定的解析
	var privateKey interface{}
	switch block.Type {
	case "RSA PRIVATE KEY":
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	case "EC PRIVATE KEY":
		privateKey, err = x509.ParseECPrivateKey(block.Bytes)
	default:
		return privateKey, nil
	}
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
