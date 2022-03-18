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

// TopicOrderKey the model 'TopicOrderKey'
type TopicOrderKey string

// List of TopicOrderKey
const (
	TOPICORDERKEY_NAME TopicOrderKey = "name"
	TOPICORDERKEY_PARTITIONS TopicOrderKey = "partitions"
	TOPICORDERKEY_RETENTION_MS TopicOrderKey = "retention.ms"
	TOPICORDERKEY_RETENTION_BYTES TopicOrderKey = "retention.bytes"
)

var allowedTopicOrderKeyEnumValues = []TopicOrderKey{
	"name",
	"partitions",
	"retention.ms",
	"retention.bytes",
}

func (v *TopicOrderKey) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TopicOrderKey(value)
	for _, existing := range allowedTopicOrderKeyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TopicOrderKey", value)
}

// NewTopicOrderKeyFromValue returns a pointer to a valid TopicOrderKey
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTopicOrderKeyFromValue(v string) (*TopicOrderKey, error) {
	ev := TopicOrderKey(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TopicOrderKey: valid values are %v", v, allowedTopicOrderKeyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TopicOrderKey) IsValid() bool {
	for _, existing := range allowedTopicOrderKeyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopicOrderKey value
func (v TopicOrderKey) Ptr() *TopicOrderKey {
	return &v
}

type NullableTopicOrderKey struct {
	value *TopicOrderKey
	isSet bool
}

func (v NullableTopicOrderKey) Get() *TopicOrderKey {
	return v.value
}

func (v *NullableTopicOrderKey) Set(val *TopicOrderKey) {
	v.value = val
	v.isSet = true
}

func (v NullableTopicOrderKey) IsSet() bool {
	return v.isSet
}

func (v *NullableTopicOrderKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopicOrderKey(val *TopicOrderKey) *NullableTopicOrderKey {
	return &NullableTopicOrderKey{value: val, isSet: true}
}

func (v NullableTopicOrderKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopicOrderKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

