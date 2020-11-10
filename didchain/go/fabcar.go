/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

/*
// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}

// Car describes basic details of what makes up a car
type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *Car
}

// InitLedger adds a base set of cars to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	cars := []Car{
		Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko"},
		Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad"},
		Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo"},
		Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max"},
		Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana"},
		Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel"},
		Car{Make: "Chery", Model: "S22L", Colour: "white", Owner: "Aarav"},
		Car{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "Pari"},
		Car{Make: "Tata", Model: "Nano", Colour: "indigo", Owner: "Valeria"},
		Car{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro"},
	}

	for i, car := range cars {
		carAsBytes, _ := json.Marshal(car)
		err := ctx.GetStub().PutState("CAR"+strconv.Itoa(i), carAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

// CreateCar adds a new car to the world state with given details
func (s *SmartContract) CreateCar(ctx contractapi.TransactionContextInterface, carNumber string, make string, model string, colour string, owner string) error {
	car := Car{
		Make:   make,
		Model:  model,
		Colour: colour,
		Owner:  owner,
	}

	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carNumber, carAsBytes)
}

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) QueryCar(ctx contractapi.TransactionContextInterface, carNumber string) (*Car, error) {
	carAsBytes, err := ctx.GetStub().GetState(carNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if carAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", carNumber)
	}

	car := new(Car)
	_ = json.Unmarshal(carAsBytes, car)

	return car, nil
}

// QueryAllDIDs returns all cars found in world state
func (s *SmartContract) QueryAllDIDs(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		car := new(Car)
		_ = json.Unmarshal(queryResponse.Value, car)

		queryResult := QueryResult{Key: queryResponse.Key, Record: car}
		results = append(results, queryResult)
	}

	return results, nil
}

// ChangeCarOwner updates the owner field of car with given id in world state
func (s *SmartContract) ChangeCarOwner(ctx contractapi.TransactionContextInterface, carNumber string, newOwner string) error {
	car, err := s.QueryCar(ctx, carNumber)

	if err != nil {
		return err
	}

	car.Owner = newOwner

	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carNumber, carAsBytes)
}


*/

// SmartContract of this fabric sample
type SmartContract struct {
	contractapi.Contract
}

// Did describes main Did details that are visible to all organizations
type Did struct {
	Context        string    `json:"context"`
	ID             string    `json:"id"`
	Created        int64     `json:"created"`
	Updated        int64     `json:"updated"`
	Publickey      string    `json:"publicKey"`
	Authentication string    `json:"authenticaiton"`
	Service        []Service `json:"service"`
}

//Service is kind of use the id
type Service struct {
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndPoint"`
}

//QueryResult set query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *Did
}

//InitDID initialize did
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	id := "did:wul:123593020"

	services := make([]Service, 2)
	services[0].Type = "signedContract"
	services[0].ServiceEndpoint = "wuldid.ddns.net/signed"
	services[1].Type = "verifyContract"
	services[1].ServiceEndpoint = "wuldid.ddns.net/verify"
	did := &Did{
		Context:        "https://www.did.com",
		ID:             id,
		Created:        1603343627,
		Service:        services,
		Publickey:      "13n4s5tFAmoCYHLsnJ9k1nspszbuQgvjaFrmJ8cSbfLmHDGNDkc69XCExX9PpbDBLA25VK2GsvYXvXEi9xr1DWEbVfUJu8u",
		Authentication: "13n4s5tFAmoCYHLsnJ9k1nspszbuQgvjaFrmJ8cSbfLmHDGNDkc69XCExX9PpbDBLA25VK2GsvYXvXEi9xr1DWEbVfUJu8u",
	}
	didJSON, err := json.Marshal(did)
	if err != nil {
		return fmt.Errorf("Unexpected Error Converting JSON!! : %q", err)
	}
	err = ctx.GetStub().PutState(id, didJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}

	return nil
}

// CreateDID creates a new Did by placing the main Did details in the DidCollection
// that can be read by both organizations. The appraisal value is stored in the owners org specific collection.
func (s *SmartContract) CreateDID(ctx contractapi.TransactionContextInterface, id string, created int64, publickey string, auth string) error {
	exists, err := s.DidExists(ctx, id)
	services := make([]Service, 2)
	services[0].Type = "signedContract"
	services[0].ServiceEndpoint = "wuldid.ddns.net/signed"
	services[1].Type = "verifyContract"
	services[1].ServiceEndpoint = "wuldid.ddns.net/verify"
	did := Did{
		Context:        "http://wuldid.ddns.net",
		ID:             id,
		Created:        created,
		Publickey:      publickey,
		Authentication: auth,
		Service:        services,
	}
	if err != nil {
		fmt.Errorf("Unexpected error!! : %q", err)
	}
	if !exists {
		didJSON, _ := json.Marshal(did)
		return ctx.GetStub().PutState(id, didJSON)
	} else {
		return fmt.Errorf("Don't exsit did!")
	}

}

//UpdatedDID updated publickey in did
func (s *SmartContract) UpdatedDID(ctx contractapi.TransactionContextInterface, publickey string, auth string, id string) error {
	exists, err := s.DidExists(ctx, id)
	if !exists && err != nil {
		return fmt.Errorf("DID didn't exisits")
	} else {
		var did Did
		didJSON, err := ctx.GetStub().GetState(id)
		if err != nil {
			return fmt.Errorf("DID have problem")
		}
		err = json.Unmarshal(didJSON, &did)
		did.Authentication = auth
		did.Publickey = publickey
		didJSON, err = json.Marshal(did)
		if err != nil {
			return fmt.Errorf("Unexpected error : %q", err)
		}
		return ctx.GetStub().PutState(id, didJSON)
	}
}

//ReadDID find did in chaincode and watch the information
//when makes application after change the function
func (s *SmartContract) ReadDID(ctx contractapi.TransactionContextInterface, id string) (string, error) {

	didJSON, err := ctx.GetStub().GetState(id)
	fmt.Print(string(didJSON))
	if err != nil {
		return string(didJSON), fmt.Errorf("Unexpected error : %q", err)
	}
	return string(didJSON), nil
}

// QueryAllDIDs returns all cars found in world state
func (s *SmartContract) QueryAllDIDs(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		did := new(Did)
		_ = json.Unmarshal(queryResponse.Value, did)

		queryResult := QueryResult{Key: queryResponse.Key, Record: did}
		results = append(results, queryResult)
	}

	return results, nil
}

//DidExists check the did exist in the chaincode
func (s *SmartContract) DidExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	didJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return didJSON != nil, fmt.Errorf("failed to read from world state: %v", err)
	}

	return didJSON != nil, nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
}
