package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//=============== PROJECT RELATED FUNCTION'S START HERE ===============================================================

//addProject
func addProject(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("starting - project creation")

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}
	log.Println(certname)

	if len(args) != 23 {
		return shim.Error("Incorrect number of arguments. Expecting 23")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	//set project details
	var project Project
	project.ProjectID = args[0]

	// get the project
	proj, err := getProject(stub, project.ProjectID)
	if err == nil {
		fmt.Println("Project is already present " + proj.ProjectID)
		return shim.Error(err.Error())
	}

	sdgString := args[18]
	var list []string
	dec := json.NewDecoder(strings.NewReader(sdgString))
	errs := dec.Decode(&list)
	log.Println(errs, list)
	log.Println(dec.Decode(&list))
	var sdg []SDG
	for i := range list {
		var s SDG
		s.SDGType = list[i]
		sdg = append(sdg, s)
	}

	project.ObjectType = "Project"
	project.Organization = args[1]
	project.NGOCompany = args[2]
	project.ProjectName = args[3]
	project.FundGoal = parseFloat(args[4])
	project.ProjectType = args[5]
	project.StartDate = args[6]
	project.EndDate = args[7]
	project.Description = args[8]
	project.Currency = args[9]
	project.FundRaised = parseFloat(args[10])
	project.FundAllocated = parseFloat(args[11])
	project.ProjectBudget = parseFloat(args[12])
	project.ProjectOwner = args[13]
	project.FundAllocationType = args[14]
	project.IsPublished = parseBool(args[15])
	project.Status = args[16]
	project.Flag = args[17]
	project.SDG = sdg

	var location Location
	location.Latitude = args[19]
	location.Longitude = args[20]
	project.Country = args[21]
	project.FundNotAllocated = parseFloat(args[22])
	project.ProjectLoc = location

	log.Println("project object is creataed ", project)

	//store project
	projectAsBytes, _ := json.Marshal(project) //convert to array of bytes
	errz := stub.PutState(project.ProjectID, projectAsBytes)

	if errz != nil {
		fmt.Println("Could not store project")
		return shim.Error(errz.Error())
	}

	log.Println("- end - Project creation")
	return shim.Success(nil)
}

//updateProject
func updateProject(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("starting - update project")

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}
	log.Println(certname)

	if len(args) != 23 {
		return shim.Error("Incorrect number of arguments. Expecting 23")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	// get the project
	project, err := getProject(stub, args[0])
	if err != nil {
		fmt.Println("Project is missing " + args[0])
		return shim.Error(err.Error())
	}

	sdgString := args[18]
	var list []string
	dec := json.NewDecoder(strings.NewReader(sdgString))
	errs := dec.Decode(&list)
	log.Println(errs, list)
	log.Println(dec.Decode(&list))
	var sdg []SDG
	for i := range list {
		var s SDG
		s.SDGType = list[i]
		sdg = append(sdg, s)
	}

	project.Organization = args[1]
	project.NGOCompany = args[2]
	project.ProjectName = args[3]
	project.FundGoal = parseFloat(args[4])
	project.ProjectType = args[5]
	project.StartDate = args[6]
	project.EndDate = args[7]
	project.Description = args[8]
	project.Currency = args[9]
	project.FundRaised = parseFloat(args[10])
	project.FundAllocated = parseFloat(args[11])
	project.ProjectBudget = parseFloat(args[12])
	project.ProjectOwner = args[13]
	project.FundAllocationType = args[14]
	project.IsPublished = parseBool(args[15])
	project.Status = args[16]
	project.Flag = args[17]
	project.SDG = sdg

	var location Location
	location.Latitude = args[19]
	location.Longitude = args[20]
	project.Country = args[21]
	project.FundNotAllocated = parseFloat(args[22])
	project.ProjectLoc = location

	log.Println("update project object is creataed ", project)

	//store project
	projectAsBytes, _ := json.Marshal(project) //convert to array of bytes
	errz := stub.PutState(project.ProjectID, projectAsBytes)

	if errz != nil {
		log.Println("Could not update project")
		return shim.Error(errz.Error())
	}

	log.Println("- end - update project")
	return shim.Success(nil)
}

//update project status and flag
func updateProjectStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("starting - update Project project status and flag")

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}
	log.Println(certname)

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	// get the project
	project, err := getProject(stub, args[0])
	if err != nil {
		fmt.Println("Project is missing " + args[0])
		return shim.Error(err.Error())
	}

	project.Status = args[1]
	project.Flag = args[2]
	project.IsPublished = parseBool(args[3])
	project.IsApproved = parseBool(args[4])
	project.Remarks = args[5]

	log.Println("update Project project status and flag object is creataed ", project)

	//store project
	projectAsBytes, _ := json.Marshal(project) //convert to array of bytes
	errz := stub.PutState(project.ProjectID, projectAsBytes)

	if errz != nil {
		log.Println("Could not update the status and flag of project")
		return shim.Error(errz.Error())
	}

	log.Println("- end - update Project project status and flag")

	return shim.Success(nil)
}

//delete project
func deleteProject(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("starting - delete project")

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}
	log.Println(certname)

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	// get the project
	project, err := getProject(stub, args[0])
	if err != nil {
		fmt.Println("Project is missing " + args[0])
		return shim.Error(err.Error())
	}

	log.Println("delete project ", project)

	err = stub.DelState(args[0]) //remove the key from chaincode state
	if err != nil {
		return shim.Error("Failed to delete project")
	}

	log.Println("- end - delete Project")

	return shim.Success(nil)
}

//=============== MILESTONE RELATED FUNCTION'S START HERE ===============================================================

//addMilestone
func addMilestone(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("starting - milestone creation")

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}
	log.Println(certname)

	if len(args) != 10 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	// check the milestone
	mil, err := getMilestone(stub, args[1])

	if err == nil {
		fmt.Println("MilestoneID is already present " + mil.MilestoneID)
		return shim.Error(err.Error())
	}

	// get the project
	project, err := getProject(stub, args[0])
	if err != nil {
		fmt.Println("Project is missing " + args[0])
		return shim.Error(err.Error())
	}

	var milestone Milestone

	milestone.ObjectType = "Milestone"
	milestone.ProjectID = args[0]
	milestone.MilestoneID = args[1]
	milestone.MilestoneName = args[2]
	milestone.StartDate = args[3]
	milestone.EndDate = args[4]
	milestone.Description = args[5]
	milestone.Status = args[6]
	milestone.IsApproved = parseBool(args[7])

	project.Status = args[8]
	project.Flag = args[9]

	//update project
	projectAsBytes, _ := json.Marshal(project) //convert to array of bytes
	errp := stub.PutState(project.ProjectID, projectAsBytes)

	if errp != nil {
		log.Println("Could not update the status and flag of project")
		return shim.Error(errp.Error())
	}

	//store project
	milestoneAsBytes, _ := json.Marshal(milestone) //convert to array of bytes
	errz := stub.PutState(milestone.MilestoneID, milestoneAsBytes)
	if errz != nil {
		fmt.Println("Could not store milestone")
		return shim.Error(errz.Error())
	}

	log.Println("- end - milestone creation")

	return shim.Success(nil)
}

//updateMilestone
func updateMilestone(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("starting - updatae milestone")

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}
	log.Println(certname)

	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	// check the milestone
	milestone, err := getMilestone(stub, args[0])

	if err != nil {
		fmt.Println("Milestone is not present " + milestone.MilestoneID)
		return shim.Error(err.Error())
	}

	// get the project
	project, err := getProject(stub, milestone.ProjectID)
	if err != nil {
		fmt.Println("Project is missing " + milestone.ProjectID)
		return shim.Error(err.Error())
	}

	milestone.MilestoneName = args[1]
	milestone.StartDate = args[2]
	milestone.EndDate = args[3]
	milestone.Description = args[4]
	milestone.Status = args[5]

	project.Status = args[6]
	project.Flag = args[7]

	//update project
	projectAsBytes, _ := json.Marshal(project) //convert to array of bytes
	errp := stub.PutState(project.ProjectID, projectAsBytes)

	if errp != nil {
		log.Println("Could not update the milestone")
		return shim.Error(errp.Error())
	}

	//store project
	milestoneAsBytes, _ := json.Marshal(milestone) //convert to array of bytes
	errz := stub.PutState(milestone.MilestoneID, milestoneAsBytes)
	if errz != nil {
		fmt.Println("Could not update milestone")
		return shim.Error(errz.Error())
	}

	log.Println("- end - update milestone")

	return shim.Success(nil)
}

//update milestone status
func updateMilestoneStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("starting - update milestone status")

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}
	log.Println(certname)

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	// check the milestone
	milestone, err := getMilestone(stub, args[0])

	if err != nil {
		fmt.Println("Milestone is not present " + milestone.MilestoneID)
		return shim.Error(err.Error())
	}

	// get the project
	project, err := getProject(stub, milestone.ProjectID)
	if err != nil {
		fmt.Println("Project is missing " + milestone.ProjectID)
		return shim.Error(err.Error())
	}

	// upate milestone
	milestone.Status = args[1]
	milestone.IsApproved = parseBool(args[2])

	// update project
	project.Status = args[3]
	project.Flag = args[4]
	project.IsApproved = parseBool(args[5])

	log.Println("update milestone status object is creataed ", project)

	//store project
	projectAsBytes, _ := json.Marshal(project) //convert to array of bytes
	errz := stub.PutState(project.ProjectID, projectAsBytes)

	if errz != nil {
		log.Println("Could not update the status and flag of project")
		return shim.Error(errz.Error())
	}

	//store project
	milestoneAsBytes, _ := json.Marshal(milestone) //convert to array of bytes
	errm := stub.PutState(milestone.MilestoneID, milestoneAsBytes)
	if errm != nil {
		fmt.Println("Could not update milestone")
		return shim.Error(errm.Error())
	}

	log.Println("- end - update milestone status")

	return shim.Success(nil)
}

//delete milestone
func deleteMilestone(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("starting - delete milestone")

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}
	log.Println(certname)

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	// get the milestone
	milestone, err := getMilestone(stub, args[0])
	if err != nil {
		fmt.Println("milestone is missing " + args[0])
		return shim.Error(err.Error())
	}

	// get the project
	project, err := getProject(stub, milestone.ProjectID)
	if err != nil {
		fmt.Println("Project is missing " + milestone.ProjectID)
		return shim.Error(err.Error())
	}

	project.Status = args[1]
	project.Flag = args[2]

	log.Println("delete milestone ", milestone)

	err = stub.DelState(args[0]) //remove the key from chaincode state
	if err != nil {
		return shim.Error("Failed to delete milestone")
	}

	log.Println("- end - delete milestone")

	return shim.Success(nil)
}

//=============== ACTIVITY RELATED FUNCTION'S START HERE ===============================================================

//addActivity
func addActivity(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("starting - activity creation")

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}
	log.Println(certname)

	if len(args) != 18 {
		return shim.Error("Incorrect number of arguments. Expecting 18")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	// check the milestone
	act, err := getActivity(stub, args[2])
	if err == nil {
		fmt.Println("ActivityID is already present " + act.MilestoneID)
		return shim.Error(err.Error())
	}

	// get the milestone
	milestone, err := getMilestone(stub, args[1])
	if err != nil {
		fmt.Println("Milestone is not present " + milestone.MilestoneID)
		return shim.Error(err.Error())
	}

	// get the project
	project, err := getProject(stub, args[0])
	if err != nil {
		fmt.Println("Project is missing " + args[0])
		return shim.Error(err.Error())
	}

	var activity Activity

	activity.ObjectType = "Activity"
	activity.ProjectID = args[0]
	activity.MilestoneID = args[1]
	activity.ActivityID = args[2]
	activity.ActivityName = args[3]
	activity.StartDate = args[4]
	activity.EndDate = args[5]
	activity.ActivityBudget = parseFloat(args[6])
	activity.Description = args[7]
	activity.SecondaryValidation = parseBool(args[8])
	activity.Remarks = args[9]
	activity.IsApproved = parseBool(args[10])
	activity.ValidatorID = args[11]
	activity.Status = args[12]
	activity.TechnicalCriteria = args[13]
	activity.FinancialCriteria = args[14]
	//update milstone status
	milestone.Status = args[15]

	//update project status
	project.Status = args[16]
	project.Flag = args[17]
	//update project
	projectAsBytes, _ := json.Marshal(project) //convert to array of bytes
	errp := stub.PutState(project.ProjectID, projectAsBytes)

	if errp != nil {
		log.Println("Could not update the status and flag of project")
		return shim.Error(errp.Error())
	}

	//update milestone
	milestoneAsBytes, _ := json.Marshal(milestone) //convert to array of bytes
	errz := stub.PutState(milestone.MilestoneID, milestoneAsBytes)
	if errz != nil {
		fmt.Println("Could not store milestone")
		return shim.Error(errz.Error())
	}

	//update activity
	activityAsBytes, _ := json.Marshal(activity) //convert to array of bytes
	erra := stub.PutState(activity.ActivityID, activityAsBytes)
	if erra != nil {
		fmt.Println("Could not store milestone")
		return shim.Error(erra.Error())
	}

	log.Println("- end - activity creation")

	return shim.Success(nil)
}

//updateActivity
func updateActivity(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("starting - activity creation")

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}
	log.Println(certname)

	if len(args) != 16 {
		return shim.Error("Incorrect number of arguments. Expecting 16")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	// check the activity
	activity, err := getActivity(stub, args[0])
	if err != nil {
		fmt.Println("ActivityID is not present " + activity.ActivityID)
		return shim.Error(err.Error())
	}

	// get the milestone
	milestone, err := getMilestone(stub, activity.MilestoneID)
	if err != nil {
		fmt.Println("Milestone is not present " + activity.MilestoneID)
		return shim.Error(err.Error())
	}

	// get the project
	project, err := getProject(stub, activity.ProjectID)
	if err != nil {
		fmt.Println("Project is missing " + activity.ProjectID)
		return shim.Error(err.Error())
	}

	activity.ActivityName = args[1]
	activity.StartDate = args[2]
	activity.EndDate = args[3]
	activity.ActivityBudget = parseFloat(args[4])
	activity.Description = args[5]
	activity.SecondaryValidation = parseBool(args[6])
	activity.Remarks = args[7]
	activity.IsApproved = parseBool(args[8])
	activity.ValidatorID = args[9]
	activity.Status = args[10]
	activity.TechnicalCriteria = args[11]
	activity.FinancialCriteria = args[12]
	//update milstone status
	milestone.Status = args[13]

	//update project status
	project.Status = args[14]
	project.Flag = args[15]

	//update project
	projectAsBytes, _ := json.Marshal(project) //convert to array of bytes
	errp := stub.PutState(project.ProjectID, projectAsBytes)

	if errp != nil {
		log.Println("Could not update the project")
		return shim.Error(errp.Error())
	}

	//update milestone
	milestoneAsBytes, _ := json.Marshal(milestone) //convert to array of bytes
	errz := stub.PutState(milestone.MilestoneID, milestoneAsBytes)
	if errz != nil {
		fmt.Println("Could not update milestone")
		return shim.Error(errz.Error())
	}

	//update activity
	activityAsBytes, _ := json.Marshal(activity) //convert to array of bytes
	erra := stub.PutState(activity.ActivityID, activityAsBytes)
	if erra != nil {
		fmt.Println("Could not update activity")
		return shim.Error(erra.Error())
	}

	log.Println("- end - update activity")

	return shim.Success(nil)
}

//update activity status
func updateActivityStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("starting - update activity status")

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}
	log.Println(certname)

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	// check the activity
	activity, err := getActivity(stub, args[0])
	if err != nil {
		fmt.Println("ActivityID is not present " + activity.ActivityID)
		return shim.Error(err.Error())
	}

	// check the milestone
	milestone, err := getMilestone(stub, activity.MilestoneID)

	if err != nil {
		fmt.Println("Milestone is not present " + activity.MilestoneID)
		return shim.Error(err.Error())
	}

	// get the project
	project, err := getProject(stub, activity.ProjectID)
	if err != nil {
		fmt.Println("Project is missing " + activity.ProjectID)
		return shim.Error(err.Error())
	}

	//update activity
	activity.Status = args[1]
	activity.IsApproved = parseBool(args[2])
	activity.Remarks = args[3]
	// upate milestone
	milestone.Status = args[4]

	// update project
	project.Status = args[5]
	project.Flag = args[6]

	log.Println("update milestone status object is creataed ", project)

	//store project
	projectAsBytes, _ := json.Marshal(project) //convert to array of bytes
	errz := stub.PutState(project.ProjectID, projectAsBytes)

	if errz != nil {
		log.Println("Could not update the status and flag of project")
		return shim.Error(errz.Error())
	}

	//store project
	milestoneAsBytes, _ := json.Marshal(milestone) //convert to array of bytes
	errm := stub.PutState(milestone.MilestoneID, milestoneAsBytes)
	if errm != nil {
		fmt.Println("Could not update milestone")
		return shim.Error(errm.Error())
	}

	activityAsBytes, _ := json.Marshal(activity) //convert to array of bytes
	erra := stub.PutState(activity.ActivityID, activityAsBytes)
	if erra != nil {
		fmt.Println("Could not update activity")
		return shim.Error(erra.Error())
	}

	log.Println("- end - update activity status")

	return shim.Success(nil)
}

//delete activity
func deleteActivity(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("starting - delete activity")

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}
	log.Println(certname)

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	// check the activity
	activity, err := getActivity(stub, args[0])
	if err != nil {
		fmt.Println("ActivityID is not present " + activity.ActivityID)
		return shim.Error(err.Error())
	}

	// get the milestone
	milestone, err := getMilestone(stub, activity.MilestoneID)
	if err != nil {
		fmt.Println("milestone is missing " + activity.MilestoneID)
		return shim.Error(err.Error())
	}

	// get the project
	project, err := getProject(stub, activity.ProjectID)
	if err != nil {
		fmt.Println("Project is missing " + activity.ProjectID)
		return shim.Error(err.Error())
	}

	//update milestone
	milestone.Status = args[1]

	//update project
	project.Status = args[2]
	project.Flag = args[3]

	log.Println("delete activity ", activity)

	err = stub.DelState(args[0]) //remove the key from chaincode state
	if err != nil {
		return shim.Error("Failed to delete activity")
	}

	log.Println("- end - delete activity")

	return shim.Success(nil)
}

//=============== DONATION RELATED FUNCTION'S START HERE ===============================================================

//fund project
func fundProject(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	log.Println("starting - fund project")

	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}
	log.Println(certname)

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	// get the project
	project, err := getProject(stub, args[0])
	if err != nil {
		fmt.Println("Project is missing ", args[0])
		return shim.Error(err.Error())
	}
	var donationAmt = parseFloat(args[1])
	donationAmt += project.FundNotAllocated
	project.FundRaised = donationAmt
	project.Flag = args[2]
	Order := "asc"
	if project.FundAllocationType == "2" { // auto fund allocate
		//get all activities whose activity budget is >= donation amount (sort by date= chronologically)
		queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Activity\",\"projectId\":\"%s\"}}",args[0])
		act, err := getQueryResultInBytesForQueryStringCouch(stub, queryString)
		if err != nil {
			return shim.Error(err.Error())
		}
		var activities []Activity
		json.Unmarshal([]byte(act), &activities)
		fmt.Println("Activities are: ", activities)
		actBalAmt := donationAmt
		for i := range activities {
			actFundRem := activities[i].ActivityBudget - activities[i].FundAllocated
			if actBalAmt >= actFundRem {
				activities[i].FundAllocated += actFundRem
				activities[i].Status = "Fund Allocated"
				project.FundAllocated += actFundRem
				project.Status = "Fund Allocated"
				project.FundNotAllocated = project.FundNotAllocated - project.FundNotAllocated
			} else {
				// add donation amount in fundNotAllocated field
				project.FundNotAllocated += actBalAmt
				break
			}

			if actBalAmt >= actFundRem {
				actBalAmt = actBalAmt - actFundRem
			}

			if i == len(activities)-1 {
				project.FundNotAllocated = actBalAmt
			}

			//update actvity
			actAsBytes, _ := json.Marshal(activities[i])                  //convert to array of bytes
			errAct := stub.PutState(activities[i].ActivityID, actAsBytes) //rewrite the project with id as key
			if errAct != nil {
				fmt.Println(errAct)
			}
		}
	} else {
		project.FundNotAllocated += parseFloat(args[1])
	}
	log.Println("project object after donation ", project)

	//store project
	projectAsBytes, _ := json.Marshal(project) //convert to array of bytes
	errz := stub.PutState(project.ProjectID, projectAsBytes)
	if errz != nil {
		log.Println("could not fund project")
		return shim.Error(errz.Error())
	}

	log.Println("- end - fund project")

	return shim.Success(nil)
}

//fund
func fund(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

//==============PROOF RELATED FUNCTION'S START HERE ===================================================
func submitProof(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	fmt.Println("starting submit_proof")

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}
	certname, err := get_cert(stub)
	if err != nil {
		fmt.Printf("INVOKE: Error retrieving cert: %s", err)
		return shim.Error("Error retrieving cert")
	}

	fmt.Println("certname ", string(certname))

	return shim.Success(nil)
}