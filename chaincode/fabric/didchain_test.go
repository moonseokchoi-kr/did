package did

//TODO  look at the example file and change code to create fucntion
import (
	"did/chaincode/fabric/mocks"
	"encoding/json"
	"os"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-samples/asset-transfer-private-data/chaincode-go/chaincode"
	"github.com/stretchr/testify/require"
)

type transactionContext interface {
	contractapi.TransactionContextInterface
}

//go:generate counterfeiter -o mocks/chaincodestub.go -fake-name ChaincodeStub . chaincodeStub
type chaincodeStub interface {
	shim.ChaincodeStubInterface
}
type DidInput struct {
	context        string           `json:"context"` //Type is used to distinguish the various types of objects in state database
	id             string           `json:"id"`
	created        string           `json:"created"`
	publickey      []PublicKey      `json:"publicKey"`
	authentication []Authentication `json:"authenticaiton"`
	service        []Service        `json:"service"`
}

const assetCollectionName = "assetCollection"
const myOrg1Msp = "Org1Testmsp"
const myOrg1Clientid = "myOrg1Userid"
const myOrg1PrivCollection = "Org1TestmspPrivateCollection"
const myOrg2Msp = "Org2Testmsp"
const myOrg2Clientid = "myOrg2Userid"
const myOrg2PrivCollection = "Org2TestmspPrivateCollection"

func TestCreateDid(t *testing.T) {
	transactionContext, chaincodeStub := prepMocksAsOrg1()
	assetTransferCC := SmartContract{}

	// No transient map
	err := assetTransferCC.CreateDid(transactionContext)
	require.EqualError(t, err, "Did not found in the transient map input")

	// transient map with incomplete asset data
	assetPropMap := map[string][]byte{
		"Did_properties": []byte("ill formatted property"),
	}
	chaincodeStub.GetTransientReturns(assetPropMap, nil)
	err = assetTransferCC.CreateDid(transactionContext)
	require.Error(t, err, "Expected error: transient map with incomplete asset data")
	require.Contains(t, err.Error(), "failed to unmarshal JSON")

	testAsset := &DidInput{
		id:      "id1",
		created: "7382920",
	}
	setReturnAssetPropsInTransientMap(t, chaincodeStub, testAsset)
	err = assetTransferCC.CreateDid(transactionContext)
	require.EqualError(t, err, "context field must be a non-empty string")

	// case when asset exists, GetPrivateData returns a valid data from ledger
	testAsset = &DidInput{
		context: "https://www.w3c.org.go",
		id:      "id1",
		created: "2830949",
	}
	setReturnAssetPropsInTransientMap(t, chaincodeStub, testAsset)
	err = assetTransferCC.CreateDid(transactionContext)
	require.EqualError(t, err, "this asset already exists: id1")
}

func TestCreateAssetSuccessful(t *testing.T) {
	transactionContext, chaincodeStub := prepMocksAsOrg1()
	assetTransferCC := chaincode.SmartContract{}
	testAsset := &DidInput{
		context:        "https://www.w3c.org.go",
		id:             "id1",
		created:        "210219384",
		publickey:      nil,
		authentication: nil,
		service:        nil,
	}
	setReturnAssetPropsInTransientMap(t, chaincodeStub, testAsset)
	err := assetTransferCC.CreateAsset(transactionContext)
	require.NoError(t, err)
	//Validate PutPrivateData calls
	calledCollection, calledID, _ := chaincodeStub.PutPrivateDataArgsForCall(0)
	require.Equal(t, assetCollectionName, calledCollection)
	require.Equal(t, "id1", calledID)

	expectedPrivateDetails := &chaincode.AssetPrivateDetails{
		ID:             "id1",
		AppraisedValue: 500,
	}
	assetBytes, err := json.Marshal(expectedPrivateDetails)
	calledCollection, calledID, calledAssetBytes := chaincodeStub.PutPrivateDataArgsForCall(1)
	require.Equal(t, myOrg1PrivCollection, calledCollection)
	require.Equal(t, "id1", calledID)
	require.Equal(t, assetBytes, calledAssetBytes)
}

func setReturnAssetPropsInTransientMap(t *testing.T, chaincodeStub *mocks.ChaincodeStub, testAsset *DidInput) []byte {
	assetBytes := []byte{}
	if testAsset != nil {
		var err error
		assetBytes, err = json.Marshal(testAsset)
		require.NoError(t, err)
	}
	assetPropMap := map[string][]byte{
		"Did_properties": assetBytes,
	}
	chaincodeStub.GetTransientReturns(assetPropMap, nil)
	return assetBytes
}

func prepMocksAsOrg1() (*mocks.TransactionContext, *mocks.ChaincodeStub) {
	return prepMocks(myOrg1Msp, myOrg1Clientid)
}
func prepMocksAsOrg2() (*mocks.TransactionContext, *mocks.ChaincodeStub) {
	return prepMocks(myOrg2Msp, myOrg2Clientid)
}
func prepMocks(orgMSP, clientID string) (*mocks.TransactionContext, *mocks.ChaincodeStub) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	clientIdentity := &mocks.ClientIdentity{}
	clientIdentity.GetMSPIDReturns(orgMSP, nil)
	clientIdentity.GetIDReturns(clientID, nil)
	//set matching msp ID using peer shim env variable
	os.Setenv("CORE_PEER_LOCALMSPID", orgMSP)
	transactionContext.GetClientIdentityReturns(clientIdentity)
	return transactionContext, chaincodeStub
}
