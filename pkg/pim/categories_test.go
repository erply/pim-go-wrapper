package pim

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCategories_CreateBulk(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/product/category/bulk", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		_, err := fmt.Fprint(w, `
{
  "results": [
    {
      "resultId": 0,
      "resourceId": 2230
    },
    {
      "resultId": 2,
      "resourceId": 2231
    },
    {
      "resultId": 1,
      "resourceId": 2232
    }
  ]
}
`)
		assert.NoError(t, err)
	})

	categories := []CategoryRequest{{}}
	results, _, err := client.Categories.CreateBulk(context.Background(), categories)
	if err != nil {
		t.Error(err)
	}

	for _, r := range results.Results {
		if r.ResultID == 1 {
			assert.Equal(t, 2232, r.ResourceID)
		}
	}
}

func TestCategories_ReadBulk(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/product/category/bulk/get", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		_, err := fmt.Fprint(w, `
{
  "results": [
    {
      "resultId": 1,
      "categories": null,
      "totalCount": 0
    },
    {
      "resultId": 0,
      "categories": [
        {
          "id": 1,
          "parent_id": 0,
          "order": 0,
          "name": {
            "de": "Standardkategorie",
            "el": "????????????? ?????",
            "en": "Default category",
            "es": "Predeterminado serie",
            "et": "Vaikimisi seeria",
            "fi": "Default category",
            "lt": "Standartin? serija",
            "lv": "Default category",
            "ru": "????? ?? ?????????",
            "sv": "Defaultserier"
          },
          "added": 0,
          "addedby": "",
          "changed": 0,
          "changedby": ""
        },
        {
          "id": 2,
          "parent_id": 0,
          "order": 0,
          "name": {
            "de": "Standardkategorie",
            "el": "????????????? ?????",
            "en": "Default category",
            "es": "Predeterminado serie",
            "et": "Vaikimisi seeria",
            "fi": "Default category",
            "lt": "Standartin? serija",
            "lv": "Default category",
            "ru": "????? ?? ?????????",
            "sv": "Defaultserier"
          },
          "added": 0,
          "addedby": "",
          "changed": 0,
          "changedby": ""
        },
        {
          "id": 5,
          "parent_id": 0,
          "order": 0,
          "name": {
            "de": "Standardkategorie",
            "el": "????????????? ?????",
            "en": "Default category",
            "es": "Predeterminado serie",
            "et": "Vaikimisi seeria",
            "fi": "Default category",
            "lt": "Standartin? serija",
            "lv": "Default category",
            "ru": "????? ?? ?????????",
            "sv": "Defaultserier"
          },
          "added": 0,
          "addedby": "",
          "changed": 0,
          "changedby": ""
        },
        {
          "id": 6,
          "parent_id": 0,
          "order": 0,
          "name": {
            "de": "Standardkategorie",
            "el": "????????????? ?????",
            "en": "Default category",
            "es": "Predeterminado serie",
            "et": "Vaikimisi seeria",
            "fi": "Default category",
            "lt": "Standartin? serija",
            "lv": "Default category",
            "ru": "????? ?? ?????????",
            "sv": "Defaultserier"
          },
          "added": 0,
          "addedby": "",
          "changed": 0,
          "changedby": ""
        },
        {
          "id": 7,
          "parent_id": 0,
          "order": 0,
          "name": {
            "de": "Standardkategorie",
            "el": "????????????? ?????",
            "en": "Default category",
            "es": "Predeterminado serie",
            "et": "Vaikimisi seeria",
            "fi": "Default category",
            "lt": "Standartin? serija",
            "lv": "Default category",
            "ru": "????? ?? ?????????",
            "sv": "Defaultserier"
          },
          "added": 0,
          "addedby": "",
          "changed": 0,
          "changedby": ""
        }
      ],
      "totalCount": 5
    }
  ]
}
`)
		assert.NoError(t, err)
	})

	opts := NewListOptions(nil, nil, nil, false)
	results, _, err := client.Categories.ReadBulk(context.Background(), []ListOptions{*opts})
	if err != nil {
		t.Error(err)
	}

	for _, r := range results.Results {
		if r.ResultID == 1 {
			assert.Nil(t, r.Records)
		} else {
			assert.Equal(t, "Defaultserier", r.Records[4].Name["sv"])
		}
	}
}
