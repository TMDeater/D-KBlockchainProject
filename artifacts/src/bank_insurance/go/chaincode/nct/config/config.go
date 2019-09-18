package config

import "github.com/hyperledger/fabric/core/chaincode/shim"

// Add Golang imports here
// Add Hyperledger imports here
// "github.com/hyperledger/fabric/core/chaincode/shim"
// Add 3rd part imports here
// Add local imports here

//Issuer default issuer
const Issuer = "MNL"

// //AgreementComponentKeyName the key type to store agreement component
// const AgreementComponentKeyName = "AgreementComponentKey"

// // CalculateBillingAmountODMUrl the URL of ODM for calculate billing amount, expose to let testing code can override it with non docker service link
// var (
// 	CalculateBillingAmountODMUrl = "http://odm-runtime.manulife:9060/DecisionService/rest/JPNCT/BillingAmount"
// 	CalculateBenefitsODMUrl      = "http://odm-runtime.manulife:9060/DecisionService/rest/JPNCT/PreviewBenefit"
// 	SurrenderODMUrl              = "http://odm-runtime.manulife:9060/DecisionService/rest/JPNCT/Surrender"
// 	MaturityDateODMUrl           = "http://odm-runtime.manulife:9060/DecisionService/rest/JPNCT/MaturityDate"
// )

const LogLevel = shim.LogDebug

// // IdentifierType for agreementTransaction
// const IdentifierType = "AgreementComponent"

// // DivisionPrecision decimal numbers
// const DivisionPrecision = 8

// // TimeFormat time format for NCT
// const TimeFormat = "2006-01-02T15:04:05-07:00"
