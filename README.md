# blockchain
BlockChain simple implementaion in Go & MongoDB

## Parts covered
- Add new Block in mongo db
- Genesis block
- Hash computation
- Previous Hash Matching
- Merkle root computation
- Merkle root Validation
- Proof of work with difficulty
- Nonce computation

## Terminal output
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

## mongo data
```
> db.Chain.find()
{ "_id" : ObjectId("5a6ebf16d99a76356e9c7487"), "Index" : 0, "Data" : "Genesis", "Time" : "Monday, 29-Jan-18 11:58:38 IST", "ThisHash" : "e455e150dac95c2ca0bdd6e4b0840e413c78afdda970a57924157a4c5ad6342b", "PrevHash" : "5feceb66ffc86f38d952786c6d696c79c2dbc239dd4e91b46729d73a27fb57e9", "Nonce" : 1, "Merkle" : "e455e150dac95c2ca0bdd6e4b0840e413c78afdda970a57924157a4c5ad6342b" }
{ "_id" : ObjectId("5a6ebf1cd99a76356e9c748f"), "Index" : 1, "Data" : "{data: 1002}", "Time" : "Monday, 29-Jan-18 11:58:44 IST", "ThisHash" : "000cd78ae1efe84446e00a9d00c6d96bb424b28216276497839f5b99ac790fda", "PrevHash" : "e455e150dac95c2ca0bdd6e4b0840e413c78afdda970a57924157a4c5ad6342b", "Nonce" : 277, "Merkle" : "c751ac686221bbdd2996690c166471f88696641c239d37184dac7ae28e27525a" }
{ "_id" : ObjectId("5a6ebf26d99a76356e9c7497"), "Data" : "{data: 1002}", "Time" : "Monday, 29-Jan-18 11:58:54 IST", "ThisHash" : "000eae4f23acdd3a0e50fc53d86669e88a5fbd25ee970d646b56394b5cd62f71", "PrevHash" : "000cd78ae1efe84446e00a9d00c6d96bb424b28216276497839f5b99ac790fda", "Nonce" : 1944, "Merkle" : "107bfed6b5a65958c526d34434bbe96964ff76b2c1e1fd4b5032bb67bcf80533", "Index" : 2 }
{ "_id" : ObjectId("5a6ebf3fd99a76356e9c749f"), "Index" : 3, "Data" : "{data: 1002}", "Time" : "Monday, 29-Jan-18 11:59:19 IST", "ThisHash" : "000089459d6e0b50965276f1ea4d7a9364f0582646620ae0d446f3e06069d679", "PrevHash" : "000eae4f23acdd3a0e50fc53d86669e88a5fbd25ee970d646b56394b5cd62f71", "Nonce" : 17106, "Merkle" : "daf726921436aaa43ebd0aa4cd4c8d8c179e1c6855e49bd1c78d49d763d7f8c3" }
{ "_id" : ObjectId("5a6ebfc2d99a76356e9c74b3"), "Merkle" : "47d63af22ec4c9da944677d5275c9094d2492c59570820aa06acecb9e655a274", "Index" : 4, "Data" : "{data: 1002}", "Time" : "Monday, 29-Jan-18 11:59:50 IST", "ThisHash" : "00000e1b50355d753248fc75153f8f4bc5f6acd34679c2d42b2652c839a59262", "PrevHash" : "000089459d6e0b50965276f1ea4d7a9364f0582646620ae0d446f3e06069d679", "Nonce" : 3473583 }
{ "_id" : ObjectId("5a6ec199d99a76356e9c74c1"), "ThisHash" : "00000c8376e44cf1f208e3439942f8faecb1fc5c3e8fd8d40c5045cc12a733a7", "PrevHash" : "00000e1b50355d753248fc75153f8f4bc5f6acd34679c2d42b2652c839a59262", "Nonce" : 957265, "Merkle" : "19e563b0d72e0362b9b226cae1dbe4323c0de9b72dd03ac4782c504a2e0d977c", "Index" : 5, "Data" : "{data: 1002}", "Time" : "Monday, 29-Jan-18 12:08:56 IST" }
{ "_id" : ObjectId("5a6ec1b5d99a76356e9c74c9"), "Index" : 6, "Data" : "{data: 1002}", "Time" : "Monday, 29-Jan-18 12:09:43 IST", "ThisHash" : "000003ba2c445fb180a8c08a3e232fdbf9af83e2aa95c4a06b253b0c81c5e1a0", "PrevHash" : "00000c8376e44cf1f208e3439942f8faecb1fc5c3e8fd8d40c5045cc12a733a7", "Nonce" : 201925, "Merkle" : "b2d0bcfdd1eca357abb3299446eb05255bd4ef08ce1787fdcbe676157c015641" }
{ "_id" : ObjectId("5a6ec1f6d99a76356e9c74d3"), "Data" : "{data: 1002}", "Time" : "Monday, 29-Jan-18 12:10:34 IST", "ThisHash" : "00000c92f03c6087fb61864c85e72f58188c1258977e7be1046ad07f099e857c", "PrevHash" : "000003ba2c445fb180a8c08a3e232fdbf9af83e2aa95c4a06b253b0c81c5e1a0", "Nonce" : 749155, "Merkle" : "b32dbf5dddb3f9f5c66599b8d715de7053430d5756cf7405b6612071c560009e", "Index" : 7 }
{ "_id" : ObjectId("5a6ec420d99a76356e9c74dd"), "Data" : "{data: 1002}", "Time" : "Monday, 29-Jan-18 12:19:58 IST", "ThisHash" : "0000063ededa6066bb0990e94a5b7ecaefc20f642a31c20dc76f31da674a75b6", "PrevHash" : "00000c92f03c6087fb61864c85e72f58188c1258977e7be1046ad07f099e857c", "Nonce" : 354955, "Merkle" : "cdde0f99aa13c75f8ca44638947251152bb982d1bded031399e3f772d99654d6", "Index" : 8 }
{ "_id" : ObjectId("5a6ec752d99a76356e9c74e8"), "Index" : 9, "Data" : "{data: 1002}", "Time" : "Monday, 29-Jan-18 12:33:45 IST", "ThisHash" : "000f164f42917af4e7e072560cabf1dd0ffa975dafb4fff4dd45bd4fb1ce9f23", "PrevHash" : "0000063ededa6066bb0990e94a5b7ecaefc20f642a31c20dc76f31da674a75b6", "Nonce" : 8913, "Merkle" : "e0bc61b8449251ad48e71ef1c19974442daa612f113e12ddb299e1d4e6dfc4c3" }
{ "_id" : ObjectId("5a6ec78fd99a76356e9c74f0"), "Index" : 10, "Data" : "{data: 1002}", "Time" : "Monday, 29-Jan-18 12:34:47 IST", "ThisHash" : "0009c7b0f24f81b424a623e9254d179019be1b23b7de45d448ae4d2cbcb77f1d", "PrevHash" : "000f164f42917af4e7e072560cabf1dd0ffa975dafb4fff4dd45bd4fb1ce9f23", "Nonce" : 16627, "Merkle" : "ccab8b646ff470bad1ee1bc02de5c61769f454c4efd06a903c67239046ea8259" }
{ "_id" : ObjectId("5a6ec8ebd99a76356e9c74f9"), "Time" : "Monday, 29-Jan-18 12:40:34 IST", "ThisHash" : "000873f1cbb56644c57cdb5766a1c5cc816afe9bf781e9a12527a7bb38b02a6c", "PrevHash" : "0009c7b0f24f81b424a623e9254d179019be1b23b7de45d448ae4d2cbcb77f1d", "Nonce" : 8021, "Merkle" : "5c69a6fdecc5a2235fea6632c074b14ec784e18cd2d972f7c37b735c2ea4df2d", "Index" : 11, "Data" : "{data: 1002}" }

```
;)
