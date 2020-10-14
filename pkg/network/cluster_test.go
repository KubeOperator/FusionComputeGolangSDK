package network

import (
	"fmt"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/site"
	"log"
	"testing"
)

func TestManager_List(t *testing.T) {
	c := client.NewFusionComputeClient("https://100.199.16.208:7443", "fit2cloud", "Huawei@1234")
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
		cs, err := cm.ListDVSwitch()
		if err != nil {
			log.Fatal(err)
		}
		for _, c := range cs {
			pg, err := cm.ListPortGroup(c.Uri)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(pg)
		}
		fmt.Println(cs)
	}
}
