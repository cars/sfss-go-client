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

// checks if the PutRedfishV1SFSSAppRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutRedfishV1SFSSAppRequest{}

// PutRedfishV1SFSSAppRequest 
type PutRedfishV1SFSSAppRequest struct {
	// Username to access the remote server
	ImageServerUserName string `json:"ImageServerUserName"`
	// Password to access the image in the remote server
	ImageServerPassword string `json:"ImageServerPassword"`
	// Location of the SFSS image in the remote server; SFSS supports only IPv4 communication with remote servers
	ImageServerLocation string `json:"ImageServerLocation"`
	// Transport type used to copy the image from the remote repository; possible values include SCP, SFTP, HTTP, and HTTPS
	TransportType string `json:"TransportType"`
}

type _PutRedfishV1SFSSAppRequest PutRedfishV1SFSSAppRequest

// NewPutRedfishV1SFSSAppRequest instantiates a new PutRedfishV1SFSSAppRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutRedfishV1SFSSAppRequest(imageServerUserName string, imageServerPassword string, imageServerLocation string, transportType string) *PutRedfishV1SFSSAppRequest {
	this := PutRedfishV1SFSSAppRequest{}
	this.ImageServerUserName = imageServerUserName
	this.ImageServerPassword = imageServerPassword
	this.ImageServerLocation = imageServerLocation
	this.TransportType = transportType
	return &this
}

// NewPutRedfishV1SFSSAppRequestWithDefaults instantiates a new PutRedfishV1SFSSAppRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutRedfishV1SFSSAppRequestWithDefaults() *PutRedfishV1SFSSAppRequest {
	this := PutRedfishV1SFSSAppRequest{}
	return &this
}

// GetImageServerUserName returns the ImageServerUserName field value
func (o *PutRedfishV1SFSSAppRequest) GetImageServerUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageServerUserName
}

// GetImageServerUserNameOk returns a tuple with the ImageServerUserName field value
// and a boolean to check if the value has been set.
func (o *PutRedfishV1SFSSAppRequest) GetImageServerUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageServerUserName, true
}

// SetImageServerUserName sets field value
func (o *PutRedfishV1SFSSAppRequest) SetImageServerUserName(v string) {
	o.ImageServerUserName = v
}

// GetImageServerPassword returns the ImageServerPassword field value
func (o *PutRedfishV1SFSSAppRequest) GetImageServerPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageServerPassword
}

// GetImageServerPasswordOk returns a tuple with the ImageServerPassword field value
// and a boolean to check if the value has been set.
func (o *PutRedfishV1SFSSAppRequest) GetImageServerPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageServerPassword, true
}

// SetImageServerPassword sets field value
func (o *PutRedfishV1SFSSAppRequest) SetImageServerPassword(v string) {
	o.ImageServerPassword = v
}

// GetImageServerLocation returns the ImageServerLocation field value
func (o *PutRedfishV1SFSSAppRequest) GetImageServerLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageServerLocation
}

// GetImageServerLocationOk returns a tuple with the ImageServerLocation field value
// and a boolean to check if the value has been set.
func (o *PutRedfishV1SFSSAppRequest) GetImageServerLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageServerLocation, true
}

// SetImageServerLocation sets field value
func (o *PutRedfishV1SFSSAppRequest) SetImageServerLocation(v string) {
	o.ImageServerLocation = v
}

// GetTransportType returns the TransportType field value
func (o *PutRedfishV1SFSSAppRequest) GetTransportType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransportType
}

// GetTransportTypeOk returns a tuple with the TransportType field value
// and a boolean to check if the value has been set.
func (o *PutRedfishV1SFSSAppRequest) GetTransportTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransportType, true
}

// SetTransportType sets field value
func (o *PutRedfishV1SFSSAppRequest) SetTransportType(v string) {
	o.TransportType = v
}

func (o PutRedfishV1SFSSAppRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutRedfishV1SFSSAppRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ImageServerUserName"] = o.ImageServerUserName
	toSerialize["ImageServerPassword"] = o.ImageServerPassword
	toSerialize["ImageServerLocation"] = o.ImageServerLocation
	toSerialize["TransportType"] = o.TransportType
	return toSerialize, nil
}

func (o *PutRedfishV1SFSSAppRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ImageServerUserName",
		"ImageServerPassword",
		"ImageServerLocation",
		"TransportType",
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

	varPutRedfishV1SFSSAppRequest := _PutRedfishV1SFSSAppRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPutRedfishV1SFSSAppRequest)

	if err != nil {
		return err
	}

	*o = PutRedfishV1SFSSAppRequest(varPutRedfishV1SFSSAppRequest)

	return err
}

type NullablePutRedfishV1SFSSAppRequest struct {
	value *PutRedfishV1SFSSAppRequest
	isSet bool
}

func (v NullablePutRedfishV1SFSSAppRequest) Get() *PutRedfishV1SFSSAppRequest {
	return v.value
}

func (v *NullablePutRedfishV1SFSSAppRequest) Set(val *PutRedfishV1SFSSAppRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutRedfishV1SFSSAppRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutRedfishV1SFSSAppRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutRedfishV1SFSSAppRequest(val *PutRedfishV1SFSSAppRequest) *NullablePutRedfishV1SFSSAppRequest {
	return &NullablePutRedfishV1SFSSAppRequest{value: val, isSet: true}
}

func (v NullablePutRedfishV1SFSSAppRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutRedfishV1SFSSAppRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


