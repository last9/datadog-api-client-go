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

// IFrameWidgetDefinition The iframe widget allows you to embed a portion of any other web page on your dashboard. Only available on FREE layout dashboards.
type IFrameWidgetDefinition struct {
	Type IFrameWidgetDefinitionType `json:"type"`
	// URL of the iframe.
	Url string `json:"url"`
}

// NewIFrameWidgetDefinition instantiates a new IFrameWidgetDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIFrameWidgetDefinition(type_ IFrameWidgetDefinitionType, url string) *IFrameWidgetDefinition {
	this := IFrameWidgetDefinition{}
	this.Type = type_
	this.Url = url
	return &this
}

// NewIFrameWidgetDefinitionWithDefaults instantiates a new IFrameWidgetDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIFrameWidgetDefinitionWithDefaults() *IFrameWidgetDefinition {
	this := IFrameWidgetDefinition{}
	var type_ IFrameWidgetDefinitionType = IFRAMEWIDGETDEFINITIONTYPE_IFRAME
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *IFrameWidgetDefinition) GetType() IFrameWidgetDefinitionType {
	if o == nil {
		var ret IFrameWidgetDefinitionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IFrameWidgetDefinition) GetTypeOk() (*IFrameWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IFrameWidgetDefinition) SetType(v IFrameWidgetDefinitionType) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *IFrameWidgetDefinition) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IFrameWidgetDefinition) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *IFrameWidgetDefinition) SetUrl(v string) {
	o.Url = v
}

func (o IFrameWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

func (o *IFrameWidgetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Type *IFrameWidgetDefinitionType `json:"type"`
		Url  *string                     `json:"url"`
	}{}
	all := struct {
		Type IFrameWidgetDefinitionType `json:"type"}`
		Url  string                     `json:"url"}`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Type == nil {
		return fmt.Errorf("Required field type missing")
	}
	if required.Url == nil {
		return fmt.Errorf("Required field url missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Type = all.Type
	o.Url = all.Url
	return nil
}

type NullableIFrameWidgetDefinition struct {
	value *IFrameWidgetDefinition
	isSet bool
}

func (v NullableIFrameWidgetDefinition) Get() *IFrameWidgetDefinition {
	return v.value
}

func (v *NullableIFrameWidgetDefinition) Set(val *IFrameWidgetDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableIFrameWidgetDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableIFrameWidgetDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIFrameWidgetDefinition(val *IFrameWidgetDefinition) *NullableIFrameWidgetDefinition {
	return &NullableIFrameWidgetDefinition{value: val, isSet: true}
}

func (v NullableIFrameWidgetDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIFrameWidgetDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
