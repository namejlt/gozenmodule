package mcommon

import (
	"errors"
	"github.com/speps/go-hashids"
	"math/rand"
	"time"
)

func (p *TokenSession) getSalt() string {
	return "Hu87TGkd65"
}

func (p *TokenSession) getTokenLen() int {
	return 10
}

func NewTokenSession(tenantId uint64, accountId uint64) *TokenSession {
	now := time.Now()
	rand.Seed(now.UnixNano())
	t := new(TokenSession)
	t.TenantId = tenantId
	t.AccountId = accountId
	t.Uuid = 1000 + rand.Int63n(10000) //1000 - 19999
	t.LoginTimestamp = now.Unix()
	return t
}

type TokenSession struct {
	TenantId       uint64 `json:"tenant_id"`       //租户号
	AccountId      uint64 `json:"account_id"`      //账号id
	Uuid           int64  `json:"uuid"`            //登录随机号 1-100
	LoginTimestamp int64  `json:"login_timestamp"` //登录时间
}

func (p *TokenSession) Encode() (str string, err error) {
	hd := hashids.NewData()
	hd.Salt = p.getSalt()
	hd.MinLength = p.getTokenLen()
	var h *hashids.HashID
	h, err = hashids.NewWithData(hd)
	if err != nil {
		return
	}
	str, err = h.EncodeInt64([]int64{int64(p.TenantId), int64(p.AccountId), p.Uuid, p.LoginTimestamp})
	return
}

func (p *TokenSession) Decode(str string) (err error) {
	hd := hashids.NewData()
	hd.Salt = p.getSalt()
	hd.MinLength = p.getTokenLen()
	var h *hashids.HashID
	h, err = hashids.NewWithData(hd)
	if err != nil {
		return
	}
	var params []int64
	params, err = h.DecodeInt64WithError(str)
	if err != nil {
		return
	}
	if len(params) == 4 {
		p.TenantId = uint64(params[0])
		p.AccountId = uint64(params[1])
		p.Uuid = params[2]
		p.LoginTimestamp = params[3]
	} else {
		err = errors.New("decode len is wrong")
		return
	}
	return
}
