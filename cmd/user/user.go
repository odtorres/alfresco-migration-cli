package user

import (
	"alfmigcli/config"
	"alfmigcli/services/cluster"
	"alfmigcli/services/fetch"
	"alfmigcli/services/pipe"
	"alfmigcli/services/user"
	"alfmigcli/services/util"
	"flag"
	"fmt"
	"os"
)

//Exec execute User command line
func Exec(commands []string) {
	userCmd := flag.NewFlagSet("user", flag.ExitOnError)
	login := userCmd.Bool("login", false, "Login in current cluster")
	describe := userCmd.Bool("describe", false, "Describe current user")

	userCmd.Parse(commands) //os.Args[2:])

	l := &cluster.List{}
	// Use the Get method to read to do items from file
	pipe.StopIfErrorArg(l.Get())

	switch {
	case *login:
		t := pipe.StopIfErrorReturnArg(util.GetParams(2, os.Stdin, userCmd.Args()...)).([]string)
		currentCluster := pipe.StopIfErrorReturnArg(l.GetCurrent()).(cluster.Item)
		result := pipe.StopIfErrorReturnArg(fetch.Get(currentCluster.ClusterURL + fmt.Sprintf(config.Login, t[0], t[1]))).([]byte)
		loginResp := &user.LoginResponse{}
		pipe.StopIfErrorArg(loginResp.Decode(result))
		fmt.Println(loginResp.Data.Ticket)
		if loginResp.Data.Ticket != "access_denied" {
			//Set current Ticket
			pipe.StopIfErrorArg(l.SetClusterTICKET(loginResp.Data.Ticket))
			pipe.StopIfErrorArg(l.Save())
		}
	case *describe:
		currentCluster := pipe.StopIfErrorReturnArg(l.GetCurrent()).(cluster.Item)

		if currentCluster.ClusterTICKET != "" {
			result := pipe.StopIfErrorReturnArg(fetch.Get(currentCluster.ClusterURL + fmt.Sprintf(config.UserInfo, currentCluster.ClusterTICKET))).([]byte)
			util.PrintIdentJson(result)
		} else {
			fmt.Fprintln(os.Stderr, "Please login first")
			os.Exit(1)
		}

	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
