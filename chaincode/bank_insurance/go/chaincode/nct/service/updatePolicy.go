package service

import ( // Add Golang imports here

	// Add Hyperledger imports here
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	// Add 3rd part imports here
	// Add local imports here

	nct "D-KBlockchainProject/chaincode/bank_insurance/go/chaincode/nct"
	"D-KBlockchainProject/chaincode/bank_insurance/go/chaincode/nct/config"
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
	bankRating := args[4]

	record, _ := stub.GetState(bankRefNo)
	if record == nil {
		//return empty response for non-exists record
		logger.Infof("Record not found")
		// return nil, nil
	}

	var ledgerInfo nct.Policy
	json.Unmarshal(record, &ledgerInfo)

	ledgerInfo.InsurancePolicyNo = insurancePolicyNo
	ledgerInfo.Status = status
	ledgerInfo.StatusRemark = statusRemark

	ledgerInfoBytesAsJSON, _ := json.Marshal(&ledgerInfo)
	stub.PutState(bankRefNo, ledgerInfoBytesAsJSON)

	// === Update the private information for insurance
	policyInsurancePrivateAsBytes, err := stub.GetPrivateData("collectionInsurancePrivate", bankRefNo)
	if err != nil {
		fmt.Println("Failed to get private insurance policy information: " + err.Error())
		return nil ,nil
	} else if policyInsurancePrivateAsBytes == nil {
		fmt.Println("Private Policy does not exist: " + bankRefNo)
		return nil, nil
	}
	policyInsurancePrivate := nct.PolicyInsurancePrivate{}
	err = json.Unmarshal(policyInsurancePrivateAsBytes, &policyInsurancePrivate)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	policyInsurancePrivate.BankRating = bankRating
	policyInsurancePrivateJSONasByte, _ := json.Marshal(policyInsurancePrivate)
	err = stub.PutPrivateData("collectionInsurancePrivate", bankRefNo, policyInsurancePrivateJSONasByte)
	if err != nil {
		fmt.Println(err.Error())
		return nil,nil
	}
	fmt.Println("- Private Insurance Information Updated")
	// === End updaate private information for insurance

	stub.SetEvent("policyUpdated", []byte(bankRefNo))
	return nil, nil
}
