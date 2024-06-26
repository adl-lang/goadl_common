// Code generated by goadlc v3 - DO NOT EDIT.
package ui

import (
	"github.com/adl-lang/goadl_rt/v3/sys/types"
)

type FormGroupKey = string

type FormGroups struct {
	_FormGroups
}

type _FormGroups struct {
	DefaultKey FormGroupKey                       `json:"defaultKey"`
	Labels     []types.Pair[FormGroupKey, string] `json:"labels"`
}

func MakeAll_FormGroups(
	defaultkey FormGroupKey,
	labels []types.Pair[FormGroupKey, string],
) FormGroups {
	return FormGroups{
		_FormGroups{
			DefaultKey: defaultkey,
			Labels:     labels,
		},
	}
}

func Make_FormGroups(
	defaultkey FormGroupKey,
	labels []types.Pair[FormGroupKey, string],
) FormGroups {
	ret := FormGroups{
		_FormGroups{
			DefaultKey: defaultkey,
			Labels:     labels,
		},
	}
	return ret
}

type FormLabel = string

type ValidRegex struct {
	_ValidRegex
}

type _ValidRegex struct {
	Regex       string `json:"regex"`
	Description string `json:"description"`
	ReturnGroup int8   `json:"returnGroup"`
}

func MakeAll_ValidRegex(
	regex string,
	description string,
	returngroup int8,
) ValidRegex {
	return ValidRegex{
		_ValidRegex{
			Regex:       regex,
			Description: description,
			ReturnGroup: returngroup,
		},
	}
}

func Make_ValidRegex(
	regex string,
	description string,
) ValidRegex {
	ret := ValidRegex{
		_ValidRegex{
			Regex:       regex,
			Description: description,
			ReturnGroup: ((*ValidRegex)(nil)).Default_returnGroup(),
		},
	}
	return ret
}

func (*ValidRegex) Default_returnGroup() int8 {
	return 0
}

type ValidValues struct {
	_ValidValues
}

type _ValidValues struct {
	Values      []string `json:"values"`
	Description string   `json:"description"`
}

func MakeAll_ValidValues(
	values []string,
	description string,
) ValidValues {
	return ValidValues{
		_ValidValues{
			Values:      values,
			Description: description,
		},
	}
}

func Make_ValidValues(
	values []string,
	description string,
) ValidValues {
	ret := ValidValues{
		_ValidValues{
			Values:      values,
			Description: description,
		},
	}
	return ret
}
