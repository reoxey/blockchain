package main

import (
	"log"

	block "github.com/reoxey/blockchain"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	chn, e := block.Init()
	if e != nil {
		log.Fatalln(e)
	}

	if e = chn.Add("reoxey", "johhny", "Enjoy!", 100); e != nil {
		log.Fatalln(e)
	}
}
