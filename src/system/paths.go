package system

const (
	BATTERYINFORMATIONPATH string = "/sys/class/power_supply/BAT0/"
	CPUINFORMATIONPATH     string = "/sys/class/hwmon/hwmon5/"
	FANINFORMATIONPATH     string = "/sys/class/hwmon/hwmon3/"
	DEVICEINFORMATIONPATH  string = "/sys/devices/virtual/dmi/id/"
	IDEASYSTEMDRIVERPATH   string = "/sys/module/legion_laptop/drivers/platform:legion/PNP0C09:00/VPC2004:00/"
	LEGIONSYSTEMDRIVERPATH string = "/sys/module/legion_laptop/drivers/platform:legion/PNP0C09:00/"
)
