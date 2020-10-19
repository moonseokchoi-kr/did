package did

//TODO  look at the example file and change code to create fucntion
import (
	"did/chaincode/fabric/mocks"
	"os"
	"testing"

	"github.com/did/chaincode/encrypte"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/stretchr/testify/require"
)

type transactionContext interface {
	contractapi.TransactionContextInterface
}

//go:generate counterfeiter -o mocks/chaincodestub.go -fake-name ChaincodeStub . chaincodeStub
type chaincodeStub interface {
	shim.ChaincodeStubInterface
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

	// Init did
	id, err := assetTransferCC.InitDID(transactionContext)
	require.NoError(t, err)

	didJSON, err := chaincodeStub.GetState(id)
	if err != nil {
		t.Errorf("DID didn't exist : %q \n didJSON :  %q", err, didJSON)
	}
	id = "did:wul:" + encrypte.GetSpecificID()
	claims := encrypte.CreateMsg(id, encrypte.CreatePubkey(id+"#key1", "PublickeyECDSABase58"))
	msg, err := encrypte.GetJWT(*claims)
	if err != nil {
		t.Errorf("Unexpected Error : %q", err)
	}
	//create DID
	err = assetTransferCC.CreateDID(transactionContext, msg)

	require.NoError(t, err)
}

func TestUpdatedDID(t *testing.T) {
	transactionContext, _ := prepMocksAsOrg1()
	assetTransferCC := SmartContract{}
	id := "did:wul:" + encrypte.GetSpecificID()
	claims := encrypte.CreateMsg(id, encrypte.CreatePubkey(id+"#key1", "PublickeyECDSABase58"))
	msg, err := encrypte.GetJWT(*claims)
	if err != nil {
		t.Errorf("Unexpected Error : %q", err)
	}
	//create DID
	err = assetTransferCC.CreateDID(transactionContext, msg)
	did, err1 := assetTransferCC.ReadDID(transactionContext, id)
	if err1 != nil {
		t.Error(did)
	}

	//Update DID
	claims = encrypte.CreateMsg(id, encrypte.CreatePubkey(id+"#key1", "PublickeyECDSABase58"))
	msg, err = encrypte.GetJWT(*claims)
	err = assetTransferCC.UpdatedDID(transactionContext, msg)
	did, _ = assetTransferCC.ReadDID(transactionContext, id)
	if err == nil {
		t.Log(did)
	}
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
