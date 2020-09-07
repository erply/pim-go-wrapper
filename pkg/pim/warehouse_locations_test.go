package pim

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestWarehouseLocations_Read(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/warehouse/locations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, err := fmt.Fprint(w, `[
  {
    "id": 1,
    "name": "loc",
    "added": 12312312,
    "addedby": "123123",
    "changed": 123123,
    "changedby": "12312312"
  },
  {
    "id": 2,
    "name": "loc2",
    "added": 12312312,
    "addedby": "123123",
    "changed": 123123,
    "changedby": "12312312"
  },
  {
    "id": 3,
    "name": "loc3",
    "added": 12312312,
    "addedby": "123123",
    "changed": 123123,
    "changedby": "12312312"
  }
]`)
		assert.NoError(t, err)
	})

	opts := NewListOptions(nil, nil, nil, false)
	locations, _, err := client.WarehouseLocations.Read(context.Background(), opts)
	if err != nil {
		t.Errorf("GET warehouse locations returned error: %v", err)
	}

	t.Log(locations)

}
