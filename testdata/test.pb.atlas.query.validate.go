// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/partitio/atlas-gentool/testdata/test.proto

package test

import options "github.com/infobloxopen/protoc-gen-atlas-query-validate/options"
import query "github.com/infobloxopen/atlas-app-toolkit/query"
import _ "github.com/infobloxopen/protoc-gen-atlas-validate/options"
import _ "github.com/infobloxopen/protoc-gen-preprocess/options"
import _ "github.com/partitio/protoc-gen-gorm/options"

// Reference imports to suppress errors if they are not otherwise used.

var TestMethodsRequireFilteringValidation = map[string]map[string]options.FilteringOption{}
var TestMethodsRequireSortingValidation = map[string][]string{}
var TestMethodsRequireFieldSelectionValidation = map[string][]string{}

func TestValidateFiltering(methodName string, f *query.Filtering) error {
	info, ok := TestMethodsRequireFilteringValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateFiltering(f, info)
}
func TestValidateSorting(methodName string, s *query.Sorting) error {
	info, ok := TestMethodsRequireSortingValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateSorting(s, info)
}
func TestValidateFieldSelection(methodName string, s *query.FieldSelection) error {
	info, ok := TestMethodsRequireFieldSelectionValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateFieldSelection(s, info)
}
