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
	Suppliers       service
	SupplierRequest struct {
		BankAccountNumber string `json:"bank_account_number,omitempty"`
		BankIban          string `json:"bank_iban,omitempty"`
		BankName          string `json:"bank_name,omitempty"`
		BankSwift         string `json:"bank_swift,omitempty"`
		Code              string `json:"code,omitempty"`
		CurrencyID        int    `json:"currency_id,omitempty"`
		Fax               string `json:"fax,omitempty"`
		FirstName         string `json:"first_name,omitempty"`
		Mail              string `json:"mail,omitempty"`
		Mobile            string `json:"mobile,omitempty"`
		Name              string `json:"name,omitempty"`
		Notes             string `json:"notes,omitempty"`
		Phone             string `json:"phone,omitempty"`
		Skype             string `json:"skype,omitempty"`
		SupplierTypeID    int    `json:"supplier_type_id,omitempty"`
		Type              int    `json:"type,omitempty"`
		VatNumber         string `json:"vat_number,omitempty"`
		VatRateID         int    `json:"vatrate_id,omitempty"`
		Website           string `json:"website,omitempty"`
	}
	Supplier struct {
		DbRecord
		SupplierRequest
		DisplayedName  string `json:"displayed_name"`
		DisplayedName2 string `json:"displayed_name_2"`
		AddedAddedBy
	}
)

func (s *Suppliers) Read(ctx context.Context, opts *ListOptions) (*[]Supplier, *http.Response, error) {
	urlStr := "supplier"
	u, err := addOptions(urlStr, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	dataResp := new([]Supplier)
	resp, err := s.client.Do(ctx, req, dataResp)
	if err != nil {
		return nil, resp, err
	}

	return dataResp, resp, nil
}

func (s *Suppliers) Create(ctx context.Context, request *SupplierRequest) (*IDResponse, *http.Response, error) {
	u := "supplier"

	req, err := s.client.NewRequest(http.MethodPost, u, request)
	if err != nil {
		return nil, nil, err
	}

	id := new(IDResponse)
	resp, err := s.client.Do(ctx, req, id)
	return id, resp, err
}

func (s *Suppliers) Delete(ctx context.Context, ids []int) (*BulkResponse, *http.Response, error) {
	if len(ids) < 1 {
		return nil, nil, errors.New("need at least one ID to delete")
	}
	var strIDs []string
	for _, id := range ids {
		strIDs = append(strIDs, strconv.Itoa(id))
	}

	u := fmt.Sprintf("supplier/%s", strings.Join(strIDs, ";"))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, nil, err
	}

	bulkIds := new(BulkResponse)
	resp, err := s.client.Do(ctx, req, bulkIds)
	return bulkIds, resp, err
}
