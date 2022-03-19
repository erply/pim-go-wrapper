package pim

import (
	"context"
	"net/http"
)

type (
	ProductUnits service
	ProductUnit  struct {
		AddedByChangedBy
		ID int `json:"id"`
		TranslatableNameJSON
		Order int `json:"order"`
	}
)

func (s *ProductUnits) Read(ctx context.Context, opts *ListOptions) (*[]ProductUnit, *http.Response, error) {
	urlStr := "product/unit"
	u, err := addOptions(urlStr, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	dataResp := new([]ProductUnit)
	resp, err := s.client.Do(ctx, req, dataResp)
	return dataResp, resp, err
}
