package pim

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestProductUnits_Read(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/product/unit", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, err := fmt.Fprint(w, `[
{
    "id":16
},
{
	"id":18
}
]`)
		assert.NoError(t, err)
	})

	opts := NewListOptions(nil, nil, nil, false)
	result, _, err := client.ProductUnits.Read(context.Background(), opts)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 2, len(*result))
	res := *result
	assert.Equal(t, 16, res[0].ID)
	assert.Equal(t, 18, res[1].ID)
}
