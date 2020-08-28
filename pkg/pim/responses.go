package pim

type (
	IDResponse struct {
		ID int `json:"id,omitempty" `
	}
	BulkResponse struct {
		IDs []int `json:"ids,omitempty"`
	}
	MessageResponse struct {
		Message string `json:"message" example:"some error"`
	}
)
