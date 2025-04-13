package api

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"gbc/backend/bridge"
	"gbc/backend/entity"
	"gbc/backend/response"
	"math/big"
	mrand "math/rand"
	"os"
	"time"
)

func KeyGenRequest(info string) interface{} {
	var license entity.License
	infoBytes := []byte(info)
	json.Unmarshal(infoBytes, &license)
	license.CreateLis()
	if license.LicenseId == "" {
		license.LicenseId = randomString(8)
	}
	key, err := keyGen(&license)
	if err != nil {
		return response.NewErrorResponse[string](500, "failed to generation a key")
	}
	return response.NewSuccessResponse(key)

}

func keyGen(license *entity.License) (string, error) {
	caPath := "resources/ca.crt"
    // // 使用 bridge.FS 检查文件是否存在
    // if _, err := bridge.Fs.Open(caPath); err != nil {
    //     if err == fs.ErrNotExist {
    //         // 文件不存在，创建证书
    //         createCA()
    //     }
    //     // 其他错误
    //     return "", err
    // }
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

func randomString(length int) string {
	charset := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	seededRand := mrand.New(mrand.NewSource(time.Now().UnixNano()))

	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

func createCA() {

	caPath := "../bridge/resources"
	// 生成RSA私钥
    privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
    if err != nil {
        panic(err)
    }

    // 创建证书模板
    template := x509.Certificate{
        SerialNumber: big.NewInt(1),
        Subject: pkix.Name{
            CommonName: "Activity-JetBrains",
        },
        NotBefore: time.Now(),
        NotAfter:  time.Now().AddDate(50, 0, 0), // 10年有效期
        KeyUsage:  x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
        BasicConstraintsValid: true,
    }

    // 创建自签名证书
    derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
    if err != nil {
        panic(err)
    }

    // 保存私钥
    privateKeyFile, err := os.Create(caPath+"/ca.key")
    if err != nil {
        panic(err)
    }
    defer privateKeyFile.Close()

    // 使用PKCS#1格式而不是PKCS#8
    privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
    if err := pem.Encode(privateKeyFile, &pem.Block{
        Type:  "RSA PRIVATE KEY",  // 使用RSA PRIVATE KEY而不是PRIVATE KEY
        Bytes: privateKeyBytes,
    }); err != nil {
        panic(err)
    }

    // 保存证书
    certFile, err := os.Create(caPath+"/ca.crt")
    if err != nil {
        panic(err)
    }
    defer certFile.Close()

    certBlock := &pem.Block{
        Type:  "CERTIFICATE",
        Bytes: derBytes,
    }

    if err := pem.Encode(certFile, certBlock); err != nil {
        panic(err)
    }
}
