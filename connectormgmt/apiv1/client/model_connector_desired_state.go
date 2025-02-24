/*
 * Connector Service Fleet Manager
 *
 * Connector Service Fleet Manager is a Rest API to manage connectors.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
	"fmt"
)

// ConnectorDesiredState the model 'ConnectorDesiredState'
type ConnectorDesiredState string

// List of ConnectorDesiredState
const (
	CONNECTORDESIREDSTATE_READY ConnectorDesiredState = "ready"
	CONNECTORDESIREDSTATE_STOPPED ConnectorDesiredState = "stopped"
	CONNECTORDESIREDSTATE_DELETED ConnectorDesiredState = "deleted"
)

var allowedConnectorDesiredStateEnumValues = []ConnectorDesiredState{
	"ready",
	"stopped",
	"deleted",
}

func (v *ConnectorDesiredState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConnectorDesiredState(value)
	for _, existing := range allowedConnectorDesiredStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConnectorDesiredState", value)
}

// NewConnectorDesiredStateFromValue returns a pointer to a valid ConnectorDesiredState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConnectorDesiredStateFromValue(v string) (*ConnectorDesiredState, error) {
	ev := ConnectorDesiredState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConnectorDesiredState: valid values are %v", v, allowedConnectorDesiredStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConnectorDesiredState) IsValid() bool {
	for _, existing := range allowedConnectorDesiredStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectorDesiredState value
func (v ConnectorDesiredState) Ptr() *ConnectorDesiredState {
	return &v
}

type NullableConnectorDesiredState struct {
	value *ConnectorDesiredState
	isSet bool
}

func (v NullableConnectorDesiredState) Get() *ConnectorDesiredState {
	return v.value
}

func (v *NullableConnectorDesiredState) Set(val *ConnectorDesiredState) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorDesiredState) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorDesiredState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorDesiredState(val *ConnectorDesiredState) *NullableConnectorDesiredState {
	return &NullableConnectorDesiredState{value: val, isSet: true}
}

func (v NullableConnectorDesiredState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorDesiredState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

