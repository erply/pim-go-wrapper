package pim

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAttributes_Read(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/attribute", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, err := fmt.Fprint(w, `[
  {
     "entity": "product",
    "id": 1,
    "name": "has_warranty",
    "record_id": 7,
    "type": "text",
    "value": {}
  },
  {
     "entity": "product",
    "id": 2,
    "name": "has_warranty",
    "record_id": 7,
    "type": "text",
    "value": {}
  },
  {
     "entity": "product",
    "id": 3,
    "name": "has_warranty",
    "record_id": 7,
    "type": "text",
    "value": {}
  }
]`)
		assert.NoError(t, err)
	})

	opts := NewListOptions(nil, nil, nil, false)
	attributes, _, err := client.Attributes.Read(context.Background(), opts, 1, 2, 3)
	if err != nil {
		t.Error(err)
	}

	for _, a := range *attributes {
		assert.Equal(t, "has_warranty", a.Name)
		assert.Equal(t, 7, a.RecordID)
	}

}

func TestAttributes_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/attribute/1;2;3", func(w http.ResponseWriter, r *http.Request) {
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

	response, _, err := client.Attributes.Delete(context.Background(), []int{1, 2, 3})
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, []int{1, 2, 3}, response.IDs)
}
