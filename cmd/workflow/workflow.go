package workflow

import (
	"alfmigcli/config"
	"alfmigcli/services/cluster"
	"alfmigcli/services/fetch"
	"alfmigcli/services/pipe"
	"alfmigcli/services/util"
	"alfmigcli/services/workflow"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"
)

//Exec execute cluster command line
func Exec(commands []string) {
	workflowCmd := flag.NewFlagSet("workflow", flag.ExitOnError)
	// Parsing command line flags
	getalldef := workflowCmd.Bool("getalldef", false, "Get all workflows definitions")
	getallandsave := workflowCmd.Bool("getallandsave", false, "Get all workflows and save it")
	getinst := workflowCmd.Bool("getinst", false, "Get workflows intances")
	getinstask := workflowCmd.Bool("getinstask", false, "Get inst with task")
	describe := workflowCmd.Bool("describe", false, "List all workflow")
	create := workflowCmd.Bool("create", false, "create a workflow")
	createByJSON := workflowCmd.Bool("createByJSON", false, "create a workflow by givin a json file")
	createByJSONForm := workflowCmd.Bool("createByJSONForm", false, "create a workflow by givin a json file")
	updateTask := workflowCmd.Bool("updateTask", false, "update a workflow task")
	workflowCmd.Parse(commands) //os.Args[2:])

	l := &cluster.List{}
	workflowList := &workflow.List{}
	pipe.StopIfErrorArg(l.Get())
	currentCluster := pipe.StopIfErrorReturnArg(l.GetCurrent()).(cluster.Item)

	switch {
	case *getalldef:
		result := pipe.StopIfErrorReturnArg(fetch.GetCookies(currentCluster.ClusterURL+fmt.Sprintf(config.GetAllDef, currentCluster.ClusterTICKET), currentCluster.ClusterTICKET)).([]byte)
		wfAllDefResp := &workflow.WfDefAllResponse{}
		pipe.StopIfErrorArg(wfAllDefResp.Decode(result))
		wfAllDefResp.PrintToTable()
	case *getallandsave:
		t := pipe.StopIfErrorReturnArg(util.GetParams(1, os.Stdin, workflowCmd.Args()...)).([]string)
		result := pipe.StopIfErrorReturnArg(fetch.GetCookies(currentCluster.ClusterURL+fmt.Sprintf(config.GetWorkfInst, t[0], currentCluster.ClusterTICKET), currentCluster.ClusterTICKET)).([]byte)
		wfAllDefResp := &workflow.WfDefAllResponse{}
		pipe.StopIfErrorArg(wfAllDefResp.Decode(result))
		//wfAllDefResp.PrintToTable()
		bar := progressbar.Default(int64(len(wfAllDefResp.Data)))
		for _, t := range *&wfAllDefResp.Data {
			//fmt.Println(k)
			result2 := pipe.StopIfErrorReturnArg(fetch.GetCookies(currentCluster.ClusterURL+fmt.Sprintf(config.GetWorkfInstTask, t.Id, currentCluster.ClusterTICKET), currentCluster.ClusterTICKET)).([]byte)
			wfResponse := &workflow.WfResponse{}
			pipe.StopIfErrorArg(wfResponse.Decode(result2))
			workflowList.Add(wfResponse.Data)
			bar.Add(1)
		}
		workflowList.Save(strings.Split(t[0], "$")[1])
	case *createByJSON:
		t := pipe.StopIfErrorReturnArg(util.GetParams(1, os.Stdin, workflowCmd.Args()...)).([]string)
		workflowName := strings.Split(t[0], "$")[1]
		fmt.Println(t[0])
		workflowListToCreate := &workflow.List{}
		workflowListToCreate.GetByFileName(workflowName)

		fmt.Println(fmt.Sprintf("%s %d", "Cantidad de procesos: ", len(*workflowListToCreate)))
		for _, workflowElement := range *workflowListToCreate {
			fmt.Println(workflowElement.Id)
			TaskProperties := workflowElement.Tasks[len(workflowElement.Tasks)-1].Properties
			util.RemoveNulls(TaskProperties)
			workflowProperties, _ := json.Marshal(TaskProperties)
			JsonPostData := fmt.Sprintf("[{ \"processDefinitionKey\": \"%s\", \"variables\": %s }]", workflowName, workflowProperties)
			fmt.Println(JsonPostData) //workflowProperties
			result := pipe.StopIfErrorReturnArg(fetch.PostCookies(currentCluster.ClusterURL+fmt.Sprintf(config.PostCreateWorkflow, currentCluster.ClusterTICKET), currentCluster.ClusterTICKET, []byte(JsonPostData))).([]byte)
			fmt.Println(result)
			responseECWF := &workflow.ResponseEnvelopeCreateWF{}
			pipe.StopIfErrorArg(responseECWF.Decode(result))
			responseECWF.PrintToTable()
		}
	case *createByJSONForm:
		t := pipe.StopIfErrorReturnArg(util.GetParams(1, os.Stdin, workflowCmd.Args()...)).([]string)
		workflowName := strings.Split(t[0], "$")[1]
		fmt.Println(t[0])
		workflowListToCreate := &workflow.List{}
		workflowListToCreate.GetByFileName(workflowName)

		fmt.Println(fmt.Sprintf("%s %d", "Cantidad de procesos: ", len(*workflowListToCreate)))
		for _, workflowElement := range *workflowListToCreate {
			fmt.Println(workflowElement.Id)
			TaskProperties := workflowElement.Tasks[len(workflowElement.Tasks)-1].Properties
			util.RemoveNulls(TaskProperties)
			//workflowProperties, _ := json.Marshal(TaskProperties)
			JsonPostData := fmt.Sprintf("{\"assoc_packageItems_added\": \"%s\",\"prop_bpowf_bpoStartDateEmail\": \"2019-10-15T10:46:00.000Z\", \"prop_bpowf_nombreComprador\": \"%s\" }", TaskProperties["bpm_package"], TaskProperties["bpowf_nombreComprador"])
			fmt.Println(JsonPostData) //workflowProperties
			result := pipe.StopIfErrorReturnArg(fetch.PostCookies(currentCluster.ClusterURL+fmt.Sprintf(config.PostCreateWorkflowForm, workflowName, currentCluster.ClusterTICKET), currentCluster.ClusterTICKET, []byte(JsonPostData))).([]byte)
			fmt.Println(result)
			/*responseECWF := &workflow.ResponseEnvelopeCreateWF{}
			pipe.StopIfErrorArg(responseECWF.Decode(result))
			responseECWF.PrintToTable()*/
		}

	case *getinst:
		t := pipe.StopIfErrorReturnArg(util.GetParams(1, os.Stdin, workflowCmd.Args()...)).([]string)
		result := pipe.StopIfErrorReturnArg(fetch.GetCookies(currentCluster.ClusterURL+fmt.Sprintf(config.GetWorkfInst, t[0], currentCluster.ClusterTICKET), currentCluster.ClusterTICKET)).([]byte)
		wfAllDefResp := &workflow.WfDefAllResponse{}
		pipe.StopIfErrorArg(wfAllDefResp.Decode(result))
		wfAllDefResp.PrintToTable()
	case *getinstask:
		t := pipe.StopIfErrorReturnArg(util.GetParams(1, os.Stdin, workflowCmd.Args()...)).([]string)
		result := pipe.StopIfErrorReturnArg(fetch.GetCookies(currentCluster.ClusterURL+fmt.Sprintf(config.GetWorkfInstTask, t[0], currentCluster.ClusterTICKET), currentCluster.ClusterTICKET)).([]byte)
		wfResponse := &workflow.WfResponse{}
		pipe.StopIfErrorArg(wfResponse.Decode(result))
		wfResponse.PrintToTable()
		fmt.Println(t[0])
	case *create:
		t := pipe.StopIfErrorReturnArg(util.GetParams(1, os.Stdin, workflowCmd.Args()...)).([]string)
		result := pipe.StopIfErrorReturnArg(fetch.PostCookies(currentCluster.ClusterURL+fmt.Sprintf(config.PostCreateWorkflow, currentCluster.ClusterTICKET), currentCluster.ClusterTICKET, []byte(t[0]))).([]byte)
		fmt.Println(result)
		responseECWF := &workflow.ResponseEnvelopeCreateWF{}
		pipe.StopIfErrorArg(responseECWF.Decode(result))
		responseECWF.PrintToTable()
		fmt.Println(t[0])
	case *updateTask:
		t := pipe.StopIfErrorReturnArg(util.GetParams(2, os.Stdin, workflowCmd.Args()...)).([]string)
		result := pipe.StopIfErrorReturnArg(fetch.PutCookies(currentCluster.ClusterURL+fmt.Sprintf(config.PutUpdateTask, t[0], currentCluster.ClusterTICKET), currentCluster.ClusterTICKET, []byte(t[1]))).([]byte)
		fmt.Println(result)
		responseECWF := &workflow.ResponseEnvelopeCreateWF{}
		pipe.StopIfErrorArg(responseECWF.Decode(result))
		responseECWF.PrintToTable()
		fmt.Println(t[0])
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

	if *describe {
		l.PrintToTable()
	}

}
