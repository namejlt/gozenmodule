package msso

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

/**

token
eyJhbGciOiJSUzI1NiJ9.eyJ1aWQiOjQwMDAyOTgxLCJhYyI6IlJNQyIsImNvZGUiOiIwMDI5ODIiLCJleHBpcmUiOjE3Mjk1NjE3OTAsIm5hbWUiOiLotL7pvpnlpKkiLCJlbnYiOiJ4amprIiwidGlkIjoxNTM3MjZ9.E0zuhhVcCjPRxsKkA_ScChtGieD3bb1OzkZv2ojb2yJTrcrjZjCXsX74RBK1MM03r7jro2hjbCree1iqO_iGLzzx6EKf1EGSxQbpBCCgWwGc02m_gsMSNHt5R_mAb-AHCWPAPTEIcP_tuOluAxJrxaSvkz2whUlp2OaFfpR5X3xVG4vGx1pK3vR-feElk-lHzjqQk6OXZr4oDXYovXNE1SV-h54vVn7f4fdGVO4b_uKRbLspJ56n1YpfGJUnZAdjP2AyWz9EENtZrr_p_rVQB0TAVqmU3SdywANE3m_46Cqnge_ZM5Ot8gH3gq0WvCoaTPwDUTj1tSNf6p8jQwvnKg

公钥-需base64解码
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAirtJEhSC0x6SHC0LymJp8HvB+pT5Lt+kq3LPQAEZiu+IvPCmzdTnpgQjk5vjioqEW3bBJk+U8nJa/wlvmFB94jQnjeRZZRAfqRTtYGJIuSkBO9m9cuT24SqJZ6MpkHcXE6oPvp4YPrn6Ac+lYSPx+n3PQ3lowHw9HDLk5iayg4f/pfevqjLOmCaP+1nlru7p4gXmeS0n+a9JGYuntqAx6bTsxLdLkTwJRSgIOkD26TW1zSFLNLkL6b+YuVPvjBE6ikfNDU0SHPFYMj1rIoDfdzNknIrtqbEwdWhCeUZVj7R6I4OOYqpNCY8+6cNNneiaeivtQNWv38elNgrs1xgiowIDAQAB

nacos

dataid: common-client-ident
group: base-service

*/

type SsoCustomClaims struct {
	Uid      uint64 `json:"uid,omitempty"`
	Ac       string `json:"ac,omitempty"`
	Code     string `json:"code,omitempty"`
	Expire   int64  `json:"expire,omitempty"`
	Name     string `json:"name,omitempty"`
	Env      string `json:"env,omitempty"`
	TenantId uint64 `json:"tid,omitempty"`
	jwt.StandardClaims
}

func JWTCheck(tokenString string, pubKey string) (check bool, data *SsoCustomClaims, err error) {
	pubKeyByte, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return
	}
	publicKey, err := x509.ParsePKIXPublicKey(pubKeyByte)
	if err != nil {
		return
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		err = errors.New("public key is not of type RSA")
		return
	}
	// 解析 JWT
	token, err := jwt.ParseWithClaims(tokenString, &SsoCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok = token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回用于验证签名的公钥
		return rsaPublicKey, nil
	})

	if err != nil {
		return
	}

	// 验证 JWT 是否有效
	if data, ok = token.Claims.(*SsoCustomClaims); ok && token.Valid {
		now := time.Now().Unix()
		if data.Expire < now { //过期
			return
		}
		check = true
		return
	} else {
		// 无效
		return
	}
}
