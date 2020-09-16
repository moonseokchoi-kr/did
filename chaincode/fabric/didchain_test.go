package did

//TODO  look at the example file and change code to create fucntion
import (
	"did/chaincode/encrypte"
	"did/chaincode/fabric/mocks"
	"os"
	"testing"

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
	msg, err := encrypte.GetJWT()
	if err != nil {
		t.Errorf("Unexpected Error : %q", err)
	}
	//create DID
	err = assetTransferCC.CreateDid(transactionContext, msg)
	require.NoError(t, err)
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
