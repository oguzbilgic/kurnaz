package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/piotrnar/gocoin/btc"
	"github.com/steakknife/Golang-Koblitz-elliptic-curve-DSA-library/bitelliptic"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type AddressInfo struct {
	Word          string
	Hash160       string
	Address       string
	Key           string // Hex encoded
	NTx           int    `json:"n_tx"`
	TotalReceived int    `json:"total_received"`
	TotalSent     int    `json:"total_send"`
	FinalBalance  int    `json:"final_balance"`
}

type Response struct {
	AddressInfos []AddressInfo `json:"addresses"`
}

func main() {
	fileName := flag.String("file", "words.txt", "Path of the file containing words")
	firstWord := flag.String("first-word", "", "First word to start scanning")
	firstLetter := flag.String("first-letter", "", "First letter to start scanning")
	lastLetter := flag.String("last-letter", "", "First letter to stop scanning")
	frequency := flag.String("frequency", "1333ms", "Check sleep duration")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	pid := strconv.Itoa(os.Getpid())
	err = os.Mkdir(pid, os.ModePerm)
	if err != nil {
		panic(err)
	}

	configFile, err := os.Create(pid + "/conf")
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(configFile, "First Word: "+*firstWord)
	fmt.Fprintln(configFile, "First Letter: "+*firstLetter)
	fmt.Fprintln(configFile, "Last Letter: "+*lastLetter)
	fmt.Fprintln(configFile, "Frequency: "+*frequency)

	allRestulsFile, err := os.Create(pid + "/all.csv")
	if err != nil {
		panic(err)
	}

	usedRestulsFile, err := os.Create(pid + "/used.csv")
	if err != nil {
		panic(err)
	}

	activeRestulsFile, err := os.Create(pid + "/actve.csv")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	if *firstWord != "" {
		for scanner.Scan() {
			if scanner.Text() == *firstWord {
				break
			}
		}
	}

	if *firstLetter != "" {
		for scanner.Scan() {
			if []rune(scanner.Text())[0] == []rune(*firstLetter)[0] {
				break
			}
		}
	}

	sleepDuration, err := time.ParseDuration(*frequency)
	if err != nil {
		panic(err)
	}

	for scanner.Scan() {
		word := scanner.Text()

		if *lastLetter != "" && []rune(word)[0] == []rune(*lastLetter)[0] {
			break
		}

		privateKey := generatePrivateKeyFromString(word)
		publicKey := generatePublicKey(privateKey)
		address := generateHash160FromPublicKey(publicKey)
		addressInfo := newAddress(address)
		addressInfo.Word = word
		addressInfo.Key = hex.EncodeToString(privateKey)

		recordAddressInfo(allRestulsFile, addressInfo)

		if addressInfo.TotalReceived > 0 {
			recordAddressInfo(usedRestulsFile, addressInfo)
		}

		if addressInfo.FinalBalance > 0 {
			recordAddressInfo(activeRestulsFile, addressInfo)
		}

		time.Sleep(sleepDuration)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func recordAddressInfo(writer io.Writer, addressInfo *AddressInfo) {
	fmt.Fprintf(writer, "%v, %v, %v, %v, %v\n", addressInfo.Key, addressInfo.Address, addressInfo.Word, addressInfo.TotalReceived, addressInfo.FinalBalance)
}

func generatePrivateKeyFromString(word string) []byte {
	hash := sha256.New()
	hash.Write([]byte(word))
	return hash.Sum(nil)
}

func generatePublicKey(privateKey []byte) []byte {
	bitcurve := bitelliptic.S256()

	x, y := bitcurve.ScalarBaseMult(privateKey)
	return bitcurve.Marshal(x, y)
}

func generateHash160FromPublicKey(publicKey []byte) string {
	btcAddress := btc.NewAddrFromPubkey(publicKey, 0)
	return hex.EncodeToString(btcAddress.Hash160[:])
}

// Using Blockexplorer's api
func generateHashFromPublicKey(publicKey []byte) string {
	publicKeyString := hex.EncodeToString(publicKey)
	//resp, err := http.Get("http://blockchain.info/q/hashpubkey/" + publicKeyString)
	resp, err := http.Get("http://blockexplorer.com/q/hashpubkey/" + publicKeyString)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	hash, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func newAddress(addressHash string) *AddressInfo {
	resp, err := http.Get("http://blockchain.info/address/" + addressHash + "?format=json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var addressInfo AddressInfo

	err = json.Unmarshal(body, &addressInfo)
	if err != nil {
		panic(err)
	}

	return &addressInfo
}
