package drivers

type (
	DriverModuleFunction struct {
		File          string
		SysFSLocation string
		GetStatus     func(s *DriverModuleFunction) string
		UpdateStatus  func(s *DriverModuleFunction, value string)
		DriverModes   map[int]string
	}
)
