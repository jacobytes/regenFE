package regenfe

import (
	"context"
	"testing"
)

func TestDriverService_GetDriverById(t *testing.T) {

	// arrange
	client := NewClient()
	ctx := context.Background()

	driver, _, err := client.Drivers.GetDriverById(ctx, "2543b91c-2543-4835-b9fc-17b5fbc79cb1")

	if err != nil {
		t.Errorf("DriverService.GetDriverById() returned an error: %v", err.Error())
	}

	if *driver.Id != "2543b91c-2543-4835-b9fc-17b5fbc79cb1" {
		t.Errorf("DriverService.GetDriverById() returned an invalid driver. Expected %v, got %v", "2543b91c-2543-4835-b9fc-17b5fbc79cb1", driver.Id)
	}

}

// TODO: Look into govcr for better integration testing
func TestDriverService_ListDrivers(t *testing.T) {

	// arrange
	client := NewClient()
	ctx := context.Background()
	options := ListOptions{
		Page: 1,
	}

	results, _, err := client.Drivers.ListDrivers(ctx, options)

	if err != nil {
		t.Errorf("DriverService.ListDrivers returned an error: %v", err.Error())
	}

	if len(results.Drivers) != 20 {
		t.Errorf("DriverService.ListDrivers results does not match default of 20")
	}

	if results.PageInfo == nil {
		t.Errorf("DriverService.ListDrivers did not include pagination info")
	}

}
