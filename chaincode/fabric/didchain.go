/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-identifier: Apache-2.0
*/

package did

import (
	"encoding/json"
	"fmt"
	"time"

	"did/chaincode/encrypte"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract of this fabric sample
type SmartContract struct {
	contractapi.Contract
}

// Did describes main Did details that are visible to all organizations
type Did struct {
	Context        string           `json:"context"`
	ID             string           `json:"id"`
	Created        int64            `json:"created"`
	Updated        int64            `json:updated`
	Publickey      []PublicKey      `json:"publicKey"`
	Authentication []Authentication `json:"authenticaiton"`
	Service        []Service        `json:"service"`
}

//PublicKey is save the key for authenfication
type PublicKey struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	PublicKey string `json:"publicKeybase58"`
	Created   int64  `json:"created"`
	Revoked   int64  `json:"revoked"`
}

//Authentication id useing authentication information when verify to id
type Authentication struct {
	ID         string `json:"id"`
	Credential string `json:"credentialDefinition"`
	Publickey  string `json:"publicKeyBase58"`
	Type       string `json:"type"`
}

//Service is kind of use the id
type Service struct {
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndPoint`
}

//InitDID initialize did
func (s *SmartContract) InitDID(ctx contractapi.TransactionContextInterface) (string, error) {
	did := &Did{
		Context: "https://www.did.com",
		ID:      encrypte.GetSpecificID(),
		Created: time.Now().Unix(),
		Service: initService(),
	}
	didJSON, err := json.Marshal(did)
	if err != nil {
		fmt.Errorf("Unexpected Error Converting JSON!! : %q", err)
	}
	err = ctx.GetStub().PutState(did.ID, didJSON)
	if err != nil {
		return "", fmt.Errorf("failed to put to world state. %v", err)
	}

	return did.ID, nil
}

//initService make base service
func initService() []Service {
	services := make([]Service, 2)
	services[0].Type = "Authentification"
	services[0].ServiceEndpoint = "did.com/auth"
	services[1].Type = "LogIn"
	services[1].ServiceEndpoint = "did.com/login"

	return services
}

//initAuth make base Auth info
func initAuth(info string, pubkeyID string) Authentication {
	auth := Authentication{
		ID:         encrypte.GetSpecificID(),
		Credential: info,
		Publickey:  pubkeyID,
		Type:       "ECDSA",
	}

	return auth
}

// CreateDID creates a new Did by placing the main Did details in the DidCollection
// that can be read by both organizations. The appraisal value is stored in the owners org specific collection.
func (s *SmartContract) CreateDID(ctx contractapi.TransactionContextInterface, msg string) error {
	pubkey, id, credential := encrypte.DecodeJwt(msg)
	didJSON, exists, err := s.DidExists(ctx, id)
	if err != nil {
		fmt.Errorf("Unexpected error!!")
	}
	if !exists {
		var did Did
		var pubkeyJSON PublicKey

		err = json.Unmarshal(pubkey, &pubkeyJSON)
		err = json.Unmarshal(didJSON, &did)

		did.Publickey[0] = pubkeyJSON
		did.Authentication[0] = initAuth(credential, pubkeyJSON.ID)
		didJSON, err := json.Marshal(did)

		if err != nil {
			fmt.Errorf("Unexpected Error : %q", err)
		}
		return ctx.GetStub().PutState(id, didJSON)
	} else {
		return fmt.Errorf("Don't exsit did!")
	}

}

//UpdatedDID updated publickey in did
func (s *SmartContract) UpdatedDID(ctx contractapi.TransactionContextInterface, msg string) error {
	pubkey, id, _ := encrypte.DecodeJwt(msg)
	didJSON, exists, err := s.DidExists(ctx, id)
	if !exists && err != nil {
		return fmt.Errorf("DID didn't exisits")
	} else {
		var did Did
		var pubkeyJSON PublicKey

		err = json.Unmarshal(pubkey, &pubkeyJSON)
		err = json.Unmarshal(didJSON, &did)

		did.Publickey[0].Revoked = time.Now().Unix()
		did.Publickey[1] = pubkeyJSON
		did.Authentication[0].Publickey = pubkeyJSON.ID
		did.Updated = time.Now().Unix()
		didJSON, err := json.Marshal(did)

		if err != nil {
			fmt.Errorf("Unexpected Error : %q", err)
		}
		return ctx.GetStub().PutState(id, didJSON)
	}
}

//ReadDID find did in chaincode and watch the information
//when makes application after change the function
func (s *SmartContract) ReadDID(ctx contractapi.TransactionContextInterface, id, query string) (Did, error) {
	didJSON, err := ctx.GetStub().GetState(id)

	var did Did
	err = json.Unmarshal(didJSON, &did)
	if err != nil {
		return did, fmt.Errorf("Unexpected error : %q", err)
	}
	return did, nil
}

//DidExists check the did exist in the chaincode
func (s *SmartContract) DidExists(ctx contractapi.TransactionContextInterface, id string) ([]byte, bool, error) {
	didJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return didJSON, didJSON != nil, nil
}
