package regenfe

import (
	"context"
	"testing"
)

func TestTeamService_GetTeamById(t *testing.T) {

	// arrange
	client := NewClient()
	ctx := context.Background()
	expected := makeExpectedTeam()

	// act
	team, _, err := client.Teams.GetTeamByID(ctx, *expected.Id)

	// assert
	if err != nil {
		t.Errorf("TeamService.GetTeamById returned an error: %v", err.Error())
	}

	if *team.Id != *expected.Id {
		t.Errorf("TeamService.GetTeamById returned an invalid team. Expected %v. Got: %v", expected.Id, team.Id)
	}
}

// TODO: Look into govcr for better integration testing
func TestTeamService_ListTeams(t *testing.T) {
	client := NewClient()

	ctx := context.Background()
	options := ListOptions{
		Page: 1,
	}

	result, _, err := client.Teams.ListTeams(ctx, options)

	if err != nil {
		t.Errorf("TeamService.ListTeams returned an error: %v", err.Error())
	}

	if len(result.Teams) != 20 {
		t.Errorf("TeamService.ListTeams results does not match default of 20")
	}

	if result.PageInfo == nil {
		t.Errorf("TeamService.ListTeams did not include pagination info")
	}

}

func makeExpectedTeam() Team {

	// Only expect data that doesn't change every race
	t := Team{
		Id:           Ptr("05dab754-2899-411b-9c4e-72311a36cc9c"),
		Name:         Ptr("JAGUAR TCS RACING"),
		TLA:          Ptr("JAGUAR"),
		CountryCode:  Ptr("GB"),
		City:         Ptr("Coventry, UK"),
		Manufacturer: Ptr("JAGUAR"),
		Powertrain:   Ptr("Jaguar I-Type 7"),
		Principal:    Ptr("James Barclay"),
	}

	return t
}
