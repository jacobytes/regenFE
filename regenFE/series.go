package regenfe

import (
	"context"
	"fmt"
)

type Series struct {
	Id   *string
	Name *string
	Type *string
}

type SeriesService service

type ListSeriesResponse struct {
	Series []*Series `json:"series"`
}

func (s *SeriesService) ListSeries(ctx context.Context) (*ListSeriesResponse, *Response, error) {

	url := "series"

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	listSeriesResponse := new(ListSeriesResponse)

	resp, err := s.client.Do(ctx, req, listSeriesResponse)

	if err != nil {
		return nil, resp, err
	}

	return listSeriesResponse, resp, nil
}

func (s *SeriesService) GetSeriesById(ctx context.Context, seriesId string) (*Series, *Response, error) {

	url := fmt.Sprintf("series/%v", seriesId)

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	series := new(Series)

	resp, err := s.client.Do(ctx, req, series)

	if err != nil {
		return nil, resp, err
	}

	return series, resp, nil
}
