package pim

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"strings"
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

type AttributeRequest struct {
	//ID of the attribute - used only to udpate
	ID int `json:"id"`
	//ID of the record
	RecordID int `json:"record_id" example:"7"`
	//Entity name of the record
	Entity string `json:"entity" example:"product"`
	//3 types are available - int, double, text
	Type string `json:"type" example:"int"`
	//name of the attribute
	Name string `json:"name" example:"has_warranty"`
	//based on the type value should be number of text
	Value interface{} `json:"value"`
}

func (s *Attributes) Attach(ctx context.Context, request *AttributeRequest) (*IDResponse, *http.Response, error) {
	u := "attribute"
	req, err := s.client.NewRequest(http.MethodPut, u, request)
	if err != nil {
		return nil, nil, err
	}

	id := new(IDResponse)
	resp, err := s.client.Do(ctx, req, id)
	return id, resp, err
}

func (s *Attributes) Delete(ctx context.Context, ids []int) (*BulkResponse, *http.Response, error) {
	if len(ids) < 1 {
		return nil, nil, errors.New("need at least one ID to delete")
	}
	var strIDs []string
	for _, id := range ids {
		strIDs = append(strIDs, strconv.Itoa(id))
	}

	u := fmt.Sprintf("attribute/%s", strings.Join(strIDs, ";"))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, nil, err
	}

	bulkIds := new(BulkResponse)
	resp, err := s.client.Do(ctx, req, bulkIds)
	return bulkIds, resp, err
}
