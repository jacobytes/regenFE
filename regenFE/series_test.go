package regenfe

import (
	"context"
	"testing"
)

func TestSeriesService_GetSeriesById(t *testing.T) {

	client := NewClient()
	ctx := context.Background()
	expectedId := "5e960681-310e-422d-973b-52ba1b25702e"

	// act
	series, _, err := client.Series.GetSeriesById(ctx, expectedId)

	// assert
	if err != nil {
		t.Errorf("Got an error while getting a series by id: %v", err.Error())
	}

	if *series.Id != expectedId {
		t.Errorf("Expected series %v. Got: %v when getting by id", expectedId, series.Id)
	}
}

// TODO: Look into govcr for better integration testing
func TestSeriesService_ListSeries(t *testing.T) {
	client := NewClient()

	ctx := context.Background()

	result, _, err := client.Series.ListSeries(ctx)

	if err != nil {
		t.Errorf("Got an error when retrieving series: %v", err.Error())
	}

	if len(result.Series) == 0 {
		t.Errorf("Got no results when retrieving series")
	}

}
