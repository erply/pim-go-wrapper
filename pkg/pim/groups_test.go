package pim

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGroups_Read(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/product/group", func(w http.ResponseWriter, r *http.Request) {
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
	result, _, err := client.Groups.Read(context.Background(), opts)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 2, len(*result))
	res := *result
	assert.Equal(t, 16, res[0].ID)
	assert.Equal(t, 18, res[1].ID)
}

func TestGroups_ReadAdditionalGroups(t *testing.T) {
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
	result, _, err := client.Groups.ReadAdditionalGroups(context.Background(), []string{"1", "2"}, *opts)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 2, len(*result))
	res := *result
	assert.Equal(t, 2, res[0].ID)
	assert.Equal(t, 3, res[1].ID)
}
