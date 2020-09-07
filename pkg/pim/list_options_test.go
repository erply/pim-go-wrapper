package pim

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"strings"
	"testing"
)

func TestAddOptions(t *testing.T) {
	u := "http://example.com"
	filter1, err := NewFilter("status", "=", "ACTIVE", "and")
	assert.NoError(t, err)
	filter2, err := NewFilter("type", "=", "BUNDLE", "")
	assert.NoError(t, err)
	opts := &ListOptions{
		Filters: []Filter{
			*filter1,
			*filter2,
		},
		PaginationParameters: NewPaginationParameters(5, 2),
		SortingParameter:     NewSortingParameter("added", true, "gr"),
	}
	_, err = addOptions(u, opts)
	assert.NoError(t, err)

	t.Run("not valid operator", func(t *testing.T) {
		opts.Filters[1].Operation = "==="
		_, err := addOptions(u, opts)
		assert.EqualError(t, err, "could not parse filtering parameter: unknown column filter operation ===, accepted values are [= >= <= contains startswith]")
	})
	t.Run("not valid column operand", func(t *testing.T) {
		opts.Filters[0].OperandAfter = "with"
		_, err := addOptions(u, opts)
		assert.EqualError(t, err, "could not parse filtering parameter: unknown filtering operand with, accepted values are [and or]")
	})
}

func TestAddIDs(t *testing.T) {
	u, err := url.Parse("http://example.com")
	assert.NoError(t, err)
	urlWithIDs, err := addIDs(u, "recordIDs", 1, 2, 3)
	assert.NoError(t, err)
	ids := urlWithIDs.Query().Get("recordIDs")
	assert.Equal(t, 3, len(strings.Split(ids, ";")))
}
