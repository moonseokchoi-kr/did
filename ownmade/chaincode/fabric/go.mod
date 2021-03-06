module did/chaincode/fabric

go 1.15

require (
	github.com/did/chaincode/encrypte v0.0.0
	github.com/hyperledger/fabric-chaincode-go v0.0.0-20200511190512-bcfeb58dd83a
	github.com/hyperledger/fabric-contract-api-go v1.1.0
	github.com/hyperledger/fabric-protos-go v0.0.0-20200707132912-fee30f3ccd23
	github.com/hyperledger/fabric-samples/asset-transfer-private-data/chaincode-go v0.0.0-20200911133247-c8703df425a8
	github.com/stretchr/testify v1.5.1 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/did/chaincode/encrypte v0.0.0 => ../encrypte
