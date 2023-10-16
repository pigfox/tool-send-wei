package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
)

var connection *ethclient.Client
var currentConfig Config

func setUp() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	currentConfig = getTestConfig()
	getConnection()
}

func main() {
	setUp()
	weisToSend := big.NewInt(0)
	// Replace the wallet addresses and private key with the actual values
	privateKey := "02198d878891ee72bca881660a25f82c877f237aa713be1657965f4646fa1489"
	fromAddress := "0x3b9359c491d65f7b50cfe8ca8417c424a1229748"
	//destination wallet
	toAddress := "0x6b25765e09a303294649a190c436e3317c2fa56b"

	// Get the balance of the from address
	balance, err := getBalance(fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Balance of %s: %s ETH", fromAddress, balance)

	zero := big.NewInt(0)
	comparison := balance.Cmp(zero)

	if comparison == 0 {
		adios()
	}

	// Transfer all funds from the from address to the to address
	err = transferFunds(privateKey, toAddress, weisToSend)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Funds transferred from %s to %s", fromAddress, toAddress)
	adios()
}

func getConnection() {
	client, err := ethclient.Dial("https://" + currentConfig.NetType + ".infura.io/v3/8cfea7aaa1384f1a87b6d5aa65119ea3")
	if err != nil {
		log.Fatal(err)
	}
	connection = client
}

func adios() {
	log.Fatal("Adios")
	os.Exit(0)
}
