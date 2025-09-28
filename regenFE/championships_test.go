package regenfe

import (
	"context"
	"testing"
)

func TestChampionshipService_GetChampionshipById(t *testing.T) {

	client := NewClient()
	ctx := context.Background()

	championship, _, err := client.Championships.GetChampionshipById(ctx, "88a88a4b-a48d-4d06-9e52-d609bb7824a3")

	if err != nil {
		t.Errorf("ChampionshipService.GetChampionshipById returned an error: %v", err.Error())
	}

	if *championship.Id != "88a88a4b-a48d-4d06-9e52-d609bb7824a3" {
		t.Errorf("ChampionshipService.GetChampionshipById returned an invalid championship. Expected %v. Got %v", "88a88a4b-a48d-4d06-9e52-d609bb7824a3", championship.Id)
	}
}

// TODO: Look into govcr for better integration testing
func TestChampionshipService_ListChampionships(t *testing.T) {

	client := NewClient()
	ctx := context.Background()

	result, _, err := client.Championships.ListChampionships(ctx)

	if err != nil {
		t.Errorf("Got an error while listing championships: %v", err.Error())
	}

	if len(result.Championships) == 0 {
		t.Errorf("Results did not include any items")
	}

}

// TODO: Look into govcr for better integration testing
func TestChampionshipService_GetLatestChampionship(t *testing.T) {

	client := NewClient()
	ctx := context.Background()

	result, _, err := client.Championships.GetLatestChampionship(ctx)

	if err != nil {
		t.Errorf("Got an error while listing championships: %v", err.Error())
	}

	if result == nil {
		t.Errorf("Results did not include any championship")
	}

}
