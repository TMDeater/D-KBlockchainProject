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

// GetInsurencePolicyByBankRefID to retreieve JP NCT AC
func GetPolicyByBankRefID(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var logger = shim.NewLogger("GetPolicyByBankRefID")
	logger.SetLevel(config.LogLevel)
	logger.Infof("GetPolicyByBankRefID called")
	logger.Debugf("Args: %+v", args)

	// Check if the correct number of arguments were passed
	// if len(args) != 3 {
	// 	return nil, errors.New("Incorrect number of arguments. Expecting 3 - country, agreementComponentNo and agreementNumber")
	// }

	bankRefNo := args[0]

	record, _ := stub.GetState(bankRefNo)

	if record == nil {
		//return empty response for non-exists record
		logger.Infof("Record not found")
		return nil, nil
	}

	return record, nil
}

// // GetAgreementComponentWithCompositeKey to retreieve JP NCT AC with composite key
// func GetAgreementComponentWithCompositeKey(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

// 	var logger = shim.NewLogger("GetAgreementComponent")
// 	logger.SetLevel(config.LogLevel)
// 	logger.Infof("GetAgreementComponentWithCompositeKey called")
// 	logger.Debugf("Args: %+v", args)

// 	country := args[0]
// 	policyID := args[1]

// 	key, _ := stub.CreateCompositeKey("policyKey", []string{country, policyID})
// 	record, _ := stub.GetState(key)

// 	if record == nil {
// 		//return empty response for non-exists record
// 		logger.Infof("Record not found")
// 		return nil, nil
// 	}

// 	return record, nil
// }

// GetPolicyByInsurencePolicyNo to retreieve JP NCT AC with composite key
func GetPolicyByInsurancePolicyNo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var logger = shim.NewLogger("GetPolcyByInsurancePolicyNo")
	logger.SetLevel(config.LogLevel)
	logger.Infof("GetPolicyByInsurancePolicyNo called")
	logger.Debugf("Args: %+v", args)

	insurancePolicyNo := args[0]

	results := []nct.AgreementComponent{}
	records, err := stub.GetStateByPartialCompositeKey("policyKey", []string{insurancePolicyNo})

	if record == nil {
		//return empty response for non-exists record
		logger.Infof("Record not found")
		return nil, nil
	}

	return record, nil

	// if !records.HasNext() {
	// 	logger.Infof("Record does not exist")
	// 	return nil, err
	// }

	// for records.HasNext() {
	// 	iteratorValue, _ := records.Next()

	// 	var tempOutput nct.AgreementComponent
	// 	json.Unmarshal(iteratorValue.Value, &tempOutput)
	// 	results = append(results, tempOutput)
	// }
	// resultsAsJSON, _ := json.Marshal(&results)

	// return resultsAsJSON, nil
}

// // GetAgreementComponentByGender to retreieve JP NCT AC with composite key
// func GetAgreementComponentByGender(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

// 	var logger = shim.NewLogger("GetAgreementComponent")
// 	logger.SetLevel(config.LogLevel)
// 	logger.Infof("GetAgreementComponentByGender called")
// 	logger.Debugf("Args: %+v", args)

// 	gender := args[0]

// 	results := []nct.AgreementComponent{}
// 	queryString := fmt.Sprintf("{\"selector\":{\"gender\":\"%s\"}}", gender)
// 	records, err := stub.GetQueryResult(queryString)

// 	if !records.HasNext() {
// 		logger.Infof("Record does not exist")
// 		return nil, err
// 	}

// 	for records.HasNext() {
// 		iteratorValue, _ := records.Next()

// 		var tempOutput nct.AgreementComponent
// 		json.Unmarshal(iteratorValue.Value, &tempOutput)
// 		results = append(results, tempOutput)
// 	}
// 	resultsAsJSON, _ := json.Marshal(&results)

// 	return resultsAsJSON, nil
// }
