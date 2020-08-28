package pim

import (
	"context"
	"fmt"
	"net/http"
)

type (
	ProductService service
	Products       []Product
	Product        struct {
		ID int `json:"id"`
		// Product type, possible types are 'PRODUCT', 'BUNDLE', 'MATRIX', 'ASSEMBLY'. By default 'PRODUCT'.
		Type string `json:"type"`
		// ID of product group. To get the list of product groups, use getProductGroups.
		GroupID int `json:"group_id"`
		// ID of product unit. To get the list of units, use getProductUnits.
		UnitID int `json:"unit_id"`

		TranslatableNameJSON

		// //PlainTextDescription is translatable plain text description
		// PlainTextDescription map[string]string `json:"plain_text_description"`
		// //HTMLDescription is translatable html description
		// HTMLDescription map[string]string `json:"html_description"`
		// Description is translatable and of 2 types: plain text, HTML. Languages should be in ISO 639-1 Code
		TranslatableDescriptionJSON

		// Product's code. Must be UNIQUE, unless the account is configured otherwise.
		Code string `json:"code"`
		// Product's second code (by convention, EAN barcode). Must be UNIQUE, unless the account is configured otherwise.
		Code2 string `json:"code2"`
		// Third code of the item (note that this field may not be visible on product card by default).
		Code3 string `json:"code3"`
		// Supplier's product code
		SupplierCode string `json:"supplier_code"`
		//TaxRateID is just the default tax rate of a product and the actual tax applied in a particular location depends on multiple rules: https://learn-api.erply.com/concepts/taxes.
		TaxRateID int `json:"tax_rate_id"`
		//Price is just the default price of a product and the actual price applied in a particular location, to a particular customer, depends on price lists and promotions: https://learn-api.erply.com/concepts/pricing
		Price float64 `json:"price"`

		// NetWeight is Item's net weight. Unit depends on region, check your Erply account (typically lbs or kg).
		NetWeight float64 `json:"net_weight"`

		Physical
		//0 or 1
		IsGiftCard int `json:"is_gift_card"`
		//0 or 1
		NonDiscountable int `json:"non_discountable"`
		//0 or 1
		NonRefundable int `json:"non_refundable"`

		// Volume is Item's fluid volume, eg. for beverages or perfumery. Unit depends on locale, check your Erply account (typically mL or fl oz).
		Volume     int `json:"volume"`
		CategoryID int `json:"category_id"`
		// ID of product brand. To get the list of brands, use getBrands.
		BrandID           int    `json:"brand_id"`
		SupplierID        int    `json:"supplier_id"`
		PriorityGroupID   int    `json:"priority_group_id"`
		CountryOfOriginID int    `json:"country_of_origin_id"`
		ManufacturerName  string `json:"manufacturer_name"`
		// Cost is Product cost
		Cost float64 `json:"cost"`
		Status
		//0 or 1
		DisplayedInWebshop *int `json:"displayed_in_webshop"`
		// LocationInWarehouseID is ID of selected location in warehouse.
		LocationInWarehouseID int `json:"location_in_warehouse_id"`
		// LocationInWarehouseText is Product's specific text added to location in warehouse.
		LocationInWarehouseText string `json:"location_in_warehouse_text"`
		// Parent product ID. Only for matrix variations (specific colors/sizes of a matrix item). See guidelines below.
		ParentProductID int `json:"parent_product_id"`
		// ContainerID is ID of another product, a beverage container that is always sold together with this item.
		DepositFeeID int `json:"deposit_fee_id"`

		FamilyID int64 `json:"family_id"`
		AddedByChangedBy
		*Attributes
	}
)

func (s *ProductService) Get(ctx context.Context, opts *ListOptions) (*Products, *http.Response, error) {
	u := fmt.Sprintf("product")
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	dataResp := new(Products)
	resp, err := s.client.Do(ctx, req, dataResp)
	return dataResp, resp, err
}

/*func (s *ProductService) Post(ctx context.Context, product *Product) () {
	u := fmt.Sprintf("product")

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	dataResp := new(Products)
	resp, err := s.client.Do(ctx, req, dataResp)
	return dataResp, resp, err
}*/
