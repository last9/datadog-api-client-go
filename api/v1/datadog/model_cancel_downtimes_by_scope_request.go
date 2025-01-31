/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// CancelDowntimesByScopeRequest Cancel downtimes according to scope.
type CancelDowntimesByScopeRequest struct {
	// The scope(s) to which the downtime applies. For example, `host:app2`. Provide multiple scopes as a comma-separated list like `env:dev,env:prod`. The resulting downtime applies to sources that matches ALL provided scopes (`env:dev` **AND** `env:prod`).
	Scope string `json:"scope"`
}

// NewCancelDowntimesByScopeRequest instantiates a new CancelDowntimesByScopeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelDowntimesByScopeRequest(scope string) *CancelDowntimesByScopeRequest {
	this := CancelDowntimesByScopeRequest{}
	this.Scope = scope
	return &this
}

// NewCancelDowntimesByScopeRequestWithDefaults instantiates a new CancelDowntimesByScopeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelDowntimesByScopeRequestWithDefaults() *CancelDowntimesByScopeRequest {
	this := CancelDowntimesByScopeRequest{}
	return &this
}

// GetScope returns the Scope field value
func (o *CancelDowntimesByScopeRequest) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *CancelDowntimesByScopeRequest) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *CancelDowntimesByScopeRequest) SetScope(v string) {
	o.Scope = v
}

func (o CancelDowntimesByScopeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

func (o *CancelDowntimesByScopeRequest) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Scope *string `json:"scope"`
	}{}
	all := struct {
		Scope string `json:"scope"}`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Scope == nil {
		return fmt.Errorf("Required field scope missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Scope = all.Scope
	return nil
}

type NullableCancelDowntimesByScopeRequest struct {
	value *CancelDowntimesByScopeRequest
	isSet bool
}

func (v NullableCancelDowntimesByScopeRequest) Get() *CancelDowntimesByScopeRequest {
	return v.value
}

func (v *NullableCancelDowntimesByScopeRequest) Set(val *CancelDowntimesByScopeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelDowntimesByScopeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelDowntimesByScopeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelDowntimesByScopeRequest(val *CancelDowntimesByScopeRequest) *NullableCancelDowntimesByScopeRequest {
	return &NullableCancelDowntimesByScopeRequest{value: val, isSet: true}
}

func (v NullableCancelDowntimesByScopeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelDowntimesByScopeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
