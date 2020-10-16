package pim

import (
	"context"
	"fmt"
	"net/http"
)

type (
	ExtraFields service

	ExtraField struct {
		ID int64 `json:"id,omitempty"`
		ExtraFieldRequest
		TranslatableNameJSON
	}

	ExtraFieldRequest struct {
		//code of the extra field
		Code string `json:"code" example:"code"`
		//Is the field active?
		IsActive bool `json:"is_active" example:"true"`
		TranslatableNameJSON
	}

	BulkReadExtraFieldResponse struct {
		Results []BulkReadExtraFieldResponseItem `json:"results"`
	}

	BulkReadExtraFieldResponseItem struct {
		//id of the response item
		ResultID int
		//in case of error
		MessageResponse

		//total number of records (ignores skip & take parameters)
		TotalCount int `json:"totalCount"`
		//resulting records
		ExtraFields []ExtraField `json:"extraFields"`
	}
)

func (s *ExtraFields) Read(ctx context.Context, fieldNumber uint, opts *ListOptions) (*[]ExtraField, *http.Response, error) {
	if err := s.isFieldNumberValid(fieldNumber); err != nil {
		return nil, nil, err
	}
	urlStr := fmt.Sprintf("product/extra/field-%d", fieldNumber)
	u, err := addOptions(urlStr, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	dataResp := new([]ExtraField)
	resp, err := s.client.Do(ctx, req, dataResp)
	return dataResp, resp, err
}

func (s *ExtraFields) Create(ctx context.Context, fieldNumber uint, f *ExtraFieldRequest) (*IDResponse, *http.Response, error) {
	if err := s.isFieldNumberValid(fieldNumber); err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("product/extra/field-%d", fieldNumber)

	req, err := s.client.NewRequest(http.MethodPost, u, f)
	if err != nil {
		return nil, nil, err
	}

	id := new(IDResponse)
	resp, err := s.client.Do(ctx, req, id)
	return id, resp, err
}

func (s *ExtraFields) CreateBulk(ctx context.Context, fieldNumber uint, fs []ExtraFieldRequest) (*BulkResponseWithResults, *http.Response, error) {
	if err := s.isFieldNumberValid(fieldNumber); err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("product/extra/field-%d/bulk", fieldNumber)

	type BulkExtraFieldRequest struct {
		Requests []ExtraFieldRequest `json:"requests"`
	}
	req, err := s.client.NewRequest(http.MethodPost, u, BulkExtraFieldRequest{Requests: fs})
	if err != nil {
		return nil, nil, err
	}

	res := new(BulkResponseWithResults)
	resp, err := s.client.Do(ctx, req, res)
	return res, resp, err
}

func (s *ExtraFields) ReadBulk(ctx context.Context, fieldNumber uint, requests []ListOptions) (*BulkReadExtraFieldResponse, *http.Response, error) {
	if err := s.isFieldNumberValid(fieldNumber); err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("product/extra/field-%d/bulk/get", fieldNumber)

	req, err := s.client.NewRequest(http.MethodPost, u, BulkReadRequest{Requests: requests})
	if err != nil {
		return nil, nil, err
	}

	res := new(BulkReadExtraFieldResponse)
	resp, err := s.client.Do(ctx, req, res)
	return res, resp, err
}

func (s *ExtraFields) Update(ctx context.Context, fieldNumber uint, fieldID int, f *ExtraFieldRequest) (*IDResponse, *http.Response, error) {
	if err := s.isFieldNumberValid(fieldNumber); err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("product/extra/field-%d/%d", fieldNumber, fieldID)

	req, err := s.client.NewRequest(http.MethodPatch, u, f)
	if err != nil {
		return nil, nil, err
	}

	id := new(IDResponse)
	resp, err := s.client.Do(ctx, req, id)
	return id, resp, err
}

func (s *ExtraFields) Move(ctx context.Context, fieldNumber uint, fieldID int, moveReq *MoveRequest) (*IDResponse, *http.Response, error) {
	if err := s.isFieldNumberValid(fieldNumber); err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("product/extra/field-%d/%d/move", fieldNumber, fieldID)

	req, err := s.client.NewRequest(http.MethodPatch, u, moveReq)
	if err != nil {
		return nil, nil, err
	}

	id := new(IDResponse)
	resp, err := s.client.Do(ctx, req, id)
	return id, resp, err
}

func (s *ExtraFields) Delete(ctx context.Context, fieldNumber uint, fieldID int) (*IDResponse, *http.Response, error) {
	if err := s.isFieldNumberValid(fieldNumber); err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("product/extra/field-%d/%d", fieldNumber, fieldID)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, nil, err
	}

	id := new(IDResponse)
	resp, err := s.client.Do(ctx, req, id)
	return id, resp, err
}

func (s *ExtraFields) isFieldNumberValid(num uint) error {
	validNumbers := []uint{1, 2, 3, 4}
	for _, validNum := range validNumbers {
		if num == validNum {
			return nil
		}
	}
	return fmt.Errorf("field number not valid, valid numbers are %v", validNumbers)
}
