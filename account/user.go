package account

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	block "github.com/reoxey/blockchain"
	"github.com/reoxey/blockchain/redis"
)

type Address string

type User struct {
	db   *redis.Conn
	Name string
	Addr string
	Date string
}

type Account interface {
	Balance() float64
	Transfer(Address)
}

func New(name string) (Account, error) {

	d := []byte(time.Now().String())

	addr := fmt.Sprintf("%x", sha256.Sum224(d))

	u := User{
		Name: name,
		Date: time.Now().String(),
		Addr: addr,
	}

	b, e := json.Marshal(&u)
	if e != nil {
		return nil, e
	}
	db := redis.Connect("127.0.0.1")
	if e = db.Set(addr, string(b)); e != nil {
		return nil, e
	}

	u.db = db

	return u, nil
}

func NewWithAddress(addr string, name string) (Account, error) {

	u := User{
		Name: name,
		Date: time.Now().String(),
		Addr: addr,
	}

	b, e := json.Marshal(&u)
	if e != nil {
		return nil, e
	}
	db := redis.Connect("127.0.0.1")
	if e = db.Set(addr, string(b)); e != nil {
		return nil, e
	}

	u.db = db

	return u, nil
}

func Get(addr string) (Account, bool) {

	db := redis.Connect("127.0.0.1")

	var u *User
	if e := db.GetJSON(addr, &u); e != nil {
		return nil, false
	}

	u.db = db

	return u, true
}

func (u User) Balance() float64 {
	b := block.Block{}
	var amt float64
	k := "LAST"

	for b.Index != 1 {
		if e := u.db.GetJSON(k, &b); e != nil {
			return amt
		}
		k = b.PrevHash
		if b.From == u.Addr && b.To == u.Addr {
			continue
		}
		if b.From == u.Addr && b.To != "" {
			amt -= b.Amount
		}
		if b.To == u.Addr && b.From != "" {
			amt += b.Amount
		}
	}

	return amt
}

func (u User) Transfer(addr Address) {
	// TODO with gRPC
}

func (u User) TransferWithData(addr Address) {
	// TODO with gRPC
}
