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
type Contract struct {
	Context    string `json:"context"`
	SellerID   string `json:"sellerid"`
	ConsumerID string `json:"consumerid"`
	Created    int64  `json:"created"`
	Contract   string `json:"contract"`
	Signature  string `json:"signature"`
}

type QueryResult struct {
	Key    string `json:"Key"`
	Record *Contract
}

//InitLedger initialize Ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	ctc := Contract{
		Context:    "http://wuldid..ddns.net",
		SellerID:   "did:wul:ZTBlMDEyZDItMzc5OS00YjJiLWI3NmYtNTUyNGVmZjI0YWRl",
		ConsumerID: "did:wul:N2FiY2IwY2QtZDMxNi00YzcwLTg3NjUtNDE1MDRhZDA0Y2Y4",
		Created:    1604972580,
		Contract:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJDb250cmFjdCIsInNlbGxlciI6IkFwcGxlIiwiY29uc3VtZXIiOiJNb29uIENob2kiLCJwcm9kdWN0IjoiaXBob24xMiIsInByaWNlIjoiMSwzMDAsMDAwIiwicGF5IjoiQ3JlZGl0Q2FyZCIsImlhdCI6MTYwNDk3MjU4MH0.jQB5jSh_8U-iL89UMhhLM5cBtP-oime6zkqSr_rI4Ho",
		Signature:  "kPZ0yrOx+IoFPmYlgfdGwa6mEjDLl9ShPgq7DKPzc7ivAIbju6SA5KhvhGi4IhfoudA1mFkb4zP9ycKXlc0XFrRYWXdWcylY6so1kvSRX5G+Ni15V3DJl517vsbI6ZlA4PGiAcWJV1DgBIECgbRrRxopGp38G7UZe3XzN2CAnwM=",
	}
	ctcJSON, err := json.Marshal(ctc)
	if err != nil {
		return fmt.Errorf("Unexpected Error Converting JSON!! : %q", err)
	}
	err = ctx.GetStub().PutState("2b1e6766-2310-11eb-adc1-0242ac120002", ctcJSON)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}

	return nil
}

// CreateDID creates a new Did by placing the main Did details in the DidCollection
// that can be read by both organizations. The appraisal value is stored in the owners org specific collection.
func (s *SmartContract) CreateContract(ctx contractapi.TransactionContextInterface, id string, sellerID string, consumerID string, created int64, contract string, signature string) error {
	ctc := Contract{
		Context:    "http://wuldid.ddns.net",
		SellerID:   sellerID,
		ConsumerID: consumerID,
		Created:    created,
		Contract:   contract,
		Signature:  signature,
	}

	exists, err := s.DidExists(ctx, id)
	if err != nil {
		fmt.Errorf("Unexpected error!! : %q", err)
	}
	if !exists {
		ctcJSON, _ := json.Marshal(ctc)
		return ctx.GetStub().PutState(id, ctcJSON)
	} else {
		return fmt.Errorf("Don't exsit did!")
	}

}

//ReadSignature Contract block get Signature
func (s *SmartContract) ReadSignature(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	ctcJSON, err := ctx.GetStub().GetState(id)
	fmt.Print(string(ctcJSON))
	if err != nil {
		return string(ctcJSON), fmt.Errorf("Unexpected error : %q", err)
	}
	ctc := new(Contract)
	_ = json.Unmarshal(ctcJSON, ctc)
	return ctc.Signature, nil
}

//ReadDID find did in chaincode and watch the information
//when makes application after change the function
func (s *SmartContract) ReadContract(ctx contractapi.TransactionContextInterface, id string) (string, error) {

	ctcJSON, err := ctx.GetStub().GetState(id)
	fmt.Print(string(ctcJSON))
	if err != nil {
		return string(ctcJSON), fmt.Errorf("Unexpected error : %q", err)
	}
	return string(ctcJSON), nil
}

// QueryAllContracts returns all cars found in world state
func (s *SmartContract) QueryAllContracts(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
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

		did := new(Contract)
		_ = json.Unmarshal(queryResponse.Value, did)

		queryResult := QueryResult{Key: queryResponse.Key, Record: did}
		results = append(results, queryResult)
	}

	return results, nil
}

//DidExists check the did exist in the chaincode
func (s *SmartContract) DidExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	ctcJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return ctcJSON != nil, fmt.Errorf("failed to read from world state: %v", err)
	}

	return ctcJSON != nil, nil
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
