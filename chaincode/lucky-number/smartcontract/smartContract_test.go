package smartcontract_test

import (
	"log"
	"lucky-number/smartcontract"
	"os"

	"github.com/gogo/protobuf/proto"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/msp"
)

var Stub *shimtest.MockStub

func NewStub() {
	Scc, err := contractapi.NewChaincode(new(smartcontract.SmartContract))
	if err != nil {
		log.Println("NewChaincode failed", err)
		os.Exit(0)
	}

	Stub = shimtest.NewMockStub("main", Scc)

	SetMspID("Org1")
}

func SetMspID(mspID string) {
	if Stub == nil {
		log.Println("Stub not initial")
		os.Exit(0)
	}

	sid := &msp.SerializedIdentity{Mspid: mspID}
	data, err := proto.Marshal(sid)
	if err != nil {
		log.Println("Marshal false")
		os.Exit(0)
	}

	Stub.Creator = data
}

func StubPutState(key string, value string) {
	Stub.MockTransactionStart("txid")
	_ = Stub.PutState(key, []byte(value))
	Stub.MockTransactionEnd("txid")
}

func StubGetState(key string) string {
	valueBytes, _ := Stub.GetState(key)
	return string(valueBytes)
}
