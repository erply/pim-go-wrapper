package pim

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestBrands_Read(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/brand", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, err := fmt.Fprint(w, `
[{
		"id": 1,
		"name": "1"
	},
	{
		"id": 3,
		"name": "3"
	},
	{
		"id": 2,
		"name": "2"
	}
]`)
		assert.NoError(t, err)
	})

	opts := NewListOptions(nil, nil, nil, false)
	results, _, err := client.Brands.Read(context.Background(), opts)
	if err != nil {
		t.Error(err)
	}

	for _, r := range *results {
		if r.ID == int64(1) {
			assert.Equal(t, "1", r.Name)
		}
	}
}

func TestBrands_ReadBulk(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/brand/bulk/get", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		_, err := fmt.Fprint(w, `
{
  "results": [
    {
      "resultId": 0,
      "brands": [
        {
          "id": 1,
          "name": "brand1",
          "added": 3586870217,
          "addedby": "user",
          "changed": 0,
          "changedby": ""
        },
        {
          "id": 6,
          "name": "brand2",
          "added": 3586870221,
          "addedby": "user",
          "changed": 0,
          "changedby": ""
        },
        {
          "id": 8,
          "name": "string",
          "added": 3586937922,
          "addedby": "user",
          "changed": 0,
          "changedby": ""
        },
        {
          "id": 9,
          "name": "string",
          "added": 3599132075,
          "addedby": "user",
          "changed": 0,
          "changedby": ""
        }
      ]
    }
  ]
}`)
		assert.NoError(t, err)
	})

	opts := NewListOptions(nil, nil, nil, false)
	results, _, err := client.Brands.ReadBulk(context.Background(), []ListOptions{*opts})
	if err != nil {
		t.Error(err)
	}

	for _, r := range results.Results {
		if r.ResultID == 0 {
			assert.Equal(t, "string", r.Records[3].Name)
		}
	}
}

func TestBrands_CreateBulk(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/brand/bulk", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		_, err := fmt.Fprint(w, `
{
  "results": [
    {
      "resultId": 1,
      "resourceId": 16
    },
    {
      "resultId": 0,
      "resourceId": 17
    },
    {
      "resultId": 2,
      "resourceId": 18
    }
  ]
}`)
		assert.NoError(t, err)
	})

	brands := []Brand{{}, {}, {}}
	results, _, err := client.Brands.CreateBulk(context.Background(), brands)
	if err != nil {
		t.Error(err)
	}

	for _, r := range results.Results {
		if r.ResultID == 2 {
			assert.Equal(t, 18, r.ResourceID)
		}
	}
}
