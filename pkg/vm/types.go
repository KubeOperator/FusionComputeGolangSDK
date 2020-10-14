package vm

type Vm struct {
	Urn               string `json:"urn,omitempty,omitempty"`
	Uri               string `json:"uri,omitempty"`
	Uuid              string `json:"uuid,omitempty"`
	Name              string `json:"name,omitempty"`
	Arch              string `json:"arch,omitempty"`
	Description       string `json:"description,omitempty"`
	Group             string `json:"group,omitempty"`
	Location          string `json:"location,omitempty"`
	LocationName      string `json:"locationName,omitempty"`
	HostUrn           string `json:"hostUrn,omitempty"`
	Status            string `json:"status,omitempty"`
	PvDriverStatus    string `json:"pvDriverStatus,omitempty"`
	ToolInstallStatus string `json:"toolInstallStatus,omitempty"`
	CdRomStatus       string `json:"cdRomStatus,omitempty"`
	IsTemplate        bool   `json:"isTemplate,omitempty"`
	IsLinkClone       bool   `json:"isLinkClone,omitempty"`
	IsBindingHost     bool   `json:"isBindingHost,omitempty"`
	CreateTime        string `json:"createTime,omitempty"`
	ToolsVersion      string `json:"toolsVersion,omitempty"`
	HostName          string `json:"hostName,omitempty"`
	ClusterName       string `json:"clusterName,omitempty"`
	HugePage          string `json:"hugePage,omitempty"`
	Idle              int    `json:"idle,omitempty"`
	VmType            int    `json:"vmType,omitempty"`
	DrStatus          int    `json:"drStatus,omitempty"`
	RpoStatus         int    `json:"rpoStatus,omitempty"`
	InitSyncStatus    int    `json:"initSyncStatus,omitempty"`
}

type ListVmResponse struct {
	Total int  `json:"total"`
	Vms   []Vm `json:"vms"`
}

type CloneVmRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Group         string `json:"group"`
	Location      string `json:"location"`
	IsBindingHost bool   `json:"isBindingHost"`
	Config        Config `json:"vmConfig"`
}

type Config struct {
	Cpu    Cpu    `json:"cpu"`
	Memory Memory `json:"memory"`
	Disks  []Disk `json:"disks"`
	Nics   []Nic  `json:"nics"`
}
type Cpu struct {
	Quantity    int `json:"quantity"`
	Reservation int `json:"reservation"`
	Weight      int `json:"weight"`
	Limit       int `json:"limit"`
}

type Memory struct {
	QuantityMB  int `json:"quantityMb"`
	Reservation int `json:"reservation"`
	Weight      int `json:"weight"`
	Limit       int `json:"limit"`
}

type Disk struct {
	SequenceNum  int    `json:"sequenceNum"`
	QuantityGB   int    `json:"quantityGb"`
	IsDataCopy   bool   `json:"isDataCopy"`
	DatastoreUrn string `json:"datastoreUrn"`
	IsThin       bool   `json:"isThin"`
}

type Nic struct {
	Name         string `json:"name"`
	PortGroupUrn string `json:"portGroupUrn"`
	Mac          string `json:"mac"`
}

type CloneVmResponse struct {
	Urn     string `json:"urn"`
	Uri     string `json:"uri"`
	TaskUrn string `json:"taskUrn"`
	TaskUri string `json:"taskUri"`
}
