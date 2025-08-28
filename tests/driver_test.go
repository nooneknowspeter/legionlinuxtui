package tests

import (
	"fmt"
	"legionlinuxtui/src/drivers"
	"testing"
)

func TestCameraPower(t *testing.T) {
	driver := drivers.CameraPower

	if driver.GetStatus(&driver) == "" {
		t.Errorf("%v failed to retrieve current %v status", Fail, driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("Toggling %v\n", driver.File)
	driver.ToggleDriverMode()

	if driver.GetStatus(&driver) == "" {
		t.Errorf("%v failed to retrieve current %v status", Fail, driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("%v ", Success)
}

func TestConservationMode(t *testing.T) {
	driver := drivers.ConservationMode

	if driver.GetStatus(&driver) == "" {
		t.Errorf("%v failed to retrieve current %v status", Fail, driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("Toggling %v\n", driver.File)
	driver.ToggleDriverMode()

	if driver.GetStatus(&driver) == "" {
		t.Errorf("%v failed to retrieve current %v status", Fail, driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("%v ", Success)
}

func TestFNLock(t *testing.T) {
	driver := drivers.FNLock

	if driver.GetStatus(&driver) == "" {
		t.Errorf("%v failed to retrieve current %v status", Fail, driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("Toggling %v\n", driver.File)
	driver.ToggleDriverMode()

	if driver.GetStatus(&driver) == "" {
		t.Errorf("%v failed to retrieve current %v status", Fail, driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("%v ", Success)
}

func TestPowerMode(t *testing.T) {
	driver := drivers.PowerMode

	if driver.GetStatus(&driver) == "" {
		t.Errorf("%v failed to retrieve current %v status", Fail, driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("Toggling %v\n", driver.File)
	driver.ToggleDriverMode()

	if driver.GetStatus(&driver) == "" {
		t.Errorf("%v failed to retrieve current %v status", Fail, driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("%v ", Success)
}

func TestRapidCharge(t *testing.T) {
	driver := drivers.RapidCharge

	if driver.GetStatus(&driver) == "" {
		t.Errorf("%v failed to retrieve current %v status", Fail, driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("Toggling %v\n", driver.File)
	driver.ToggleDriverMode()

	if driver.GetStatus(&driver) == "" {
		t.Errorf("%v failed to retrieve current %v status", Fail, driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("%v ", Success)
}

func TestTouchPad(t *testing.T) {
	driver := drivers.TouchPad

	if driver.GetStatus(&driver) == "" {
		t.Errorf("%v failed to retrieve current %v status", Fail, driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("Toggling %v\n", driver.File)
	driver.ToggleDriverMode()

	if driver.GetStatus(&driver) == "" {
		t.Errorf("%v failed to retrieve current %v status", Fail, driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("%v ", Success)
}

func TestUSBCharging(t *testing.T) {
	driver := drivers.USBCharging

	if driver.GetStatus(&driver) == "" {
		t.Errorf("%v failed to retrieve current %v status", Fail, driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("Toggling %v\n", driver.File)
	driver.ToggleDriverMode()

	if driver.GetStatus(&driver) == "" {
		t.Errorf("%v failed to retrieve current %v status", Fail, driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("%v ", Success)
}
