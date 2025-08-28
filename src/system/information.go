package system

import (
	"legionlinuxtui/src/helpers"
)

type (
	systemInformation struct {
		Name        string
		BiosVersion string
		Family      string
		Vendor      string
	}
)

var (
	System systemInformation = systemInformation{
		Name:        helpers.ReadFile("/etc/hostname"),
		BiosVersion: helpers.ReadFile(DEVICEINFORMATIONPATH + "bios_version"),
		Family:      helpers.ReadFile(DEVICEINFORMATIONPATH + "product_family"),
		Vendor:      helpers.ReadFile(DEVICEINFORMATIONPATH + "sys_vendor"),
	}
)
