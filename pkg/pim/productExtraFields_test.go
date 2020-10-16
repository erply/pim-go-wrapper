package pim

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestExtraFields_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/product/extra/field-1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		_, err := fmt.Fprint(w, `
		{
		  "id":1
		}
	`)
		assert.NoError(t, err)
	})

	ExtraFields := &ExtraFieldRequest{}
	idResult, _, err := client.ExtraFields.Create(context.Background(), 1, ExtraFields)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, idResult.ID)
}

func TestExtraFields_Read(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/product/extra/field-1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, err := fmt.Fprint(w, `
		[
{
    "added": 1595230999,
    "addedby": "k@erp.xyz",
    "changed": 1595230999,
    "changedby": "k@erp.xyz",
    "code": "code",
    "id": 1234,
    "is_active": true,
    "name": {
      "en": "string"
    },
    "order": 15
  },
{
    "added": 1595230999,
    "addedby": "k@erp.xyz",
    "changed": 1595230999,
    "changedby": "k@erp.xyz",
    "code": "code",
    "id": 1234,
    "is_active": true,
    "name": {
      "et": "string"
    },
    "order": 15
  }
]
`)
		assert.NoError(t, err)
	})

	opts := NewListOptions(nil, nil, nil, false)
	result, _, err := client.ExtraFields.Read(context.Background(), 1, opts)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 2, len(*result))
	res := *result
	assert.Equal(t, "string", res[0].Name["en"])
	assert.Equal(t, "string", res[1].Name["et"])
}

func TestExtraFields_Update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/product/extra/field-1/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		_, err := fmt.Fprint(w, `
		{
		  "id":1
		}
	`)
		assert.NoError(t, err)
	})

	f := &ExtraFieldRequest{}
	idResult, _, err := client.ExtraFields.Update(context.Background(), 1, 1, f)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, idResult.ID)
}

func TestExtraFields_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/product/extra/field-1/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		_, err := fmt.Fprint(w, `
		{
		  "id":1
		}
	`)
		assert.NoError(t, err)
	})

	idResult, _, err := client.ExtraFields.Delete(context.Background(), 1, 1)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, idResult.ID)
}

func TestExtraFields_Move(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/product/extra/field-1/1/move", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		_, err := fmt.Fprint(w, `
		{
		  "id":1
		}
	`)
		assert.NoError(t, err)
	})

	moveRequest := &MoveRequest{
		TargetID: 5,
		Position: "after",
	}
	idResult, _, err := client.ExtraFields.Move(context.Background(), 1, 1, moveRequest)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, idResult.ID)
}

func TestExtraFields_CreateBulk(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/product/extra/field-1/bulk", func(w http.ResponseWriter, r *http.Request) {
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

	var req []ExtraFieldRequest
	results, _, err := client.ExtraFields.CreateBulk(context.Background(), 1, req)
	if err != nil {
		t.Error(err)
	}

	for _, r := range results.Results {
		if r.ResultID == 1 {
			assert.Equal(t, 2232, r.ResourceID)
		}
	}
}

func TestCExtraFields_ReadBulk(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/product/extra/field-1/bulk/get", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		_, err := fmt.Fprint(w, `{
    "results": [
        {
            "resultId": 1,
            "extraFields": null,
            "totalCount": 0
        },
        {
            "resultId": 0,
            "extraFields": [
                {
                    "id": 1,
                    "order": 0,
                    "name": {
                        "el": "????????????? ?????"
                    }
                },
                {
                    "id": 2,
                    "order": 0,
                    "name": {
                        "de": "Standardkategorie"
                    }
                }
            ],
            "totalCount": 5
        }
    ]
}`)
		assert.NoError(t, err)
	})

	opts := NewListOptions(nil, nil, nil, false)
	results, _, err := client.ExtraFields.ReadBulk(context.Background(), 1, []ListOptions{*opts})
	if err != nil {
		t.Error(err)
	}

	for _, r := range results.Results {
		if r.ResultID == 1 {
			assert.Nil(t, r.ExtraFields)
		} else {
			assert.Equal(t, "Standardkategorie", r.ExtraFields[1].Name["de"])
		}
	}
}
