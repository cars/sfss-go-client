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

// checks if the LicensesGET type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicensesGET{}

// LicensesGET This model lists all the licenses that are installed in SFSS.
type LicensesGET struct {
	Licenses []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner `json:"Licenses"`
	// Number of licenses installed on the SFSS VM
	LicensesodataCount float32 `json:"Licenses@odata.count"`
	OdataId string `json:"@odata.id"`
	OdataContext string `json:"@odata.context"`
	OdataType string `json:"@odata.type"`
}

type _LicensesGET LicensesGET

// NewLicensesGET instantiates a new LicensesGET object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicensesGET(licenses []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner, licensesodataCount float32, odataId string, odataContext string, odataType string) *LicensesGET {
	this := LicensesGET{}
	this.Licenses = licenses
	this.LicensesodataCount = licensesodataCount
	this.OdataId = odataId
	this.OdataContext = odataContext
	this.OdataType = odataType
	return &this
}

// NewLicensesGETWithDefaults instantiates a new LicensesGET object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicensesGETWithDefaults() *LicensesGET {
	this := LicensesGET{}
	return &this
}

// GetLicenses returns the Licenses field value
func (o *LicensesGET) GetLicenses() []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner {
	if o == nil {
		var ret []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner
		return ret
	}

	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value
// and a boolean to check if the value has been set.
func (o *LicensesGET) GetLicensesOk() ([]GetRedfishV1SFSSAppLicenses200ResponseLicensesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Licenses, true
}

// SetLicenses sets field value
func (o *LicensesGET) SetLicenses(v []GetRedfishV1SFSSAppLicenses200ResponseLicensesInner) {
	o.Licenses = v
}

// GetLicensesodataCount returns the LicensesodataCount field value
func (o *LicensesGET) GetLicensesodataCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LicensesodataCount
}

// GetLicensesodataCountOk returns a tuple with the LicensesodataCount field value
// and a boolean to check if the value has been set.
func (o *LicensesGET) GetLicensesodataCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicensesodataCount, true
}

// SetLicensesodataCount sets field value
func (o *LicensesGET) SetLicensesodataCount(v float32) {
	o.LicensesodataCount = v
}

// GetOdataId returns the OdataId field value
func (o *LicensesGET) GetOdataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value
// and a boolean to check if the value has been set.
func (o *LicensesGET) GetOdataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataId, true
}

// SetOdataId sets field value
func (o *LicensesGET) SetOdataId(v string) {
	o.OdataId = v
}

// GetOdataContext returns the OdataContext field value
func (o *LicensesGET) GetOdataContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataContext
}

// GetOdataContextOk returns a tuple with the OdataContext field value
// and a boolean to check if the value has been set.
func (o *LicensesGET) GetOdataContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataContext, true
}

// SetOdataContext sets field value
func (o *LicensesGET) SetOdataContext(v string) {
	o.OdataContext = v
}

// GetOdataType returns the OdataType field value
func (o *LicensesGET) GetOdataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataType
}

// GetOdataTypeOk returns a tuple with the OdataType field value
// and a boolean to check if the value has been set.
func (o *LicensesGET) GetOdataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataType, true
}

// SetOdataType sets field value
func (o *LicensesGET) SetOdataType(v string) {
	o.OdataType = v
}

func (o LicensesGET) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicensesGET) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Licenses"] = o.Licenses
	toSerialize["Licenses@odata.count"] = o.LicensesodataCount
	toSerialize["@odata.id"] = o.OdataId
	toSerialize["@odata.context"] = o.OdataContext
	toSerialize["@odata.type"] = o.OdataType
	return toSerialize, nil
}

func (o *LicensesGET) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Licenses",
		"Licenses@odata.count",
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

	varLicensesGET := _LicensesGET{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLicensesGET)

	if err != nil {
		return err
	}

	*o = LicensesGET(varLicensesGET)

	return err
}

type NullableLicensesGET struct {
	value *LicensesGET
	isSet bool
}

func (v NullableLicensesGET) Get() *LicensesGET {
	return v.value
}

func (v *NullableLicensesGET) Set(val *LicensesGET) {
	v.value = val
	v.isSet = true
}

func (v NullableLicensesGET) IsSet() bool {
	return v.isSet
}

func (v *NullableLicensesGET) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicensesGET(val *LicensesGET) *NullableLicensesGET {
	return &NullableLicensesGET{value: val, isSet: true}
}

func (v NullableLicensesGET) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicensesGET) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


