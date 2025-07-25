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

// checks if the FoundationalConfigsGET type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FoundationalConfigsGET{}

// FoundationalConfigsGET This model lists the foundational configuration information.
type FoundationalConfigsGET struct {
	// Lists all the foundational configurations from the SFSS VM
	FoundationalConfigs []interface{} `json:"FoundationalConfigs"`
	// Number of foundational configurations present in the SFSS VM
	FoundationalConfigsodataCount float32 `json:"FoundationalConfigs@odata.count"`
	OdataId string `json:"@odata.id"`
	OdataContext string `json:"@odata.context"`
	OdataType string `json:"@odata.type"`
}

type _FoundationalConfigsGET FoundationalConfigsGET

// NewFoundationalConfigsGET instantiates a new FoundationalConfigsGET object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFoundationalConfigsGET(foundationalConfigs []interface{}, foundationalConfigsodataCount float32, odataId string, odataContext string, odataType string) *FoundationalConfigsGET {
	this := FoundationalConfigsGET{}
	this.FoundationalConfigs = foundationalConfigs
	this.FoundationalConfigsodataCount = foundationalConfigsodataCount
	this.OdataId = odataId
	this.OdataContext = odataContext
	this.OdataType = odataType
	return &this
}

// NewFoundationalConfigsGETWithDefaults instantiates a new FoundationalConfigsGET object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFoundationalConfigsGETWithDefaults() *FoundationalConfigsGET {
	this := FoundationalConfigsGET{}
	return &this
}

// GetFoundationalConfigs returns the FoundationalConfigs field value
func (o *FoundationalConfigsGET) GetFoundationalConfigs() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.FoundationalConfigs
}

// GetFoundationalConfigsOk returns a tuple with the FoundationalConfigs field value
// and a boolean to check if the value has been set.
func (o *FoundationalConfigsGET) GetFoundationalConfigsOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.FoundationalConfigs, true
}

// SetFoundationalConfigs sets field value
func (o *FoundationalConfigsGET) SetFoundationalConfigs(v []interface{}) {
	o.FoundationalConfigs = v
}

// GetFoundationalConfigsodataCount returns the FoundationalConfigsodataCount field value
func (o *FoundationalConfigsGET) GetFoundationalConfigsodataCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.FoundationalConfigsodataCount
}

// GetFoundationalConfigsodataCountOk returns a tuple with the FoundationalConfigsodataCount field value
// and a boolean to check if the value has been set.
func (o *FoundationalConfigsGET) GetFoundationalConfigsodataCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FoundationalConfigsodataCount, true
}

// SetFoundationalConfigsodataCount sets field value
func (o *FoundationalConfigsGET) SetFoundationalConfigsodataCount(v float32) {
	o.FoundationalConfigsodataCount = v
}

// GetOdataId returns the OdataId field value
func (o *FoundationalConfigsGET) GetOdataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value
// and a boolean to check if the value has been set.
func (o *FoundationalConfigsGET) GetOdataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataId, true
}

// SetOdataId sets field value
func (o *FoundationalConfigsGET) SetOdataId(v string) {
	o.OdataId = v
}

// GetOdataContext returns the OdataContext field value
func (o *FoundationalConfigsGET) GetOdataContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataContext
}

// GetOdataContextOk returns a tuple with the OdataContext field value
// and a boolean to check if the value has been set.
func (o *FoundationalConfigsGET) GetOdataContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataContext, true
}

// SetOdataContext sets field value
func (o *FoundationalConfigsGET) SetOdataContext(v string) {
	o.OdataContext = v
}

// GetOdataType returns the OdataType field value
func (o *FoundationalConfigsGET) GetOdataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdataType
}

// GetOdataTypeOk returns a tuple with the OdataType field value
// and a boolean to check if the value has been set.
func (o *FoundationalConfigsGET) GetOdataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdataType, true
}

// SetOdataType sets field value
func (o *FoundationalConfigsGET) SetOdataType(v string) {
	o.OdataType = v
}

func (o FoundationalConfigsGET) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FoundationalConfigsGET) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["FoundationalConfigs"] = o.FoundationalConfigs
	toSerialize["FoundationalConfigs@odata.count"] = o.FoundationalConfigsodataCount
	toSerialize["@odata.id"] = o.OdataId
	toSerialize["@odata.context"] = o.OdataContext
	toSerialize["@odata.type"] = o.OdataType
	return toSerialize, nil
}

func (o *FoundationalConfigsGET) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"FoundationalConfigs",
		"FoundationalConfigs@odata.count",
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

	varFoundationalConfigsGET := _FoundationalConfigsGET{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFoundationalConfigsGET)

	if err != nil {
		return err
	}

	*o = FoundationalConfigsGET(varFoundationalConfigsGET)

	return err
}

type NullableFoundationalConfigsGET struct {
	value *FoundationalConfigsGET
	isSet bool
}

func (v NullableFoundationalConfigsGET) Get() *FoundationalConfigsGET {
	return v.value
}

func (v *NullableFoundationalConfigsGET) Set(val *FoundationalConfigsGET) {
	v.value = val
	v.isSet = true
}

func (v NullableFoundationalConfigsGET) IsSet() bool {
	return v.isSet
}

func (v *NullableFoundationalConfigsGET) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFoundationalConfigsGET(val *FoundationalConfigsGET) *NullableFoundationalConfigsGET {
	return &NullableFoundationalConfigsGET{value: val, isSet: true}
}

func (v NullableFoundationalConfigsGET) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFoundationalConfigsGET) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


