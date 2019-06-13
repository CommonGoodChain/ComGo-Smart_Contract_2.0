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

	FoundationID       string   `json:"foundationId"`
	FoundationUsername string   `json:"foundationUsername"`
	FoundationCompany  string   `json:"foundationCompany"`
	Role               string   `json:"role"`
	ProjectPermission  string   `json:"projectPermission"`
	UserPermission     string   `json:"userPermission"`
	Location           Location `json:"location"`
}

//NGO ss
type NGO struct {
	ObjectType string `json:"docType"` //field for couchdb

	NGOID             string      `json:"ngoId"`
	NGOUsername       string      `json:"ngoUsername"`
	NGOCompany        string      `json:"ngoCompany"`
	Role              string      `json:"role"`
	Permissions       Permissions `json:"permissions"`
	Location          Location    `json:"location"`
	ProjectPermission string      `json:"projectPermission"`
	UserPermission    string      `json:"userPermission"`
}

//Donor ss
type Donor struct {
	ObjectType string `json:"docType"` //field for couchdb

	DonorID           string      `json:"donorId"`
	DonorUsername     string      `json:"donorUsername"`
	DonorCompany      string      `json:"donorcompany"`
	Donations         []string    `json:"donations"`
	Role              string      `json:"role"`
	Permissions       Permissions `json:"permissions"`
	Location          Location    `json:"location"`
	ProjectPermission string      `json:"projectPermission"`
	UserPermission    string      `json:"userPermission"`
}

//Board as
type Board struct {
	ObjectType string `json:"docType"` //field for couchdb

	BoardID           string `json:"boardID"`
	BoardUsername     string `json:"boardUsername"`
	BoardCompany      string `json:"boardCompany"`
	Role              string `json:"role"`
	ProjectPermission string `json:"projectPermission"`
	UserPermission    string `json:"userPermission"`
}

//Validator as
type Validator struct {
	ProjectPermission string `json:"projectPermission"`
	UserPermission    string `json:"userPermission"`
}

//CRM as
type CRM struct {
	ObjectType string `json:"docType"` //field for couchdb

	CRMID             string `json:"crmID"`
	CRMUsername       string `json:"crmUsername"`
	CRMCompany        string `json:"crmCompany"`
	Role              string `json:"role"`
	ProjectPermission string `json:"projectPermission"`
	UserPermission    string `json:"userPermission"`
}

//Project as
type Project struct {
	ObjectType         string   `json:"docType"` //field for couchdb
	ProjectID          string   `json:"projectId"`
	ProjectName        string   `json:"projectName"`
	ProjectType        string   `json:"projectType"`
	Flag               string   `json:"flag"`
	FundGoal           float64  `json:"fundGoal"`
	Currency           string   `json:"currency"`
	FundRaised         float64  `json:"fundRaised"`
	FundAllocated      float64  `json:"fundAllocated"`
	FundNotAllocated   float64  `json:"fundNotAllocated"`
	ProjectBudget      float64  `json:"projectBudget"`
	ProjectOwner       string   `json:"projectOwner"`
	FoundationCompany  string   `json:"foundationCompany"`
	NGOCompany         string   `json:"ngoCompany"`
	Donations          []string `json:"donations"`
	Status             string   `json:"status"`
	FundAllocationType string   `json:"fundAllocationType"` // 1 = Manual, 2 = Automated, 3 = On Proof Submission, 4 = On Validation
	TransactionLoc     Location `json:"transactionLoc"`
	SDG                []SDG    `json:"SDG"`
	ProjectLoc         Location `json:"projectLoc"`
	CreatedBy          string   `json:"createdBy"`
	SubRole            string   `json:"subRole"`
	IsPublished        bool     `json:"isPublished"`
	IsApproved         bool     `json:"isApproved"`
	Remarks            string   `json:"remarks"`
	StartDate          string   `json:"startDate"`
	EndDate            string   `json:"endDate"`
	Description        string   `json:"description"`
	Country            string   `json:"country"`
}

//Milestone as
type Milestone struct {
	ObjectType       string   `json:"docType"` //field for couchdb
	MilestoneName    string   `json:"milestoneName"`
	StartDate        string   `json:"startDate"`
	EndDate          string   `json:"endDate"`
	MilestoneID      string   `json:"milestoneId"`
	ProjectID        string   `json:"projectId"`
	MilBudget        float64  `json:"milestoneBudget"`
	MilFundAllocated float64  `json:"milFundAllocated"`
	MilFundRequested float64  `json:"milFundRequested"`
	MilFundReleased  float64  `json:"MilFundReleased"`
	MilestoneOwner   string   `json:"milestoneOwner"`
	ActivityCount    int      `json:"activityCount"`
	Status           string   `json:"status"`
	TransactionLoc   Location `json:"transactionLoc"`
	IsApproved       bool     `json:"isApproved"`
	Description      string   `json:"description"`
}

//Activity as
type Activity struct {
	ObjectType          string   `json:"docType"` //field for couchdb
	ActivityName        string   `json:"activityName"`
	StartDate           string   `json:"startDate"`
	EndDate             string   `json:"endDate"`
	ActivityBudget      float64  `json:"activityBudget"`
	FundAllocated       float64  `json:"fundAllocated"`
	FundReleased        float64  `json:"fundReleased"`
	FundRequested       float64  `json:"fundRequested"`
	ActivityID          string   `json:"activityId"`
	MilestoneID         string   `json:"milestoneId"`
	ProjectID           string   `json:"projectId"`
	Validation          bool     `json:"validation"`
	Status              string   `json:"status"`
	ValidatorID         string   `json:"validatorId"`
	SecondaryValidation bool     `json:"secondaryValidation"`
	PartialValidation   bool     `json:"partialValidation"`
	TransactionLoc      Location `json:"transactionLoc"`
	IsApproved          bool     `json:"isApproved"`
	Remarks             string   `json:"remarks"`
	Description         string   `json:"description"`
}

//SDG as
type SDG struct {
	SDGType string `json:"SDGType"`
}

//ProjectFunds as
type ProjectFunds struct {
	DonorName       string  `json:"donorName"`
	DonoationAmount float64 `json:"donationAmount"`
}

//Location as
type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

//Permissions for users
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
	} else if function == "query_all" {
		return query_all(stub, args)
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
