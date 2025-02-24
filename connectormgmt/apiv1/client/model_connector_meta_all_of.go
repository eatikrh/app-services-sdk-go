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
)

// ConnectorMetaAllOf struct for ConnectorMetaAllOf
type ConnectorMetaAllOf struct {

	ResourceVersion *int64 `json:"resource_version,omitempty"`

}

// NewConnectorMetaAllOf instantiates a new ConnectorMetaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorMetaAllOf() *ConnectorMetaAllOf {
	this := ConnectorMetaAllOf{}
	return &this
}

// NewConnectorMetaAllOfWithDefaults instantiates a new ConnectorMetaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorMetaAllOfWithDefaults() *ConnectorMetaAllOf {
	this := ConnectorMetaAllOf{}


	return &this
}


// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *ConnectorMetaAllOf) GetResourceVersion() int64 {
	if o == nil || o.ResourceVersion == nil {
		var ret int64
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorMetaAllOf) GetResourceVersionOk() (*int64, bool) {
	if o == nil || o.ResourceVersion == nil {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *ConnectorMetaAllOf) HasResourceVersion() bool {
	if o != nil && o.ResourceVersion != nil {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given int64 and assigns it to the ResourceVersion field.
func (o *ConnectorMetaAllOf) SetResourceVersion(v int64) {
	o.ResourceVersion = &v
}


func (o ConnectorMetaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.ResourceVersion != nil {
		toSerialize["resource_version"] = o.ResourceVersion
	}
    
	return json.Marshal(toSerialize)
}

type NullableConnectorMetaAllOf struct {
	value *ConnectorMetaAllOf
	isSet bool
}

func (v NullableConnectorMetaAllOf) Get() *ConnectorMetaAllOf {
	return v.value
}

func (v *NullableConnectorMetaAllOf) Set(val *ConnectorMetaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorMetaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorMetaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorMetaAllOf(val *ConnectorMetaAllOf) *NullableConnectorMetaAllOf {
	return &NullableConnectorMetaAllOf{value: val, isSet: true}
}

func (v NullableConnectorMetaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorMetaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

