package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/steakknife/Golang-Koblitz-elliptic-curve-DSA-library/bitelliptic"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
	file, err := os.Open("words.txt")
	if err != nil {
		panic(err)
	}

	allRestulsFile, err := os.Create("all.csv")
	if err != nil {
		panic(err)
	}

	usedRestulsFile, err := os.Create("used.csv")
	if err != nil {
		panic(err)
	}

	activeRestulsFile, err := os.Create("actve.csv")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		privateKey := generatePrivateKeyFromString(word)
		publicKey := generatePublicKey(privateKey)
		address := generateHashFromPublicKey(publicKey)
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

		time.Sleep(1333 * time.Millisecond)
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

func generateHashFromPublicKey(publicKey []byte) string {
	publicKeyString := hex.EncodeToString(publicKey)
	resp, err := http.Get("http://blockchain.info/q/hashpubkey/" + publicKeyString)
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
