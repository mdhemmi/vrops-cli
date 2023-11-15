package vrops_cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Action_SearchvROPs(vropsdata *VROPsData) {
	if Check_Service_Availability(vropsdata.FQDN, 443) {

		token := Get_vrops_token(vropsdata.FQDN, vropsdata.Username, vropsdata.Password, vropsdata.Auth, vropsdata.Debug, vropsdata.Insecure)
		//fmt.Println(token)
		resource := Get_Resource(vropsdata.FQDN, token, vropsdata.Search, vropsdata.Debug, vropsdata.Insecure)
		if resource != "" {
			fmt.Println("vROPs: " + vropsdata.FQDN)
			fmt.Println("vROPS Resource ID: " + resource)
			//fmt.Println(" ")
			//stats := Get_Resource_Stats(item.Name, token, resource, vropsdata.Debug, vropsdata.Insecure)
			//spew.Dump(stats)
			properties := Get_Resource_Properties(vropsdata.FQDN, token, resource, vropsdata.Debug, vropsdata.Insecure)
			//spew.Dump(properties.ResourcePropertiesList)
			for item, _ := range properties.ResourcePropertiesList {
				//spew.Dump(properties.ResourcePropertiesList[item].ResourceID)
				//spew.Dump(properties.ResourcePropertiesList[item].Property)
				for _, v := range properties.ResourcePropertiesList[item].Property {
					if !vropsdata.Extended {
						switch {
						case strings.Contains(v.Name, "parentVcenter"):
							fmt.Println("vCenter: " + v.Value)
						case strings.Contains(v.Name, "parentHost"):
							fmt.Println("ESXi: " + v.Value)
						case strings.Contains(v.Name, "config|name"):
							fmt.Println("Resource name: " + v.Value)
						case strings.Contains(v.Name, "summary|runtime|powerState"):
							fmt.Println("PowerState: " + v.Value)
						case strings.Contains(v.Name, "summary|MOID"):
							fmt.Println("MoID: " + v.Value)
						case strings.Contains(v.Name, "summary|parentCluster"):
							fmt.Println("Cluster: " + v.Value)
						case strings.Contains(v.Name, "summary|parentDatacenter"):
							fmt.Println("Datacenter: " + v.Value)
						case strings.Contains(v.Name, "config|hardware|diskSpace"):
							fmt.Println("Diskspace: " + v.Value)
						case strings.Contains(v.Name, "config|hardware|memoryKB"):
							s, _ := strconv.ParseFloat(v.Value, 64)
							s = s * 1024
							fs := HumanFileSize(s)
							fmt.Println("Memory: ", fs)
						case strings.Contains(v.Name, "config|hardware|numCoresPerSocket"):
							fmt.Println("numCoresPerSocket: " + v.Value)
						case strings.Contains(v.Name, "config|hardware|numCpu"):
							fmt.Println("numCPU: " + v.Value)
						case strings.Contains(v.Name, "config|hardware|numSockets"):
							fmt.Println("numSockets: " + v.Value)
						case strings.Contains(v.Name, "summary|datastore"):
							fmt.Println("Datastore: " + v.Value)
						case strings.Contains(v.Name, "summary|datastoreClusters"):
							fmt.Println("DatastoreCluster: " + v.Value)
						case strings.Contains(v.Name, "config|hyperThread|active"):
							fmt.Println("Hyperthreading active: " + v.Value)
						case strings.Contains(v.Name, "config|hyperThread|available"):
							fmt.Println("Hyperthreading active: " + v.Value)
						case strings.Contains(v.Name, "config|network|dnsserver"):
							fmt.Println("DNS Servers: " + v.Value)
						case strings.Contains(v.Name, "config|security|service:NTP Daemon|isRunning"):
							fmt.Println("NTP running: " + v.Value)
						case strings.Contains(v.Name, "config|security|service:NTP Daemon|policy"):
							fmt.Println("Hyperthreading active: " + v.Value)
						case strings.Contains(v.Name, "config|security|service:SSH|isRunning"):
							fmt.Println("SSH: " + v.Value)
						case strings.Contains(v.Name, "config|security|service:SSH|policy"):
							fmt.Println("SSH Policy: " + v.Value)
						case strings.Contains(v.Name, "cpu|cpuModel"):
							fmt.Println("cpuModel: " + v.Value)
						case strings.Contains(v.Name, "hardware|biosVersion"):
							fmt.Println("Bios Version: " + v.Value)
						case strings.Contains(v.Name, "hardware|cpuInfo|numCpuCores"):
							fmt.Println("numCpuCores: " + v.Value)
						case strings.Contains(v.Name, "hardware|cpuInfo|numCpuPackages"):
							fmt.Println("numCpuPackages: " + v.Value)
						case strings.Contains(v.Name, "hardware|memorySize"):
							fmt.Println("Memory: " + v.Value)
						case strings.Contains(v.Name, "hardware|powerManagementPolicy"):
							fmt.Println("Power Mgmt Policy: " + v.Value)
						case strings.Contains(v.Name, "hardware|serialNumberTag"):
							fmt.Println("SerialNumber: " + v.Value)
						case strings.Contains(v.Name, "hardware|serviceTag"):
							fmt.Println("ServiceTag: " + v.Value)
						case strings.Contains(v.Name, "hardware|vendor"):
							fmt.Println("Vendor: " + v.Value)
						case strings.Contains(v.Name, "hardware|vendorModel"):
							fmt.Println("Vendor Model: " + v.Value)
						case strings.Contains(v.Name, "net:vmk0|ip_address"):
							fmt.Println("vmk0 IP: " + v.Value)
						case strings.Contains(v.Name, "net:vmk1|ip_address"):
							fmt.Println("vmk1 IP: " + v.Value)
						case strings.Contains(v.Name, "net:vmk2|ip_address"):
							fmt.Println("vmk2 IP: " + v.Value)
						case strings.Contains(v.Name, "net|mgmt_address"):
							fmt.Println("Mgmt IP: " + v.Value)
						case strings.Contains(v.Name, "runtime|connectionState"):
							fmt.Println("Connection State: " + v.Value)
						case strings.Contains(v.Name, "runtime|maintenanceState"):
							fmt.Println("Maintenance State: " + v.Value)
						case strings.Contains(v.Name, "runtime|powerState"):
							fmt.Println("Powerstate: " + v.Value)
						case strings.Contains(v.Name, "summary|hostuuid"):
							fmt.Println("Host UUID: " + v.Value)
						case strings.Contains(v.Name, "summary|version"):
							fmt.Println("Version: " + v.Value)
						case strings.Contains(v.Name, "sys|build"):
							fmt.Println("Build: " + v.Value)

						}
					} else {
						//name := strings.ReplaceAll(v.Name, "|", " - ")

						parts := strings.Split(v.Name, "|")
						last := parts[len(parts)-1]

						fmt.Println(cases.Title(language.Und).String(last) + ": " + v.Value)
					}

				}
			}
			os.Exit(0)
		}
	}
}


