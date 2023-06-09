package vrops_cli

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get_vrops_token(vrops string, vropsusername string, vropspassword string, authSource string, debug bool, insecure bool) string {
	var token string
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure}
	url := "https://" + vrops + "/suite-api/api/auth/token/acquire"
	loginjson := "{\"username\": \"" + vropsusername + "\",\"password\": \"" + vropspassword + "\",\"authSource\": \"" + authSource + "\"}"
	var jsonStr = []byte(loginjson)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	xmlstring := string(body)

	token = Between(xmlstring, "<ops:token>", "</ops:token>")
	if debug {
		funcname := GetCurrentFuncName()
		Debug("", "", "start")
		Debug(funcname, "string", "print")
		Debug(loginjson, "string", "print")
		Debug(xmlstring, "string", "print")
		Debug(token, "string", "print")
		Debug("", "", "end")
	}

	return token
}

// ==============================================================================================================================

func Get_Resource(vropshost string, token string, item string, debug bool, insecure bool) string {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure}
	//url := "https://" + vropshost + "/suite-api/api/resources?name=" + item + "&resourceKind=VirtualMachine&adapterKind=VMWARE"
	url := "https://" + vropshost + "/suite-api/api/resources?name=" + item + "&adapterKind=VMWARE"
	//fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	vropstoken := "vRealizeOpsToken " + token
	req.Header.Set("Authorization", vropstoken)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	jsonstring := string(body)
	//spew.Dump(jsonstring)
	resource := Between(jsonstring, "identifier=\"", "\"><ops:resourceKey>")
	//fmt.Println(resource)
	if debug {
		funcname := GetCurrentFuncName()
		Debug("", "", "start")
		Debug(funcname, "string", "print")
		Debug(jsonstring, "string", "print")
		Debug(resource, "string", "print")
		Debug(token, "string", "print")
		Debug("", "", "end")
	}
	return resource
}

// ==============================================================================================================================

func Get_Resource_Stats(vropshost string, token string, item string, debug bool, insecure bool) StatsOfResources {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure}
	url := "https://" + vropshost + "/suite-api/api/resources/stats/latest?resourceId=" + item
	//fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	vropstoken := "vRealizeOpsToken " + token
	req.Header.Set("Authorization", vropstoken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	jsonstring := string(body)
	//spew.Dump(xmlstring)
	stats := new(StatsOfResources)
	err = json.Unmarshal([]byte(jsonstring), &stats)
	if err != nil {
		fmt.Println(err.Error())
	}
	//spew.Dump(stats)
	if debug {
		funcname := GetCurrentFuncName()
		Debug("", "", "start")
		Debug(funcname, "string", "print")
		Debug(jsonstring, "string", "print")
		Debug(token, "string", "print")
		Debug("", "", "end")
	}
	return *stats
}

// ==============================================================================================================================

func Get_Resource_Properties(vropshost string, token string, item string, debug bool, insecure bool) PropertiesOfResources {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure}
	url := "https://" + vropshost + "/suite-api/api/resources/properties?resourceId=" + item
	//fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	vropstoken := "vRealizeOpsToken " + token
	req.Header.Set("Authorization", vropstoken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	jsonstring := string(body)
	//spew.Dump(xmlstring)
	properties := new(PropertiesOfResources)
	err = json.Unmarshal([]byte(jsonstring), &properties)
	if err != nil {
		fmt.Println(err.Error())
	}
	//spew.Dump(stats)
	if debug {
		funcname := GetCurrentFuncName()
		Debug("", "", "start")
		Debug(funcname, "string", "print")
		Debug(jsonstring, "string", "print")
		Debug(token, "string", "print")
		Debug("", "", "end")
	}
	return *properties
}
