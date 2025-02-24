/*
 * Service Accounts API Documentation
 *
 * This is the API documentation for Service Accounts
 *
 * API version: 1.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceaccountsclient

import (
	"encoding/json"
)

// InlineResponse400CauseStackTrace struct for InlineResponse400CauseStackTrace
type InlineResponse400CauseStackTrace struct {

	MethodName *string `json:"methodName,omitempty"`

	FileName *string `json:"fileName,omitempty"`

	LineNumber *int32 `json:"lineNumber,omitempty"`

	ClassName *string `json:"className,omitempty"`

	NativeMethod *bool `json:"nativeMethod,omitempty"`

}

// NewInlineResponse400CauseStackTrace instantiates a new InlineResponse400CauseStackTrace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse400CauseStackTrace() *InlineResponse400CauseStackTrace {
	this := InlineResponse400CauseStackTrace{}
	return &this
}

// NewInlineResponse400CauseStackTraceWithDefaults instantiates a new InlineResponse400CauseStackTrace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse400CauseStackTraceWithDefaults() *InlineResponse400CauseStackTrace {
	this := InlineResponse400CauseStackTrace{}






	return &this
}


// GetMethodName returns the MethodName field value if set, zero value otherwise.
func (o *InlineResponse400CauseStackTrace) GetMethodName() string {
	if o == nil || o.MethodName == nil {
		var ret string
		return ret
	}
	return *o.MethodName
}

// GetMethodNameOk returns a tuple with the MethodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400CauseStackTrace) GetMethodNameOk() (*string, bool) {
	if o == nil || o.MethodName == nil {
		return nil, false
	}
	return o.MethodName, true
}

// HasMethodName returns a boolean if a field has been set.
func (o *InlineResponse400CauseStackTrace) HasMethodName() bool {
	if o != nil && o.MethodName != nil {
		return true
	}

	return false
}

// SetMethodName gets a reference to the given string and assigns it to the MethodName field.
func (o *InlineResponse400CauseStackTrace) SetMethodName(v string) {
	o.MethodName = &v
}


// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *InlineResponse400CauseStackTrace) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400CauseStackTrace) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *InlineResponse400CauseStackTrace) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *InlineResponse400CauseStackTrace) SetFileName(v string) {
	o.FileName = &v
}


// GetLineNumber returns the LineNumber field value if set, zero value otherwise.
func (o *InlineResponse400CauseStackTrace) GetLineNumber() int32 {
	if o == nil || o.LineNumber == nil {
		var ret int32
		return ret
	}
	return *o.LineNumber
}

// GetLineNumberOk returns a tuple with the LineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400CauseStackTrace) GetLineNumberOk() (*int32, bool) {
	if o == nil || o.LineNumber == nil {
		return nil, false
	}
	return o.LineNumber, true
}

// HasLineNumber returns a boolean if a field has been set.
func (o *InlineResponse400CauseStackTrace) HasLineNumber() bool {
	if o != nil && o.LineNumber != nil {
		return true
	}

	return false
}

// SetLineNumber gets a reference to the given int32 and assigns it to the LineNumber field.
func (o *InlineResponse400CauseStackTrace) SetLineNumber(v int32) {
	o.LineNumber = &v
}


// GetClassName returns the ClassName field value if set, zero value otherwise.
func (o *InlineResponse400CauseStackTrace) GetClassName() string {
	if o == nil || o.ClassName == nil {
		var ret string
		return ret
	}
	return *o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400CauseStackTrace) GetClassNameOk() (*string, bool) {
	if o == nil || o.ClassName == nil {
		return nil, false
	}
	return o.ClassName, true
}

// HasClassName returns a boolean if a field has been set.
func (o *InlineResponse400CauseStackTrace) HasClassName() bool {
	if o != nil && o.ClassName != nil {
		return true
	}

	return false
}

// SetClassName gets a reference to the given string and assigns it to the ClassName field.
func (o *InlineResponse400CauseStackTrace) SetClassName(v string) {
	o.ClassName = &v
}


// GetNativeMethod returns the NativeMethod field value if set, zero value otherwise.
func (o *InlineResponse400CauseStackTrace) GetNativeMethod() bool {
	if o == nil || o.NativeMethod == nil {
		var ret bool
		return ret
	}
	return *o.NativeMethod
}

// GetNativeMethodOk returns a tuple with the NativeMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400CauseStackTrace) GetNativeMethodOk() (*bool, bool) {
	if o == nil || o.NativeMethod == nil {
		return nil, false
	}
	return o.NativeMethod, true
}

// HasNativeMethod returns a boolean if a field has been set.
func (o *InlineResponse400CauseStackTrace) HasNativeMethod() bool {
	if o != nil && o.NativeMethod != nil {
		return true
	}

	return false
}

// SetNativeMethod gets a reference to the given bool and assigns it to the NativeMethod field.
func (o *InlineResponse400CauseStackTrace) SetNativeMethod(v bool) {
	o.NativeMethod = &v
}


func (o InlineResponse400CauseStackTrace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.MethodName != nil {
		toSerialize["methodName"] = o.MethodName
	}
    
	if o.FileName != nil {
		toSerialize["fileName"] = o.FileName
	}
    
	if o.LineNumber != nil {
		toSerialize["lineNumber"] = o.LineNumber
	}
    
	if o.ClassName != nil {
		toSerialize["className"] = o.ClassName
	}
    
	if o.NativeMethod != nil {
		toSerialize["nativeMethod"] = o.NativeMethod
	}
    
	return json.Marshal(toSerialize)
}

type NullableInlineResponse400CauseStackTrace struct {
	value *InlineResponse400CauseStackTrace
	isSet bool
}

func (v NullableInlineResponse400CauseStackTrace) Get() *InlineResponse400CauseStackTrace {
	return v.value
}

func (v *NullableInlineResponse400CauseStackTrace) Set(val *InlineResponse400CauseStackTrace) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse400CauseStackTrace) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse400CauseStackTrace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse400CauseStackTrace(val *InlineResponse400CauseStackTrace) *NullableInlineResponse400CauseStackTrace {
	return &NullableInlineResponse400CauseStackTrace{value: val, isSet: true}
}

func (v NullableInlineResponse400CauseStackTrace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse400CauseStackTrace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

