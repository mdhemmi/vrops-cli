package main

import (
	"flag"
	"fmt"

	vrops_cli "vrops-cli/lib"
)

var version = "0.0.1"
var extendedFlag bool

func main() {
	//fmt.Println("Version: " + version)

	vropsdata := new(vrops_cli.VROPsData)
	flag.StringVar(&vropsdata.Action, "action", "null", "Action: query, search")
	flag.StringVar(&vropsdata.FQDN, "fqdn", "null", "vROPs FQDN or IP")
	flag.StringVar(&vropsdata.Auth, "auth", "local", "Auth source local, vIDMAuthSource")
	flag.StringVar(&vropsdata.Search, "search", "null", "Search string")
	flag.BoolVar(&extendedFlag, "e", false, "Extended output true / false")
	flag.StringVar(&vropsdata.Username, "u", "null", "vROPs username")
	flag.StringVar(&vropsdata.Password, "p", "null", "vROPs user password")
	flag.BoolVar(&vropsdata.Insecure, "i", true, "Insecure SSL connection")
	flag.BoolVar(&vropsdata.Debug, "d", false, "Debug output")

	flag.Parse()

	vropsdata.Extended = vrops_cli.IsFlagPassed("e")
	//vropsdata.Username = vrops_cli.Get_Username()

	//vropsdata.Debug = false
	//vropsdata.Insecure = true
	switch vropsdata.Action {
	case "setup":
		fmt.Println("Setup keyring")
		//vrops_cli.Setup_keyring(vropsdata)
	case "updatepw":
		fmt.Println("Update Password")
	default:
		vrops_cli.Action_SearchvROPs(vropsdata)

	}

}
