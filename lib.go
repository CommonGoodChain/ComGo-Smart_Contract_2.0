package main

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ==============================================================
// Input Sanitation - dumb input checking, look for empty strings
// ==============================================================
func sanitize_arguments(strs []string) error {
	for i, val := range strs {
		if len(val) <= 0 {
			return errors.New("Argument " + strconv.Itoa(i) + " must be a non-empty string")
		}
		if len(val) > 2000 {
			return errors.New("Argument " + strconv.Itoa(i) + " must be <= 2000 characters")
		}
	}
	return nil
}

//parseFloat will parse the string values into float
func parseFloat(str string) float64 {

	parsedValue, _ := strconv.ParseFloat(str, 64)

	return parsedValue
}

//parseBoolean will parse the string values into boolean
func parseBool(str string) bool {

	parsedValue, _ := strconv.ParseBool(str)

	return parsedValue
}

// ============================================================================================================================
// Get User Certificate Common Name - get the seller asset from ledger
// ============================================================================================================================
func get_cert(stub shim.ChaincodeStubInterface) ([]byte, error) {

	creator, err := stub.GetCreator() //get the var from ledger
	if err != nil {
		//jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return creator, errors.New("failed to get creator")
	}

	certStart := bytes.IndexAny(creator, "----BEGIN CERTIFICATE-----")
	if certStart == -1 {
		//logger.Debug("No certificate found")
		return creator, errors.New("no certificate found")
	}
	certText := creator[certStart:]
	block, _ := pem.Decode(certText)
	if block == nil {
		//logger.Debug("Error received on pem.Decode of certificate",  certText)
		return creator, errors.New("Error received on pem.Decode of certificate")
	}

	ucert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		//logger.Debug("Error received on ParseCertificate", err)
		return creator, errors.New("Error received on ParseCertificate")
	}

	fmt.Println([]byte(ucert.Subject.CommonName))
	fmt.Println("- end read")

	return []byte(ucert.Subject.CommonName), nil //send it onward
}

// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// =========================================================================================
// getQueryResultForQueryStringCouch executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultInStringForQueryStringCouch(stub shim.ChaincodeStubInterface, queryString string) (string, error) {

	fmt.Printf("- getQueryResultForQueryStringCouch queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return "nil", err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return "nil", err
	}

	fmt.Printf("- getQueryResultForQueryStringCouch queryResult:\n%s\n", buffer.String())

	return buffer.String(), nil
}

// =========================================================================================
// getQueryResultForQueryStringCouch executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultInBytesForQueryStringCouch(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryStringCouch queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryStringCouch queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// =========================================================================================
// getQueryResultForQueryStringCouch executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func myfunction(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryStringCouch queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := myFunctions(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryStringCouch queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// ===========================================================================================
// constructQueryResponseFromIterator constructs a JSON array containing query results from
// a given result iterator
// ===========================================================================================
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}

//-------------GET USER FUNCTION'S START HERE  ------------------------------------------------------------

// ============================================================================================================================
// Get Foundation - get the foundation user from ledger
// ============================================================================================================================
func getPrivateUser(stub shim.ChaincodeStubInterface, id string) (PrivateUser, error) {
	var user PrivateUser
	userAsBytes, err := stub.GetState(id) //getState retreives a key/value from the ledger
	if err != nil {                       //this seems to always succeed, even if key didn't exist
		return user, errors.New("Failed to get private user - " + id)
	}
	json.Unmarshal(userAsBytes, &user) //un stringify it aka JSON.parse()

	if len(user.Username) == 0 { //test if user is actually here or just nil
		return user, errors.New("foundation user does not exist - " + id + ", '" + user.Username + "' '")
	}

	return user, nil
}

// ============================================================================================================================
// Get Foundation - get the foundation user from ledger
// ============================================================================================================================
func getDonor(stub shim.ChaincodeStubInterface, id string) (Donor, error) {
	var donorUser Donor
	userAsBytes, err := stub.GetState(id) //getState retreives a key/value from the ledger
	if err != nil {                       //this seems to always succeed, even if key didn't exist
		return donorUser, errors.New("Failed to get foundation user - " + id)
	}
	json.Unmarshal(userAsBytes, &donorUser) //un stringify it aka JSON.parse()

	if len(donorUser.DonorUsername) == 0 { //test if user is actually here or just nil
		return donorUser, errors.New("foundation user does not exist - " + id + ", '" + donorUser.DonorUsername + "' '" + donorUser.DonorCompany + "'")
	}

	return donorUser, nil
}

// ============================================================================================================================
// Get organization - get the organization user from ledger
// ============================================================================================================================
func getOrg(stub shim.ChaincodeStubInterface, id string) (Organization, error) {
	var organizationUser Organization
	userAsBytes, err := stub.GetState(id) //getState retreives a key/value from the ledger
	if err != nil {                       //this seems to always succeed, even if key didn't exist
		return organizationUser, errors.New("Failed to get organization user - " + id)
	}
	json.Unmarshal(userAsBytes, &organizationUser) //un stringify it aka JSON.parse()

	if len(organizationUser.OrgUsername) == 0 { //test if user is actually here or just nil
		return organizationUser, errors.New("organization does not exist - " + id + ", '" + organizationUser.OrgUsername + "' '" + organizationUser.OrgCompany + "'")
	}

	return organizationUser, nil
}

// =================================================================================================
// query_all_invoice - Query records using a (partial) composite key named by first argument
// =================================================================================================
func query_all(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	log.Println("***********Entering query_invoice_by_status***********")
	// if len(args) < 1 {
	// 	return shim.Error("Incorrect number of arguments. Expecting 1")
	// }
	log.Println(args[1])

	// docType := strings.Replace(args[1], "\"", "", -1)
	// invoiceStatus := status
	docType := args[1]
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"%s\"}}", docType)

	queryResults, err := getQueryResultInBytesForQueryStringCouch(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// =========================================================================================
// getQueryResultForQueryStringCouch executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryStringCouch(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	log.Printf("- getQueryResultForQueryStringCouch queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	log.Printf("- getQueryResultForQueryStringCouch queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// ===========================================================================================
// constructQueryResponseFromIterator constructs a JSON array containing query results from
// a given result iterator
// ===========================================================================================
func myFunctions(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		// buffer.WriteString("{\"Key\":")
		// buffer.WriteString("\"")
		// buffer.WriteString(queryResponse.Key)
		// buffer.WriteString("\"")

		// buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		// buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}
