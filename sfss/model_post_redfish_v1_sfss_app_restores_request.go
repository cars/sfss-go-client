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

// checks if the PostRedfishV1SFSSAppRestoresRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostRedfishV1SFSSAppRestoresRequest{}

// PostRedfishV1SFSSAppRestoresRequest 
type PostRedfishV1SFSSAppRestoresRequest struct {
	// Remote server location where the backup file is available; only IPv4 communication is supported
	ImageServerLocation string `json:"ImageServerLocation"`
	// Password to access the remote server
	ImageServerPassword string `json:"ImageServerPassword"`
	// Transport type used to copy the backup file from the remote server; possible values include SCP, HTTP, HTTPS, and SFTP
	TransportType string `json:"TransportType"`
	// Username to access the remote server
	ImageServerUserName string `json:"ImageServerUserName"`
}

type _PostRedfishV1SFSSAppRestoresRequest PostRedfishV1SFSSAppRestoresRequest

// NewPostRedfishV1SFSSAppRestoresRequest instantiates a new PostRedfishV1SFSSAppRestoresRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostRedfishV1SFSSAppRestoresRequest(imageServerLocation string, imageServerPassword string, transportType string, imageServerUserName string) *PostRedfishV1SFSSAppRestoresRequest {
	this := PostRedfishV1SFSSAppRestoresRequest{}
	this.ImageServerLocation = imageServerLocation
	this.ImageServerPassword = imageServerPassword
	this.TransportType = transportType
	this.ImageServerUserName = imageServerUserName
	return &this
}

// NewPostRedfishV1SFSSAppRestoresRequestWithDefaults instantiates a new PostRedfishV1SFSSAppRestoresRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostRedfishV1SFSSAppRestoresRequestWithDefaults() *PostRedfishV1SFSSAppRestoresRequest {
	this := PostRedfishV1SFSSAppRestoresRequest{}
	return &this
}

// GetImageServerLocation returns the ImageServerLocation field value
func (o *PostRedfishV1SFSSAppRestoresRequest) GetImageServerLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageServerLocation
}

// GetImageServerLocationOk returns a tuple with the ImageServerLocation field value
// and a boolean to check if the value has been set.
func (o *PostRedfishV1SFSSAppRestoresRequest) GetImageServerLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageServerLocation, true
}

// SetImageServerLocation sets field value
func (o *PostRedfishV1SFSSAppRestoresRequest) SetImageServerLocation(v string) {
	o.ImageServerLocation = v
}

// GetImageServerPassword returns the ImageServerPassword field value
func (o *PostRedfishV1SFSSAppRestoresRequest) GetImageServerPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageServerPassword
}

// GetImageServerPasswordOk returns a tuple with the ImageServerPassword field value
// and a boolean to check if the value has been set.
func (o *PostRedfishV1SFSSAppRestoresRequest) GetImageServerPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageServerPassword, true
}

// SetImageServerPassword sets field value
func (o *PostRedfishV1SFSSAppRestoresRequest) SetImageServerPassword(v string) {
	o.ImageServerPassword = v
}

// GetTransportType returns the TransportType field value
func (o *PostRedfishV1SFSSAppRestoresRequest) GetTransportType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransportType
}

// GetTransportTypeOk returns a tuple with the TransportType field value
// and a boolean to check if the value has been set.
func (o *PostRedfishV1SFSSAppRestoresRequest) GetTransportTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransportType, true
}

// SetTransportType sets field value
func (o *PostRedfishV1SFSSAppRestoresRequest) SetTransportType(v string) {
	o.TransportType = v
}

// GetImageServerUserName returns the ImageServerUserName field value
func (o *PostRedfishV1SFSSAppRestoresRequest) GetImageServerUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageServerUserName
}

// GetImageServerUserNameOk returns a tuple with the ImageServerUserName field value
// and a boolean to check if the value has been set.
func (o *PostRedfishV1SFSSAppRestoresRequest) GetImageServerUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageServerUserName, true
}

// SetImageServerUserName sets field value
func (o *PostRedfishV1SFSSAppRestoresRequest) SetImageServerUserName(v string) {
	o.ImageServerUserName = v
}

func (o PostRedfishV1SFSSAppRestoresRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostRedfishV1SFSSAppRestoresRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ImageServerLocation"] = o.ImageServerLocation
	toSerialize["ImageServerPassword"] = o.ImageServerPassword
	toSerialize["TransportType"] = o.TransportType
	toSerialize["ImageServerUserName"] = o.ImageServerUserName
	return toSerialize, nil
}

func (o *PostRedfishV1SFSSAppRestoresRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ImageServerLocation",
		"ImageServerPassword",
		"TransportType",
		"ImageServerUserName",
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

	varPostRedfishV1SFSSAppRestoresRequest := _PostRedfishV1SFSSAppRestoresRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostRedfishV1SFSSAppRestoresRequest)

	if err != nil {
		return err
	}

	*o = PostRedfishV1SFSSAppRestoresRequest(varPostRedfishV1SFSSAppRestoresRequest)

	return err
}

type NullablePostRedfishV1SFSSAppRestoresRequest struct {
	value *PostRedfishV1SFSSAppRestoresRequest
	isSet bool
}

func (v NullablePostRedfishV1SFSSAppRestoresRequest) Get() *PostRedfishV1SFSSAppRestoresRequest {
	return v.value
}

func (v *NullablePostRedfishV1SFSSAppRestoresRequest) Set(val *PostRedfishV1SFSSAppRestoresRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostRedfishV1SFSSAppRestoresRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostRedfishV1SFSSAppRestoresRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostRedfishV1SFSSAppRestoresRequest(val *PostRedfishV1SFSSAppRestoresRequest) *NullablePostRedfishV1SFSSAppRestoresRequest {
	return &NullablePostRedfishV1SFSSAppRestoresRequest{value: val, isSet: true}
}

func (v NullablePostRedfishV1SFSSAppRestoresRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostRedfishV1SFSSAppRestoresRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


