/*
SmartFabric Storage Software Application REST APIs

REST APIs used for managing SFSS application are captured in this section.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sfss

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PostRedfishV1SFSSAppAuthenticationSequenceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostRedfishV1SFSSAppAuthenticationSequenceRequest{}

// PostRedfishV1SFSSAppAuthenticationSequenceRequest 
type PostRedfishV1SFSSAppAuthenticationSequenceRequest struct {
	// Defines a set of authentication servers (local and remote) configured in SFSS and the order in which SFSS looks up user information in these servers
	AuthenticationSequence []string `json:"AuthenticationSequence"`
}

type _PostRedfishV1SFSSAppAuthenticationSequenceRequest PostRedfishV1SFSSAppAuthenticationSequenceRequest

// NewPostRedfishV1SFSSAppAuthenticationSequenceRequest instantiates a new PostRedfishV1SFSSAppAuthenticationSequenceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostRedfishV1SFSSAppAuthenticationSequenceRequest(authenticationSequence []string) *PostRedfishV1SFSSAppAuthenticationSequenceRequest {
	this := PostRedfishV1SFSSAppAuthenticationSequenceRequest{}
	this.AuthenticationSequence = authenticationSequence
	return &this
}

// NewPostRedfishV1SFSSAppAuthenticationSequenceRequestWithDefaults instantiates a new PostRedfishV1SFSSAppAuthenticationSequenceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostRedfishV1SFSSAppAuthenticationSequenceRequestWithDefaults() *PostRedfishV1SFSSAppAuthenticationSequenceRequest {
	this := PostRedfishV1SFSSAppAuthenticationSequenceRequest{}
	return &this
}

// GetAuthenticationSequence returns the AuthenticationSequence field value
func (o *PostRedfishV1SFSSAppAuthenticationSequenceRequest) GetAuthenticationSequence() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuthenticationSequence
}

// GetAuthenticationSequenceOk returns a tuple with the AuthenticationSequence field value
// and a boolean to check if the value has been set.
func (o *PostRedfishV1SFSSAppAuthenticationSequenceRequest) GetAuthenticationSequenceOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationSequence, true
}

// SetAuthenticationSequence sets field value
func (o *PostRedfishV1SFSSAppAuthenticationSequenceRequest) SetAuthenticationSequence(v []string) {
	o.AuthenticationSequence = v
}

func (o PostRedfishV1SFSSAppAuthenticationSequenceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostRedfishV1SFSSAppAuthenticationSequenceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AuthenticationSequence"] = o.AuthenticationSequence
	return toSerialize, nil
}

func (o *PostRedfishV1SFSSAppAuthenticationSequenceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"AuthenticationSequence",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostRedfishV1SFSSAppAuthenticationSequenceRequest := _PostRedfishV1SFSSAppAuthenticationSequenceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostRedfishV1SFSSAppAuthenticationSequenceRequest)

	if err != nil {
		return err
	}

	*o = PostRedfishV1SFSSAppAuthenticationSequenceRequest(varPostRedfishV1SFSSAppAuthenticationSequenceRequest)

	return err
}

type NullablePostRedfishV1SFSSAppAuthenticationSequenceRequest struct {
	value *PostRedfishV1SFSSAppAuthenticationSequenceRequest
	isSet bool
}

func (v NullablePostRedfishV1SFSSAppAuthenticationSequenceRequest) Get() *PostRedfishV1SFSSAppAuthenticationSequenceRequest {
	return v.value
}

func (v *NullablePostRedfishV1SFSSAppAuthenticationSequenceRequest) Set(val *PostRedfishV1SFSSAppAuthenticationSequenceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostRedfishV1SFSSAppAuthenticationSequenceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostRedfishV1SFSSAppAuthenticationSequenceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostRedfishV1SFSSAppAuthenticationSequenceRequest(val *PostRedfishV1SFSSAppAuthenticationSequenceRequest) *NullablePostRedfishV1SFSSAppAuthenticationSequenceRequest {
	return &NullablePostRedfishV1SFSSAppAuthenticationSequenceRequest{value: val, isSet: true}
}

func (v NullablePostRedfishV1SFSSAppAuthenticationSequenceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostRedfishV1SFSSAppAuthenticationSequenceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


