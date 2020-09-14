/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-identifier: Apache-2.0
*/

package did

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	
)

// SmartContract of this fabric sample
type SmartContract struct {
	contractapi.Contract
}

// Did describes main Did details that are visible to all organizations
type Did struct {
	context        string           `json:"context"`
	id             string           `json:"id"`
	created        string           `json:"created"`
	publickey      []PublicKey      `json:"publicKey"`
	authentication []Authentication `json:"authenticaiton"`
	service        []Service        `json:"service"`
}

//PublicKey is save the key for authenfication
type PublicKey struct {
	id        string `json:"id"`
	Type      string `json:"type"`
	publicKey string `json:"publicKeybase58"`
	created   string `json:"created"`
	revoked   string `json:"revoked"`
}

//Authentication id useing authentication information when verify to id
type Authentication struct {
	id         string `json:"id"`
	credential string `json:"credentialDefinition"`
	publickey  string `json:"publicKeyBase58"`
	Type       string `json:"type"`
}

//Service is kind of use the id
type Service struct {
	Type            string `json:"type"`
	serviceEndpoint string `json:"serviceEndPoint`
}

func (s *SmartContract) InitDID(ctx contractapi.TransactionContextInterface) error {
	did := &Did{
		context : "https://www.did.com"
		id: encrypte.getSpecificID()
		created : time.Now().Unix()
	}
	didJSON, err := json.Marshal(did)

	err := ctx.GetStub().PutState(did.id, didJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}

	return nil
}

// CreateDid creates a new Did by placing the main Did details in the DidCollection
// that can be read by both organizations. The appraisal value is stored in the owners org specific collection.
func (s *SmartContract) CreateDid(ctx contractapi.TransactionContextInterface, msg string) error {
	pubkey, id := encrypte.decodeJwt(msg)
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		var pubkeyJSON PublicKey
		err = json.UnMarshal(pubkeyJSON, &PublicKey)
		did := &Did{
			context : "https://www.did.com",
			id: id,
			created: time.Now().Unix(),
			publicKey: pubkeyJSON
		}
		didJSON, err := json.Marshal(did)
		return ctx.GetStub().PutState(id, didJSON)
	}else{
		InitDID(ctx)
	}

}
