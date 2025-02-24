/*
 * Apicurio Registry API [v2]
 *
 * Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 
 *
 * API version: 2.1.0-SNAPSHOT
 * Contact: apicurio@lists.jboss.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryinstanceclient

import (
	"encoding/json"
)

// Rule struct for Rule
type Rule struct {

	Config string `json:"config"`

	Type *RuleType `json:"type,omitempty"`

}

// NewRule instantiates a new Rule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRule(config string) *Rule {
	this := Rule{}
	this.Config = config
	return &this
}

// NewRuleWithDefaults instantiates a new Rule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleWithDefaults() *Rule {
	this := Rule{}



	return &this
}


// GetConfig returns the Config field value
func (o *Rule) GetConfig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *Rule) GetConfigOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *Rule) SetConfig(v string) {
	o.Config = v
}


// GetType returns the Type field value if set, zero value otherwise.
func (o *Rule) GetType() RuleType {
	if o == nil || o.Type == nil {
		var ret RuleType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetTypeOk() (*RuleType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Rule) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given RuleType and assigns it to the Type field.
func (o *Rule) SetType(v RuleType) {
	o.Type = &v
}


func (o Rule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if true {
		toSerialize["config"] = o.Config
	}
    
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
    
	return json.Marshal(toSerialize)
}

type NullableRule struct {
	value *Rule
	isSet bool
}

func (v NullableRule) Get() *Rule {
	return v.value
}

func (v *NullableRule) Set(val *Rule) {
	v.value = val
	v.isSet = true
}

func (v NullableRule) IsSet() bool {
	return v.isSet
}

func (v *NullableRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRule(val *Rule) *NullableRule {
	return &NullableRule{value: val, isSet: true}
}

func (v NullableRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

