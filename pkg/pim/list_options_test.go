package pim

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddOptions(t *testing.T) {
	url := "http://example.com"
	opts := &ListOptions{
		Filters: []Filter{
			*NewFilter(NewColumnFilter("status", "=", "ACTIVE"), "and"),
			*NewFilter(NewColumnFilter("type", "=", "BUNDLE"), ""),
		},
		PaginationParameters: NewPaginationParameters(5, 2),
		SortingParameter:     NewSortingParameter("added", true, "gr"),
	}
	_, err := addOptions(url, opts)
	assert.NoError(t, err)

	t.Run("not valid operator", func(t *testing.T) {
		opts.Filters[1].ColumnFilter.Operation = "==="
		_, err := addOptions(url, opts)
		assert.EqualError(t, err, "could not parse filtering parameter: unknown column filter operation ===, accepted values are [= >= <= contains startswith]")
	})
	t.Run("not valid column operand", func(t *testing.T) {
		opts.Filters[0].Operand = "with"
		_, err := addOptions(url, opts)
		assert.EqualError(t, err, "could not parse filtering parameter: unknown filtering operand with, accepted values are [and or]")
	})
}
