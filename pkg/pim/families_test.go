package pim

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFamilies_CreateBulk(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/product/family/bulk", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		_, err := fmt.Fprint(w, `
{
  "results": [
    {
      "resultId": 0,
      "resourceId": 320
    },
    {
      "resultId": 1,
      "resourceId": 321
    }
  ]
}
`)
		assert.NoError(t, err)
	})

	families := []FamilyRequest{}
	results, _, err := client.Families.CreateBulk(context.Background(), families)
	if err != nil {
		t.Error(err)
	}

	for _, r := range results.Results {
		if r.ResultID == 1 {
			assert.Equal(t, 321, r.ResourceID)
		}
	}
}

func TestFamilies_ReadBulk(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/product/family/bulk/get", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		_, err := fmt.Fprint(w, `
{
  "results": [
    {
      "resultId": 1,
      "families": [
        {
          "id": 66,
          "name": {
            "en": "for links",
            "et": "",
            "gr": "greek",
            "ru": ""
          },
          "archived": 0,
          "added": 1579269991,
          "addedby_id": 3,
          "changed": 1580392680,
          "changedby_id": 3
        },
        {
          "id": 80,
          "name": {
            "en": "Souvenirs",
            "et": "",
            "ru": ""
          },
          "archived": 1,
          "added": 1579603655,
          "addedby_id": 3,
          "changed": 0,
          "changedby_id": 0
        }
      ]
    },
    {
      "resultId": 0,
      "families": [
        {
          "id": 66,
          "name": {
            "en": "for links",
            "et": "",
            "gr": "greek",
            "ru": ""
          },
          "archived": 0,
          "added": 1579269991,
          "addedby_id": 3,
          "changed": 1580392680,
          "changedby_id": 3
        },
        {
          "id": 80,
          "name": {
            "en": "Souvenirs",
            "et": "",
            "ru": ""
          },
          "archived": 1,
          "added": 1579603655,
          "addedby_id": 3,
          "changed": 0,
          "changedby_id": 0
        }
      ]
    }
  ]
}
`)
		assert.NoError(t, err)
	})

	opts := NewListOptions(nil, nil, nil, false)
	results, _, err := client.Families.ReadBulk(context.Background(), []FamilyListOptions{{
		ListOptions:  *opts,
		WithArchived: false,
	}})

	if err != nil {
		t.Error(err)
	}

	for _, r := range results.Results {
		if r.ResultID == 0 {
			assert.Equal(t, "Souvenirs", r.Records[1].Name["en"])
			assert.Equal(t, 3, r.Records[1].AddedByID)
		}
	}
}
