//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcegraph

const (
	moduleName    = "armresourcegraph"
	moduleVersion = "v0.2.0"
)

// AuthorizationScopeFilter - Defines what level of authorization resources should be returned based on the which subscriptions
// and management groups are passed as scopes.
type AuthorizationScopeFilter string

const (
	AuthorizationScopeFilterAtScopeAndBelow      AuthorizationScopeFilter = "AtScopeAndBelow"
	AuthorizationScopeFilterAtScopeAndAbove      AuthorizationScopeFilter = "AtScopeAndAbove"
	AuthorizationScopeFilterAtScopeExact         AuthorizationScopeFilter = "AtScopeExact"
	AuthorizationScopeFilterAtScopeAboveAndBelow AuthorizationScopeFilter = "AtScopeAboveAndBelow"
)

// PossibleAuthorizationScopeFilterValues returns the possible values for the AuthorizationScopeFilter const type.
func PossibleAuthorizationScopeFilterValues() []AuthorizationScopeFilter {
	return []AuthorizationScopeFilter{
		AuthorizationScopeFilterAtScopeAndBelow,
		AuthorizationScopeFilterAtScopeAndAbove,
		AuthorizationScopeFilterAtScopeExact,
		AuthorizationScopeFilterAtScopeAboveAndBelow,
	}
}

// ToPtr returns a *AuthorizationScopeFilter pointing to the current value.
func (c AuthorizationScopeFilter) ToPtr() *AuthorizationScopeFilter {
	return &c
}

// ColumnDataType - Data type of a column in a table.
type ColumnDataType string

const (
	ColumnDataTypeString   ColumnDataType = "string"
	ColumnDataTypeInteger  ColumnDataType = "integer"
	ColumnDataTypeNumber   ColumnDataType = "number"
	ColumnDataTypeBoolean  ColumnDataType = "boolean"
	ColumnDataTypeObject   ColumnDataType = "object"
	ColumnDataTypeDatetime ColumnDataType = "datetime"
)

// PossibleColumnDataTypeValues returns the possible values for the ColumnDataType const type.
func PossibleColumnDataTypeValues() []ColumnDataType {
	return []ColumnDataType{
		ColumnDataTypeString,
		ColumnDataTypeInteger,
		ColumnDataTypeNumber,
		ColumnDataTypeBoolean,
		ColumnDataTypeObject,
		ColumnDataTypeDatetime,
	}
}

// ToPtr returns a *ColumnDataType pointing to the current value.
func (c ColumnDataType) ToPtr() *ColumnDataType {
	return &c
}

// FacetSortOrder - The sorting order by the selected column (count by default).
type FacetSortOrder string

const (
	FacetSortOrderAsc  FacetSortOrder = "asc"
	FacetSortOrderDesc FacetSortOrder = "desc"
)

// PossibleFacetSortOrderValues returns the possible values for the FacetSortOrder const type.
func PossibleFacetSortOrderValues() []FacetSortOrder {
	return []FacetSortOrder{
		FacetSortOrderAsc,
		FacetSortOrderDesc,
	}
}

// ToPtr returns a *FacetSortOrder pointing to the current value.
func (c FacetSortOrder) ToPtr() *FacetSortOrder {
	return &c
}

// ResultFormat - Defines in which format query result returned.
type ResultFormat string

const (
	ResultFormatTable       ResultFormat = "table"
	ResultFormatObjectArray ResultFormat = "objectArray"
)

// PossibleResultFormatValues returns the possible values for the ResultFormat const type.
func PossibleResultFormatValues() []ResultFormat {
	return []ResultFormat{
		ResultFormatTable,
		ResultFormatObjectArray,
	}
}

// ToPtr returns a *ResultFormat pointing to the current value.
func (c ResultFormat) ToPtr() *ResultFormat {
	return &c
}

// ResultTruncated - Indicates whether the query results are truncated.
type ResultTruncated string

const (
	ResultTruncatedTrue  ResultTruncated = "true"
	ResultTruncatedFalse ResultTruncated = "false"
)

// PossibleResultTruncatedValues returns the possible values for the ResultTruncated const type.
func PossibleResultTruncatedValues() []ResultTruncated {
	return []ResultTruncated{
		ResultTruncatedTrue,
		ResultTruncatedFalse,
	}
}

// ToPtr returns a *ResultTruncated pointing to the current value.
func (c ResultTruncated) ToPtr() *ResultTruncated {
	return &c
}
