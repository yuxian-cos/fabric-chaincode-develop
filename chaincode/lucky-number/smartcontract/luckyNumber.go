package smartcontract

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type LuckyNumberStruct struct {
	CreatedAt int64  `json:"created_at"`
	CreatedBy string `json:"created_by"`
	Value int64      `json:"value"`
}


type SetLuckyNumberRequest struct {
	Value int64 `json:"value"`
}


func (s *SmartContract) SetLuckyNumber(ctx contractapi.TransactionContextInterface, request SetLuckyNumberRequest) (*Response, error) {
	log.Printf("SetLuckyNumber")

	timestamp, err := getTxTimestamp(ctx)
	if err != nil {
		return nil, fmt.Errorf("GetTxTimestamp error: %v", err)
	}

	myMspID, err := getMyMspID(ctx)
	if err != nil {
		return nil, fmt.Errorf("getMyMspID error: %v", err)
	}

	owner := "Org1MSP"

	if (myMspID != owner) {
		return nil, fmt.Errorf("Only %s can set lucky number", owner)
	}

	data := LuckyNumberStruct{
		CreatedAt: timestamp,
		CreatedBy: myMspID,
		Value: request.Value,
	}

	dataJsonBytes, err :=  json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Marshal error: %v", err)
	}

	if err := ctx.GetStub().PutState("luckyNumber", dataJsonBytes); err != nil {
		return nil, fmt.Errorf("PutState error: %v", err)
	}

	return &Response{Status: "success", Message: fmt.Sprintf("%s success set lucky number to %d in %d", myMspID, data.Value, timestamp)}, nil
}

func (s *SmartContract) GetLuckyNumber(ctx contractapi.TransactionContextInterface) (*Response, error) {
	log.Printf("GetLuckyNumber")

	dataJsonBytes, err := ctx.GetStub().GetState("luckyNumber")
	if err != nil {
		return nil, fmt.Errorf("PutState error: %v", err)
	}

	if dataJsonBytes == nil {
		return nil, fmt.Errorf("Lucky number not exist")
	}

	var data LuckyNumberStruct

	if err := json.Unmarshal(dataJsonBytes, &data); err != nil {
		return nil, fmt.Errorf("Unmarshal error: %v", err)
	}

	return &Response{Status: "success", Value: data}, nil
}
