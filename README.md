# kurnaz

brute force bitcoin brain wallets

## usage

Compile the program using `go build` command. Then run the command bellow.
Or you can just download the compiled binary, but you will need to copy the
word files.

```bash
$ ./kurnaz -first-word issue -last-letter v -file data/words.txt -frequency 200ms
```

This will generate 4 files unider `PID/` directory.

* all.csv: all the words scanned
* used.csv: all of the used bitcoin wallets
* active.csv: bitcoin wallets with positive balance
* conf: command line flags used when process was executed

## log

Log files use csv format with the fields bellow

```csv
private_key, bitcoin_address, word, total_received, final_balance
```

## data

I have included few data files with words in them in `data/`directory.

* words.txt: ~60,000 common english words
* hacker-jargon.txt: hakcer jargon
* names.txt: common given names
* places.txt: common place names
* urban.txt: popular words from urban dictionary

## results

Here are the results from initial run:

```bash
$ wc -l results/*/*.csv 

		0 results/jargon/actve.csv
	 2324 results/jargon/all.csv
	   60 results/jargon/used.csv
		0 results/urban/actve.csv
	 3147 results/urban/all.csv
	   53 results/urban/used.csv
		0 results/words/actve.csv
	27589 results/words/all.csv
	 6941 results/words/used.csv
	40114 total
```

Here are few interesting examples:

```bash
$ cat results/*/used.csv | grep -v 5460

77af778b51abd4a3c51c5ddd97204a9c3ae614ebccb75a606c3b6865aed6744e, 162TRPRZvdgLVNksMoMyGJsYBfYtB4Q8tM, cat, 15000000, 0
44bd7ae60f478fae1061e11a7739f4b94d1daf917982d33b6fc8a01a63f89c21, 1A1EFfQmUYVEZWXrfeDWVhXiwa1oNUynqr, H, 100000, 0
3ef81cb18bdaac2f67a114146b7f9c8da4bf8ceef8021dfc2da4daa8c1416e52, 1JRC6dGSfC5LAtX482sk5hsXyv3qWUF2m3, hammer, 1000000, 0
b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9, 1CS8g7nwaxPPprb4vqcTVdLCuCRirsbsMb, hello world, 1000000, 0
4813494d137e1631bba301d5acab6e7bb7aa74ce1185d456565ef51d737677b2, 148qEts4TkouGRwvUMRFM8dB9MjxM6iCuN, root, 100000, 0
382132701c4733c3402706cfdd3c8fc7f41f80a88dce5428d145259a41c5f12f, 1J18GoeAeJnsCd9j46pmM1HJw8hrZG2A9i, superuser, 10920, 0
9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08, 1HKqKTMpBTZZ8H5zcqYEWYBaaWELrDEXeE, test, 4118760, 0
f0e4c2f76c58916ec258f246851bea091d14d4247a2fc3e18694461b1816e13b, 12SU5JgVwfR5bA7NGKEfZRT1Zi5yvPR4An, asdf, 500000, 0
ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb, 1HUBHMij46Hae75JPdWjeZ5Q7KaL7EFRSD, a, 1000000, 0
cd6357efdd966de8c0cb2f876cc89ec74ce35f0968e11743987084bd42fb8944, 19MxhZPumMt9ntfszzCTPmWNQeh6j6QqP2, dog, 1000000, 0
8d586f8b60df1ad7502b52fc50d2e0b6d0d9b5643875c054d821dfd8e6672bed, 13uuHfJ2s7eCyVL96DTVguvzg24gtAwys3, Hell, 100000, 0
62fe4843ff9d2d79c52749cb0073c8490e7e490f03b98911d3acd4661ea69b5b, 1NegSeECmz7xFnFQ4yE8QKPDmkyKLKZVEc, Jesus, 1000000, 0
686f746a95b6f836d7d70567c302c3f9ebb5ee0def3d1220ee9d4e9f34f5e131, 1Mm6ouhpHqbtahCRNYfTo7Art1fbmk7PcR, love, 1200000, 0
12e90b8e74f20fc0a7274cff9fcbae14592db12292757f1ea0d7503d30799fd2, 1LVL6qEhMQTbNtSBDfBkmzo5ZS1PwaKZWs, poop, 100000, 0
8e5eb603482f00768b60cb17f947e273d6aa7c82ffaf8e589a06f6e841c3cef8, 17ac4moXPanV5QzuXRCiBs8uzRSjeSos3h, qwertyuiopasdfghjklzxcvbnm, 100000000, 0
bc5fb9abe8d5e72eb49cf00b3dbd173cbf914835281fadd674d5a2b680e47d50, 1FnPDy9Dtke7PTyD2jmasZV54ozSD2tNpQ, aberdeen, 164000, 0
277089d91c0bdf4f2e6862ba7e4a07605119431f5d13f726dd352b06f1b206a9, 157ySnBjQBBsryeq3yfcQYANX15Xtav2T7, bytes, 200000, 0
77af778b51abd4a3c51c5ddd97204a9c3ae614ebccb75a606c3b6865aed6744e, 162TRPRZvdgLVNksMoMyGJsYBfYtB4Q8tM, cat, 15000000, 0
811eb81b9d11d65a36c53c3ebdb738ee303403cb79d781ccf4b40764e0a9d12a, 15Z16yvxv3oH6FBd83qkgo8AmzYcaSy2vX, chicken, 100000, 0
cd6357efdd966de8c0cb2f876cc89ec74ce35f0968e11743987084bd42fb8944, 19MxhZPumMt9ntfszzCTPmWNQeh6j6QqP2, dog, 1000000, 0
3ef81cb18bdaac2f67a114146b7f9c8da4bf8ceef8021dfc2da4daa8c1416e52, 1JRC6dGSfC5LAtX482sk5hsXyv3qWUF2m3, hammer, 1000000, 0
fd70ad909b94deb27b460692084d9f2b1dbc9df3c6bcfd3caee571e707031e3f, 1EYz2AhbVe2GJ1th1j5czNkAeSBViQfrUW, hitler, 100000, 0
4a07a4310034102668a862f2ec7d3ba7416937b2f85c90b38257cf5b13093b0c, 1CwjHYsPUc4Du8dx7AkdBJj4ebWC8bxkF3, icecream, 13491343, 0
```
