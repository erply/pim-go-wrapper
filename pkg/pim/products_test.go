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

func TestProducts_ReadAdditionalGroups(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/product/1;2/additional-groups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, err := fmt.Fprint(w, `  {"results": [
    {
      "id": 2,
      "product_id": 1
    },
    {
      "id": 3,
      "product_id": 2
    }
  ]}`)
		assert.NoError(t, err)
	})

	opts := NewPaginationParameters(0, 0)
	result, _, err := client.Products.ReadAdditionalGroups(context.Background(), []string{"1", "2"}, *opts)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 2, len(*result))
	res := *result
	assert.Equal(t, 2, res[0].ID)
	assert.Equal(t, 3, res[1].ID)
}
