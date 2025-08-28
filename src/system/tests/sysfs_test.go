package tests

import (
	"fmt"
	"legionlinuxtui/src/system"
	"testing"
)

func TestCameraPower(t *testing.T) {
	driver := system.CameraPower

	if driver.GetStatus(&driver) == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("Toggling %v\n", driver.File)
	driver.ToggleDriverMode()

	if driver.GetStatus(&driver) == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))
}

func TestConservationMode(t *testing.T) {
	driver := system.ConservationMode

	if driver.GetStatus(&driver) == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("Toggling %v\n", driver.File)
	driver.ToggleDriverMode()

	if driver.GetStatus(&driver) == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))
}

func TestFNLock(t *testing.T) {
	driver := system.FNLock

	if driver.GetStatus(&driver) == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("Toggling %v\n", driver.File)
	driver.ToggleDriverMode()

	if driver.GetStatus(&driver) == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))
}

func TestPowerMode(t *testing.T) {
	driver := system.PowerMode

	if driver.GetStatus(&driver) == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("Toggling %v\n", driver.File)
	driver.ToggleDriverMode()

	if driver.GetStatus(&driver) == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))
}

func TestRapidCharge(t *testing.T) {
	driver := system.RapidCharge

	if driver.GetStatus(&driver) == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("Toggling %v\n", driver.File)
	driver.ToggleDriverMode()

	if driver.GetStatus(&driver) == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))
}

func TestTouchPad(t *testing.T) {
	driver := system.TouchPad

	if driver.GetStatus(&driver) == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("Toggling %v\n", driver.File)
	driver.ToggleDriverMode()

	if driver.GetStatus(&driver) == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))
}

func TestUSBCharging(t *testing.T) {
	driver := system.USBCharging

	if driver.GetStatus(&driver) == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("Toggling %v\n", driver.File)
	driver.ToggleDriverMode()

	if driver.GetStatus(&driver) == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))
}

func TestWinKey(t *testing.T) {
	driver := system.WinKey

	if driver.GetStatus(&driver) == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))

	fmt.Printf("Toggling %v\n", driver.File)
	driver.ToggleDriverMode()

	if driver.GetStatus(&driver) == "" {
		t.Errorf("  failed to retrieve current %v status", driver.File)
		return
	}
	fmt.Printf("Current %v Status: %v\n", driver.File, driver.GetStatus(&driver))
}
