package pim

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type (
	Groups       service
	ProductGroup struct {
		ID int `json:"id,omitempty"`
		ProductGroupRequest
	}
	ProductGroupRequest struct {
		LocationTaxRates []struct {
			TaxRateID   int `json:"tax_rate_id"`
			WarehouseID int `json:"warehouse_id"`
		} `json:"location_tax_rates,omitempty"`
		TranslatableNameJSON
		TranslatableDescriptionJSON
		NonDiscountable         int    `json:"non_discountable,omitempty"`
		Order                   int    `json:"order,omitempty"`
		ParentID                int    `json:"parent_id,omitempty"`
		QuickBooksCreditAccount string `json:"quick_books_credit_account,omitempty"`
		QuickBooksDebitAccount  string `json:"quick_books_debit_account,omitempty"`
		RewardPoints            int    `json:"reward_points,omitempty"`
		ShowInWebShop           int    `json:"show_in_webshop,omitempty"`
		//These fields are not editable
		AddedByChangedBy
	}

	ProductAdditionalGroup struct {
		AddedByChangedBy
		GroupID   int `json:"group_id"`
		ID        int `json:"id"`
		ProductID int `json:"product_id"`
	}

	ProductAdditionalGroupsRequest struct {
		Results []ProductAdditionalGroup `json:"results"`
	}
)

func (s *Groups) Read(ctx context.Context, opts *ListOptions) (*[]ProductGroup, *http.Response, error) {
	urlStr := "product/group"
	u, err := addOptions(urlStr, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	dataResp := new([]ProductGroup)
	resp, err := s.client.Do(ctx, req, dataResp)
	return dataResp, resp, err
}

func (s *Groups) ReadAdditionalGroups(ctx context.Context, ids []string, opts PaginationParameters) (*[]ProductAdditionalGroup, *http.Response, error) {
	urlStr := fmt.Sprintf("product/%s/additional-groups", strings.Join(ids, ";"))
	u, err := addOptions(urlStr, &ListOptions{PaginationParameters: &opts})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	dataResp := new(ProductAdditionalGroupsRequest)
	resp, err := s.client.Do(ctx, req, dataResp)
	return &dataResp.Results, resp, err
}
