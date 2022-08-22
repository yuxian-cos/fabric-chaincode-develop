package smartcontract

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}

func getMyMspID(ctx contractapi.TransactionContextInterface) (string, error) {
	mspID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", fmt.Errorf("GetMSPID error: %v", err)
	}
	return mspID, nil
}

func getTxTimestamp(ctx contractapi.TransactionContextInterface) (int64, error) {
	txTimestamp, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return 0, fmt.Errorf("getTxTimestamp error: %v", err)
	}
	return txTimestamp.Seconds, nil
}
