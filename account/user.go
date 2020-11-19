package account

import (
	"crypto/sha256"
	"fmt"
	"time"

	block "github.com/reoxey/blockchain"
	"github.com/reoxey/blockchain/redis"
)

type Address string

type User struct {
	db   *redis.Conn
	Addr Address
}

type Account interface {
	Balance() float64
	Transfer(Address)
}

func New() (Account, error) {

	d := []byte(time.Now().String())

	u := User{
		db:   redis.Connect("127.0.0.1"),
		Addr: Address(fmt.Sprintf("%x", sha256.Sum224(d))),
	}

	return u, nil
}

func NewWithAddress(addr Address) (Account, error) {

	u := User{
		db:   redis.Connect("127.0.0.1"),
		Addr: addr,
	}

	return u, nil
}

func (u User) Balance() float64 {
	b := block.Block{}
	var amt float64
	k := "LAST"

	for k != "GENESIS" {
		if e := u.db.GetJSON(k, &b); e != nil {
			return amt
		}
		k = b.PrevHash
		if b.From == string(u.Addr) && b.To == string(u.Addr) {
			continue
		}
		if b.From == string(u.Addr) && b.To != "" {
			amt -= b.Amount
		}
		if b.To == string(u.Addr) && b.From != "" {
			amt += b.Amount
		}
	}

	return amt
}

func (u User) Transfer(addr Address) {
	// TODO
}
