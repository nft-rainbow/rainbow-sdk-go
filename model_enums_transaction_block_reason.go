/*
Rainbow-API

The responses of the open api in swagger focus on the data field rather than the code and the message fields

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rainbowsdk

import (
	"encoding/json"
	"fmt"
)

// EnumsTransactionBlockReason the model 'EnumsTransactionBlockReason'
type EnumsTransactionBlockReason int32

// List of enums.TransactionBlockReason
const (
	ENUMSTRANSACTIONBLOCKREASON_TX_BLOCK_REASON_DB_ERR EnumsTransactionBlockReason = 1
	ENUMSTRANSACTIONBLOCKREASON_TX_BLOCK_REASON_SPONSORING EnumsTransactionBlockReason = 2
	ENUMSTRANSACTIONBLOCKREASON_TX_BLOCK_REASON_IO_ERR EnumsTransactionBlockReason = 3
	ENUMSTRANSACTIONBLOCKREASON_TX_BLOCK_REASON_NOT_ENOUGH_CASH EnumsTransactionBlockReason = 4
	ENUMSTRANSACTIONBLOCKREASON_TX_BLOCK_REASON_FAILED_GET_SPONSOR_INFO EnumsTransactionBlockReason = 5
	ENUMSTRANSACTIONBLOCKREASON_TX_BLOCK_REASON_SPONSOR_NOT_ENOUGH EnumsTransactionBlockReason = 6
)

// All allowed values of EnumsTransactionBlockReason enum
var AllowedEnumsTransactionBlockReasonEnumValues = []EnumsTransactionBlockReason{
	1,
	2,
	3,
	4,
	5,
	6,
}

func (v *EnumsTransactionBlockReason) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumsTransactionBlockReason(value)
	for _, existing := range AllowedEnumsTransactionBlockReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumsTransactionBlockReason", value)
}

// NewEnumsTransactionBlockReasonFromValue returns a pointer to a valid EnumsTransactionBlockReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumsTransactionBlockReasonFromValue(v int32) (*EnumsTransactionBlockReason, error) {
	ev := EnumsTransactionBlockReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumsTransactionBlockReason: valid values are %v", v, AllowedEnumsTransactionBlockReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumsTransactionBlockReason) IsValid() bool {
	for _, existing := range AllowedEnumsTransactionBlockReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enums.TransactionBlockReason value
func (v EnumsTransactionBlockReason) Ptr() *EnumsTransactionBlockReason {
	return &v
}

type NullableEnumsTransactionBlockReason struct {
	value *EnumsTransactionBlockReason
	isSet bool
}

func (v NullableEnumsTransactionBlockReason) Get() *EnumsTransactionBlockReason {
	return v.value
}

func (v *NullableEnumsTransactionBlockReason) Set(val *EnumsTransactionBlockReason) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumsTransactionBlockReason) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumsTransactionBlockReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumsTransactionBlockReason(val *EnumsTransactionBlockReason) *NullableEnumsTransactionBlockReason {
	return &NullableEnumsTransactionBlockReason{value: val, isSet: true}
}

func (v NullableEnumsTransactionBlockReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumsTransactionBlockReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

