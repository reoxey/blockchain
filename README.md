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

0
Hash Re-computation Passed: 000c0b908f37fd643c1e2a6e66cbc868a9e7ee670f1bd146374f38cee5f471a0
1
Hash Re-computation Passed: 000a395471de996cb78841ddf8aa5e09a342e61c04d21e54039c5689eacec060
Hash Matched with previous block: 000a395471de996cb78841ddf8aa5e09a342e61c04d21e54039c5689eacec060
Merkle Root Matched: f1a978d329b60e30d1eb07d136ae8de2fa9be08645008d374d76a91bc16fdcea
2
Hash Re-computation Passed: 000ae35b143cabed4d221b1b7a2a062659731a74b778e2b0f97518a30e1c8c43
Hash Matched with previous block: 000ae35b143cabed4d221b1b7a2a062659731a74b778e2b0f97518a30e1c8c43
Merkle Root Matched: bc00af800a7dbcd71807ddfbd7cc5ce72a5ab4f7aefc6bd493c83d722ccdc73c
3
Hash Re-computation Passed: 000ac87bc041ecdb66bf41d282170d9d8692b328f1737b9299789bfc812e6ec3
Hash Matched with previous block: 000ac87bc041ecdb66bf41d282170d9d8692b328f1737b9299789bfc812e6ec3
Merkle Root Matched: 38970f410f4299e21ae4910ff5727e9d0cb929f97caad49b5f07cf291022616f
4
Hash Re-computation Passed: 0001c16750cc84e5738e3a487ca630319d1f719094134dc957ffd3d66d4ad690
Hash Matched with previous block: 0001c16750cc84e5738e3a487ca630319d1f719094134dc957ffd3d66d4ad690
Merkle Root Matched: 372de6dd73c8c55094570754a484ab70ee0f5075c784fffd24e217aaa0601258
5
Hash Re-computation Passed: 000d4f0dc6aa2f81b5fa9f7cc17c1018e3bba64847b5215528b25081d521182d
Hash Matched with previous block: 000d4f0dc6aa2f81b5fa9f7cc17c1018e3bba64847b5215528b25081d521182d
Merkle Root Matched: 9db2abc33d4f3f7b8a22af2e97d558374e4709181b661b174fd0188114106c2f
6
Hash Re-computation Passed: d7f0a3be884ddd5a30252ef8dbb6cbfd69bfed6f62a91193ee001a70703266ee
Hash Matched with previous block: d7f0a3be884ddd5a30252ef8dbb6cbfd69bfed6f62a91193ee001a70703266ee
Merkle Root Matched: 9cc11bc2c1459276656adc4efeadca459939c6f83e526b65b534398156755c81
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
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining....................................................................................................
Mining......................
Block Mined with Hash: 0003369f18f9f7e123a9e3e34b9039bc322f61665337322714712353c0f2d43c
Nonce: 2622
Time Taken to mine: 71.95244ms
New Block Added

```

## mongo data
```
> db.Chain.find()
{ "_id" : ObjectId("5a6ee8c9d99a76356e9c7541"), "Time" : "Monday, 29-Jan-18 14:56:33 IST", "ThisHash" : "d7f0a3be884ddd5a30252ef8dbb6cbfd69bfed6f62a91193ee001a70703266ee", "PrevHash" : "5feceb66ffc86f38d952786c6d696c79c2dbc239dd4e91b46729d73a27fb57e9", "Nonce" : 1, "Merkle" : "d7f0a3be884ddd5a30252ef8dbb6cbfd69bfed6f62a91193ee001a70703266ee", "Index" : 0, "Data" : "Genesis" }
{ "_id" : ObjectId("5a6ee8d7d99a76356e9c754a"), "Merkle" : "9cc11bc2c1459276656adc4efeadca459939c6f83e526b65b534398156755c81", "Index" : 1, "Data" : "{data: 1002}", "Time" : "Monday, 29-Jan-18 14:56:47 IST", "ThisHash" : "000d4f0dc6aa2f81b5fa9f7cc17c1018e3bba64847b5215528b25081d521182d", "PrevHash" : "d7f0a3be884ddd5a30252ef8dbb6cbfd69bfed6f62a91193ee001a70703266ee", "Nonce" : 2239 }
{ "_id" : ObjectId("5a6ee8ded99a76356e9c7552"), "Index" : 2, "Data" : "{data: 1002}", "Time" : "Monday, 29-Jan-18 14:56:54 IST", "ThisHash" : "0001c16750cc84e5738e3a487ca630319d1f719094134dc957ffd3d66d4ad690", "PrevHash" : "000d4f0dc6aa2f81b5fa9f7cc17c1018e3bba64847b5215528b25081d521182d", "Nonce" : 9737, "Merkle" : "9db2abc33d4f3f7b8a22af2e97d558374e4709181b661b174fd0188114106c2f" }
{ "_id" : ObjectId("5a6eed14d99a76356e9c757f"), "PrevHash" : "0001c16750cc84e5738e3a487ca630319d1f719094134dc957ffd3d66d4ad690", "Nonce" : 5673, "Merkle" : "372de6dd73c8c55094570754a484ab70ee0f5075c784fffd24e217aaa0601258", "Index" : 3, "Data" : "{data: 1002}", "Time" : "Monday, 29-Jan-18 15:14:52 IST", "ThisHash" : "000ac87bc041ecdb66bf41d282170d9d8692b328f1737b9299789bfc812e6ec3" }
{ "_id" : ObjectId("5a6eed1bd99a76356e9c7587"), "ThisHash" : "000ae35b143cabed4d221b1b7a2a062659731a74b778e2b0f97518a30e1c8c43", "PrevHash" : "000ac87bc041ecdb66bf41d282170d9d8692b328f1737b9299789bfc812e6ec3", "Nonce" : 2541, "Merkle" : "38970f410f4299e21ae4910ff5727e9d0cb929f97caad49b5f07cf291022616f", "Index" : 4, "Data" : "{data: 1002}", "Time" : "Monday, 29-Jan-18 15:14:59 IST" }
{ "_id" : ObjectId("5a6eed20d99a76356e9c758f"), "Index" : 5, "Data" : "{data: 1002}", "Time" : "Monday, 29-Jan-18 15:15:04 IST", "ThisHash" : "000a395471de996cb78841ddf8aa5e09a342e61c04d21e54039c5689eacec060", "PrevHash" : "000ae35b143cabed4d221b1b7a2a062659731a74b778e2b0f97518a30e1c8c43", "Nonce" : 1386, "Merkle" : "bc00af800a7dbcd71807ddfbd7cc5ce72a5ab4f7aefc6bd493c83d722ccdc73c" }
{ "_id" : ObjectId("5a6eed26d99a76356e9c7597"), "Nonce" : 1927, "Merkle" : "f1a978d329b60e30d1eb07d136ae8de2fa9be08645008d374d76a91bc16fdcea", "Index" : 6, "Data" : "{data: 1002}", "Time" : "Monday, 29-Jan-18 15:15:09 IST", "ThisHash" : "000c0b908f37fd643c1e2a6e66cbc868a9e7ee670f1bd146374f38cee5f471a0", "PrevHash" : "000a395471de996cb78841ddf8aa5e09a342e61c04d21e54039c5689eacec060" }
{ "_id" : ObjectId("5a6eedf5d99a76356e9c75a0"), "Time" : "Monday, 29-Jan-18 15:18:37 IST", "ThisHash" : "0003369f18f9f7e123a9e3e34b9039bc322f61665337322714712353c0f2d43c", "PrevHash" : "000c0b908f37fd643c1e2a6e66cbc868a9e7ee670f1bd146374f38cee5f471a0", "Nonce" : 2622, "Merkle" : "73f220fced342310e0526dcdf0ecf17276f5453c6504780f2163705379f296c8", "Index" : 7, "Data" : "{data: 1002}" }
{ "_id" : ObjectId("5a6eee04d99a76356e9c75a8"), "Time" : "Monday, 29-Jan-18 15:18:51 IST", "ThisHash" : "00003ce4ffa263b467a3ce062a10ae0a7b0604a291271f30ac304220634f1f07", "PrevHash" : "0003369f18f9f7e123a9e3e34b9039bc322f61665337322714712353c0f2d43c", "Nonce" : 37868, "Merkle" : "7d71da0de7cda05414d077e5f96182a9f65ce1593a36785e227389a49a49bda0", "Index" : 8, "Data" : "{data: 1002}" }

```
;)
