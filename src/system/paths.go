package system

const (
	BATTERYINFORMATIONPATH string = "/sys/class/power_supply/BAT0/"
	// FIX: use another dir instead of hwmon or find paths at runtime; hwmon allocates dirs dynamically
	CPUINFORMATIONPATH    string = "/sys/class/hwmon/hwmon5/"
	FANINFORMATIONPATH    string = "/sys/class/hwmon/hwmon3/"
	DEVICEINFORMATIONPATH string = "/sys/devices/virtual/dmi/id/"
	// IDEASYSTEMDRIVERPATH -> sysfs ideapad module drivers
	IDEASYSTEMDRIVERPATH string = "/sys/module/legion_laptop/drivers/platform:legion/PNP0C09:00/VPC2004:00/"
	// LEGIONSYSTEMDRIVERPATH -> sysfs legion linux module drivers
	LEGIONSYSTEMDRIVERPATH string = "/sys/module/legion_laptop/drivers/platform:legion/PNP0C09:00/"
)
