package tests

import (
	"fmt"
	"legionlinuxtui/src/system"
	"testing"
)

func testDriverFunctionGetStatus(t *testing.T, driver *system.DriverModuleFunction) {
	if driver.GetDriverFunctionState() == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetDriverFunctionState())
}

func testDriverFunctionToggle(driver *system.DriverModuleFunction) {
	fmt.Printf("Toggling %v\n", driver.File)
	driver.ToggleDriverState()
}

func TestConservationMode(t *testing.T) {
	driver := &system.ConservationMode

	testDriverFunctionGetStatus(t, driver)

	testDriverFunctionToggle(driver)

	testDriverFunctionGetStatus(t, driver)
}

func TestFNLock(t *testing.T) {
	driver := &system.FNLock

	testDriverFunctionGetStatus(t, driver)

	testDriverFunctionToggle(driver)

	testDriverFunctionGetStatus(t, driver)
}

func TestPowerMode(t *testing.T) {
	driver := &system.PowerMode

	testDriverFunctionGetStatus(t, driver)

	testDriverFunctionToggle(driver)

	testDriverFunctionGetStatus(t, driver)
}

func TestRapidCharge(t *testing.T) {
	driver := &system.RapidCharge

	testDriverFunctionGetStatus(t, driver)

	testDriverFunctionToggle(driver)

	testDriverFunctionGetStatus(t, driver)
}

func TestTouchPad(t *testing.T) {
	driver := &system.TouchPad

	testDriverFunctionGetStatus(t, driver)

	testDriverFunctionToggle(driver)

	testDriverFunctionGetStatus(t, driver)
}

func TestUSBCharging(t *testing.T) {
	driver := &system.USBCharging

	testDriverFunctionGetStatus(t, driver)

	testDriverFunctionToggle(driver)

	testDriverFunctionGetStatus(t, driver)
}

func TestWinKey(t *testing.T) {
	driver := &system.WinKey

	testDriverFunctionGetStatus(t, driver)

	testDriverFunctionToggle(driver)

	testDriverFunctionGetStatus(t, driver)
}
