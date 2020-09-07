package pim

import (
	"context"
	"net/http"
)

type (
	Attributes service
	Attribute  struct {
		//ID of the attribute
		ID int `json:"id" example:"33"`
		//ID of the database record
		RecordID int `json:"record_id" example:"7"`
		//Entity name of the record
		Entity string `json:"entity" example:"product"`
		//3 types are available - int, double, text
		Type string `json:"type" example:"text"`
		Name string `json:"name" example:"has_warranty"`
		//based on the type value can be number of text
		Value interface{} `json:"value"`
	}
)

func (s *Attributes) Read(ctx context.Context, opts *ListOptions, recordIDs ...int) (*[]Attribute, *http.Response, error) {
	urlStr := "attribute"
	u, err := addOptions(urlStr, opts)
	if err != nil {
		return nil, nil, err
	}
	u, err = addIDs(u, "recordIDs", recordIDs...)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	dataResp := new([]Attribute)
	resp, err := s.client.Do(ctx, req, dataResp)
	return dataResp, resp, err
}
