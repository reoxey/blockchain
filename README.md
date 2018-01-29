# blockchain
BlockChain simple implementaion in GoLang & MongoDB

# Parts covered
- Add new Block in mongo db
- Genesis block
- Hash computation
- Previous Hash Matching
- Merkle root computation
- Merkle root Validation
- Proof of work with difficulty
- Nonce computation

Terminal output
```
reoxey@HELL:~/go/src/github.com/reoxey/block$ go run index.go 

Hash Re-computation Passed: 0000063ededa6066bb0990e94a5b7ecaefc20f642a31c20dc76f31da674a75b6
Hash Matched with previous block: 0000063ededa6066bb0990e94a5b7ecaefc20f642a31c20dc76f31da674a75b6
Merkle Root Matched: e0bc61b8449251ad48e71ef1c19974442daa612f113e12ddb299e1d4e6dfc4c3
...................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining...........................
Block Mined with Hash: 0009c7b0f24f81b424a623e9254d179019be1b23b7de45d448ae4d2cbcb77f1d 
Nonce: 16627
Time Taken to mine: 327.298753ms
New Block Added

```

;)
