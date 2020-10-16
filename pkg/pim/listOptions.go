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
	PaginationParameters *PaginationParameters `json:"paginationParameters,omitempty"`

	//SortingParameter is a JSON object with 3 fields -
	//1) selector - field name, example value:added,
	//2) desc - short for descending, value: true or false,
	//3) language: example value: gr, is used to sort by translatable fields such as name.
	//Example value: {selector:added,desc:true,language:gr}
	SortingParameter *SortingParameter `json:"sortingParameters,omitempty"`

	//WithTotalCount is a boolean parameter to optionally return total number of records in the X-Total-Count response header
	WithTotalCount bool `json:"withTotalCount,omitempty"`
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
	//Selector represents the name for a specific field. For example status, group_id
	Selector string `json:"fieldName,omitempty"`
	//The possible filtering operations are: "=", ">=", "<=", "contains" and "startswith".
	Operation string `json:"operator,omitempty"`
	//Value is the filter value of any type
	Value interface{} `json:"value,omitempty"`
	//Operand represents the connection of the filter to the next filter. Supported operands: and,or.
	OperandBefore string `json:"operandBefore,omitempty"`
}

//NewFilter will validate the operation, operandAfter and return the filter
func NewFilter(selector, operation string, value interface{}, operandBefore string) (*Filter, error) {
	if operandBefore != "" {
		if err := validateFilteringOperand(operandBefore); err != nil {
			return nil, err
		}
	}
	if err := validateColumnFilterOperation(operation); err != nil {
		return nil, err
	}
	return &Filter{Selector: selector, Operation: operation, Value: value, OperandBefore: operandBefore}, nil
}

type SortingParameter struct {
	//JSON field. For description: description_plain or description_html.
	Selector string `json:"selector,omitempty" example:"id"`
	//Descending or Ascending direction switch
	Desc     bool   `json:"desc,omitempty"`
	Language string `json:"language,omitempty" example:"gr"`
}

func NewSortingParameter(selector string, desc bool, language string) *SortingParameter {
	return &SortingParameter{Selector: selector, Desc: desc, Language: language}
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts *ListOptions) (*url.URL, error) {
	if opts == nil {
		return nil, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return nil, err
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
			return nil, errors.Wrap(err, "could not parse sorting parameter")
		}
		q.Add("sort", string(bytes))
	}

	//apply filters
	if opts.Filters != nil {
		filters, err := parseFilters(opts.Filters)
		if err != nil {
			return nil, errors.Wrap(err, "could not parse filtering parameter")
		}
		q.Add("filter", filters)
	}

	if opts.WithTotalCount {
		q.Add("withTotalCount", "true")
	}

	u.RawQuery = q.Encode()
	return u, nil
}

//add semicolon separated list of IDs to the request
func addIDs(url *url.URL, parameterName string, recordIDs ...int) (*url.URL, error) {
	if len(recordIDs) > 100 {
		return nil, errors.New("limit reached: up to 100 ids are allowed in the request")
	}
	if len(recordIDs) == 0 {
		return url, nil
	}
	var recordIDsQueryString string
	for i, id := range recordIDs {
		if i == 0 {
			recordIDsQueryString += strconv.Itoa(id)
		} else {
			recordIDsQueryString += ";" + strconv.Itoa(id)
		}
	}
	q := url.Query()
	q.Add(parameterName, recordIDsQueryString)
	url.RawQuery = q.Encode()
	return url, nil
}

func parseFilters(filters []Filter) (string, error) {
	var f []interface{}
	for _, filterOrOperator := range filters {
		fo := filterOrOperator
		if err := validateColumnFilterOperation(fo.Operation); err != nil {
			return "", err
		}
		if fo.OperandBefore != "" {
			if err := validateFilteringOperand(fo.OperandBefore); err != nil {
				return "", err
			}
			f = append(f, fo.OperandBefore)
		}

		f = append(f, [3]interface{}{fo.Selector, fo.Operation, fo.Value})
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
