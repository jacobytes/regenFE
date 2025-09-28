package regenfe

import (
	"context"
	"fmt"
)

type Driver struct {
	Id           *string
	FirstName    *string
	LastName     *string
	DriverNumber *string
	TLA          *string
	Color        *string
	Team         *DriverTeamRelation
	Metadata     *any
}

type DriverTeamRelation struct {
	Id   *string
	Name *string
}

type DriverService service

type ListDriverResponse struct {
	PaginationResponse
	Drivers        []*Driver `json:"drivers"`
	ChampionshipId *int      `json:"championship"`
}

func (s *DriverService) ListDrivers(ctx context.Context, params ListOptions) (*ListDriverResponse, *Response, error) {

	url := "drivers"

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	listDriverResponse := new(ListDriverResponse)

	resp, err := s.client.Do(ctx, req, listDriverResponse)

	if err != nil {
		return nil, resp, err
	}

	return listDriverResponse, resp, nil
}

func (s *DriverService) GetDriverById(ctx context.Context, teamId string) (*Driver, *Response, error) {

	url := fmt.Sprintf("drivers/%v", teamId)

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	driver := new(Driver)

	resp, err := s.client.Do(ctx, req, driver)

	if err != nil {
		return nil, resp, err
	}

	return driver, resp, nil
}
