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

// checks if the GetRedfishV1SFSSAppLicensesLicenseId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRedfishV1SFSSAppLicensesLicenseId200Response{}

// GetRedfishV1SFSSAppLicensesLicenseId200Response 
type GetRedfishV1SFSSAppLicensesLicenseId200Response struct {
	// License service tag
	ServiceTag string `json:"ServiceTag"`
	// Total number of endpoints that the license supports
	TotalNumOfEndPoints float32 `json:"TotalNumOfEndPoints"`
	// License identifier
	Identifier string `json:"Identifier"`
	// Type of license; possible values include Evaluation, Enterprise, Partner, and Expansion licenses
	LicenseType string `json:"LicenseType"`
	// License expiry date
	LicenseExpiry string `json:"LicenseExpiry"`
	OdataId string `json:"@odata.id"`
	OdataType string `json:"@odata.type"`
	OdataContext string `json:"@odata.context"`
	// A unique NVMe Qualified Name (NQN) that is used to identify the SFSS VM
	DeviceId string `json:"DeviceId"`
}

type _GetRedfishV1SFSSAppLicensesLicenseId200Response GetRedfishV1SFSSAppLicensesLicenseId200Response

// NewGetRedfishV1SFSSAppLicensesLicenseId200Response instantiates a new GetRedfishV1SFSSAppLicensesLicenseId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRedfishV1SFSSAppLicensesLicenseId200Response(serviceTag string, totalNumOfEndPoints float32, identifier string, licenseType string, licenseExpiry string, odataId string, odataType string, odataContext string, deviceId string) *GetRedfishV1SFSSAppLicensesLicenseId200Response {
	this := GetRedfishV1SFSSAppLicensesLicenseId200Response{}
	this.ServiceTag = serviceTag
	this.TotalNumOfEndPoints = totalNumOfEndPoints
	this.Identifier = identifier
	this.LicenseType = licenseType
	this.LicenseExpiry = licenseExpiry
	this.OdataId = odataId
	this.OdataType = odataType
	this.OdataContext = odataContext
	this.DeviceId = deviceId
	return &this
}

// NewGetRedfishV1SFSSAppLicensesLicenseId200ResponseWithDefaults instantiates a new GetRedfishV1SFSSAppLicensesLicenseId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRedfishV1SFSSAppLicensesLicenseId200ResponseWithDefaults() *GetRedfishV1SFSSAppLicensesLicenseId200Response {
	this := GetRedfishV1SFSSAppLicensesLicenseId200Response{}
	return &this
}

// GetServiceTag returns the ServiceTag field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetServiceTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceTag
}

// GetServiceTagOk returns a tuple with the ServiceTag field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetServiceTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceTag, true
}

// SetServiceTag sets field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) SetServiceTag(v string) {
	o.ServiceTag = v
}

// GetTotalNumOfEndPoints returns the TotalNumOfEndPoints field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetTotalNumOfEndPoints() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalNumOfEndPoints
}

// GetTotalNumOfEndPointsOk returns a tuple with the TotalNumOfEndPoints field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetTotalNumOfEndPointsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalNumOfEndPoints, true
}

// SetTotalNumOfEndPoints sets field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) SetTotalNumOfEndPoints(v float32) {
	o.TotalNumOfEndPoints = v
}

// GetIdentifier returns the Identifier field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) SetIdentifier(v string) {
	o.Identifier = v
}

// GetLicenseType returns the LicenseType field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetLicenseType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetLicenseTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseType, true
}

// SetLicenseType sets field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) SetLicenseType(v string) {
	o.LicenseType = v
}

// GetLicenseExpiry returns the LicenseExpiry field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetLicenseExpiry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseExpiry
}

// GetLicenseExpiryOk returns a tuple with the LicenseExpiry field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetLicenseExpiryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseExpiry, true
}

// SetLicenseExpiry sets field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) SetLicenseExpiry(v string) {
	o.LicenseExpiry = v
}

// GetOdataId returns the OdataId field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetOdataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetOdataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataId, true
}

// SetOdataId sets field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) SetOdataId(v string) {
	o.OdataId = v
}

// GetOdataType returns the OdataType field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetOdataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataType
}

// GetOdataTypeOk returns a tuple with the OdataType field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetOdataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataType, true
}

// SetOdataType sets field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) SetOdataType(v string) {
	o.OdataType = v
}

// GetOdataContext returns the OdataContext field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetOdataContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataContext
}

// GetOdataContextOk returns a tuple with the OdataContext field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetOdataContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataContext, true
}

// SetOdataContext sets field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) SetOdataContext(v string) {
	o.OdataContext = v
}

// GetDeviceId returns the DeviceId field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) SetDeviceId(v string) {
	o.DeviceId = v
}

func (o GetRedfishV1SFSSAppLicensesLicenseId200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRedfishV1SFSSAppLicensesLicenseId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ServiceTag"] = o.ServiceTag
	toSerialize["TotalNumOfEndPoints"] = o.TotalNumOfEndPoints
	toSerialize["Identifier"] = o.Identifier
	toSerialize["LicenseType"] = o.LicenseType
	toSerialize["LicenseExpiry"] = o.LicenseExpiry
	toSerialize["@odata.id"] = o.OdataId
	toSerialize["@odata.type"] = o.OdataType
	toSerialize["@odata.context"] = o.OdataContext
	toSerialize["DeviceId"] = o.DeviceId
	return toSerialize, nil
}

func (o *GetRedfishV1SFSSAppLicensesLicenseId200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ServiceTag",
		"TotalNumOfEndPoints",
		"Identifier",
		"LicenseType",
		"LicenseExpiry",
		"@odata.id",
		"@odata.type",
		"@odata.context",
		"DeviceId",
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

	varGetRedfishV1SFSSAppLicensesLicenseId200Response := _GetRedfishV1SFSSAppLicensesLicenseId200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRedfishV1SFSSAppLicensesLicenseId200Response)

	if err != nil {
		return err
	}

	*o = GetRedfishV1SFSSAppLicensesLicenseId200Response(varGetRedfishV1SFSSAppLicensesLicenseId200Response)

	return err
}

type NullableGetRedfishV1SFSSAppLicensesLicenseId200Response struct {
	value *GetRedfishV1SFSSAppLicensesLicenseId200Response
	isSet bool
}

func (v NullableGetRedfishV1SFSSAppLicensesLicenseId200Response) Get() *GetRedfishV1SFSSAppLicensesLicenseId200Response {
	return v.value
}

func (v *NullableGetRedfishV1SFSSAppLicensesLicenseId200Response) Set(val *GetRedfishV1SFSSAppLicensesLicenseId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRedfishV1SFSSAppLicensesLicenseId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRedfishV1SFSSAppLicensesLicenseId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRedfishV1SFSSAppLicensesLicenseId200Response(val *GetRedfishV1SFSSAppLicensesLicenseId200Response) *NullableGetRedfishV1SFSSAppLicensesLicenseId200Response {
	return &NullableGetRedfishV1SFSSAppLicensesLicenseId200Response{value: val, isSet: true}
}

func (v NullableGetRedfishV1SFSSAppLicensesLicenseId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRedfishV1SFSSAppLicensesLicenseId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


