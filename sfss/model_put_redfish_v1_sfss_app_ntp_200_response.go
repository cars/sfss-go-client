/*
SmartFabric Storage Software Application REST APIs

REST APIs used for managing SFSS application are captured in this section.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sfss

import (
	"encoding/json"
)

// checks if the PutRedfishV1SFSSAppNTP200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutRedfishV1SFSSAppNTP200Response{}

// PutRedfishV1SFSSAppNTP200Response struct for PutRedfishV1SFSSAppNTP200Response
type PutRedfishV1SFSSAppNTP200Response struct {
	// Status of the NTP service; possible values include Enable and Disable
	Status *string `json:"Status,omitempty"`
}

// NewPutRedfishV1SFSSAppNTP200Response instantiates a new PutRedfishV1SFSSAppNTP200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutRedfishV1SFSSAppNTP200Response() *PutRedfishV1SFSSAppNTP200Response {
	this := PutRedfishV1SFSSAppNTP200Response{}
	return &this
}

// NewPutRedfishV1SFSSAppNTP200ResponseWithDefaults instantiates a new PutRedfishV1SFSSAppNTP200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutRedfishV1SFSSAppNTP200ResponseWithDefaults() *PutRedfishV1SFSSAppNTP200Response {
	this := PutRedfishV1SFSSAppNTP200Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PutRedfishV1SFSSAppNTP200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutRedfishV1SFSSAppNTP200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PutRedfishV1SFSSAppNTP200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PutRedfishV1SFSSAppNTP200Response) SetStatus(v string) {
	o.Status = &v
}

func (o PutRedfishV1SFSSAppNTP200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutRedfishV1SFSSAppNTP200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	return toSerialize, nil
}

type NullablePutRedfishV1SFSSAppNTP200Response struct {
	value *PutRedfishV1SFSSAppNTP200Response
	isSet bool
}

func (v NullablePutRedfishV1SFSSAppNTP200Response) Get() *PutRedfishV1SFSSAppNTP200Response {
	return v.value
}

func (v *NullablePutRedfishV1SFSSAppNTP200Response) Set(val *PutRedfishV1SFSSAppNTP200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePutRedfishV1SFSSAppNTP200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePutRedfishV1SFSSAppNTP200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutRedfishV1SFSSAppNTP200Response(val *PutRedfishV1SFSSAppNTP200Response) *NullablePutRedfishV1SFSSAppNTP200Response {
	return &NullablePutRedfishV1SFSSAppNTP200Response{value: val, isSet: true}
}

func (v NullablePutRedfishV1SFSSAppNTP200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutRedfishV1SFSSAppNTP200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


