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

//addPrivateUser
func addPrivateUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("started to register private user")

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
	_, err = getPrivateUser(stub, string(certname))
	if err == nil {
		fmt.Println("This Private user already exists - " + string(certname))
		return shim.Error("This private user already exists - " + string(certname))
	}

	var user PrivateUser
	user.ObjectType = "PrivateUser"
	user.UserID = string(certname)
	user.Username = args[0]
	user.Permissions = nil
	user.UserPermission = args[3]
	user.Role = args[4]

	var location Location
	location.Latitude = args[5]
	location.Longitude = args[6]

	log.Println("final private user ", user)

	//store user
	userAsBytes, _ := json.Marshal(user)
	err = stub.PutState(user.UserID, userAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	log.Println("- end registration of private user")
	return shim.Success(nil)
}

//updatePrivateUser
func updatePrivateUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
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
	privateUser, err := getPrivateUser(stub, string(certname))
	if err != nil {
		fmt.Println("This foundation user already exists - " + string(certname))
		return shim.Error("This foundation user already exists - " + string(certname))
	}

	privateUser.Username = args[0]
	privateUser.Permissions = nil
	// foundationUser.ProjectPermission = args[2]
	privateUser.UserPermission = args[3]
	privateUser.Role = args[4]

	var location Location
	location.Latitude = args[5]
	location.Longitude = args[6]

	log.Println("final updated foundation user ", privateUser)

	//store user
	userAsBytes, _ := json.Marshal(privateUser)
	err = stub.PutState(privateUser.UserID, userAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	log.Println("- end update foundation")
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

//addAdmin
func addOrg(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	return shim.Success(nil)
}

//updateAdmin
func updateOrg(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	return shim.Success(nil)
}
