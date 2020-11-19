package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
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

	fmt.Printf("%+v\n", b)

	if !c.checkIntegrity() {
		return ErrIntegrityFailed
	}

	proofOfWork(&b)

	if e := put(b, c.db, b.ThisHash); e != nil {
		return e
	}
	if e := put(b, c.db, "LAST"); e != nil {
		return e
	}
	c.Block = b

	return nil
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
		Data:       "GENESIS",
		Time:       g.Date,
		ThisHash:   hash,
		Difficulty: g.Difficulty,
		Incent:     g.Incent,
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
			Index:      x,
			From:       "LORD",
			To:         k,
			Data:       k + " Balance",
			Time:       time.Now().String(),
			PrevHash:   hash,
			Amount:     amt,
			Difficulty: g.Difficulty,
			Incent:     g.Incent,
		}

		hash = hashThis(b)
		b.ThisHash = hash

		if e = put(b, c.db, hash); e != nil {
			return e
		}
		x++
	}

	b = Block{
		Index:      x,
		Time:       time.Now().String(),
		PrevHash:   hash,
		Difficulty: g.Difficulty,
		Incent:     g.Incent,
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
	_b.ThisHash = ""
	_b.Duration = 0
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

	b := Block{}

	if e := c.db.GetJSON("LAST", &b); e != nil {
		return false
	}
	k := b.ThisHash

	for b.Index > 1 {
		if e := c.db.GetJSON(k, &b); e != nil {
			fmt.Println("---", k, e)
			return false
		}

		hash := b.ThisHash
		if k != hash {
			return false
		}
		k = b.PrevHash

		computedHash := hashThis(b)
		fmt.Println(computedHash, hash)
		if hash != computedHash {
			return false
		}
	}

	return true
}

func proofOfWork(b *Block) {
	start := time.Now()
	hash := ""

	if b.Duration < 1_000_000_000 {
		b.Difficulty++
	} else {
		b.Difficulty--
	}

	for {
		b.Nonce++
		hash = hashThis(*b) // TODO goroutine and channels
		if hash[:b.Difficulty] == zeros(b.Difficulty) {
			break
		}
		if b.Nonce%100 == 0 {
			fmt.Println()
			fmt.Print("Mining")
		}
		fmt.Print(".")
	}
	b.Duration = time.Since(start)
	b.ThisHash = hash
}

func zeros(level int) string {
	var z strings.Builder
	for level != 0 {
		z.WriteString("0")
		level--
	}
	return z.String()
}
