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
)
