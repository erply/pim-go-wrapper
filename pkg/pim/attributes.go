package pim

import (
	"context"
	"net/http"
)

type (
	Attributes service
	Attribute  struct {
		//ID of the attribute
		ID int `json:"id,omitempty" example:"33"`
		//ID of the database record
		RecordID int `json:"record_id,omitempty" example:"7"`
		//Entity name of the record
		Entity string `json:"entity,omitempty" example:"product"`
		//3 types are available - int, double, text
		Type string `json:"type,omitempty" example:"text"`
		Name string `json:"name,omitempty" example:"has_warranty"`
		//based on the type value can be number of text
		Value interface{} `json:"value,omitempty"`
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
