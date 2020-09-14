package encrypte

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/dgrijalva/jwt-go"
	"github.com/segmentio/ksuid"
)

var msg = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
var pubkeyCurve = elliptic.P256()

/**
* make id for did
*author : choimoonseok
*date : 2020-09-06
 */
func getSpecificID() string {
	id := ksuid.New()
	return id.String()
}

/**
* string to byte array
* author : choimoonseok
* date : 2020-09-08
**/
func stringTobyte(msg string) []byte {
	return []byte(msg)
}

/**
* created ecdsaKey
*author : choimoonseok
*date : 2020-09-06
 */
func makeECDSAKey(msg string) (string, string, string) {

	privatekey := new(ecdsa.PrivateKey)
	privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)

	if err != nil {
		fmt.Println(err)
	}

	pubkey := privatekey.PublicKey
	siginature := signHash(msg, privatekey)
	return encodeKey(privatekey, pubkey, siginature)
}

/**
* encode base 58 character key
*author : choimoonseok
*date : 2020-09-06
 */
func encodeKey(privateKey *ecdsa.PrivateKey, publickey ecdsa.PublicKey, signiture []byte) (string, string, string) {
	privateByte := privateKey.D.Bytes()
	priBase58 := base58.CheckEncode(privateByte, 0)
	publicByte := elliptic.Marshal(publickey, publickey.X, publickey.Y)
	pubBase58 := base58.CheckEncode(publicByte, 0)
	sigBase58 := base58.CheckEncode(signiture, 0)

	return priBase58, pubBase58, sigBase58
}

/**
* hash to message usgin sha256
*author : choimoonseok
*date : 2020-09-06
 */
func hash(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	return h.Sum(nil)
}

/**
* created signature
*author : choimoonseok
*date : 2020-09-06
 */
func signHash(message string, privateKey *ecdsa.PrivateKey) []byte {
	signhash := hash(stringTobyte(message))
	signature, serr := ecdsa.SignASN1(rand.Reader, privateKey, signhash)
	if serr != nil {
		fmt.Println(serr)
	}

	return signature
}

/**
* decode base 58
*author : choimoonseok
*date : 2020-09-06
 */
func decodepubkey(publickey string) ecdsa.PublicKey {
	pubByte, _, err := base58.CheckDecode(publickey)
	if err != nil {
		fmt.Println(err)
	}
	x, y := elliptic.Unmarshal(pubkeyCurve, pubByte)
	pubkey := ecdsa.PublicKey{
		Curve: pubkeyCurve,
		X:     x,
		Y:     y,
	}
	return pubkey

}

/**
* decode base 58 for signature
*author : choimoonseok
*date : 2020-09-06
 */
func decodeSignature(signature string) []byte {
	sigbyte, _, err := base58.CheckDecode(signature)

	if err != nil {
		fmt.Println(err)
	}
	return sigbyte
}

/**
* verify key
*author : choimoonseok
*date : 2020-09-06
 */
func verify(publickey string, signature string, message string) bool {
	signhash := hash(stringTobyte(message))
	pubkey := decodepubkey(publickey)
	sig := decodeSignature(signature)
	verify := ecdsa.VerifyASN1(&pubkey, signhash, sig)

	return verify
}

func getJwt() (string, error) {

	expirationTime := time.Now().Add(5 * time.Minute)
	type Publickey struct {
		Type      string `json:"type"`
		publicKey string `json:"publicKeybase58"`
		created   int64  `json:"created"`
		revoked   int64  `json:"revoked"`
	}
	type Message struct {
		id        string
		publickey Publickey
		jwt.StandardClaims
	}
	pubkey := &Message{
		id: getSpecificID(),
		publickey: Publickey{
			Type:      "ecdsa",
			publicKey: "13n4s5tFAmoCYHLsnJ9k1nspszbuQgvjaFrmJ8cSbfLmHDGNDkc69XCExX9PpbDBLA25VK2GsvYXvXEi9xr1DWEbVfUJu8u",
			created:   time.Now().Unix(),
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, pubkey)

	tokenString, err := token.SignedString("JwtKey")

	if err != nil {
		return "", fmt.Errorf("Unexpected Error! : %q", err)
	} else {
		return tokenString, nil
	}
}

/**
* decode json web token
* author : choimoonseok
* date : 2020-09-13
 */
func decodeJwt(tokenString string) ([]byte, string) {
	type Publickey struct {
		Type      string `json:"type"`
		publicKey string `json:"publicKeybase58"`
		created   int64  `json:"created"`
		revoked   int64  `json:"revoked"`
	}
	type Message struct {
		id        string
		publickey Publickey
		jwt.StandardClaims
	}
	token, err := jwt.ParseWithClaims(tokenString, &Message{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("Jwtkey"), nil
	})
	if err != nil {
		fmt.Errorf("Errsign!!")
		return nil, ""
	}
	if message, ok := token.Claims.(*Message); ok && token.Valid {
		id := message.id
		pubkey, err := json.Marshal(message.publickey)
		if err != nil {
			return nil, ""
		}
		return pubkey, id

	} else {
		return nil
	}
}
