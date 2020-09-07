package pim

type (
	DbRecord struct {
		//object ID
		ID int `json:"id"`
	}
	AddedAddedBy struct {
		//Unix timestamp
		Added int64 `json:"added"`
		//username
		AddedBy string `json:"addedby"`
	}
	ChangedChangedBy struct {
		//Unix timestamp
		Changed int64 `json:"changed"`
		//username
		ChangedBy string `json:"changedby"`
	}

	AddedByChangedBy struct {
		AddedAddedBy
		ChangedChangedBy
	}

	TranslatableNameJSON struct {
		Name map[string]string `json:"name"`
	}
	TranslatableDescriptionJSON struct {
		Description map[string]ProductDescription `json:"description"`
	}
	ProductDescription struct {
		PlainText string `json:"plain_text"`
		HTML      string `json:"html"`
	}
	Physical struct {
		// NetWeight is Item's net weight. Unit depends on region, check your Erply account (typically lbs or kg).
		NetWeight float64 `json:"net_weight"`
		// GrossWeight is Item's gross weight (with packaging). Unit depends on region, check your Erply account (typically lbs or kg).
		GrossWeight float64 `json:"gross_weight"`
		// Length is Item's physical dimensions.
		Length float64 `json:"length"`
		Width  float64 `json:"width"`
		Height float64 `json:"height"`
	}
	Status struct {
		//Status is a classifier with four possible values: 'ACTIVE' (DEFAULT), 'NO_LONGER_ORDERED', 'NOT_FOR_SALE' and 'ARCHIVED'.
		Status string `json:"status" example:"ACTIVE"`
	}

	ProductAttributes struct {
		DeliveryTime            string `json:"delivery_time"`
		PackagingType           string `json:"packaging_type" `
		AlcoholRegistryNumber   string `json:"alcohol_registry_number"`
		AlcoholPercentage       string `json:"alcohol_percentage"`
		Batches                 string `json:"batches"`
		ExciseDeclarationNumber string `json:"excise_declaration_number"`

		//boolean flag 0 or 1
		TaxFree int `json:"tax_free"`
		//boolean flag 0 or 1
		IsRegularGiftCard int `json:"is_regular_gift_card"`
		//boolean flag 0 or 1
		RewardPointsNotAllowed int `json:"reward_points_not_allowed"`
		//boolean flag 0 or 1
		NonStockProduct int `json:"non_stock_product"`
		//boolean flag 0 or 1
		CashierMustEnterPrice int `json:"cashier_must_enter_price"`
		//boolean flag 0 or 1
		LabelsNotNeeded int `json:"labels_not_needed"`

		DepositFeeAmount int `json:"deposit_fee_amount"`
	}
)
