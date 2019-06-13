/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding  ownership.  The ASF licenses this file
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
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ============================================================================================================================
// write() - genric write variable into ledger
//
// Shows Off PutState() - writting a key/value into the ledger
//
// Inputs - Array of strings
//    0   ,    1
//   key  ,  value
//  "abc" , "test"
// ============================================================================================================================
func write(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var key, value string
	var err error
	fmt.Println("starting write")

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2. key of the variable and value to set")
	}

	// input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	key = args[1] //rename for funsies
	value = args[2]
	err = stub.PutState(key, []byte(value)) //write the variable into the ledger
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end write")
	return shim.Success(nil)
}

// ============================================================================================================================
// invke() - fetches JSON from mongodb and creates ASSET struct
// ============================================================================================================================

func invke(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	return shim.Success(nil)
}

//=============== USER REGISTARTION RELATED FUNCTION'S START HERE ===============================================================

//addFoundation
func addFoundation(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("started to register new foundation")

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}

	//check if user already exists
	_, err = getFoundation(stub, string(certname))
	if err == nil {
		fmt.Println("This foundation user already exists - " + string(certname))
		return shim.Error("This foundation user already exists - " + string(certname))
	}

	var user Foundation
	user.ObjectType = "Foundation"
	user.FoundationID = string(certname)
	user.FoundationUsername = args[0]
	user.FoundationCompany = args[1]
	user.ProjectPermission = args[2]
	user.UserPermission = args[3]
	user.Role = args[4]

	var location Location
	location.Latitude = args[5]
	location.Longitude = args[6]

	log.Println("final foundation user ", user)

	//store user
	userAsBytes, _ := json.Marshal(user)
	err = stub.PutState(user.FoundationID, userAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	log.Println("- end init foundation")
	return shim.Success(nil)
}

//updateFoundation
func updateFoundation(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("started to update the foundation")

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}

	//check if user already exists
	foundationUser, err := getFoundation(stub, string(certname))
	if err != nil {
		fmt.Println("This foundation user already exists - " + string(certname))
		return shim.Error("This foundation user already exists - " + string(certname))
	}

	foundationUser.FoundationUsername = args[0]
	foundationUser.FoundationCompany = args[1]
	foundationUser.ProjectPermission = args[2]
	foundationUser.UserPermission = args[3]
	foundationUser.Role = args[4]

	var location Location
	location.Latitude = args[5]
	location.Longitude = args[6]

	log.Println("final updated foundation user ", foundationUser)

	//store user
	userAsBytes, _ := json.Marshal(foundationUser)
	err = stub.PutState(foundationUser.FoundationID, userAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	log.Println("- end update foundation")
	return shim.Success(nil)
}

//addNGO
func addNGO(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("started to register new NGO")

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}

	//check if user already exists
	_, err = getNGO(stub, string(certname))
	if err == nil {
		fmt.Println("This NGO user already exists - " + string(certname))
		return shim.Error("This NGO user already exists - " + string(certname))
	}

	var user NGO
	user.ObjectType = "NGO"
	user.NGOID = string(certname)
	user.NGOUsername = args[0]
	user.NGOCompany = args[1]
	user.ProjectPermission = args[2]
	user.UserPermission = args[3]
	user.Role = args[4]

	var location Location
	location.Latitude = args[5]
	location.Longitude = args[6]

	log.Println("final NGO user object ", user)

	//store user
	userAsBytes, _ := json.Marshal(user)
	err = stub.PutState(user.NGOID, userAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	log.Println("- end init NGO")
	return shim.Success(nil)
}

//updateNGO
func updateNGO(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("started to update the foundation")

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}

	//check if user already exists
	ngoUser, err := getNGO(stub, string(certname))
	if err != nil {
		fmt.Println("This NGO user already exists - " + string(certname))
		return shim.Error("This NGO user already exists - " + string(certname))
	}

	ngoUser.NGOUsername = args[0]
	ngoUser.NGOCompany = args[1]
	ngoUser.ProjectPermission = args[2]
	ngoUser.UserPermission = args[3]
	ngoUser.Role = args[4]

	var location Location
	location.Latitude = args[5]
	location.Longitude = args[6]

	log.Println("final updated the ngo user ", ngoUser)

	//store user
	userAsBytes, _ := json.Marshal(ngoUser)
	err = stub.PutState(ngoUser.NGOID, userAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	log.Println("- end update NGO")
	return shim.Success(nil)
}

//addDonor
func addDonor(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("started to register new donor")

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}

	//check if user already exists
	_, err = getDonor(stub, string(certname))
	if err == nil {
		fmt.Println("This donor user already exists - " + string(certname))
		return shim.Error("This donor user already exists - " + string(certname))
	}

	var user Donor
	user.ObjectType = "Donor"
	user.DonorID = string(certname)
	user.DonorUsername = args[0]
	user.DonorCompany = args[1]
	user.ProjectPermission = args[2]
	user.UserPermission = args[3]
	user.Role = args[4]

	var location Location
	location.Latitude = args[5]
	location.Longitude = args[6]

	log.Println("final DONOR user object ", user)

	//store user
	userAsBytes, _ := json.Marshal(user)
	err = stub.PutState(user.DonorID, userAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	log.Println("- end init Donor")
	return shim.Success(nil)
}

//updateDonor
func updateDonor(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("started to update the donor")

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}

	//check if user already exists
	donorUser, err := getDonor(stub, string(certname))
	if err != nil {
		fmt.Println("This Donor user already exists - " + string(certname))
		return shim.Error("This Donor user already exists - " + string(certname))
	}

	donorUser.DonorUsername = args[0]
	donorUser.DonorCompany = args[1]
	donorUser.ProjectPermission = args[2]
	donorUser.UserPermission = args[3]
	donorUser.Role = args[4]

	var location Location
	location.Latitude = args[5]
	location.Longitude = args[6]

	log.Println("final updated the donor user ", donorUser)

	//store user
	userAsBytes, _ := json.Marshal(donorUser)
	err = stub.PutState(donorUser.DonorID, userAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	log.Println("- end update Donor")
	return shim.Success(nil)
}

//addBoard
func addBoard(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("started to register new board")

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}

	//check if user already exists
	_, err = getBoard(stub, string(certname))
	if err == nil {
		fmt.Println("This board user already exists - " + string(certname))
		return shim.Error("This board user already exists - " + string(certname))
	}

	var user Board
	user.ObjectType = "Board"
	user.BoardID = string(certname)
	user.BoardUsername = args[0]
	user.BoardCompany = args[1]
	user.ProjectPermission = args[2]
	user.UserPermission = args[3]
	user.Role = args[4]

	var location Location
	location.Latitude = args[5]
	location.Longitude = args[6]

	log.Println("final Board user object ", user)

	//store user
	userAsBytes, _ := json.Marshal(user)
	err = stub.PutState(user.BoardID, userAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	log.Println("- end init Board")
	return shim.Success(nil)
}

//updateBoard
func updateBoard(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("started to update the board")

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}

	//check if user already exists
	boardUser, err := getBoard(stub, string(certname))
	if err != nil {
		fmt.Println("This board user already exists - " + string(certname))
		return shim.Error("This board user already exists - " + string(certname))
	}

	boardUser.BoardUsername = args[0]
	boardUser.BoardCompany = args[1]
	boardUser.ProjectPermission = args[2]
	boardUser.UserPermission = args[3]
	boardUser.Role = args[4]

	var location Location
	location.Latitude = args[5]
	location.Longitude = args[6]

	log.Println("final updated the board user ", boardUser)

	//store user
	userAsBytes, _ := json.Marshal(boardUser)
	err = stub.PutState(boardUser.BoardID, userAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	log.Println("- end update Board")
	return shim.Success(nil)
}

//addCRM
func addCRM(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("started to register new crm")

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}

	//check if user already exists
	_, err = getCRM(stub, string(certname))
	if err == nil {
		fmt.Println("This board user already exists - " + string(certname))
		return shim.Error("This board user already exists - " + string(certname))
	}

	var user CRM
	user.ObjectType = "Board"
	user.CRMID = string(certname)
	user.CRMUsername = args[0]
	user.CRMCompany = args[1]
	user.ProjectPermission = args[2]
	user.UserPermission = args[3]
	user.Role = args[4]

	var location Location
	location.Latitude = args[5]
	location.Longitude = args[6]

	log.Println("final Board user object ", user)

	//store user
	userAsBytes, _ := json.Marshal(user)
	err = stub.PutState(user.CRMID, userAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	log.Println("- end init Board")
	return shim.Success(nil)
}

//updateCRM
func updateCRM(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("started to update the crm")

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}

	//check if user already exists
	crmUser, err := getCRM(stub, string(certname))
	if err != nil {
		fmt.Println("This crm user already exists - " + string(certname))
		return shim.Error("This crm user already exists - " + string(certname))
	}

	crmUser.CRMUsername = args[0]
	crmUser.CRMCompany = args[1]
	crmUser.ProjectPermission = args[2]
	crmUser.UserPermission = args[3]
	crmUser.Role = args[4]

	var location Location
	location.Latitude = args[5]
	location.Longitude = args[6]

	log.Println("final updated the crm user ", crmUser)

	//store user
	userAsBytes, _ := json.Marshal(crmUser)
	err = stub.PutState(crmUser.CRMID, userAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	log.Println("- end update crm")
	return shim.Success(nil)
}

//addAdmin
func addAdmin(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	return shim.Success(nil)
}

//updateAdmin
func updateAdmin(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	return shim.Success(nil)
}
