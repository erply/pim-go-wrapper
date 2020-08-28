package pim

import (
	"encoding/json"
	"fmt"
	"github.com/pingcap/log"
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
}

type PaginationParameters struct {
	Skip uint //skip n records
	Take uint //amount to take
}

type Filter struct {
	//ColumnFilter array represents a filter for a specific column. For example ["status","startswith","ACTIVE"] or ["group_id","<=","2"].
	//The possible filtering operations are: "=", ">=", "<=", "contains" and "startswith".
	ColumnFilter [3]interface{}
	//Operand represents the connection of filters. Supported operands: and,or.
	Operand string
}

type SortingParameter struct {
	//JSON field. For description: description_plain or description_html.
	Selector string `json:"selector" example:"id"`
	//Descending or Ascending direction switch
	Desc     bool   `json:"desc"`
	Language string `json:"language" example:"gr"`
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts *ListOptions) (string, error) {
	if opts == nil {
		return "", nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
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
			return s, errors.Wrap(err, "could not parse sorting parameter")
		}
		log.Warn(string(bytes))
		q.Add("sort", string(bytes))
	}

	//apply filters
	if opts.Filters != nil {
		var f []interface{}
		for _, filterOrOperator := range opts.Filters {
			if err := validateColumnFilterOperation(filterOrOperator.ColumnFilter[1]); err != nil {
				return "", err
			}
			f = append(f, filterOrOperator.ColumnFilter)
			if filterOrOperator.Operand != "" {
				if err := validateFilteringOperand(filterOrOperator.Operand); err != nil {
					return "", err
				}
				f = append(f, filterOrOperator.Operand)
			}
		}
		bytes, err := json.Marshal(f)
		if err != nil {
			return s, errors.Wrap(err, "could not parse filtering parameter")
		}
		log.Warn(string(bytes))
		q.Add("filter", string(bytes))
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
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
