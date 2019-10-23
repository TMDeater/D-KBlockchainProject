package service

import ( // Add Golang imports here

	// Add Hyperledger imports here
	"encoding/json"
	"strconv"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	// Add 3rd part imports here

	// Add local imports here

	nct "D-KBlockchainProject/chaincode/bank_insurance/go/chaincode/nct"
	"D-KBlockchainProject/chaincode/bank_insurance/go/chaincode/nct/config"
)

// CreateAgreementComponent to create JP NCT AC
func CreatePolicy(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	// decimal.DivisionPrecision = config.DivisionPrecision

	var logger = shim.NewLogger("CreatePolicy")
	logger.SetLevel(config.LogLevel)
	logger.Infof("CreatePolicy called")

	logger.Debugf("Args: %+v", args)

	// errorMsg := ErrorMessageResponse{}

	// // Parse the input data
	insurancePolicyNo := args[0]
	status := args[1]
	statusRemark := args[2]

	bankRefNo := args[3]
	prodCatOverview := args[4]
	agentCode := args[5]
	premium := args[6]
	currency := args[7]
	payMode := args[8]
	planRegion := args[9]

	ac := nct.Policy{}
	ac.InsurancePolicyNo = insurancePolicyNo
	ac.Status = status
	ac.StatusRemark = statusRemark
	ac.BankRefNo = bankRefNo
	ac.ProdCatOverview = prodCatOverview
	ac.AgentCode = agentCode
	ac.Premium, _ = strconv.ParseFloat(premium, 64)
	ac.Currency = currency
	ac.PayMode, _ = strconv.ParseFloat(payMode, 64)
	ac.PlanRegion = planRegion

	acBytesAsJSON, _ := json.Marshal(&ac)
	stub.PutState(bankRefNo, acBytesAsJSON)

	// create insurance private data with none bank rating first
	policyInsurancePrivate := nct.PolicyInsurancePrivate{
		ObjectType:			"PolicyInsurancePrivate",
		InsurancePolicyNo:	insurancePolicyNo,
		BankRefNo:			bankRefNo,
		BankRating:			"NA",
	}
	policyInsurancePrivateBytes, err := json.Marshal(policyInsurancePrivate)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}
	err = stub.PutPrivateData("collectionInsurancePrivate", bankRefNo, policyInsurancePrivateBytes)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}

	stub.SetEvent("policyCreated", []byte(bankRefNo))
	return nil, nil
}

// CreateAgreementComponentWithCompositeKey to create JP NCT AC with composite key
// func CreateAgreementComponentWithCompositeKey(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

// 	// decimal.DivisionPrecision = config.DivisionPrecision

// 	var logger = shim.NewLogger("CreateAgreementComponent")
// 	logger.SetLevel(config.LogLevel)
// 	logger.Infof("CreateAgreementComponentWithCompositeKey called")

// 	logger.Debugf("Args: %+v", args)

// 	// errorMsg := ErrorMessageResponse{}

// 	// // Parse the input data
// 	policyID := args[0] //Not in use as of now
// 	country := args[1]
// 	insuredPeople := args[2]
// 	gender := args[3]
// 	premium := args[4]
// 	faceamount := args[5]
// 	ac := nct.AgreementComponent{}
// 	ac.PolicyID = policyID
// 	ac.Country = country
// 	ac.InsuredPeople = insuredPeople
// 	ac.Gender = gender
// 	ac.Premium, _ = strconv.ParseFloat(premium, 64)
// 	ac.FaceAmount, _ = strconv.ParseFloat(faceamount, 64)

// 	acBytesAsJSON, _ := json.Marshal(&ac)
// 	key, _ := stub.CreateCompositeKey("policyKey", []string{country, policyID})
// 	stub.PutState(key, acBytesAsJSON)
// 	stub.SetEvent("agreementComponentCreated", []byte(policyID))
// 	return nil, nil
// }
