package vm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/common"
	"path"
	"strconv"
	"strings"
)

const (
	siteMask = "<site_uri>"
	vmUrl    = "<site_uri>/vms"
)

type Manager interface {
	ListVm(isTemplate bool) ([]Vm, error)
	GetVM(vmUri string) (*Vm, error)
	CloneVm(templateUri string, request CloneVmRequest) (*CloneVmResponse, error)
	DeleteVm(vmUri string) (*DeleteVmResponse, error)
	UploadImage(vmUri string, request ImportTemplateRequest) (*ImportTemplateResponse, error)
}

func NewManager(client client.FusionComputeClient, siteUri string) Manager {
	return &manager{client: client, siteUri: siteUri}
}

type manager struct {
	client  client.FusionComputeClient
	siteUri string
}

func (m *manager) CloneVm(templateUri string, request CloneVmRequest) (*CloneVmResponse, error) {
	for n := range request.VmCustomization.NicSpecification {
		if !strings.Contains(request.VmCustomization.NicSpecification[n].Netmask, ".") {
			b, err := strconv.Atoi(request.VmCustomization.NicSpecification[0].Netmask)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("can not parse netmask: %s", err.Error()))
			}
			mask, err := parseMask(b)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("can not parse netmask: %s", err.Error()))
			}
			request.VmCustomization.NicSpecification[n].Netmask = mask
		}
	}

	vm, err := m.GetVM(templateUri)
	if err != nil {
		return nil, err
	}
	disks := vm.VmConfig.Disks
	if len(disks) > 0 {
		if request.Config.Disks[0].QuantityGB < vm.VmConfig.Disks[0].QuantityGB {
			request.Config.Disks[0].QuantityGB = vm.VmConfig.Disks[0].QuantityGB
		}
	}
	var cloneVmResponse CloneVmResponse
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().SetBody(&request).Post(path.Join(templateUri, "action", "clone"))
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		err := json.Unmarshal(resp.Body(), &cloneVmResponse)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, common.FormatHttpError(resp)
	}
	return &cloneVmResponse, nil
}

func (m *manager) ListVm(isTemplate bool) ([]Vm, error) {
	var vms []Vm
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	request := api.R()
	if isTemplate {
		request.SetQueryParam("isTemplate", "true")
	}
	resp, err := request.Get(strings.Replace(vmUrl, siteMask, m.siteUri, -1))
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		var listVmResponse ListVmResponse
		err := json.Unmarshal(resp.Body(), &listVmResponse)
		if err != nil {
			return nil, err
		}
		vms = listVmResponse.Vms

	} else {
		return nil, common.FormatHttpError(resp)
	}
	return vms, nil
}

func (m *manager) DeleteVm(vmUri string) (*DeleteVmResponse, error) {
	var deleteVmResponse DeleteVmResponse
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().Delete(vmUri)
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		err := json.Unmarshal(resp.Body(), &deleteVmResponse)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, common.FormatHttpError(resp)
	}
	return &deleteVmResponse, nil
}

func (m *manager) GetVM(vmUri string) (*Vm, error) {
	var item Vm
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().Get(vmUri)

	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		err := json.Unmarshal(resp.Body(), &item)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, common.FormatHttpError(resp)
	}
	return &item, nil
}

func (m *manager) UploadImage(vmUri string, request ImportTemplateRequest) (*ImportTemplateResponse, error) {
	var res ImportTemplateResponse
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().SetBody(&request).Post(path.Join(vmUri, "action", "import"))
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		err := json.Unmarshal(resp.Body(), &res)
		if err != nil {
			return nil, err
		}
		return &res, nil
	} else {
		return nil, common.FormatHttpError(resp)
	}
}

func parseMask(num int) (mask string, err error) {
	var buff bytes.Buffer
	for i := 0; i < int(num); i++ {
		buff.WriteString("1")
	}
	for i := num; i < 32; i++ {
		buff.WriteString("0")
	}
	masker := buff.String()
	a, _ := strconv.ParseUint(masker[:8], 2, 64)
	b, _ := strconv.ParseUint(masker[8:16], 2, 64)
	c, _ := strconv.ParseUint(masker[16:24], 2, 64)
	d, _ := strconv.ParseUint(masker[24:32], 2, 64)
	resultMask := fmt.Sprintf("%v.%v.%v.%v", a, b, c, d)
	return resultMask, nil
}
