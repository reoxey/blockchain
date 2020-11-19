package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/reoxey/blockchain/redis"
)

type Chain struct {
	Block
	db *redis.Conn
}

type Block struct {
	Index     int
	From      string
	To        string
	Data      string
	Time      string
	Amount    float64
	Incent    float64
	ThisHash  string
	PrevHash  string
	Nonce     int
	Proof     string
	Iteration int
}

type genesis struct {
	Balances interface{}
	Nonce    int
	Date     string
	Incent   float64
}

type Chainer interface {
	Add(from, to, data string, amt float64) error
}

var (
	ErrIntegrityFailed    = errors.New("blockchain: integrity failed")
	ErrProofOfWorkFailed  = errors.New("blockchain: proof of work return false")
	ErrFromToSame         = errors.New("blockchain: from and to addresses are same")
	ErrInvalidUserAddress = errors.New("blockchain: invalid or empty user account address")
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

func (c Chain) Add(from, to, data string, amt float64) error {

	if from == to {
		return ErrFromToSame
	}
	if from == "" || to == "" {
		return ErrInvalidUserAddress
	}

	b := c.Block

	b.Index++
	b.PrevHash = b.ThisHash
	b.From = from
	b.To = to
	b.Data = data
	b.Amount = amt
	b.Time = time.Now().String()
	b.ThisHash = hashThis(b)

	fmt.Printf("%+v\n", b)

	if !c.checkIntegrity() {
		return ErrIntegrityFailed
	}

	if proofOfWork(b) {
		if e := put(b, c.db, b.ThisHash); e != nil {
			return e
		}
		if e := put(b, c.db, "LAST"); e != nil {
			return e
		}
		return nil
	}

	return ErrProofOfWorkFailed
}

func (c *Chain) genesis() error {

	j, e := ioutil.ReadFile("genesis.json")
	if e != nil {
		return e
	}

	var g genesis
	e = json.Unmarshal(j, &g)
	if e != nil {
		return e
	}

	hash := fmt.Sprintf("%x", sha256.Sum256(j))

	b := Block{
		Data:     "GENESIS",
		Time:     g.Date,
		ThisHash: hash,
		Nonce:    g.Nonce,
		Incent:   g.Incent,
	}
	if e = put(b, c.db, "GENESIS"); e != nil {
		return e
	}

	x := 1
	m, ok := g.Balances.(map[string]interface{})
	if !ok {
		return errors.New("balances type invalid")
	}

	for k, v := range m {

		amt, ok := v.(float64)
		if !ok {
			amt = 0
		}

		b := Block{
			Index:    x,
			From:     "LORD",
			To:       k,
			Data:     k + " Balance",
			Time:     time.Now().String(),
			PrevHash: hash,
			Amount:   amt,
			Nonce:    g.Nonce,
			Incent:   g.Incent,
		}

		hash = hashThis(b)
		b.ThisHash = hash

		if e = put(b, c.db, hash); e != nil {
			return e
		}
		x++
	}

	b = Block{
		Index:    x,
		Time:     time.Now().String(),
		PrevHash: hash,
		Nonce:    g.Nonce,
		Incent:   g.Incent,
	}
	b.ThisHash = hashThis(b)
	if e = put(b, c.db, b.ThisHash); e != nil {
		return e
	}
	if e = put(b, c.db, "LAST"); e != nil {
		return e
	}
	c.Block = b

	return nil
}

func put(b Block, db *redis.Conn, k string) error {

	s, e := json.Marshal(b)
	if e != nil {
		return e
	}

	fmt.Println(string(s))

	if e = db.Set(k, string(s)); e != nil {
		return e
	}

	return nil
}

func hashThis(_b Block) string {
	buff := new(bytes.Buffer)
	gob.NewEncoder(buff).Encode(_b)
	return fmt.Sprintf("%x", sha256.Sum256(buff.Bytes()))
}

func (c *Chain) getState() bool {
	var b Block
	if e := c.db.GetJSON("LAST", &b); e != nil {
		return false
	}

	if b.ThisHash == "" {
		return false
	}

	c.Block = b

	return true
}

func (c *Chain) checkIntegrity() bool {
	return true
}

func proofOfWork(b Block) bool {
	return true
}
