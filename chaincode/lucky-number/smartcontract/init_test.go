package smartcontract_test

import (
	"encoding/json"
	"lucky-number/smartcontract"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		NewStub()
		SetMspID("Org1")

		res := Stub.MockInvoke("uuid", [][]byte{[]byte("Init")})
		assert.Equal(t, int32(200), res.Status)

		resPayload := smartcontract.Response{}
		_ = json.Unmarshal(res.Payload, &resPayload)

		assert.Equal(t, "success", resPayload.Status)
		assert.Equal(t, "init chaincode", resPayload.Message)
	})
}
