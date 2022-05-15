package node

import (
	"alfmigcli/config"
	"alfmigcli/services/cluster"
	"alfmigcli/services/fetch"
	"alfmigcli/services/node"
	"alfmigcli/services/pipe"
	"alfmigcli/services/workflow"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

//Exec command line for node service
func Exec(commands []string) {
	nodeCmd := flag.NewFlagSet("node", flag.ExitOnError)
	verifywfnodes := nodeCmd.Bool("verifywfnodes", false, "verify if nodes exist in the wf")

	nodeCmd.Parse(commands) // os.Args[2:])

	l := &cluster.List{}
	lwf := &workflow.List{}
	lwfv := &workflow.ListVerify{}
	pipe.StopIfErrorArg(l.Get())
	pipe.StopIfErrorArg(lwf.GetFromFile())
	currentCluster := pipe.StopIfErrorReturnArg(l.GetCurrent()).(cluster.Item)

	switch {
	case *verifywfnodes:
		anyError := false
		errorsSize := 0
		for _, e := range *lwf {
			result := pipe.StopIfErrorReturnArg(fetch.GetCookies(currentCluster.ClusterURL+fmt.Sprintf(config.NodeSearchUUID, e.Tasks[0].Properties["cm_noderef"], currentCluster.ClusterTICKET), currentCluster.ClusterTICKET)).([]byte)
			gresponse := &node.List{}
			json.Unmarshal(result, gresponse)
			//fmt.Println(currentCluster.ClusterURL + fmt.Sprintf(config.NodeSearchUUID, e.Tasks[0].Properties["cm_noderef"], currentCluster.ClusterTICKET))

			if len(*gresponse) == 0 {
				fmt.Println(e.Tasks[0].Properties["cm_noderef"])
				anyError = true
				errorsSize++
			} else {
				fmt.Println((*gresponse)[0].Name)
				lwfv.Add(e)
			}
		}
		if anyError {
			fmt.Println("Missing nodes: " + fmt.Sprintf("%d", errorsSize))
		}
		if len(*lwfv) > 0 {
			lwfv.Save()
		}
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
