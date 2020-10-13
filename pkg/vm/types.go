package vm

type Vm struct {
	Urn               string
	Uri               string
	Uuid              string
	Name              string
	Arch              string
	Description       string
	Group             string
	Location          string
	LocationName      string
	HostUrn           string
	Status            string
	PvDriverStatus    string
	ToolInstallStatus string
	CdRomStatus       string
	IsTemplate        bool
	IsLinkClone       bool
	IsBindingHost     bool
	CreateTime        string
	ToolsVersion      string
	HostName          string
	ClusterName       string
	HugePage          string
	Idle              int
	VmType            int
	DrStatus          int
	RpoStatus         int
	InitSyncStatus    int
}

type ListVmResponse struct {
	Total int
	Vms   []Vm
}
