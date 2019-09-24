package service

import ( // Add Golang imports here

	// Add Hyperledger imports here
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	// Add 3rd part imports here
	// Add local imports here

	nct "test/chaincode/nct"
	"test/chaincode/nct/config"
)

func UpdatePolicyByBankRefID(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var logger = shim.NewLogger("UpdatePolicyByBankRefID")
	logger.SetLevel(config.LogLevel)
	logger.Infof("UpdatePolicyByBankRefID called")
	logger.Debugf("Args: %+v", args)

	//1. get state
	//2. update the leger
	//3. put state
	bankRefNo := args[0]
	insurancePolicyNo := args[1]
	status := args[2]
	statusRemark := args[3]

	record, _ := stub.GetState(bankRefNo)
	if record == nil {
		//return empty response for non-exists record
		logger.Infof("Record not found")
		// return nil, nil
	}	

	var ledgerInfo nct.Policy
	json.Unmarsal(record.Value, &ledgerInfo)

	ledgerInfo.InsurancePolicyNo = insurancePolicyNo
	ledgerInfo.Status = status
	ledgerInfo.StatusRemarkk = statusRemark

	ledgerInfoBytesAsJSON, _ := json.Marshal(&ledgerInfo)
	stub.PutState(bankRefNo, ledgerInfoBytesAsJSON)
	stub.SetEvent("policyCreated", []byte(bankRefNo))
	return nil, nil
}