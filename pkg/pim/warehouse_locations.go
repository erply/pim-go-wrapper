package pim

import (
	"context"
	"fmt"
	"net/http"
)

type (
	WarehouseLocations service
	WarehouseLocation  struct {
		DbRecord
		Name string `json:"name"`
		AddedByChangedBy
	}
)

func (s *WarehouseLocations) Read(ctx context.Context, opts *ListOptions) (*[]WarehouseLocation, *http.Response, error) {
	u := fmt.Sprintf("warehouse/locations")
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
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
