package pim

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type MatrixProducts service
type MatrixProduct struct {
	Code               string `json:"code"`
	Code2              string `json:"code2"`
	DisplayedInWebshop int64  `json:"displayed_in_webshop"`
	ID                 int64  `json:"id"`
	TranslatableNameJSON
	ParentProductID   int64 `json:"parent_product_id"`
	ParentProductName struct {
		En string `json:"en"`
		Et string `json:"et"`
		Ru string `json:"ru"`
	} `json:"parent_product_name"`
	Status               string `json:"status"`
	VariationDescription []struct {
		DimensionID         int64             `json:"dimension_id"`
		DimensionName       map[string]string `json:"dimension_name"`
		DimensionValueCode  string            `json:"dimension_value_code"`
		DimensionValueID    int64             `json:"dimension_value_id"`
		DimensionValueName  map[string]string `json:"dimension_value_name"`
		DimensionValueOrder int64             `json:"dimension_value_order"`
	} `json:"variation_description"`
}

func (s *MatrixProducts) Read(ctx context.Context, matrixProductIDs, productIDs []uint, paginationParameters *PaginationParameters) (*[]MatrixProduct, *http.Response, error) {
	urlStr := "matrix/product"
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, nil, err
	}
	q := u.Query()

	//apply pagination
	if paginationParameters != nil {
		q.Add("skip", strconv.Itoa(int(paginationParameters.Skip)))
		q.Add("take", strconv.Itoa(int(paginationParameters.Take)))
	}

	if len(matrixProductIDs) > 0 {
		q.Add("matrixProductIDs", strings.ReplaceAll(strings.Trim(fmt.Sprint(productIDs), "[]"), " ", ","))
	}

	if len(productIDs) > 0 {
		q.Add("productIDs", strings.ReplaceAll(strings.Trim(fmt.Sprint(productIDs), "[]"), " ", ","))
	}

	req, err := s.client.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	dataResp := new([]MatrixProduct)
	resp, err := s.client.Do(ctx, req, dataResp)
	return dataResp, resp, err
}
