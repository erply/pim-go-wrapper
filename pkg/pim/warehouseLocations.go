package pim

import (
	"context"
	"net/http"
)

type (
	WarehouseLocations service
	WarehouseLocation  struct {
		DbRecord
		Name string `json:"name,omitempty"`
		AddedByChangedBy
	}
)

func (s *WarehouseLocations) Read(ctx context.Context, opts *ListOptions) (*[]WarehouseLocation, *http.Response, error) {
	urlStr := "warehouse/locations"
	u, err := addOptions(urlStr, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	dataResp := new([]WarehouseLocation)
	resp, err := s.client.Do(ctx, req, dataResp)
	if err != nil {
		return nil, resp, err
	}

	return dataResp, resp, nil
}
