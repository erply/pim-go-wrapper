package pim

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddOptions(t *testing.T) {
	url := "http://example.com"
	opts := &ListOptions{
		Filters: []Filter{
			{
				ColumnFilter: [3]interface{}{"status", "=", "ACTIVE"},
				Operand:      "and",
			},
			{
				ColumnFilter: [3]interface{}{"type", "=", "BUNDLE"},
			},
		},
		PaginationParameters: &PaginationParameters{
			Skip: 5,
			Take: 2,
		},
		SortingParameter: &SortingParameter{
			Selector: "added",
			Desc:     true,
			Language: "gr",
		},
	}
	_, err := addOptions(url, opts)
	assert.NoError(t, err)

	t.Run("not valid operator", func(t *testing.T) {
		opts.Filters[1].ColumnFilter[1] = "==="
		_, err := addOptions(url, opts)
		assert.EqualError(t, err, "unknown column filter operation ===, accepted values are [= >= <= contains startswith]")
	})
	t.Run("not valid column operand", func(t *testing.T) {
		opts.Filters[0].Operand = "with"
		_, err := addOptions(url, opts)
		assert.EqualError(t, err, "unknown filtering operand with, accepted values are [and or]")
	})
}
