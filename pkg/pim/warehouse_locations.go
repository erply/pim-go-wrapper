package pim

import (
	"context"
	"fmt"
	"net/http"
)

type WarehouseLocationService service

type WarehouseLocations []WarehouseLocation
type WarehouseLocation struct {
	DbRecord
	Name string `json:"name"`
	AddedByChangedBy
}

// ListArtifacts lists all artifacts that belong to a repository.
func (s *WarehouseLocationService) Get(ctx context.Context, opts *ListOptions) (*WarehouseLocations, *http.Response, error) {
	u := fmt.Sprintf("warehouse/locations")
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	dataResp := new(WarehouseLocations)
	resp, err := s.client.Do(ctx, req, dataResp)
	if err != nil {
		return nil, resp, err
	}

	return dataResp, resp, nil
}
