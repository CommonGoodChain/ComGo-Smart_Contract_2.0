package main

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// ==============================================================
// Input Sanitation - dumb input checking, look for empty strings
// ==============================================================
func sanitize_arguments(strs []string) error {
	for i, val := range strs {
		if len(val) <= 0 {
			return errors.New("Argument " + strconv.Itoa(i) + " must be a non-empty string")
		}
		if len(val) > 200 {
			return errors.New("Argument " + strconv.Itoa(i) + " must be <= 200 characters")
		}
	}
	return nil
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
func getQueryResultForQueryStringCouch(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

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
func getFoundation(stub shim.ChaincodeStubInterface, id string) (Foundation, error) {
	var user Foundation
	userAsBytes, err := stub.GetState(id) //getState retreives a key/value from the ledger
	if err != nil {                       //this seems to always succeed, even if key didn't exist
		return user, errors.New("Failed to get foundation user - " + id)
	}
	json.Unmarshal(userAsBytes, &user) //un stringify it aka JSON.parse()

	if len(user.FoundationUsername) == 0 { //test if user is actually here or just nil
		return user, errors.New("foundation user does not exist - " + id + ", '" + user.FoundationUsername + "' '" + user.FoundationCompany + "'")
	}

	return user, nil
}
