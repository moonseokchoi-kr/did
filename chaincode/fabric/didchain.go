/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-identifier: Apache-2.0
*/

package chaincode

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

// CreateDid creates a new Did by placing the main Did details in the DidCollection
// that can be read by both organizations. The appraisal value is stored in the owners org specific collection.
func (s *SmartContract) CreateDid(ctx contractapi.TransactionContextInterface) error {

	// Get new Did from transient map
	transientMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return fmt.Errorf("error getting transient: %v", err)
	}

	// Did properties are private, therefore they get passed in transient field, instead of func args
	transientDidJSON, ok := transientMap["Did_properties"]
	if !ok {
		//log error to stdout
		return fmt.Errorf("Did not found in the transient map input")
	}
	type DidInput struct {
		context        string           `json:"context"`
		id             string           `json:"id"`
		created        string           `json:"created"`
		publickey      []PublicKey      `json:"publicKey"`
		authentication []Authentication `json:"authenticaiton"`
		service        []Service        `json:"service"`
	}
	var didInput DidInput
	err = json.Unmarshal(transientDidJSON, &didInput)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	if len(didInput.context) == 0 {
		return fmt.Errorf("context field must be a non-empty string")
	}
	if len(didInput.id) == 0 {
		return fmt.Errorf("id field must be a non-empty string")
	}

	// Make submitting client the owner
	Did := Did{
		context:        didInput.context,
		id:             didInput.id,
		created:        didInput.created,
		publickey:      didInput.publickey,
		authentication: didInput.authentication,
		service:        didInput.service,
	}
	DidJSONasBytes, err := json.Marshal(Did)
	if err != nil {
		return fmt.Errorf("failed to marshal Did into JSON: %v", err)
	}

	// Save Did to private data collection
	// Typical logger, logs to stdout/file in the fabric managed docker container, running this chaincode
	// Look for container name like dev-peer0.org1.example.com-{chaincodename_version}-xyz
	err = ctx.GetStub().PutState(didInput.id, DidJSONasBytes)
	if err != nil {
		return fmt.Errorf("failed to put Did into private data collecton: %v", err)
	}

	return nil
}
