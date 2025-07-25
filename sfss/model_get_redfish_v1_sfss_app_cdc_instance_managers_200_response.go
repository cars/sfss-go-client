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

// checks if the GetRedfishV1SFSSAppCDCInstanceManagers200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRedfishV1SFSSAppCDCInstanceManagers200Response{}

// GetRedfishV1SFSSAppCDCInstanceManagers200Response 
type GetRedfishV1SFSSAppCDCInstanceManagers200Response struct {
	// A set of CDC instances
	CDCInstanceManagers []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner `json:"CDCInstanceManagers"`
	// Number of CDC instances
	CDCInstanceManagersodataCount float32 `json:"CDCInstanceManagers@odata.count"`
	OdataId string `json:"@odata.id"`
	OdataContext string `json:"@odata.context"`
	OdataType string `json:"@odata.type"`
}

type _GetRedfishV1SFSSAppCDCInstanceManagers200Response GetRedfishV1SFSSAppCDCInstanceManagers200Response

// NewGetRedfishV1SFSSAppCDCInstanceManagers200Response instantiates a new GetRedfishV1SFSSAppCDCInstanceManagers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRedfishV1SFSSAppCDCInstanceManagers200Response(cDCInstanceManagers []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner, cDCInstanceManagersodataCount float32, odataId string, odataContext string, odataType string) *GetRedfishV1SFSSAppCDCInstanceManagers200Response {
	this := GetRedfishV1SFSSAppCDCInstanceManagers200Response{}
	this.CDCInstanceManagers = cDCInstanceManagers
	this.CDCInstanceManagersodataCount = cDCInstanceManagersodataCount
	this.OdataId = odataId
	this.OdataContext = odataContext
	this.OdataType = odataType
	return &this
}

// NewGetRedfishV1SFSSAppCDCInstanceManagers200ResponseWithDefaults instantiates a new GetRedfishV1SFSSAppCDCInstanceManagers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRedfishV1SFSSAppCDCInstanceManagers200ResponseWithDefaults() *GetRedfishV1SFSSAppCDCInstanceManagers200Response {
	this := GetRedfishV1SFSSAppCDCInstanceManagers200Response{}
	return &this
}

// GetCDCInstanceManagers returns the CDCInstanceManagers field value
func (o *GetRedfishV1SFSSAppCDCInstanceManagers200Response) GetCDCInstanceManagers() []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner {
	if o == nil {
		var ret []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner
		return ret
	}

	return o.CDCInstanceManagers
}

// GetCDCInstanceManagersOk returns a tuple with the CDCInstanceManagers field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppCDCInstanceManagers200Response) GetCDCInstanceManagersOk() ([]GetRedfishV1SFSSAppLicenses200ResponseLicensesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.CDCInstanceManagers, true
}

// SetCDCInstanceManagers sets field value
func (o *GetRedfishV1SFSSAppCDCInstanceManagers200Response) SetCDCInstanceManagers(v []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner) {
	o.CDCInstanceManagers = v
}

// GetCDCInstanceManagersodataCount returns the CDCInstanceManagersodataCount field value
func (o *GetRedfishV1SFSSAppCDCInstanceManagers200Response) GetCDCInstanceManagersodataCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CDCInstanceManagersodataCount
}

// GetCDCInstanceManagersodataCountOk returns a tuple with the CDCInstanceManagersodataCount field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppCDCInstanceManagers200Response) GetCDCInstanceManagersodataCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CDCInstanceManagersodataCount, true
}

// SetCDCInstanceManagersodataCount sets field value
func (o *GetRedfishV1SFSSAppCDCInstanceManagers200Response) SetCDCInstanceManagersodataCount(v float32) {
	o.CDCInstanceManagersodataCount = v
}

// GetOdataId returns the OdataId field value
func (o *GetRedfishV1SFSSAppCDCInstanceManagers200Response) GetOdataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppCDCInstanceManagers200Response) GetOdataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataId, true
}

// SetOdataId sets field value
func (o *GetRedfishV1SFSSAppCDCInstanceManagers200Response) SetOdataId(v string) {
	o.OdataId = v
}

// GetOdataContext returns the OdataContext field value
func (o *GetRedfishV1SFSSAppCDCInstanceManagers200Response) GetOdataContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataContext
}

// GetOdataContextOk returns a tuple with the OdataContext field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppCDCInstanceManagers200Response) GetOdataContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataContext, true
}

// SetOdataContext sets field value
func (o *GetRedfishV1SFSSAppCDCInstanceManagers200Response) SetOdataContext(v string) {
	o.OdataContext = v
}

// GetOdataType returns the OdataType field value
func (o *GetRedfishV1SFSSAppCDCInstanceManagers200Response) GetOdataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataType
}

// GetOdataTypeOk returns a tuple with the OdataType field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppCDCInstanceManagers200Response) GetOdataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataType, true
}

// SetOdataType sets field value
func (o *GetRedfishV1SFSSAppCDCInstanceManagers200Response) SetOdataType(v string) {
	o.OdataType = v
}

func (o GetRedfishV1SFSSAppCDCInstanceManagers200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRedfishV1SFSSAppCDCInstanceManagers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CDCInstanceManagers"] = o.CDCInstanceManagers
	toSerialize["CDCInstanceManagers@odata.count"] = o.CDCInstanceManagersodataCount
	toSerialize["@odata.id"] = o.OdataId
	toSerialize["@odata.context"] = o.OdataContext
	toSerialize["@odata.type"] = o.OdataType
	return toSerialize, nil
}

func (o *GetRedfishV1SFSSAppCDCInstanceManagers200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"CDCInstanceManagers",
		"CDCInstanceManagers@odata.count",
		"@odata.id",
		"@odata.context",
		"@odata.type",
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

	varGetRedfishV1SFSSAppCDCInstanceManagers200Response := _GetRedfishV1SFSSAppCDCInstanceManagers200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRedfishV1SFSSAppCDCInstanceManagers200Response)

	if err != nil {
		return err
	}

	*o = GetRedfishV1SFSSAppCDCInstanceManagers200Response(varGetRedfishV1SFSSAppCDCInstanceManagers200Response)

	return err
}

type NullableGetRedfishV1SFSSAppCDCInstanceManagers200Response struct {
	value *GetRedfishV1SFSSAppCDCInstanceManagers200Response
	isSet bool
}

func (v NullableGetRedfishV1SFSSAppCDCInstanceManagers200Response) Get() *GetRedfishV1SFSSAppCDCInstanceManagers200Response {
	return v.value
}

func (v *NullableGetRedfishV1SFSSAppCDCInstanceManagers200Response) Set(val *GetRedfishV1SFSSAppCDCInstanceManagers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRedfishV1SFSSAppCDCInstanceManagers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRedfishV1SFSSAppCDCInstanceManagers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRedfishV1SFSSAppCDCInstanceManagers200Response(val *GetRedfishV1SFSSAppCDCInstanceManagers200Response) *NullableGetRedfishV1SFSSAppCDCInstanceManagers200Response {
	return &NullableGetRedfishV1SFSSAppCDCInstanceManagers200Response{value: val, isSet: true}
}

func (v NullableGetRedfishV1SFSSAppCDCInstanceManagers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRedfishV1SFSSAppCDCInstanceManagers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


