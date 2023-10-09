package model

import "encoding/json"

type LSBLKModel struct {
	Name        string       `json:"name"`
	FsType      string       `json:"fstype"`
	Size        uint64       `json:"size"`
	FSSize      json.Number  `json:"fssize"`
	Path        string       `json:"path"`
	Model       string       `json:"model"` // Device Identifier
	RM          bool         `json:"rm"`    // Is it a portable device?
	RO          bool         `json:"ro"`    // Is it a read-only device?
	State       string       `json:"state"`
	PhySec      int          `json:"phy-sec"` // Physical sector size
	Type        string       `json:"type"`
	Vendor      string       `json:"vendor"`  // Vendor
	Rev         string       `json:"rev"`     // Revision Version
	FSAvail     json.Number  `json:"fsavail"` // Available Space
	FSUse       string       `json:"fsuse%"`  // Used Percentage
	MountPoint  string       `json:"mountpoint"`
	Format      string       `json:"format"`
	Health      string       `json:"health"`
	HotPlug     bool         `json:"hotplug"`
	UUID        string       `json:"uuid"`
	PTUUID      string       `json:"ptuuid"`
	PartUUID    string       `json:"partuuid"`
	FSUsed      json.Number  `json:"fsused"`
	Temperature int          `json:"temperature"`
	Tran        string       `json:"tran"`
	MinIO       uint64       `json:"min-io"`
	UsedPercent float64      `json:"used_percent"`
	Serial      string       `json:"serial"`
	Children    []LSBLKModel `json:"children"`
	SubSystems  string       `json:"subsystems"`
	Label       string       `json:"label"`
	// Details Specific
	StartSector uint64 `json:"start_sector,omitempty"`
	Rota        bool   `json:"rota"` // true(hhd) false(ssd)
	DiskType    string `json:"disk_type"`
	EndSector   uint64 `json:"end_sector,omitempty"`
}

type Drive struct {
	Name           string         `json:"name"`
	Size           uint64         `json:"size"`
	Model          string         `json:"model"`
	Health         string         `json:"health"`
	Temperature    int            `json:"temperature"`
	DiskType       string         `json:"disk_type"`
	NeedFormat     bool           `json:"need_format"`
	Serial         string         `json:"serial"`
	Path           string         `json:"path"`
	ChildrenNumber int            `json:"children_number"`
	Children       []DiskChildren `json:"children"`
	Supported      bool           `json:"supported"`
}
type DiskChildren struct {
	Name      string `json:"name"`
	Size      uint64 `json:"size"`
	Format    string `json:"format"`
	Supported bool   `json:"supported"`
}

type USBDriveStatus struct {
	Name     string        `json:"name"`
	Size     uint64        `json:"size"`
	Model    string        `json:"model"`
	Avail    uint64        `json:"avail"`
	Children []USBChildren `json:"children"`
}
type USBChildren struct {
	Name       string `json:"name"`
	Size       uint64 `json:"size"`
	Avail      uint64 `json:"avail"`
	MountPoint string `json:"mount_point"`
}

type Storage struct {
	UUID        string `json:"uuid"`
	MountPoint  string `json:"mount_point"`
	Size        string `json:"size"`
	Avail       string `json:"avail"`
	Used        string `json:"used"`
	Type        string `json:"type"`
	Path        string `json:"path"`
	DriveName   string `json:"drive_name"`
	Label       string `json:"label"`
	PersistedIn string `json:"persisted_in"` // none, fstab, casaos
}
type Storages struct {
	DiskName string    `json:"disk_name"`
	Size     uint64    `json:"size"`
	Path     string    `json:"path"`
	Children []Storage `json:"children"`
	Type     string    `json:"type"`
}

type DiskStatus struct {
	Size   uint64 `json:"size"`
	Avail  uint64 `json:"avail"` // Available Space
	Health bool   `json:"health"`
	Used   uint64 `json:"used"`
}
