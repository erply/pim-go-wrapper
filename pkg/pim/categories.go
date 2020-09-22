package pim

import (
	"context"
	"net/http"
)

type (
	Categories      service
	CategoryRequest struct {
		//ParentID Refers to the parent category — product categories are hierarchical.
		ParentID int `json:"main_id"`
		//Translatable name
		TranslatableNameJSON
	}

	CategoryResponse struct {
		ID       int `json:"id"`
		ParentID int `json:"parent_id"`
		Order    int `json:"order"`
		TranslatableNameJSON
		AddedByChangedBy
	}
)

type (
	BulkCategoriesResponse struct {
		Results []BulkCategoriesResponseItem `json:"results"`
	}
	BulkCategoriesResponseItem struct {
		//id of the response, if requested 3 read requests each ID represents 1 response item
		ResultID int                `json:"resultId"`
		Records  []CategoryResponse `json:"categories"`
		MessageResponse
		//total number of records (ignores skip & take parameters)
		TotalCount int `json:"totalCount"`
	}
)

func (s *Categories) ReadBulk(ctx context.Context, requests []ListOptions) (*BulkCategoriesResponse, *http.Response, error) {
	u := "product/category/bulk/get"

	type BulkReadRequest struct {
		Requests []ListOptions `json:"requests"`
	}
	req, err := s.client.NewRequest(http.MethodPost, u, BulkReadRequest{Requests: requests})
	if err != nil {
		return nil, nil, err
	}

	res := new(BulkCategoriesResponse)
	resp, err := s.client.Do(ctx, req, res)
	return res, resp, err
}

func (s *Categories) CreateBulk(ctx context.Context, categories []CategoryRequest) (*BulkResponseWithResults, *http.Response, error) {
	u := "product/category/bulk"

	type BulkRequest struct {
		Requests []CategoryRequest `json:"requests"`
	}
	req, err := s.client.NewRequest(http.MethodPost, u, BulkRequest{Requests: categories})
	if err != nil {
		return nil, nil, err
	}

	res := new(BulkResponseWithResults)
	resp, err := s.client.Do(ctx, req, res)
	return res, resp, err
}
