package pim

type (
	DbRecord struct {
		//object ID
		ID int `json:"id,omitempty"`
	}
	AddedAddedBy struct {
		//Unix timestamp
		Added int64 `json:"added,omitempty"`
		//username
		AddedBy string `json:"addedby,omitempty"`
	}
	ChangedChangedBy struct {
		//Unix timestamp
		Changed int64 `json:"changed,omitempty"`
		//username
		ChangedBy string `json:"changedby,omitempty"`
	}

	AddedByChangedBy struct {
		AddedAddedBy
		ChangedChangedBy
	}

	TranslatableNameJSON struct {
		Name map[string]string `json:"name,omitempty"`
	}
	TranslatableDescriptionJSON struct {
		Description map[string]ProductDescription `json:"description,omitempty"`
	}
	ProductDescription struct {
		PlainText string `json:"plain_text,omitempty"`
		HTML      string `json:"html,omitempty"`
	}
	Physical struct {
		// NetWeight is Item's net weight. Unit depends on region, check your Erply account (typically lbs or kg).
		NetWeight float64 `json:"net_weight,omitempty"`
		// GrossWeight is Item's gross weight (with packaging). Unit depends on region, check your Erply account (typically lbs or kg).
		GrossWeight float64 `json:"gross_weight,omitempty"`
		// Length is Item's physical dimensions.
		Length float64 `json:"length,omitempty"`
		Width  float64 `json:"width,omitempty"`
		Height float64 `json:"height,omitempty"`
	}
	Status struct {
		//Status is a classifier with four possible values: 'ACTIVE' (DEFAULT), 'NO_LONGER_ORDERED', 'NOT_FOR_SALE' and 'ARCHIVED'.
		Status string `json:"status,omitempty" example:"ACTIVE"`
	}

	ProductAttributes struct {
		DeliveryTime            string `json:"delivery_time,omitempty"`
		PackagingType           string `json:"packaging_type,omitempty"`
		AlcoholRegistryNumber   string `json:"alcohol_registry_number,omitempty"`
		AlcoholPercentage       string `json:"alcohol_percentage,omitempty"`
		Batches                 string `json:"batches,omitempty"`
		ExciseDeclarationNumber string `json:"excise_declaration_number,omitempty"`

		//boolean flag 0 or 1
		TaxFree int `json:"tax_free,omitempty"`
		//boolean flag 0 or 1
		IsRegularGiftCard int `json:"is_regular_gift_card,omitempty"`
		//boolean flag 0 or 1
		RewardPointsNotAllowed int `json:"reward_points_not_allowed,omitempty"`
		//boolean flag 0 or 1
		NonStockProduct int `json:"non_stock_product,omitempty"`
		//boolean flag 0 or 1
		CashierMustEnterPrice int `json:"cashier_must_enter_price,omitempty"`
		//boolean flag 0 or 1
		LabelsNotNeeded int `json:"labels_not_needed,omitempty"`

		DepositFeeAmount int `json:"deposit_fee_amount,omitempty"`
	}
)
