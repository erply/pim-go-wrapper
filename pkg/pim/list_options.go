package pim

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/url"
	"strconv"
)

type ListOptions struct {
	//Filters are an array of arrays of length 3 and string, each array represents a filter and the string - operand: and,or.
	//For the array 1st string is the field name, 2nd string is = and the 3rd object is the value (can be any type).
	//Supported filtering operations: =, >=, <=, contains, startswith.
	//Example value: [[group_id,=,2],and,[status,=,ACTIVE],and,[type,=,BUNDLE],or,[type,=,PRODUCT]]
	Filters []Filter

	//PaginationParameters include skip and take integer parameters
	PaginationParameters *PaginationParameters

	//SortingParameter is a JSON object with 3 fields -
	//1) selector - field name, example value:added,
	//2) desc - short for descending, value: true or false,
	//3) language: example value: gr, is used to sort by translatable fields such as name.
	//Example value: {selector:added,desc:true,language:gr}
	SortingParameter *SortingParameter

	//WithTotalCount is a boolean parameter to optionally return total number of records in the X-Total-Count response header
	WithTotalCount bool
}

func NewListOptions(filters []Filter, paginationParameters *PaginationParameters, sortingParameter *SortingParameter, withTotalCount bool) *ListOptions {
	return &ListOptions{Filters: filters, PaginationParameters: paginationParameters, SortingParameter: sortingParameter, WithTotalCount: withTotalCount}
}

type PaginationParameters struct {
	Skip uint //skip n records
	Take uint //amount to take
}

func NewPaginationParameters(skip, take uint) *PaginationParameters {
	return &PaginationParameters{Skip: skip, Take: take}
}

type Filter struct {
	//ColumnFilter array represents a filter for a specific column. For example ["status","startswith","ACTIVE"] or ["group_id","<=","2"].
	//The possible filtering operations are: "=", ">=", "<=", "contains" and "startswith".
	ColumnFilter *ColumnFilter
	//Operand represents the connection of filters. Supported operands: and,or.
	Operand string
}

type ColumnFilter struct {
	Selector  string
	Operation string
	Value     interface{}
}

func NewColumnFilter(selector, operation string, value interface{}) *ColumnFilter {
	return &ColumnFilter{Selector: selector, Operation: operation, Value: value}
}

func NewFilter(columnFilter *ColumnFilter, operand string) *Filter {
	return &Filter{ColumnFilter: columnFilter, Operand: operand}
}

type SortingParameter struct {
	//JSON field. For description: description_plain or description_html.
	Selector string `json:"selector" example:"id"`
	//Descending or Ascending direction switch
	Desc     bool   `json:"desc"`
	Language string `json:"language" example:"gr"`
}

func NewSortingParameter(selector string, desc bool, language string) *SortingParameter {
	return &SortingParameter{Selector: selector, Desc: desc, Language: language}
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts *ListOptions) (string, error) {
	if opts == nil {
		return "", nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	q := u.Query()

	//apply pagination
	if opts.PaginationParameters != nil {
		q.Add("skip", strconv.Itoa(int(opts.PaginationParameters.Skip)))
		q.Add("take", strconv.Itoa(int(opts.PaginationParameters.Take)))
	}

	//apply sorting
	if opts.SortingParameter != nil {
		bytes, err := json.Marshal(opts.SortingParameter)
		if err != nil {
			return "", errors.Wrap(err, "could not parse sorting parameter")
		}
		q.Add("sort", string(bytes))
	}

	//apply filters
	if opts.Filters != nil {
		filters, err := parseFilters(opts.Filters)
		if err != nil {
			return "", errors.Wrap(err, "could not parse filtering parameter")
		}
		q.Add("filter", filters)
	}

	if opts.WithTotalCount {
		q.Add("withTotalCount", "true")
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}

func parseFilters(filters []Filter) (string, error) {
	var f []interface{}
	for _, filterOrOperator := range filters {
		if err := validateColumnFilterOperation(filterOrOperator.ColumnFilter.Operation); err != nil {
			return "", err
		}
		cf := filterOrOperator.ColumnFilter
		f = append(f, [3]interface{}{cf.Selector, cf.Operation, cf.Value})
		if filterOrOperator.Operand != "" {
			if err := validateFilteringOperand(filterOrOperator.Operand); err != nil {
				return "", err
			}
			f = append(f, filterOrOperator.Operand)
		}
	}
	bytes, err := json.Marshal(f)
	if err != nil {
		return "", errors.Wrap(err, "could not parse filtering parameter")
	}
	return string(bytes), nil
}
func validateColumnFilterOperation(op interface{}) error {
	okValues := []string{"=", ">=", "<=", "contains", "startswith"}
	for _, v := range okValues {
		if op == v {
			return nil
		}
	}
	return fmt.Errorf("unknown column filter operation %s, accepted values are %s", op, okValues)
}

func validateFilteringOperand(op interface{}) error {
	okValues := []string{"and", "or"}
	for _, v := range okValues {
		if op == v {
			return nil
		}
	}
	return fmt.Errorf("unknown filtering operand %s, accepted values are %s", op, okValues)
}
