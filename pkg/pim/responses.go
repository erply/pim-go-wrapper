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
		IDResponse
		MessageResponse
	}
)
