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

// checks if the GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response{}

// GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response 
type GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response struct {
	// Defines a set of authentication servers (local and remote) configured in SFSS and the order in which SFSS looks up user information in these servers
	AuthenticationSequence []string `json:"AuthenticationSequence"`
}

type _GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response

// NewGetRedfishV1SFSSAppAuthenticationSequenceEnums200Response instantiates a new GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRedfishV1SFSSAppAuthenticationSequenceEnums200Response(authenticationSequence []string) *GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response {
	this := GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response{}
	this.AuthenticationSequence = authenticationSequence
	return &this
}

// NewGetRedfishV1SFSSAppAuthenticationSequenceEnums200ResponseWithDefaults instantiates a new GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRedfishV1SFSSAppAuthenticationSequenceEnums200ResponseWithDefaults() *GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response {
	this := GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response{}
	return &this
}

// GetAuthenticationSequence returns the AuthenticationSequence field value
func (o *GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response) GetAuthenticationSequence() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuthenticationSequence
}

// GetAuthenticationSequenceOk returns a tuple with the AuthenticationSequence field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response) GetAuthenticationSequenceOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationSequence, true
}

// SetAuthenticationSequence sets field value
func (o *GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response) SetAuthenticationSequence(v []string) {
	o.AuthenticationSequence = v
}

func (o GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AuthenticationSequence"] = o.AuthenticationSequence
	return toSerialize, nil
}

func (o *GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetRedfishV1SFSSAppAuthenticationSequenceEnums200Response := _GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRedfishV1SFSSAppAuthenticationSequenceEnums200Response)

	if err != nil {
		return err
	}

	*o = GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response(varGetRedfishV1SFSSAppAuthenticationSequenceEnums200Response)

	return err
}

type NullableGetRedfishV1SFSSAppAuthenticationSequenceEnums200Response struct {
	value *GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response
	isSet bool
}

func (v NullableGetRedfishV1SFSSAppAuthenticationSequenceEnums200Response) Get() *GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response {
	return v.value
}

func (v *NullableGetRedfishV1SFSSAppAuthenticationSequenceEnums200Response) Set(val *GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRedfishV1SFSSAppAuthenticationSequenceEnums200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRedfishV1SFSSAppAuthenticationSequenceEnums200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRedfishV1SFSSAppAuthenticationSequenceEnums200Response(val *GetRedfishV1SFSSAppAuthenticationSequenceEnums200Response) *NullableGetRedfishV1SFSSAppAuthenticationSequenceEnums200Response {
	return &NullableGetRedfishV1SFSSAppAuthenticationSequenceEnums200Response{value: val, isSet: true}
}

func (v NullableGetRedfishV1SFSSAppAuthenticationSequenceEnums200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRedfishV1SFSSAppAuthenticationSequenceEnums200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


