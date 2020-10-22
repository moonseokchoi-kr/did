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
func GetSpecificID() string {
	id := ksuid.New()
	return id.String()
}

/**
* string to byte array
* author : choimoonseok
* date : 2020-09-08
**/
func StringToByte(msg string) []byte {
	return []byte(msg)
}

/**
* created ecdsaKey
*author : choimoonseok
*date : 2020-09-06
 */
func MakeECDSAKey(msg string) (string, string, string) {

	privatekey := new(ecdsa.PrivateKey)
	privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)

	if err != nil {
		fmt.Println(err)
	}

	pubkey := privatekey.PublicKey
	siginature := SignHash(msg, privatekey)
	return EncodeKey(privatekey, pubkey, siginature)
}

/**
* encode base 58 character key
*author : choimoonseok
*date : 2020-09-06
 */
func EncodeKey(privateKey *ecdsa.PrivateKey, publickey ecdsa.PublicKey, Signiture []byte) (string, string, string) {
	privateByte := privateKey.D.Bytes()
	priBase58 := base58.CheckEncode(privateByte, 0)
	publicByte := elliptic.Marshal(publickey, publickey.X, publickey.Y)
	pubBase58 := base58.CheckEncode(publicByte, 0)
	sigBase58 := base58.CheckEncode(Signiture, 0)

	return priBase58, pubBase58, sigBase58
}

/**
* Hash to message usgin sha256
*author : choimoonseok
*date : 2020-09-06
 */
func Hash(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	return h.Sum(nil)
}

/**
* created Signature
*author : choimoonseok
*date : 2020-09-06
 */
func SignHash(message string, privateKey *ecdsa.PrivateKey) []byte {
	SignHash := Hash(StringToByte(message))
	Signature, serr := ecdsa.SignASN1(rand.Reader, privateKey, SignHash)
	if serr != nil {
		fmt.Println(serr)
	}

	return Signature
}

/**
* Decode base 58
*author : choimoonseok
*date : 2020-09-06
 */
func DecodePublicKey(publickey string) ecdsa.PublicKey {
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
* Decode base 58 for Signature
*author : choimoonseok
*date : 2020-09-06
 */
func DecodeSignature(Signature string) []byte {
	sigbyte, _, err := base58.CheckDecode(Signature)

	if err != nil {
		fmt.Println(err)
	}
	return sigbyte
}

/**
* Verify key
*author : choimoonseok
*date : 2020-09-06
 */
func Verify(publickey string, Signature string, message string) bool {
	SignHash := Hash(StringToByte(message))
	pubkey := DecodePublicKey(publickey)
	sig := DecodeSignature(Signature)
	verify := ecdsa.VerifyASN1(&pubkey, SignHash, sig)

	return verify
}

type Message struct {
	ID                   string
	CredentialDefinition string
	Publickey            Publickey
	jwt.StandardClaims
}
type Publickey struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	PublicKey string `json:"publicKeybase58"`
	Created   int64  `json:"created"`
	Revoked   int64  `json:"revoked"`
}
type CredentialDefinition struct {
	Name   string
	Birth  string
	Phone  string
	Age    int
	Gender string
	jwt.StandardClaims
}

func CreateCRD(name string, birth string, phone string, age int, gender string) CredentialDefinition {
	credentialDefinition := &CredentialDefinition{
		Name:           name,
		Birth:          birth,
		Phone:          phone,
		Age:            age,
		Gender:         gender,
		StandardClaims: jwt.StandardClaims{},
	}
	return *credentialDefinition
}

func CreatePubkey(id string, Type string) Publickey {
	crd := CreateCRD("MoonSeok", "1997-06-13", "010-1234-5678", 24, "M")
	t, _ := json.Marshal(crd)
	str := string(t)
	_, pubkey, _ := MakeECDSAKey(str)
	publickey := &Publickey{
		ID:        id,
		Type:      Type,
		PublicKey: pubkey,
		Created:   time.Now().Unix(),
	}

	return *publickey
}

func CreateMsg(id string, pubkey Publickey) *Message {
	expirationTime := time.Now().Add(5 * time.Hour)
	claims := &Message{
		ID:        id,
		Publickey: pubkey,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	return claims
}

/**
* GetJWT is Get JSON Web Token
 */
func GetJWT(claims Message) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("JwtKey"))

	if err != nil {
		return "", fmt.Errorf("Unexpected Error! : %q", err)
	} else {
		return tokenString, nil
	}
}

/**
* Decode json web token
* author : choimoonseok
* date : 2020-09-13
 */
func DecodeJwt(tokenString string) ([]byte, string, string) {
	token, err := jwt.ParseWithClaims(tokenString, &Message{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("JwtKey"), nil
	})
	if err != nil {
		fmt.Errorf("ErrSign!!")
		return nil, "", ""
	}
	if message, ok := token.Claims.(*Message); ok && token.Valid {
		id := message.ID
		credential := message.CredentialDefinition
		pubkey, err := json.Marshal(message.Publickey)

		if err != nil {
			return nil, "", ""
		}
		return pubkey, id, credential

	} else {
		return nil, "", ""
	}
}
