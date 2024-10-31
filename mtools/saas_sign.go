package mtools

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"time"
)

/**


标准验签和加解密

rsa SHA256

标准接口 签名验证


*/

// RsaEncrypt encrypts data using rsa public key.
func RsaEncrypt(pubkey, data []byte) ([]byte, error) {
	block, _ := pem.Decode(pubkey)
	if block == nil {
		return nil, errors.New("decode public key error")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), data)
}

// RsaDecrypt decrypts data using rsa private key.
func RsaDecrypt(prvkey, cipher []byte) ([]byte, error) {
	block, _ := pem.Decode(prvkey)
	if block == nil {
		return nil, errors.New("decode private key error")
	}
	prv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, prv, cipher)
}

// RsaSign signs using private key in PEM format 返回base64加密字符串
func RsaSign(prvkey []byte, hash crypto.Hash, data []byte) (r string, err error) {
	/*block, _ := pem.Decode(prvkey)
	if block == nil {
		return nil, errors.New("decode private key error")
	}*/
	privateKey, err := x509.ParsePKCS8PrivateKey(prvkey)
	if err != nil {
		return
	}

	pk := privateKey.(*rsa.PrivateKey)
	// MD5 and SHA1 are not supported as they are not secure.
	var hashed []byte
	switch hash {
	case crypto.SHA224:
		h := sha256.Sum224(data)
		hashed = h[:]
	case crypto.SHA256:
		h := sha256.Sum256(data)
		hashed = h[:]
	case crypto.SHA384:
		h := sha512.Sum384(data)
		hashed = h[:]
	case crypto.SHA512:
		h := sha512.Sum512(data)
		hashed = h[:]
	}
	var ret []byte
	ret, err = rsa.SignPKCS1v15(rand.Reader, pk, hash, hashed)
	if err != nil {
		return
	}
	r = base64.StdEncoding.EncodeToString(ret)
	return
}

// RsaVerifySign verifies signature using public key in PEM format.
// A valid signature is indicated by returning a nil error.
func RsaVerifySign(pubkey []byte, hash crypto.Hash, data []byte, sign string) (err error) {
	sig, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return
	}
	/*block, _ := pem.Decode(pubkey)
	if block == nil {
		return errors.New("decode public key error")
	}*/
	pub, err := x509.ParsePKIXPublicKey(pubkey)
	if err != nil {
		return err
	}

	// SHA1 and MD5 are not supported as they are not secure.
	var hashed []byte
	switch hash {
	case crypto.SHA224:
		h := sha256.Sum224(data)
		hashed = h[:]
	case crypto.SHA256:
		h := sha256.Sum256(data)
		hashed = h[:]
	case crypto.SHA384:
		h := sha512.Sum384(data)
		hashed = h[:]
	case crypto.SHA512:
		h := sha512.Sum512(data)
		hashed = h[:]
	}
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), hash, hashed, sig)
}

func GetPubKey(tenantId uint64) []byte {
	key := []byte("根据租户的固定值")
	return key
}

func GetPrvKey(tenantId uint64) []byte {
	key := []byte("根据租户的固定值")
	return key
}

func GetSignAuthorization(tenantId uint64, method string, path string, body []byte) (data string, err error) {
	/**
	构造签名串，签名串一共有五行，每一行为一个参数。行尾以 \n（换行符，ASCII编码值为0x0A）结束，包括最后一行。如果参数本身以\n结束，也需要附加一个\n。

	HTTP请求方法\n
	URL\n
	请求随机串\n
	请求时间戳\n
	请求报文主体\n

	*/
	nonceStr := RandomStr(32)
	timestamp := time.Now().Unix()
	signContent := fmt.Sprintf(`%s
%s
%s
%d
%s
`,
		method,
		path,
		nonceStr,
		timestamp,
		string(body),
	)

	signature, err := RsaSign(GetPrvKey(tenantId), crypto.SHA256, []byte(signContent))
	if err != nil {
		return
	}
	data = fmt.Sprintf(`META-SHA256-RSA2048 nonce_str="%s",timestamp="%d",signature="%s"`,
		nonceStr, timestamp, signature)
	return
}
