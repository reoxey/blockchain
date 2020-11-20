package block

import (
	"errors"
	"time"

	"github.com/reoxey/blockchain/redis"
)

type Chain struct {
	Block
	db *redis.Conn
}

type Block struct {
	Index      int
	From       string
	To         string
	Data       string
	Time       string
	Amount     float64
	Incent     float64
	ThisHash   string
	PrevHash   string
	Difficulty int
	Nonce      int
	Duration   time.Duration
}

type genesis struct {
	Balances   interface{}
	Difficulty int
	Date       string
	Incent     float64
}

type Chainer interface {
	Add(from, to, data string, amt float64) error
}

var (
	ErrIntegrityFailed    = errors.New("blockchain: integrity failed")
	ErrFromToSame         = errors.New("blockchain: from and to addresses are same")
	ErrInvalidUserAddress = errors.New("blockchain: invalid or empty user account address")
	ErrBalanceInvalid     = errors.New("balances type invalid")
)

func Init() (Chainer, error) {

	db := redis.Connect("127.0.0.1")

	c := Chain{
		db: db,
	}

	if ok := c.getState(); !ok {
		e := c.genesis()
		if e != nil {
			return nil, e
		}
	}

	return &c, nil
}
