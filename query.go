package main

import (
	"encoding/json"
	"errors"

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
func getHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	return shim.Success(nil)
}

//query
func query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	return shim.Success(nil)
}
