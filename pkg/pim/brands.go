package pim

import (
	"context"
	"net/http"
)

type (
	Brands service
	Brand  struct {
		ID   int64  `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		AddedByChangedBy
	}
)

func (s *Brands) Read(ctx context.Context, opts *ListOptions) (*[]Brand, *http.Response, error) {
	urlStr := "brand"
	u, err := addOptions(urlStr, opts)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	dataResp := new([]Brand)
	resp, err := s.client.Do(ctx, req, dataResp)
	return dataResp, resp, err
}

type (
	BulkBrandsResponse struct {
		Results []BulkBrandsResponseItem `json:"results"`
	}
	BulkBrandsResponseItem struct {
		//id of the response, if requested 3 read requests each ID represents 1 response item
		ResultID int     `json:"resultId"`
		Records  []Brand `json:"brands"`
		//in case of error
		MessageResponse
	}
)

func (s *Brands) ReadBulk(ctx context.Context, requests []ListOptions) (*BulkBrandsResponse, *http.Response, error) {
	u := "brand/bulk/get"

	req, err := s.client.NewRequest(http.MethodPost, u, BulkReadRequest{Requests: requests})
	if err != nil {
		return nil, nil, err
	}

	res := new(BulkBrandsResponse)
	resp, err := s.client.Do(ctx, req, res)
	return res, resp, err
}

func (s *Brands) CreateBulk(ctx context.Context, brands []Brand) (*BulkResponseWithResults, *http.Response, error) {
	u := "brand/bulk"

	type BulkRequest struct {
		Requests []Brand `json:"requests"`
	}
	req, err := s.client.NewRequest(http.MethodPost, u, BulkRequest{Requests: brands})
	if err != nil {
		return nil, nil, err
	}

	res := new(BulkResponseWithResults)
	resp, err := s.client.Do(ctx, req, res)
	return res, resp, err
}
