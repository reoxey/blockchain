package account

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/reoxey/blockchain/redis"
)

type Address string

type User struct {
	db   *redis.Conn
	Addr Address
}

type Account interface {
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

func (u User) Transfer(addr Address) {
	// TODO
}
