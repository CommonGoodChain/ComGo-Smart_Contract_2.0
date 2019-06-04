/*
Li
censed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding  donorship.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

//Foundation is ...
type Foundation struct {
	ObjectType string `json:"docType"` //field for couchdb

	FoundationID       string      `json:"foundationId"`
	FoundationUsername string      `json:"foundationUsername"`
	FoundationCompany  string      `json:"foundationCompany"`
	Role               string      `json:"role"`
	Permissions        Permissions `json:"permissions"`
	Location           Location    `json:"location"`
}

//NGO ss
type NGO struct {
	ObjectType string `json:"docType"` //field for couchdb

	NGOID       string      `json:"ngoId"`
	NGOUsername string      `json:"ngoUsername"`
	NGOCompany  string      `json:"ngoCompany"`
	Role        string      `json:"role"`
	Permissions Permissions `json:"permissions"`
	Location    Location    `json:"location"`
}

//Donor ss
type Donor struct {
	ObjectType string `json:"docType"` //field for couchdb

	DonorID       string      `json:"donorId"`
	DonorUsername string      `json:"donorUsername"`
	DonorCompany  string      `json:"donorcompany"`
	Donations     []string    `json:"donations"`
	Role          string      `json:"role"`
	Permissions   Permissions `json:"permissions"`
	Location      Location    `json:"location"`
}

//Board as
type Board struct {
}

//Validator as
type Validator struct {
}

//CRM as
type CRM struct {
}

//SDG as
type SDG struct {
	SDGType string `json:"SDGType"`
}

//Location as
type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

//Permissions as
type Permissions struct {
	All       string `json:"all"`
	Write     string `json:"write"`
	Read      string `json:"read"`
	ReadWrite string `json:"readWrite"`
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode - %s", err)
	}
}

// ============================================================================================================================
// Init - initialize the chaincode - projects don’t need anything initlization, so let's run a dead simple test instead
// ============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Projects Is Starting Up")
	_, args := stub.GetFunctionAndParameters()
	var Aval int
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// convert numeric string to integer
	Aval, err = strconv.Atoi(args[0])
	if err != nil {
		return shim.Error("Expecting a numeric string argument to Init()")
	}

	// store compaitible projects application version
	err = stub.PutState("projects_ui", []byte("3.5.0"))
	if err != nil {
		return shim.Error(err.Error())
	}

	// this is a very simple dumb test.  let's write to the ledger and error on any errors
	err = stub.PutState("selftest", []byte(strconv.Itoa(Aval))) //making a test var "selftest", its handy to read this right away to test the network
	if err != nil {
		return shim.Error(err.Error()) //self-test fail
	}

	fmt.Println(" - ready for action") //self-test pass
	return shim.Success(nil)
}

// ============================================================================================================================
// Invoke - Our entry point for Invocations
// ============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println(" ")
	fmt.Println("starting invoke, for - " + function)

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting at least 1")
	}

	// Handle different functions
	if function == "init" {
		return t.Init(stub)
	} else if function == "read" {
		return read(stub, args)
	} else if function == "write" {
		return write(stub, args)
	} else if function == "invke" {
		return invke(stub, args)
	} else if function == "addFoundation" { // Registration API's
		return addFoundation(stub, args)
	} else if function == "updateFoundation" {
		return updateFoundation(stub, args)
	} else if function == "addNGO" {
		return addNGO(stub, args)
	} else if function == "updateNGO" {
		return updateNGO(stub, args)
	} else if function == "addDonor" {
		return addDonor(stub, args)
	} else if function == "updateDonor" {
		return updateDonor(stub, args)
	} else if function == "addBoard" {
		return addBoard(stub, args)
	} else if function == "updateBoard" {
		return updateBoard(stub, args)
	} else if function == "addCRM" {
		return addCRM(stub, args)
	} else if function == "updateCRM" {
		return updateCRM(stub, args)
	} else if function == "addAdmin" {
		return addAdmin(stub, args)
	} else if function == "updateAdmin" {
		return updateAdmin(stub, args)
	} else if function == "addProject" { // Project API's
		return addProject(stub, args)
	} else if function == "updateProject" {
		return updateProject(stub, args)
	} else if function == "addMilestone" {
		return addMilestone(stub, args)
	} else if function == "updateMilestone" {
		return updateMilestone(stub, args)
	} else if function == "addActivity" {
		return addActivity(stub, args)
	} else if function == "updateActivity" {
		return updateActivity(stub, args)
	} else if function == "updateProjectStatus" { // Flow API's
		return updateProjectStatus(stub, args)
	} else if function == "getHistory" { // Query API's
		return getHistory(stub, args)
	} else if function == "query" {
		return query(stub, args)
	}

	// error out
	fmt.Println("Received unknown invoke function name - " + function)
	return shim.Error("Received unknown invoke function name - '" + function + "'")
}

// ============================================================================================================================
// Query - legacy function
// ============================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Error("Unknown supported call - Query()")
}
