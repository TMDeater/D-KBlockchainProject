package main

import (
	// Add Golang imports here
	// Add Hyperledger imports here
	"github.com/hyperledger/fabric/core/chaincode/shim"
	// Add 3rd part imports here
	// Add local imports here
	cc "D-KBlockchainProject/artifacts/src/bank_insurance/go/chaincode"
	"D-KBlockchainProject/artifacts/src/bank_insurance/go/chaincode/nct/config"
)

// Each Golang project can have only 1 main
// Use the following `standard` way to start the chaincode
func main() {
	var logger = shim.NewLogger("main")
	logger.SetLevel(config.LogLevel)

	err := shim.Start(new(cc.JPNCTChaincode))
	if err != nil {
		logger.Errorf("Error starting Chaincode: %s", err)
	}
}
