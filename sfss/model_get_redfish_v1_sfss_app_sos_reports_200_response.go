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

// checks if the GetRedfishV1SFSSAppSosReports200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRedfishV1SFSSAppSosReports200Response{}

// GetRedfishV1SFSSAppSosReports200Response 
type GetRedfishV1SFSSAppSosReports200Response struct {
	// A collection of system information that includes configuration details and diagnostic information
	SosReports []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner `json:"SosReports"`
	//  Number of SOS reports downloaded
	SosReportsodataCount float32 `json:"SosReports@odata.count"`
	OdataId string `json:"@odata.id"`
	OdataContext string `json:"@odata.context"`
	OdataType string `json:"@odata.type"`
}

type _GetRedfishV1SFSSAppSosReports200Response GetRedfishV1SFSSAppSosReports200Response

// NewGetRedfishV1SFSSAppSosReports200Response instantiates a new GetRedfishV1SFSSAppSosReports200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRedfishV1SFSSAppSosReports200Response(sosReports []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner, sosReportsodataCount float32, odataId string, odataContext string, odataType string) *GetRedfishV1SFSSAppSosReports200Response {
	this := GetRedfishV1SFSSAppSosReports200Response{}
	this.SosReports = sosReports
	this.SosReportsodataCount = sosReportsodataCount
	this.OdataId = odataId
	this.OdataContext = odataContext
	this.OdataType = odataType
	return &this
}

// NewGetRedfishV1SFSSAppSosReports200ResponseWithDefaults instantiates a new GetRedfishV1SFSSAppSosReports200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRedfishV1SFSSAppSosReports200ResponseWithDefaults() *GetRedfishV1SFSSAppSosReports200Response {
	this := GetRedfishV1SFSSAppSosReports200Response{}
	return &this
}

// GetSosReports returns the SosReports field value
func (o *GetRedfishV1SFSSAppSosReports200Response) GetSosReports() []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner {
	if o == nil {
		var ret []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner
		return ret
	}

	return o.SosReports
}

// GetSosReportsOk returns a tuple with the SosReports field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppSosReports200Response) GetSosReportsOk() ([]GetRedfishV1SFSSAppLicenses200ResponseLicensesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.SosReports, true
}

// SetSosReports sets field value
func (o *GetRedfishV1SFSSAppSosReports200Response) SetSosReports(v []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner) {
	o.SosReports = v
}

// GetSosReportsodataCount returns the SosReportsodataCount field value
func (o *GetRedfishV1SFSSAppSosReports200Response) GetSosReportsodataCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SosReportsodataCount
}

// GetSosReportsodataCountOk returns a tuple with the SosReportsodataCount field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppSosReports200Response) GetSosReportsodataCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SosReportsodataCount, true
}

// SetSosReportsodataCount sets field value
func (o *GetRedfishV1SFSSAppSosReports200Response) SetSosReportsodataCount(v float32) {
	o.SosReportsodataCount = v
}

// GetOdataId returns the OdataId field value
func (o *GetRedfishV1SFSSAppSosReports200Response) GetOdataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppSosReports200Response) GetOdataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataId, true
}

// SetOdataId sets field value
func (o *GetRedfishV1SFSSAppSosReports200Response) SetOdataId(v string) {
	o.OdataId = v
}

// GetOdataContext returns the OdataContext field value
func (o *GetRedfishV1SFSSAppSosReports200Response) GetOdataContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataContext
}

// GetOdataContextOk returns a tuple with the OdataContext field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppSosReports200Response) GetOdataContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataContext, true
}

// SetOdataContext sets field value
func (o *GetRedfishV1SFSSAppSosReports200Response) SetOdataContext(v string) {
	o.OdataContext = v
}

// GetOdataType returns the OdataType field value
func (o *GetRedfishV1SFSSAppSosReports200Response) GetOdataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataType
}

// GetOdataTypeOk returns a tuple with the OdataType field value
// and a boolean to check if the value has been set.
func (o *GetRedfishV1SFSSAppSosReports200Response) GetOdataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataType, true
}

// SetOdataType sets field value
func (o *GetRedfishV1SFSSAppSosReports200Response) SetOdataType(v string) {
	o.OdataType = v
}

func (o GetRedfishV1SFSSAppSosReports200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRedfishV1SFSSAppSosReports200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["SosReports"] = o.SosReports
	toSerialize["SosReports@odata.count"] = o.SosReportsodataCount
	toSerialize["@odata.id"] = o.OdataId
	toSerialize["@odata.context"] = o.OdataContext
	toSerialize["@odata.type"] = o.OdataType
	return toSerialize, nil
}

func (o *GetRedfishV1SFSSAppSosReports200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"SosReports",
		"SosReports@odata.count",
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

	varGetRedfishV1SFSSAppSosReports200Response := _GetRedfishV1SFSSAppSosReports200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRedfishV1SFSSAppSosReports200Response)

	if err != nil {
		return err
	}

	*o = GetRedfishV1SFSSAppSosReports200Response(varGetRedfishV1SFSSAppSosReports200Response)

	return err
}

type NullableGetRedfishV1SFSSAppSosReports200Response struct {
	value *GetRedfishV1SFSSAppSosReports200Response
	isSet bool
}

func (v NullableGetRedfishV1SFSSAppSosReports200Response) Get() *GetRedfishV1SFSSAppSosReports200Response {
	return v.value
}

func (v *NullableGetRedfishV1SFSSAppSosReports200Response) Set(val *GetRedfishV1SFSSAppSosReports200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRedfishV1SFSSAppSosReports200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRedfishV1SFSSAppSosReports200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRedfishV1SFSSAppSosReports200Response(val *GetRedfishV1SFSSAppSosReports200Response) *NullableGetRedfishV1SFSSAppSosReports200Response {
	return &NullableGetRedfishV1SFSSAppSosReports200Response{value: val, isSet: true}
}

func (v NullableGetRedfishV1SFSSAppSosReports200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRedfishV1SFSSAppSosReports200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


