package node

import (
	"alfmigcli/config"
	"alfmigcli/services/cluster"
	"alfmigcli/services/fetch"
	"alfmigcli/services/node"
	"alfmigcli/services/pipe"
	"alfmigcli/services/util"
	"flag"
	"fmt"
	"os"
)

//Exec command line for node service
func Exec(commands []string) {
	nodeCmd := flag.NewFlagSet("node", flag.ExitOnError)
	spath := nodeCmd.String("spath", "", "Node service search path")
	sid := nodeCmd.String("sid", "", "Node service search id")
	schildren := nodeCmd.String("schildren", "", "Get secondary childern")
	json := nodeCmd.Bool("json", false, "Describe current information in Json format")
	nodeCmd.Parse(commands) // os.Args[2:])

	l := &cluster.List{}
	// Use the Get method to read to do items from file
	pipe.StopIfErrorArg(l.Get())
	switch {
	case *spath != "":
		currentCluster := pipe.StopIfErrorReturnArg(l.GetCurrent()).(cluster.Item)

		if currentCluster.ClusterTICKET == "" {
			fmt.Fprintln(os.Stderr, "Please login first")
			os.Exit(1)
		}
		result := pipe.StopIfErrorReturnArg(fetch.Get(currentCluster.ClusterURL + fmt.Sprintf(config.NodeSearchPath, *spath, currentCluster.ClusterTICKET))).([]byte)

		if *json {
			util.PrintIdentJson(result)
		} else {
			SearchResp := &node.GeneralResponse{}
			pipe.StopIfErrorArg(SearchResp.Decode(result))
			SearchResp.Response.PropertiesPrintToTable()
			SearchResp.Response.DataPrintToTable()
		}
	case *schildren != "":
		currentCluster := pipe.StopIfErrorReturnArg(l.GetCurrent()).(cluster.Item)

		if currentCluster.ClusterTICKET == "" {
			fmt.Fprintln(os.Stderr, "Please login first")
			os.Exit(1)
		}
		result := pipe.StopIfErrorReturnArg(fetch.Get(currentCluster.ClusterURL + fmt.Sprintf(config.NodeSecondaryChildren, *schildren, currentCluster.ClusterTICKET))).([]byte)
		util.PrintIdentJson(result)
	case *sid != "":
		currentCluster := pipe.StopIfErrorReturnArg(l.GetCurrent()).(cluster.Item)

		if currentCluster.ClusterTICKET == "" {
			fmt.Fprintln(os.Stderr, "Please login first")
			os.Exit(1)
		}

		t := pipe.StopIfErrorReturnArg(util.GetParams(1, os.Stdin, nodeCmd.Args()...)).([]string)

		result := pipe.StopIfErrorReturnArg(fetch.Get(currentCluster.ClusterURL + fmt.Sprintf(config.NodeSearchId, t[0], *sid, currentCluster.ClusterTICKET))).([]byte)

		if *json {
			util.PrintIdentJson(result)
		} else {
			SearchResp := &node.MessageResponse{}
			pipe.StopIfErrorArg(SearchResp.Decode(result))
			result := SearchResp.Response["msg"]
			result.PropertiesPrintToTable()
			result.DataPrintToTable()
		}
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
