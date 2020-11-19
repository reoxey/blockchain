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
	block
	db *redis.Conn
}

type block struct {
	Index    int
	From     string
	To       string
	Data     string
	Time     string
	Amount   float64
	Incent   float64
	ThisHash string
	PrevHash string
	Nonce    int
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
	ErrIntegrityFailed   = errors.New("blockchain: integrity failed")
	ErrProofOfWorkFailed = errors.New("blockchain: proof of work return false")
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
	b := c.block

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

	if b.proofOfWork() {
		if e := c.put(b.ThisHash); e != nil {
			return e
		}
		if e := c.put("LAST"); e != nil {
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

	b := block{
		Data:     "GENESIS",
		Time:     g.Date,
		ThisHash: hash,
		Nonce:    g.Nonce,
		Incent:   g.Incent,
	}
	if e = c.put("GENESIS"); e != nil {
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

		b := block{
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

		if e = c.put(hash); e != nil {
			return e
		}
		x++
	}

	b = block{
		Index:    x,
		Time:     time.Now().String(),
		PrevHash: hash,
		Nonce:    g.Nonce,
		Incent:   g.Incent,
	}
	b.ThisHash = hashThis(b)
	if e = c.put(b.ThisHash); e != nil {
		return e
	}
	if e = c.put("LAST"); e != nil {
		return e
	}
	c.block = b

	return nil
}

func (c *Chain) put(k string) error {

	s, e := json.Marshal(c.block)
	if e != nil {
		return e
	}

	fmt.Println(string(s))

	if e = c.db.Set(k, string(s)); e != nil {
		return e
	}

	return nil
}

func hashThis(_b block) string {
	buff := new(bytes.Buffer)
	gob.NewEncoder(buff).Encode(_b)
	return fmt.Sprintf("%x", sha256.Sum256(buff.Bytes()))
}

func (c *Chain) getState() bool {
	var b block
	if e := c.db.GetJSON("LAST", &b); e != nil {
		return false
	}

	if b.ThisHash == "" {
		return false
	}

	c.block = b

	return true
}

func (c *Chain) checkIntegrity() bool {
	return true
}

func (b *block) proofOfWork() bool {
	return true
}
