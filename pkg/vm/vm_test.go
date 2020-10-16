package vm

import (
	"fmt"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/site"
	"log"
	"testing"
)

func TestManager_List(t *testing.T) {
	c := client.NewFusionComputeClient("https://100.199.16.208:7443", "kubeoperator", "Calong@2015")
	err := c.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer c.DisConnect()

	sm := site.NewManager(c)
	ss, err := sm.ListSite()
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range ss {
		cm := NewManager(c, s.Uri)
		cs, err := cm.ListVm(true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(cs)
	}
}

func TestManager_CloneVm(t *testing.T) {
	c := client.NewFusionComputeClient("https://100.199.16.208:7443", "kubeoperator", "Calong@2015")
	err := c.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer c.DisConnect()
	m := NewManager(c, "/service/sites/43BC08E8")
	task, err := m.CloneVm("/service/sites/43BC08E8/vms/i-00000034", CloneVmRequest{
		Name:          "test-1",
		Description:   "test create vm",
		Location:      "urn:sites:43BC08E8:clusters:117",
		IsBindingHost: false,
		Config: Config{
			Cpu: Cpu{
				Quantity:    2,
				Reservation: 0,
			},
			Memory: Memory{
				QuantityMB:  2048,
				Reservation: 2048,
			},
			Disks: []Disk{
				{
					SequenceNum:  1,
					QuantityGB:   50,
					IsDataCopy:   true,
					DatastoreUrn: "urn:sites:43BC08E8:datastores:41",
					IsThin:       true,
				},
			},
			Nics: []Nic{
				{
					Name:         "vmnic1",
					PortGroupUrn: "urn:sites:43BC08E8:dvswitchs:1:portgroups:1",
				},
			},
		},
		VmCustomization: Customization{
			OsType:             "Linux",
			Hostname:           "test-1",
			IsUpdateVmPassword: false,
			NicSpecification: []NicSpecification{
				{
					SequenceNum: 1,
					Ip:          "100.199.10.88",
					Netmask:     "255.255.255.0",
					Gateway:     "100.199.10.1",
					Setdns:      "114.114.114.114",
					Adddns:      "8.8.8.8",
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(task)

}
