package regenfe

import (
	"context"
	"fmt"
)

type Team struct {
	Id            *string `json:"id,omitempty"`
	Name          *string `json:"name,omitempty"`
	TLA           *string
	Color         *string
	Wins          *int
	Podiums       *int
	Starts        *int
	Points        *int
	CountryCode   *string
	City          *string
	ReserveDriver *string
	Manufacturer  *string
	Powertrain    *string
	Principal     *string
	Standing      *int
	Drivers       []Driver
	RelatedTeams  []TeamRelation
	Metadata      *any
	Championship  *any
}

type TeamOption func(*Team)

// func NewTeam(opts ...TeamOption) *Team {
// 	t := &Team{}
// 	for _, opt := range opts {
// 		opt(t)
// 	}
// 	return t
// }

// func WithId(id string) TeamOption {
// 	return func(t *Team) { t.Id = &id }
// }

// func WithName(name string) TeamOption {
// 	return func(t *Team) { t.Name = &name }
// }

// func WithTla(tla string) TeamOption {
// 	return func(t *Team) { t.TLA = &tla }
// }

// func WithColor(color string) TeamOption {
// 	return func(t *Team) { t.Color = &color }
// }

// func WithPodiums(podiums int) TeamOption {
// 	return func(t *Team) { t.Wins = &podiums }
// }

// func WithStarts(starts int) TeamOption {
// 	return func(t *Team) { t.Starts = &starts }
// }

// func WithPoints(points int) TeamOption {
// 	return func(t *Team) { t.Points = &points }
// }

// func WithCountryCode(code string) TeamOption {
// 	return func(t *Team) { t.CountryCode = &code }
// }

// func WithReserveDriver(id string) TeamOption {
// 	return func(t *Team) { t.ReserveDriver = &id }
// }

// func WithManufacturer(id string) TeamOption {
// 	return func(t *Team) { t.Manufacturer = &id }
// }

// func WithPowertrain(name string) TeamOption {
// 	return func(t *Team) { t.Powertrain = &name }
// }

// func WithPrincipal(name string) TeamOption {
// 	return func(t *Team) { t.Principal = &name }
// }

// func WithStanding(standing int) TeamOption {
// 	return func(t *Team) { t.Standing = &standing }
// }

// func WithDrivers(drivers []Driver) TeamOption {
// 	return func(t *Team) { t.Drivers = drivers }
// }

// func WithRelatedTeams(teams []TeamRelation) TeamOption {
// 	return func(t *Team) { t.RelatedTeams = teams }
// }

// func WithChampionShip(championship any) TeamOption {
// 	return func(t *Team) { t.Championship = &championship }
// }

type TeamService service

type ListTeamsResponse struct {
	PaginationResponse
	Teams        []*Team `json:"teams"`
	Championship *int    `json:"championship"`
}

func (s *TeamService) ListTeams(ctx context.Context, params ListOptions) (*ListTeamsResponse, *Response, error) {
	url := "teams"

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	listTeamResponse := new(ListTeamsResponse)

	resp, err := s.client.Do(ctx, req, listTeamResponse)

	if err != nil {
		return nil, resp, err
	}

	return listTeamResponse, resp, nil
}

func (s *TeamService) GetTeamByID(ctx context.Context, teamID string) (*Team, *Response, error) {

	url := fmt.Sprintf("teams/%v", teamID)

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	team := new(Team)

	resp, err := s.client.Do(ctx, req, team)

	if err != nil {
		return nil, resp, err
	}

	return team, resp, nil
}

type TeamRelation struct {
	Id           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Championship any
}
