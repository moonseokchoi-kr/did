package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func main() {
	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
		os.Exit(1)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			fmt.Printf("Failed to populate wallet contents: %s\n", err)
			os.Exit(1)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		fmt.Printf("Failed to get network: %s\n", err)
		os.Exit(1)
	}

	contract := network.GetContract("didchain")

	result, err := contract.EvaluateTransaction("queryAllDIDs")
	if err != nil {
		fmt.Printf("Failed to evaluate transaction: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(result))

	result, err = contract.SubmitTransaction("createCar", "did:wul:ODc3MmYyZmYtOTBkZi00YTYxLWIyZWQtZTExYjlkM2NjMWE0", "1605001765", "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCB8AQeAlHFv7Z5jckKyYVuzJkjx0sTrHGMKd9rMr1R6TPYJQoGjvBzGx4sd/OL95INHmmPYWL3+P5fRgg5pp63yK1RrTuogUuNrZuT1P2Ux0/SZfonk14OXb8SpDGmTKqzlkPGL/d0T6kCZnUHCXbifYO6evMrqpE5ngRYukKFDQIDAQAB", "Grey", "E4YU+WOXZfDBrju9XGxKZYh7eJaox+h4kf1MlSRpfXLmSSUUCSpcEUJ7esDMqPVKaKqCe7G7smmk+cHA/t5bKMm+t7umu6io5jSvbPvTJOg75w+/3D5ts6LC7IUcpLC1u6eK598DLNlntbtH7l7gPllE536oSUfq5mFv6sDhl3c=")
	if err != nil {
		fmt.Printf("Failed to submit transaction: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(result))

	result, err = contract.EvaluateTransaction("readDID", "did:wul:ODc3MmYyZmYtOTBkZi00YTYxLWIyZWQtZTExYjlkM2NjMWE0")
	if err != nil {
		fmt.Printf("Failed to evaluate transaction: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(result))

	_, err = contract.SubmitTransaction("updateDID", "asdjkfsiadocjsajkwlqjlsdasd", "qopdfaklsjvcsakvnaskldjvaslasvjiacvjlkcjl", "did:wul:123593020")
	if err != nil {
		fmt.Printf("Failed to submit transaction: %s\n", err)
		os.Exit(1)
	}

	result, err = contract.EvaluateTransaction("queryCar", "CAR10")
	if err != nil {
		fmt.Printf("Failed to evaluate transaction: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(result))
}

func populateWallet(wallet *gateway.Wallet) error {
	credPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	err = wallet.Put("appUser", identity)
	if err != nil {
		return err
	}
	return nil
}
