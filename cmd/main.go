package main

import (
	"fmt"
	"log"

	block "github.com/reoxey/blockchain"
	"github.com/reoxey/blockchain/account"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	chn, e := block.Init()
	if e != nil {
		log.Fatalln(e)
	}

	acc, e := account.NewWithAddress("reoxey")
	if e != nil {
		log.Fatalln(e)
	}

	fmt.Println(acc.Balance())

	if e = chn.Add("reoxey", "johhny", "Enjoy!", 100); e != nil {
		log.Fatalln(e)
	}
}
