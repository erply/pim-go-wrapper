package pim

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSuppliers_Read(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/supplier", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, err := fmt.Fprint(w, `[
  {
    "id": 28,
    "type": 32,
    "supplier_type_id": 5,
    "name": "Test ",
    "first_name": "",
    "code": "",
    "vat_number": "",
    "bank_name": "",
    "bank_account_number": "",
    "bank_iban": "",
    "bank_swift": "",
    "phone": "+123",
    "mobile": "+123",
    "fax": "+123",
    "mail": "h@e.com",
    "skype": "",
    "website": "",
    "notes": "",
    "vatrate_id": 0,
    "currency_id": 0,
    "displayed_name": "Test ",
    "displayed_name2": "Test",
    "added": 1625040278,
    "addedby": "l@e"
  },
  {
    "id": 31,
    "type": 0,
    "supplier_type_id": 5,
    "name": "string",
    "first_name": "",
    "code": "",
    "vat_number": "",
    "bank_name": "",
    "bank_account_number": "",
    "bank_iban": "",
    "bank_swift": "",
    "phone": "+123",
    "mobile": "+123",
    "fax": "+123",
    "mail": "h@c.com",
    "skype": "string",
    "website": "",
    "notes": "string",
    "vatrate_id": 0,
    "currency_id": 0,
    "displayed_name": "string, ",
    "displayed_name2": " string",
    "added": 1625040780,
    "addedby": "h@e"
  }
]`)
		assert.NoError(t, err)
	})

	opts := NewListOptions(nil, nil, nil, false)
	response, _, err := client.Suppliers.Read(context.Background(), opts)
	if err != nil {
		t.Error(err)
	}

	suppliers := *response
	assert.Equal(t, "Test ", suppliers[0].DisplayedName)
	assert.Equal(t, "h@c.com", suppliers[1].Mail)
	assert.Equal(t, "+123", suppliers[1].Fax)
}

func TestSuppliers_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/supplier", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		_, err := fmt.Fprint(w, `
		{
		  "id":1
		}
	`)
		assert.NoError(t, err)
	})

	p := &SupplierRequest{}
	idResult, _, err := client.Suppliers.Create(context.Background(), p)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, idResult.ID)
}

func TestSuppliers_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/supplier/1;2;3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		_, err := fmt.Fprint(w, `
{
	"ids": [
		1, 2, 3
	]
}
`)
		assert.NoError(t, err)
	})

	response, _, err := client.Suppliers.Delete(context.Background(), []int{1, 2, 3})
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, []int{1, 2, 3}, response.IDs)
}
