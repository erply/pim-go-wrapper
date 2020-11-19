package pim

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestProducts_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		_, err := fmt.Fprint(w, `
		{
		  "id":1
		}
	`)
		assert.NoError(t, err)
	})

	p := &ProductRequest{}
	idResult, _, err := client.Products.Create(context.Background(), p)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, idResult.ID)
}

func TestProducts_Read(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, err := fmt.Fprint(w, `
		[
{
    "age_restriction":16
},
{
	"age_restriction":18
}
]
`)
		assert.NoError(t, err)
	})

	opts := NewListOptions(nil, nil, nil, false)
	result, _, err := client.Products.Read(context.Background(), opts)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 2, len(*result))
	res := *result
	assert.Equal(t, 16, res[0].AgeRestriction)
	assert.Equal(t, 18, res[1].AgeRestriction)
}
