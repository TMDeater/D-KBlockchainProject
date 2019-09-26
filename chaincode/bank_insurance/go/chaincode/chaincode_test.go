package chaincode

import (
	// Add Golang imports here

	"encoding/json"
	"testing"

	// Add Hyperledger imports here
	nct "test/chaincode/nct"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	// Add 3rd part imports here
	// Add local imports here
)

func init() {
}

//TestaccountingChaincode_Init Test case for accountingChaincode Init method
func TestChaincode_Init(t *testing.T) {
	cc := new(JPNCTChaincode)
	stub := shim.NewMockStub("cc", cc)
	if stub == nil {
		t.Fatal("TestChaincode_Init", "Fail to create mock stub")
	}
}

//TestAccChaincode_Invoke_GenerateGL Test case for AccChaincode Invoke method with function generateGL
func TestChaincode_Invoke_Create(t *testing.T) {
	// Inject ODMURL for testing

	cc := new(JPNCTChaincode)
	stub := shim.NewMockStub("cc", cc)

	if stub == nil {
		t.Fatal("Fail to create mock stub")
	}
	// res := stub.MockInvoke("Test_CreateAgreementComponentWithCompositeKey_Function", [][]byte{[]byte("invoke"), []byte("createAgreementComponentWithCompositeKey"), []byte("1"), []byte("HK"), []byte("Adam"), []byte("Male"), []byte("100"), []byte("10000")})
	
	res := stub.MockInvoke("Test_CreateAgreementComponentWithCompositeKey_Function", [][]byte{[]byte("invoke"), []byte("createAgreementComponentWithCompositeKey"), []byte("2"), []byte("US"), []byte("Ben"), []byte("Male"), []byte("1000"), []byte("100000")})

	if res.Status != shim.OK {
		t.FailNow()
	}

	// result, _ := stub.GetStateByPartialCompositeKey("policyKey", []string{"US"})
	result, _ := stub.GetState("2")
	tempOutput := nct.AgreementComponent{}

	json.Unmarshal(result, &tempOutput)
	t.Logf("The record entry:%s\n", string(result))
	if tempOutput.Country != "US" {
		t.FailNow()
	}

	// if !result.HasNext() {
	// 	t.FailNow()
	// }
	// for result.HasNext() {
	// 	iteratorValue, _ := result.Next()
	// 	json.Unmarshal(iteratorValue.Value, &tempOutput)
	// 	t.Logf("The record entry:%s\n", string(iteratorValue.Value))
	// 	if tempOutput.Country != "US" {
	// 		t.FailNow()
	// 	}
	// }

}
