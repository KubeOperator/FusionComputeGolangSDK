package site

import (
	"fmt"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
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
	m := NewManager(c)
	ss, err := m.List()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ss)
}
