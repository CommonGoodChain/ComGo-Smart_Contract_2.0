package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//=============== GET DETAILS OF PROJECT ===============================================================

//getHistory
func getHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	return shim.Success(nil)
}

//query
func query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	return shim.Success(nil)
}
