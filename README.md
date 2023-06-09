# vRealize Aria Operations CLI

This small tool allows you to search for objects in vRealize Aria Operations from command line.

# Download

Download the binary for your operating system

# Parameter

| Parameter | Description |
|---|---|
| -fqdn | FQDN or IP of your vRealize Aria Operations environment |
| -u | vRealize Aria Operations username |
| -p | vRealize Aria Operations password |
| -auth | Authentication provide for example: local |
| -search | Search string |
| -e | Extended output of metrics and properties |
| -i | insecure - no ssl certificate validation for self signed certificates |
| -d | Debug output of the CLI |

# Examples:

## Short version

```
./vrops-cli -u=admin -p=VMware123! -auth=local -fqdn=vrops.your-domain.com -search="vm-name-01"
```
### Output
```
vROPs: vrops.your-domain.com
vROPS Resource ID: 13b425da-0615-4734-a123-da852120250c
Diskspace: 16.0
Memory:  768 MB
numCoresPerSocket: 1.0
numCPU: 1.0
numSockets: 1.0
Resource name: vm-name-01
Datastore: vmware01-100
Datastore: NFS-Cluster
MoID: vm-67063
Cluster: Dell-Cluster
Datacenter: Homelab
ESXi: 192.168.0.10
vCenter: vcsa.your-domain.com
Connection State: connected
PowerState: Powered On
```

## Extended version

```
./vrops-cli -u=admin -p=VMware123! -auth=local -fqdn=vrops.your-domain.com -search="vm-name-01" -e
```
### Output
```
vROPs: vrops.your-domain.com
vROPS Resource ID: 13b425da-0615-4734-a123-da852120250c
Limit: -1.0
Reservation: 0.0
Shares: 2000.0
Createdate: 1.597736620472E12
Mem_hotadd: false
Vcpu_hotadd: false
Vcpu_hotremove: false
Guestfullname: Ubuntu Linux (64-bit)
Diskspace: 16.0
Memorykb: 786432.0
Numcorespersocket: 1.0
Numcpu: 1.0
Numsockets: 1.0
Limit: -1.0
Reservation: 0.0
Shares: 15360.0
Name: vm-name-01
Numrdms: 0.0
Numvmdks: 1.0
Numvmdksonly: 1.0
Version: vmx-14
Limit: -1.0
Reservation: 0.0
Speed: 3.504000065E9
Authenticationstatus: Failure
Discoverymethod: Guest Alias
Servicediscoverystatus: Failure
Guestosmemnotcollecting: 0.0
Host_limit: -1.0
Ip_address: 192.168.0.2
Mac_address: 00:50:56:99:1a:59
Memorycap: 786432.0
Protectiongroup: N/A
Recoveryplans: N/A
Issrmplaceholder: false
Istemplate: false
Numethernetcards: 1.0
Type: default
Datastore: vmware01-100
Datastoreclusters: NFS-Cluster
Folder: 
Fullname: Ubuntu Linux (64-bit)
Guestfamily: linuxGuest
Hostname: vm-name-01
Ipaddress: 192.168.0.2
Toolsrunningstatus: Guest Tools Running
Toolsversion: 11360
Toolsversionstatus2: Guest Tools Unmanaged
Moid: vm-67063
Parentcluster: Dell-Cluster
Parentdatacenter: Homelab
Parentfolder: vm
Parenthost: 192.168.0.10
Parentvcenter: vcsa.your-domain.com
Connectionstate: connected
Isidle: 0.0
Powerstate: Powered On
Smbiosuuid: 4219275d-66ce-1aa0-0238-713ffca11257
Tag: none
Tagjson: none
Uuid: 50198bd4-5b94-698a-3854-8383bd4c6035
Resource_kind_subtype: GENERAL
Resource_kind_type: GENERAL
Configuredgb: 16.0
Datastore: vmware01-100
Isrdm: false
Label: Hard disk 1
Provisioning_type: Thin Provision
```

