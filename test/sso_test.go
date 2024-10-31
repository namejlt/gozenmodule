package test

import (
	"github.com/namejlt/gozenmodule/msso"
	"testing"
)

func TestJWTCheck(t *testing.T) {

	token := `eyJhbGciOiJSUzI1NiJ9.eyJ1aWQiOjQwMDAyOTgxLCJhYyI6IlJNQyIsImNvZGUiOiIwMDI5ODIiLCJleHBpcmUiOjE3Mjk1NjE3OTAsIm5hbWUiOiLotL7pvpnlpKkiLCJlbnYiOiJ4amprIiwidGlkIjoxNTM3MjZ9.E0zuhhVcCjPRxsKkA_ScChtGieD3bb1OzkZv2ojb2yJTrcrjZjCXsX74RBK1MM03r7jro2hjbCree1iqO_iGLzzx6EKf1EGSxQbpBCCgWwGc02m_gsMSNHt5R_mAb-AHCWPAPTEIcP_tuOluAxJrxaSvkz2whUlp2OaFfpR5X3xVG4vGx1pK3vR-feElk-lHzjqQk6OXZr4oDXYovXNE1SV-h54vVn7f4fdGVO4b_uKRbLspJ56n1YpfGJUnZAdjP2AyWz9EENtZrr_p_rVQB0TAVqmU3SdywANE3m_46Cqnge_ZM5Ot8gH3gq0WvCoaTPwDUTj1tSNf6p8jQwvnKg`

	secret := `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAirtJEhSC0x6SHC0LymJp8HvB+pT5Lt+kq3LPQAEZiu+IvPCmzdTnpgQjk5vjioqEW3bBJk+U8nJa/wlvmFB94jQnjeRZZRAfqRTtYGJIuSkBO9m9cuT24SqJZ6MpkHcXE6oPvp4YPrn6Ac+lYSPx+n3PQ3lowHw9HDLk5iayg4f/pfevqjLOmCaP+1nlru7p4gXmeS0n+a9JGYuntqAx6bTsxLdLkTwJRSgIOkD26TW1zSFLNLkL6b+YuVPvjBE6ikfNDU0SHPFYMj1rIoDfdzNknIrtqbEwdWhCeUZVj7R6I4OOYqpNCY8+6cNNneiaeivtQNWv38elNgrs1xgiowIDAQAB`

	check, data, err := msso.JWTCheck(token, secret)

	if err != nil {
		t.Error(err)
	} else {
		t.Log(check)
		t.Log(data)
	}
}
