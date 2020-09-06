package chaincode

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"hash"
	"io"

	"github.com/btcsuite/btcutil/base58"
	"github.com/segmentio/ksuid"
)

func getSpecificID() string {
	id := ksuid.New()
	return id.String()
}

func makeECDSAKey() (string, string, string) {
	pubkeyCurve := elliptic.P256()
	privatekey := new(ecdsa.PrivateKey)
	privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)

	if err != nil {
		fmt.Println(err)
	}

	pubkey := privatekey.PublicKey
	siginature := signHash(getSpecificID(), privatekey)
	return encodeKey(privatekey, pubkey, siginature)
}

func encodeKey(privateKey *ecdsa.PrivateKey, publickey ecdsa.PublicKey, signiture []byte) (string, string, string) {
	privateByte := privateKey.D.Bytes()
	priBase58 := base58.CheckEncode(privateByte, 0)
	publicByte := elliptic.Marshal(publickey, publickey.X, publickey.Y)
	pubBase58 := base58.CheckEncode(publicByte, 0)
	sigBase58 := base58.CheckEncode(signiture, 0)

	return priBase58, pubBase58, sigBase58
}

func signHash(message string, privateKey *ecdsa.PrivateKey) []byte {
	var h hash.Hash
	h = md5.New()

	io.WriteString(h, message)

	signhash := h.Sum(nil)

	signature, serr := ecdsa.SignASN1(rand.Reader, privateKey, signhash)
	if serr != nil {
		fmt.Println(serr)
	}

	return signature
}

func verifySignature(publickey string, signature string, hash []byte) {}
