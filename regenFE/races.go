package regenfe

import (
	"context"
	"fmt"
)

type Race struct {
	Id                *string
	Name              *string
	Sequence          *int
	Country           *string
	City              *string
	Date              *string
	HasSessionResults *bool
	HasRaceResults    *bool
}

type Session struct {
	Id                    *string
	Name                  *string
	Date                  *string
	StartTime             *string
	FinishTime            *string
	ContingencyStartTime  *string
	ContingencyFinishTime *string
	OffsetGMT             *string
	HasResults            *bool
	StartingGridAvailable *bool
	SessionLiveStatus     *string
	MetaData              any
}

type RaceService service

func (s *RaceService) GetRaceById(ctx context.Context, raceId string) (*Race, *Response, error) {

	url := fmt.Sprintf("races/%v", raceId)

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	race := new(Race)

	resp, err := s.client.Do(ctx, req, race)

	if err != nil {
		return nil, resp, err
	}

	return race, resp, nil
}

type ListRaceResponse struct {
	PaginationResponse
	Races []*Race `json:"races"`
}

func (s *RaceService) ListRaces(ctx context.Context, params ListOptions) (*ListRaceResponse, *Response, error) {

	url := "races"

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	listRaceResponse := new(ListRaceResponse)

	resp, err := s.client.Do(ctx, req, listRaceResponse)

	if err != nil {
		return nil, resp, err
	}

	return listRaceResponse, resp, nil
}

type ListSessionsResponse struct {
	PaginationResponse
	Sessions []*Session `json:"sessions"`
}

func (s *RaceService) ListSessions(ctx context.Context, raceId string, params ListOptions) (*ListSessionsResponse, *Response, error) {

	url := fmt.Sprintf("races/%v/sessions", raceId)

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	listSessionResponse := new(ListSessionsResponse)

	resp, err := s.client.Do(ctx, req, listSessionResponse)

	if err != nil {
		return nil, resp, err
	}

	return listSessionResponse, resp, nil
}

func (s *RaceService) GetSessionById(ctx context.Context, raceId string, sessionId string) (*Session, *Response, error) {

	url := fmt.Sprintf("races/%v/sessions/%v", raceId, sessionId)

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	session := new(Session)

	resp, err := s.client.Do(ctx, req, session)

	if err != nil {
		return nil, resp, err
	}

	return session, resp, nil
}
