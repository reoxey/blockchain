/**
BlockChain working in Go & MongoDB
reoxey
 */
package main

import(
	"time"
	"fmt"
	"crypto/sha256"
    "bytes"
    "gopkg.in/mgo.v2/bson"
	"github.com/reoxey/mongo"
	"encoding/json"
	"encoding/gob"
)

const BLOCKCHAIN = "BlockChain"
const CHAIN = "Chain"

var i int
var prevH string
var root string

type block struct {
	Index int
	Data string
	Time string
	ThisHash string
	PrevHash string
	Nonce uint64
	Merkle string
}

func NewBlock(_data string, level int) *block {

	NowTime := time.Now().Format(time.RFC850);

	_b1 := &block{
		Index: i,
		Data: _data,
		Time: NowTime,
		PrevHash: prevH,
		Nonce: 0,
	}
	_hash, _nonce := ProofOfWork(level, _b1)
	_merkle := MerkleRoot(_hash, root)
	
	_b := &block{
		Index: i,
		Data: _data,
		Time: NowTime,
		PrevHash: prevH,
		ThisHash: _hash,
		Merkle: _merkle,
		Nonce: _nonce,
	}
	
	return _b
}

func Init(){
	i = 0
	prevH = fmt.Sprintf("%x", sha256.Sum256([]byte("0")))
	root = ""
}

func HashThis(_b *block) string{
	buff := new(bytes.Buffer)
	e := gob.NewEncoder(buff).Encode(_b)
	if e != nil {
		fmt.Println(e)
	}
	return fmt.Sprintf("%x", sha256.Sum256(buff.Bytes()))
}

func MerkleRoot(_hash, _root string) string{

	if len(_root) == 0{
		return _hash
	}

	ss := _root + _hash

	return fmt.Sprintf("%x", sha256.Sum256([]byte(ss)))
}

func AddGenesisBlock(){
	Init()
	newBlock := NewBlock("Genesis",0)
	
	var ino map[string]interface{}
    inrec, _ := json.Marshal(newBlock)
    json.Unmarshal(inrec, &ino)
    
	mongo.Insert(ino)
}

func AddBlock(doc map[string]interface{}, data string, level int){

	index := int(doc["Index"].(float64))
	index = index + 1
	prev,_ := doc["ThisHash"].(string)
	merkle,_ := doc["Merkle"].(string)

	i = index
	prevH = prev
	root = merkle
	
	newBlock := NewBlock(data, level)
	
	var ino map[string]interface{}
    inrec, _ := json.Marshal(newBlock)
    json.Unmarshal(inrec, &ino)
    
	mongo.Insert(ino)
	fmt.Println("New Block Added")
}

func CheckIntegrity(docs []bson.M) bool {

	length := len(docs)

	if length == 0 {
		AddGenesisBlock()
		fmt.Println("Genesis Block Added")
		return false
	} else if length == 1 {
		return true
	} else {
		var prev, merkle, hash string
		for k, v := range docs {

			_prevH, _ := v["PrevHash"].(string)
			_merkle, _ := v["Merkle"].(string)
			_hash, _ := v["ThisHash"].(string)

			if k == 0 {
				prev = _prevH
				merkle = _merkle
				hash = _hash
				continue
			}

			_index := int(v["Index"].(float64))
			_data, _ := v["Data"].(string)
			_time, _ := v["Time"].(string)
			_nonce := uint64(v["Nonce"].(float64))

			_b1 := &block{
				Index:    _index,
				Data:     _data,
				Time:     _time,
				PrevHash: _prevH,
				Nonce: _nonce,
			}
			_computedHash := HashThis(_b1)

			if _hash == _computedHash {
				fmt.Println("Hash Re-computation Passed: " + _hash)
			} else {
				fmt.Println("Hash Re-computation Failed: " + _computedHash)
				break
			}

			if prev == _hash {
				fmt.Println("Hash Matched with previous block: " + prev)
				prev = _hash
			} else {
				fmt.Println("Hash MisMatched with previous block: " + prev + " = " + _hash)
				break
			}

			_computedMerkle := MerkleRoot(hash, _merkle)
			if merkle == _computedMerkle {
				fmt.Println("Merkle Root Matched: " + merkle)
				merkle = _merkle
			} else {
				fmt.Println("Merkle Root MisMatched: " + merkle + " = " + _computedMerkle)
				break
			}
			return true
		}
	}
	return false
}

func ProofOfWork(level int, _b *block) (string, uint64) {
	var Hash string
	start := time.Now()
	for{
		_b.Nonce++
		Hash = HashThis(_b)
		if Hash[:level] == ZeroCount(level) {
			break
		}
		if _b.Nonce % 100 == 0 {
			fmt.Println()
			fmt.Print("Mining")
		}
		fmt.Print(".")
	}
	TimeTaken := time.Since(start)
	fmt.Println()
	fmt.Print("Block Mined with Hash: "+Hash+" \nNonce: ")
	fmt.Println(_b.Nonce)
	fmt.Printf("Time Taken to mine: %s", TimeTaken)
	fmt.Println()
	return Hash, _b.Nonce
}

func ZeroCount(level int) string {
	zeros := ""
	for  {
		if level == 0 {
			break
		}
		zeros += "0"
		level--
	}
	return zeros
}

func main() {
	mongo.MongoDial(BLOCKCHAIN, CHAIN)

	o := bson.M{}
	y := bson.M{"Index": 1, "ThisHash": 1, "Merkle": 1}

	doc := mongo.FindLast(o, y, "-_id")

	fmt.Println()
	fmt.Println()

	docs := mongo.FindAll(o, "-_id")

	if CheckIntegrity(docs) {
		value := "{data: 1002}"
		AddBlock(doc, value, 3)
	}
}

//time.Now().Format(time.RFC850)
