package cluster

import (
	"alfmigcli/services/cluster"
	"alfmigcli/services/pipe"
	"alfmigcli/services/util"
	"flag"
	"fmt"
	"os"
)

//Exec execute cluster command line
func Exec(commands []string) {
	clusterCmd := flag.NewFlagSet("cluster", flag.ExitOnError)
	// Parsing command line flags
	add := clusterCmd.Bool("add", false, "Add a cluster to the list")
	describe := clusterCmd.Bool("describe", false, "List all clusters")
	visible := clusterCmd.Int("visible", 0, "Set to visible")
	invisible := clusterCmd.Int("invisible", 0, "Set to invisible")
	current := clusterCmd.Int("current", 0, "Is current")
	delete := clusterCmd.Int("delete", 0, "Remove item")
	ticket := clusterCmd.Bool("ticket", false, "Get ticket for current cluster")
	clusterCmd.Parse(commands) //os.Args[2:])

	l := &cluster.List{}
	pipe.StopIfErrorArg(l.Get())

	switch {
	case *add:
		t := pipe.StopIfErrorReturnArg(util.GetParams(2, os.Stdin, clusterCmd.Args()...)).([]string)
		l.Add(t[0], t[1])
		pipe.StopIfErrorArg(l.Save())
	case *visible > 0:
		pipe.StopIfErrorArg(l.Visible(*visible, true))
		pipe.StopIfErrorArg(l.Save())
	case *invisible > 0:
		pipe.StopIfErrorArg(l.Visible(*invisible, false))
		pipe.StopIfErrorArg(l.Save())
	case *current > 0:
		pipe.StopIfErrorArg(l.Current(*current))
		pipe.StopIfErrorArg(l.Save())
	case *delete > 0:
		pipe.StopIfErrorArg(l.Delete(*delete))
		pipe.StopIfErrorArg(l.Save())
	case *ticket:
		cluster := pipe.StopIfErrorReturnArg(l.GetCurrent()).(cluster.Item)
		fmt.Print(cluster.ClusterTICKET)
	}
	if *describe {
		l.PrintToTable()
	}

}
