package pim

type (
	IDResponse struct {
		ID int `json:"id,omitempty" `
	}
	BulkResponse struct {
		IDs []int `json:"ids,omitempty"`
	}
	MessageResponse struct {
		Message string `json:"message,omitempty" example:"some error"`
	}
	BulkResponseWithResults struct {
		Results []BulkResult `json:"results,omitempty"`
	}
	BulkResult struct {
		//identifier of the result item
		ResultID int `json:"resultId" example:"2"`
		//identifier of the REST-ful response
		ResourceID int `json:"resourceId" example:"2"`
		MessageResponse
	}

	BulkReadProductResponse struct {
		Results []BulkReadProductResponseItem `json:"results"`
	}

	BulkReadProductResponseItem struct {
		//id of the response item
		ResultID int
		//in case of error
		MessageResponse

		//total number of records (ignores skip & take parameters)
		TotalCount int `json:"totalCount"`
		//resulting records
		Products []Product `json:"products"`
	}
)
