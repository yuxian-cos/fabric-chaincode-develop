package smartcontract

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) (*Response, error) {
	log.Printf("Init")
	return &Response{Status: "success", Message: "init chaincode"}, nil
}
