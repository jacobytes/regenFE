package regenfe

import (
	"context"
	"fmt"
)

type Championship struct {
	Id                *string
	Name              *string
	Status            *string
	LastFinishedRound *int
	Series            *ChampionshipSeriesRelation
}

type ChampionshipSeriesRelation struct {
	Id   *string
	Name *string
}

type ChampionshipService service

type ListChampionshipResponse struct {
	Championships []*Championship `json:"championships"`
}

func (s *ChampionshipService) ListChampionships(ctx context.Context) (*ListChampionshipResponse, *Response, error) {

	url := "championships"

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	listChampionshipResponse := new(ListChampionshipResponse)

	resp, err := s.client.Do(ctx, req, listChampionshipResponse)

	if err != nil {
		return nil, resp, err
	}

	return listChampionshipResponse, resp, nil
}

func (s *ChampionshipService) GetChampionshipById(ctx context.Context, championshipId string) (*Championship, *Response, error) {

	url := fmt.Sprintf("championships/%v", championshipId)

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	championship := new(Championship)

	resp, err := s.client.Do(ctx, req, championship)

	if err != nil {
		return nil, resp, err
	}

	return championship, resp, nil
}

func (s *ChampionshipService) GetLatestChampionship(ctx context.Context) (*Championship, *Response, error) {

	url := "championships/latest"

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	championship := new(Championship)

	resp, err := s.client.Do(ctx, req, championship)

	if err != nil {
		return nil, resp, err
	}

	return championship, resp, nil
}
