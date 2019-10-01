package chaincode

import (
	// Add Golang imports here

	"fmt"
	//"time"
	// "strconv"
	// Add Hyperledger imports here
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"

	// Add 3rd part imports here
	//"github.com/shopspring/decimal"
	// Add local imports here
	"D-KBlockchainProject/chaincode/bank_insurance/go/chaincode/nct/config"
	nct "D-KBlockchainProject/chaincode/bank_insurance/go/chaincode/nct/service"
)

// JPNCTChaincode Chaincode struct (No field should be here)
type JPNCTChaincode struct {
}

// Define the function names that Invoke cc transaction supports here
// TODO: consider using `ENUM` in golang (which whould be not easy to code but maybe still should do)
const (
	createAgreementComponent                 string = "createAgreementComponent"
	createAgreementComponentWithCompositeKey string = "createAgreementComponentWithCompositeKey"
	getAgreementComponent                    string = "getAgreementComponent"
	getAgreementComponentWithCompositeKey    string = "getAgreementComponentWithCompositeKey"
	getAgreementComponentByCountry           string = "getAgreementComponentByCountry"
	getAgreementComponentByGender            string = "getAgreementComponentByGender"

	createPolicy                 string = "createPolicy"
	getPolicyByBankRefID         string = "getPolicyByBankRefID"
	getPolicyByInsurancePolicyNo string = "getPolicyByInsurancePolicyNo"
	updatePolicyByBankRefID      string = "updatePolicyByBankRefID"
)

// NO business functions should locate here
// Please create separate files under folder `nct` using package name `nct`

// Place the implementation of shim interface here (including Init, Invoke)

// Init : Implementing Init implementation of shim interface
func (t *JPNCTChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	args := stub.GetArgs()
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting no arguments (except `init`)")
	}
	return shim.Success(nil)
}

// Invoke : Implementing Invoke implementation of shim interface
func (t *JPNCTChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	var logger = shim.NewLogger("shim.Invoke")
	logger.SetLevel(config.LogLevel)

	function, args := stub.GetFunctionAndParameters()

	if function != "invoke" {
		shim.Error("Invalid function name received " + function + " for invoke transcation, should be invoke")
	}

	logger.Infof("function : %s, args : %s", function, args)

	action := args[0]
	args = args[1:]

	switch action {
	case createPolicy:
		result, err := nct.CreatePolicy(stub, args)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(result)
	// case createAgreementComponentWithCompositeKey:
	// 	result, err := nct.CreateAgreementComponentWithCompositeKey(stub, args)
	// 	if err != nil {
	// 		return shim.Error(err.Error())
	// 	}
	// 	return shim.Success(result)
	case getPolicyByBankRefID:
		result, err := nct.GetPolicyByBankRefID(stub, args)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(result)
	case getPolicyByInsurancePolicyNo:
		result, err := nct.GetPolicyByInsurancePolicyNo(stub, args)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(result)
	case updatePolicyByBankRefID:
		result, err := nct.UpdatePolicyByBankRefID(stub, args)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(result)
		// case getAgreementComponentByCountry:
		// 	result, err := nct.GetAgreementComponentByCountry(stub, args)
		// 	if err != nil {
		// 		return shim.Error(err.Error())
		// 	}
		// 	return shim.Success(result)
		// case getAgreementComponentByGender:
		// 	result, err := nct.GetAgreementComponentByGender(stub, args)
		// 	if err != nil {
		// 		return shim.Error(err.Error())
		// 	}
		// 	return shim.Success(result)
		// }
	}
	errMsg := fmt.Sprintln("Received unknown action invocation: [action:", action, ", args:", args, "].")
	return shim.Error(errMsg)

}
