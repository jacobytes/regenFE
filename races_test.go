package regenfe

import (
	"context"
	"testing"
)

func TestRaceService_GetRaceById(t *testing.T) {

	// arrange
	raceId := "c1dd1f8a-5112-4864-8d2d-cfcc8951d197"
	client := NewClient()
	ctx := context.Background()

	// act
	race, _, err := client.Races.GetRaceById(ctx, raceId)

	// assert
	if err != nil {
		t.Errorf("RaceService.GetRaceById returned an error: %v", err.Error())
	}

	if *race.Id != raceId {
		t.Errorf("RaceService.GetRaceById returned an invalid Race. Expected %v, got %v", "c1dd1f8a-5112-4864-8d2d-cfcc8951d197", race.Id)
	}

}

// TODO: Look into govcr for better integration testing
func TestRaceService_ListRaces(t *testing.T) {

	// arrange
	client := NewClient()
	ctx := context.Background()
	options := ListOptions{
		Page: 1,
	}

	results, _, err := client.Races.ListRaces(ctx, options)

	if err != nil {
		t.Errorf("RaceService.ListRaces returned an error: %v", err.Error())
	}

	if results.PageInfo == nil {
		t.Errorf("RaceService.ListRaces did not include pagination info")
	}

}

func TestRaceService_ListSessions(t *testing.T) {

	raceId := "c1dd1f8a-5112-4864-8d2d-cfcc8951d197"
	client := NewClient()
	ctx := context.Background()
	options := ListOptions{
		Page: 1,
	}

	results, _, err := client.Races.ListSessions(ctx, raceId, options)

	if err != nil {
		t.Errorf("RaceService.ListSessions returned an error: %v", err.Error())
	}

	if results.PageInfo == nil {
		t.Errorf("RaceService.ListSessions did not include pagination info")
	}
}

func TestRaceService_GetSessionById(t *testing.T) {

	raceId := "c1dd1f8a-5112-4864-8d2d-cfcc8951d197"
	sessionId := "75e32986-089a-47ca-b509-7498798aceeb"
	client := NewClient()
	ctx := context.Background()

	result, _, err := client.Races.GetSessionById(ctx, raceId, sessionId)

	if err != nil {
		t.Errorf("RaceService.GetSessionById returned an error: %v", err.Error())
	}

	if *result.Id != sessionId {
		t.Errorf("RaceService.GetSessionById returned an invalid session. Expected %v, got %v", sessionId, result.Id)
	}
}
