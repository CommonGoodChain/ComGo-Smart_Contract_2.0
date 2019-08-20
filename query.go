package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//=============== GET DETAILS OF PROJECT ===============================================================

// ============================================================================================================================
// Get Foundation - get the foundation user from ledger
// ============================================================================================================================
func getProject(stub shim.ChaincodeStubInterface, id string) (Project, error) {
	var project Project
	userAsBytes, err := stub.GetState(id) //getState retreives a key/value from the ledger
	if err != nil {                       //this seems to always succeed, even if key didn't exist
		return project, errors.New("Failed to get project by id - " + id)
	}
	json.Unmarshal(userAsBytes, &project) //un stringify it aka JSON.parse()

	if len(project.ProjectID) == 0 { //test if user is actually here or just nil
		return project, errors.New("project does not exist - " + id + ", '" + project.ProjectID + "' '")
	}

	return project, nil
}

// ============================================================================================================================
// Get Milestone - get the milestone user from ledger
// ============================================================================================================================
func getMilestone(stub shim.ChaincodeStubInterface, id string) (Milestone, error) {
	var milestone Milestone
	userAsBytes, err := stub.GetState(id) //getState retreives a key/value from the ledger
	if err != nil {                       //this seems to always succeed, even if key didn't exist
		return milestone, errors.New("Failed to get milestone by id - " + id)
	}
	json.Unmarshal(userAsBytes, &milestone) //un stringify it aka JSON.parse()

	if len(milestone.MilestoneID) == 0 { //test if user is actually here or just nil
		return milestone, errors.New("milestone does not exist - " + id + ", '" + milestone.MilestoneID + "' '")
	}

	return milestone, nil
}

// ============================================================================================================================
// Get Activity - get the milestone user from ledger
// ============================================================================================================================
func getActivity(stub shim.ChaincodeStubInterface, id string) (Activity, error) {
	var activity Activity
	userAsBytes, err := stub.GetState(id) //getState retreives a key/value from the ledger
	if err != nil {                       //this seems to always succeed, even if key didn't exist
		return activity, errors.New("Failed to get activity by id - " + id)
	}
	json.Unmarshal(userAsBytes, &activity) //un stringify it aka JSON.parse()

	if len(activity.ActivityID) == 0 { //test if user is actually here or just nil
		return activity, errors.New("activity does not exist - " + id + ", '" + activity.ActivityID + "' '")
	}

	return activity, nil
}

//getHistory
// func getHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {

// 	return shim.Success(nil)
// }

func getHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	/*type AuditHistory struct {
		TxId  string `json:"txId"`
		Value Asset  `json:"value"`
	}
	var history []AuditHistory
	var project Asset
	*/
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	projectId := args[0]
	fmt.Printf("- start getHistoryForProject: %s\n", projectId)

	// Get History
	resultsIterator, err := stub.GetHistoryForKey(projectId)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	fmt.Printf("- getHistory resultsIterator: ", resultsIterator)
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true

	}

	buffer.WriteString("]")

	fmt.Printf("- getHistoryForMarble returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

//query
func query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	log.Println("***********Entering getQuery***********")

	log.Println(args)

	queryString := args[0]

	queryResults, err := getQueryResultForQueryStringCouch(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}
