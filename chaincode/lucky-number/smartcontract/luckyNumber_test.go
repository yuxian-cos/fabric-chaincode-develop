package smartcontract_test

import (
	"encoding/json"
	"fmt"
	"lucky-number/smartcontract"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetLuckyNumber(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		NewStub()
		SetMspID("Org1")

		res := Stub.MockInvoke("uuid", [][]byte{[]byte("SetLuckyNumber"), []byte(`{
				"value": 9527
			}`)})

		assert.Equal(t, int32(200), res.Status)
		resPayload := smartcontract.Response{}
		_ = json.Unmarshal(res.Payload, &resPayload)
		assert.Equal(t, "success", resPayload.Status)
		assert.Regexp(t, `Org1 success set lucky number to 9527 in \d*`, resPayload.Message)
		fmt.Println(StubGetState("luckyNumber"))
		assert.Regexp(t, `{"created_at":\d*,"created_by":"Org1","value":9527}`, StubGetState("luckyNumber"))
	})
	t.Run("error", func(t *testing.T) {
		t.Run("Org2 can't set lucky number", func(t *testing.T) {
			NewStub()
			SetMspID("Org2")

			res := Stub.MockInvoke("uuid", [][]byte{[]byte("SetLuckyNumber"), []byte(`{
					"value": 9527
				}`)})

			assert.Equal(t, int32(500), res.Status)
			assert.Equal(t, "Only Org1 can set lucky number", res.Message)
		})
	})
}

func TestGetLuckyNumber(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		NewStub()
		SetMspID("Org1")

		StubPutState("luckyNumber", `{"created_at":1660886373,"created_by":"Org1","value":9527}`)

		res := Stub.MockInvoke("uuid", [][]byte{[]byte("GetLuckyNumber")})

		assert.Equal(t, int32(200), res.Status)
		resPayload := smartcontract.Response{}
		_ = json.Unmarshal(res.Payload, &resPayload)
		assert.Equal(t, "success", resPayload.Status)
		valueJsonBytes, _ := json.Marshal(resPayload.Value)
		assert.Equal(t, `{"created_at":1660886373,"created_by":"Org1","value":9527}`, string(valueJsonBytes))
	t.Run("error", func(t *testing.T) {
		t.Run("lucky number not exist", func(t *testing.T) {
			NewStub()
			SetMspID("Org1")

			res := Stub.MockInvoke("uuid", [][]byte{[]byte("GetLuckyNumber")})

			assert.Equal(t, int32(500), res.Status)
			assert.Equal(t, "Lucky number not exist", res.Message)
		})
	})
	})
}
