package pim

import (
	"context"
	"net/http"
)

type (
	Families      service
	FamilyRequest struct {
		//Names translations
		TranslatableNameJSON
		//Archived ..
		Archived bool `json:"archived" example:"false"`
	}
	FamilyResponse struct {
		ID int64 `json:"id"`
		TranslatableNameJSON
		Archived int `json:"archived"`
		AddedByIDChangedByID
	}

	BulkReadFamiliesResponse struct {
		Results []BulkReadFamiliesResponseItem `json:"results"`
	}
	BulkReadFamiliesResponseItem struct {
		//id of the response, if requested 3 read requests each ID represents 1 response item
		ResultID int              `json:"resultId"`
		Records  []FamilyResponse `json:"families"`
		MessageResponse
		//total number of records (ignores skip & take parameters)
		TotalCount int `json:"totalCount"`
	}

	FamilyListOptions struct {
		ListOptions
		//WithArchived is a boolean parameter that tells the API to fetch the archived items as well.
		WithArchived bool `json:"withArchived,omitempty"`
	}
)

func (s *Families) ReadBulk(ctx context.Context, requests []FamilyListOptions) (*BulkReadFamiliesResponse, *http.Response, error) {
	u := "product/family/bulk/get"

	type BulkReadRequest struct {
		Requests []FamilyListOptions `json:"requests"`
	}
	req, err := s.client.NewRequest(http.MethodPost, u, BulkReadRequest{Requests: requests})
	if err != nil {
		return nil, nil, err
	}

	res := new(BulkReadFamiliesResponse)
	resp, err := s.client.Do(ctx, req, res)
	return res, resp, err
}

func (s *Families) CreateBulk(ctx context.Context, families []FamilyRequest) (*BulkResponseWithResults, *http.Response, error) {
	u := "product/family/bulk"

	type BulkRequest struct {
		Requests []FamilyRequest `json:"requests"`
	}
	req, err := s.client.NewRequest(http.MethodPost, u, BulkRequest{Requests: families})
	if err != nil {
		return nil, nil, err
	}

	res := new(BulkResponseWithResults)
	resp, err := s.client.Do(ctx, req, res)
	return res, resp, err
}
