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

// checks if the GetIDRedfishV1SFSSApp200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIDRedfishV1SFSSApp200Response{}

// GetIDRedfishV1SFSSApp200Response 
type GetIDRedfishV1SFSSApp200Response struct {
	// Location of the SFSS image in the remote server; SFSS supports only IPv4 communication with remote servers
	ImageServerLocation string `json:"ImageServerLocation"`
	// A detailed status message about the add image operation
	StatusMessage string `json:"StatusMessage"`
	// Password to access the image in the remote server
	ImageServerPassword string `json:"ImageServerPassword"`
	// Status of the add image operation
	Status string `json:"Status"`
	// Transport type used to copy the image from the remote server; possible values include SCP, SFTP, HTTP, and HTTPS
	TransportType string `json:"TransportType"`
	// Username to access the remote server
	ImageServerUserName string `json:"ImageServerUserName"`
	// Version of the SFSS image
	Version string `json:"Version"`
	OdataId string `json:"@odata.id"`
	OdataType string `json:"@odata.type"`
	OdataContext string `json:"@odata.context"`
}

type _GetIDRedfishV1SFSSApp200Response GetIDRedfishV1SFSSApp200Response

// NewGetIDRedfishV1SFSSApp200Response instantiates a new GetIDRedfishV1SFSSApp200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIDRedfishV1SFSSApp200Response(imageServerLocation string, statusMessage string, imageServerPassword string, status string, transportType string, imageServerUserName string, version string, odataId string, odataType string, odataContext string) *GetIDRedfishV1SFSSApp200Response {
	this := GetIDRedfishV1SFSSApp200Response{}
	this.ImageServerLocation = imageServerLocation
	this.StatusMessage = statusMessage
	this.ImageServerPassword = imageServerPassword
	this.Status = status
	this.TransportType = transportType
	this.ImageServerUserName = imageServerUserName
	this.Version = version
	this.OdataId = odataId
	this.OdataType = odataType
	this.OdataContext = odataContext
	return &this
}

// NewGetIDRedfishV1SFSSApp200ResponseWithDefaults instantiates a new GetIDRedfishV1SFSSApp200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIDRedfishV1SFSSApp200ResponseWithDefaults() *GetIDRedfishV1SFSSApp200Response {
	this := GetIDRedfishV1SFSSApp200Response{}
	return &this
}

// GetImageServerLocation returns the ImageServerLocation field value
func (o *GetIDRedfishV1SFSSApp200Response) GetImageServerLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageServerLocation
}

// GetImageServerLocationOk returns a tuple with the ImageServerLocation field value
// and a boolean to check if the value has been set.
func (o *GetIDRedfishV1SFSSApp200Response) GetImageServerLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageServerLocation, true
}

// SetImageServerLocation sets field value
func (o *GetIDRedfishV1SFSSApp200Response) SetImageServerLocation(v string) {
	o.ImageServerLocation = v
}

// GetStatusMessage returns the StatusMessage field value
func (o *GetIDRedfishV1SFSSApp200Response) GetStatusMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value
// and a boolean to check if the value has been set.
func (o *GetIDRedfishV1SFSSApp200Response) GetStatusMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusMessage, true
}

// SetStatusMessage sets field value
func (o *GetIDRedfishV1SFSSApp200Response) SetStatusMessage(v string) {
	o.StatusMessage = v
}

// GetImageServerPassword returns the ImageServerPassword field value
func (o *GetIDRedfishV1SFSSApp200Response) GetImageServerPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageServerPassword
}

// GetImageServerPasswordOk returns a tuple with the ImageServerPassword field value
// and a boolean to check if the value has been set.
func (o *GetIDRedfishV1SFSSApp200Response) GetImageServerPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageServerPassword, true
}

// SetImageServerPassword sets field value
func (o *GetIDRedfishV1SFSSApp200Response) SetImageServerPassword(v string) {
	o.ImageServerPassword = v
}

// GetStatus returns the Status field value
func (o *GetIDRedfishV1SFSSApp200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetIDRedfishV1SFSSApp200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetIDRedfishV1SFSSApp200Response) SetStatus(v string) {
	o.Status = v
}

// GetTransportType returns the TransportType field value
func (o *GetIDRedfishV1SFSSApp200Response) GetTransportType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransportType
}

// GetTransportTypeOk returns a tuple with the TransportType field value
// and a boolean to check if the value has been set.
func (o *GetIDRedfishV1SFSSApp200Response) GetTransportTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransportType, true
}

// SetTransportType sets field value
func (o *GetIDRedfishV1SFSSApp200Response) SetTransportType(v string) {
	o.TransportType = v
}

// GetImageServerUserName returns the ImageServerUserName field value
func (o *GetIDRedfishV1SFSSApp200Response) GetImageServerUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageServerUserName
}

// GetImageServerUserNameOk returns a tuple with the ImageServerUserName field value
// and a boolean to check if the value has been set.
func (o *GetIDRedfishV1SFSSApp200Response) GetImageServerUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageServerUserName, true
}

// SetImageServerUserName sets field value
func (o *GetIDRedfishV1SFSSApp200Response) SetImageServerUserName(v string) {
	o.ImageServerUserName = v
}

// GetVersion returns the Version field value
func (o *GetIDRedfishV1SFSSApp200Response) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetIDRedfishV1SFSSApp200Response) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetIDRedfishV1SFSSApp200Response) SetVersion(v string) {
	o.Version = v
}

// GetOdataId returns the OdataId field value
func (o *GetIDRedfishV1SFSSApp200Response) GetOdataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value
// and a boolean to check if the value has been set.
func (o *GetIDRedfishV1SFSSApp200Response) GetOdataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataId, true
}

// SetOdataId sets field value
func (o *GetIDRedfishV1SFSSApp200Response) SetOdataId(v string) {
	o.OdataId = v
}

// GetOdataType returns the OdataType field value
func (o *GetIDRedfishV1SFSSApp200Response) GetOdataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataType
}

// GetOdataTypeOk returns a tuple with the OdataType field value
// and a boolean to check if the value has been set.
func (o *GetIDRedfishV1SFSSApp200Response) GetOdataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataType, true
}

// SetOdataType sets field value
func (o *GetIDRedfishV1SFSSApp200Response) SetOdataType(v string) {
	o.OdataType = v
}

// GetOdataContext returns the OdataContext field value
func (o *GetIDRedfishV1SFSSApp200Response) GetOdataContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataContext
}

// GetOdataContextOk returns a tuple with the OdataContext field value
// and a boolean to check if the value has been set.
func (o *GetIDRedfishV1SFSSApp200Response) GetOdataContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataContext, true
}

// SetOdataContext sets field value
func (o *GetIDRedfishV1SFSSApp200Response) SetOdataContext(v string) {
	o.OdataContext = v
}

func (o GetIDRedfishV1SFSSApp200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIDRedfishV1SFSSApp200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ImageServerLocation"] = o.ImageServerLocation
	toSerialize["StatusMessage"] = o.StatusMessage
	toSerialize["ImageServerPassword"] = o.ImageServerPassword
	toSerialize["Status"] = o.Status
	toSerialize["TransportType"] = o.TransportType
	toSerialize["ImageServerUserName"] = o.ImageServerUserName
	toSerialize["Version"] = o.Version
	toSerialize["@odata.id"] = o.OdataId
	toSerialize["@odata.type"] = o.OdataType
	toSerialize["@odata.context"] = o.OdataContext
	return toSerialize, nil
}

func (o *GetIDRedfishV1SFSSApp200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ImageServerLocation",
		"StatusMessage",
		"ImageServerPassword",
		"Status",
		"TransportType",
		"ImageServerUserName",
		"Version",
		"@odata.id",
		"@odata.type",
		"@odata.context",
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

	varGetIDRedfishV1SFSSApp200Response := _GetIDRedfishV1SFSSApp200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetIDRedfishV1SFSSApp200Response)

	if err != nil {
		return err
	}

	*o = GetIDRedfishV1SFSSApp200Response(varGetIDRedfishV1SFSSApp200Response)

	return err
}

type NullableGetIDRedfishV1SFSSApp200Response struct {
	value *GetIDRedfishV1SFSSApp200Response
	isSet bool
}

func (v NullableGetIDRedfishV1SFSSApp200Response) Get() *GetIDRedfishV1SFSSApp200Response {
	return v.value
}

func (v *NullableGetIDRedfishV1SFSSApp200Response) Set(val *GetIDRedfishV1SFSSApp200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIDRedfishV1SFSSApp200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIDRedfishV1SFSSApp200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIDRedfishV1SFSSApp200Response(val *GetIDRedfishV1SFSSApp200Response) *NullableGetIDRedfishV1SFSSApp200Response {
	return &NullableGetIDRedfishV1SFSSApp200Response{value: val, isSet: true}
}

func (v NullableGetIDRedfishV1SFSSApp200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIDRedfishV1SFSSApp200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


