package chaincode

import (
	// Add Golang imports here

	"encoding/json"
	"fmt"
	"testing"

	// Add Hyperledger imports here
	nct "D-KBlockchainProject/chaincode/bank_insurance/go/chaincode/nct"

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

	res := stub.MockInvoke("Test_CreatePolicyWithCompositeKey_Function", [][]byte{[]byte("invoke"), []byte("createPolicy"), []byte(""), []byte(""), []byte(""), []byte("1"), []byte("Overview"), []byte("agent code"), []byte("1000"), []byte("HKD"), []byte("6"), []byte("HKG")})

	if res.Status != shim.OK {
		t.FailNow()
	}

	res2 := stub.MockInvoke("updatePolicy", [][]byte{[]byte("invoke"), []byte("updatePolicyByBankRefID"), []byte("1"), []byte("222"), []byte("0"), []byte("remark")})
	if res2.Status != shim.OK {
		t.FailNow()
	}

	// result2, _ := stub.GetStateByPartialCompositeKey("policyKey", []string{"US"})
	result, _ := stub.GetState("1")
	tempOutput := nct.Policy{}

	json.Unmarshal(result, &tempOutput)
	t.Logf("The record entry:%s\n", string(result))
	fmt.Println("Record get from BankRefNo: ", tempOutput.BankRefNo)
	fmt.Println("InsurancePolicyNo: ", tempOutput.InsurancePolicyNo)
	fmt.Println("Status: ", tempOutput.Status)
	fmt.Println("StatusRemark: ", tempOutput.StatusRemark)
	if tempOutput.Currency != "HKD" {
		t.FailNow()
	}
	if tempOutput.InsurancePolicyNo != "222" {
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
