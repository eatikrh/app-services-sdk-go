/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.8.1-SNAPSHOT
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
	"fmt"
)

// AclOperationFilter the model 'AclOperationFilter'
type AclOperationFilter string

// List of AclOperationFilter
const (
	ACLOPERATIONFILTER_ALL AclOperationFilter = "ALL"
	ACLOPERATIONFILTER_READ AclOperationFilter = "READ"
	ACLOPERATIONFILTER_WRITE AclOperationFilter = "WRITE"
	ACLOPERATIONFILTER_CREATE AclOperationFilter = "CREATE"
	ACLOPERATIONFILTER_DELETE AclOperationFilter = "DELETE"
	ACLOPERATIONFILTER_ALTER AclOperationFilter = "ALTER"
	ACLOPERATIONFILTER_DESCRIBE AclOperationFilter = "DESCRIBE"
	ACLOPERATIONFILTER_DESCRIBE_CONFIGS AclOperationFilter = "DESCRIBE_CONFIGS"
	ACLOPERATIONFILTER_ALTER_CONFIGS AclOperationFilter = "ALTER_CONFIGS"
	ACLOPERATIONFILTER_ANY AclOperationFilter = "ANY"
)

var allowedAclOperationFilterEnumValues = []AclOperationFilter{
	"ALL",
	"READ",
	"WRITE",
	"CREATE",
	"DELETE",
	"ALTER",
	"DESCRIBE",
	"DESCRIBE_CONFIGS",
	"ALTER_CONFIGS",
	"ANY",
}

func (v *AclOperationFilter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AclOperationFilter(value)
	for _, existing := range allowedAclOperationFilterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AclOperationFilter", value)
}

// NewAclOperationFilterFromValue returns a pointer to a valid AclOperationFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAclOperationFilterFromValue(v string) (*AclOperationFilter, error) {
	ev := AclOperationFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AclOperationFilter: valid values are %v", v, allowedAclOperationFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AclOperationFilter) IsValid() bool {
	for _, existing := range allowedAclOperationFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AclOperationFilter value
func (v AclOperationFilter) Ptr() *AclOperationFilter {
	return &v
}

type NullableAclOperationFilter struct {
	value *AclOperationFilter
	isSet bool
}

func (v NullableAclOperationFilter) Get() *AclOperationFilter {
	return v.value
}

func (v *NullableAclOperationFilter) Set(val *AclOperationFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAclOperationFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAclOperationFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAclOperationFilter(val *AclOperationFilter) *NullableAclOperationFilter {
	return &NullableAclOperationFilter{value: val, isSet: true}
}

func (v NullableAclOperationFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAclOperationFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

